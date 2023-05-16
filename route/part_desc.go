/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package route

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
)

const (
	ObInvalidPartId = -1
	ObPartIdBitNum  = 28
	ObPartIdShift   = 32
	ObMask          = (1 << ObPartIdBitNum) | 1<<(ObPartIdBitNum+ObPartIdShift)
	ObSubPartIdMask = 0xffffffff & 0xfffffff
)

type obColumnIndexesPair struct {
	column *obColumn
	// indexes indicates the column index in rowkey
	// eg1: create table t(c1 int, c2 int, c3 int, primary key (c1,c2,c3)) partition by hash(c2) partitions 10;
	// ##### column -> c2
	// ##### indexes -> [1], c2 is in the first position of the primary key
	// eg2: create table t(c1 varchar(20), c2 varchar(20), gen varchar(20) generated always as (concat(c1, c2), primary key (c1,c2)) partition by key(gen) partitions 10;
	// ##### column -> gen
	// ##### indexes -> [0, 1], gen depends on c1 and c2, which are in the first and second positions of the primary key.
	indexes []int
}

func NewObColumnIndexesPair(column *obColumn, indexes []int) *obColumnIndexesPair {
	return &obColumnIndexesPair{column, indexes}
}

func (p *obColumnIndexesPair) String() string {
	columnStr := "nil"
	if p.column != nil {
		columnStr = p.column.String()
	}
	var indexesStr string
	indexesStr = indexesStr + "["
	for i := 0; i < len(p.indexes); i++ {
		if i > 0 {
			indexesStr += ", "
		}
		indexesStr += strconv.Itoa(p.indexes[i])
	}
	indexesStr += "]"
	return "obColumnIndexesPair{" +
		"column:" + columnStr + ", " +
		"indexes:" + indexesStr +
		"}"
}

type obPartDescCommon struct {
	PartFuncType obPartFuncType
	PartExpr     string
	// orderedPartColumnNames Represents all partitioned column names
	// eg:
	//    partition by range(c1, c2)
	//    orderedPartColumnNames = ["c1", "c2"]
	OrderedPartColumnNames              []string
	OrderedPartRefColumnRowKeyRelations []*obColumnIndexesPair
	PartColumns                         []*obColumn
	RowKeyElement                       *table.ObRowKeyElement
}

func newObPartDescCommon(partFuncType obPartFuncType, partExpr string, orderedPartColumnNames []string) *obPartDescCommon {
	return &obPartDescCommon{
		PartFuncType:           partFuncType,
		PartExpr:               partExpr,
		OrderedPartColumnNames: orderedPartColumnNames,
	}
}

// setCommRowKeyElement set the primary key name and constructs the relationship between the partition key and primary key.
func (c *obPartDescCommon) setCommRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	c.RowKeyElement = rowKeyElement
	if len(c.OrderedPartColumnNames) != 0 && len(c.PartColumns) != 0 {
		relations := make([]*obColumnIndexesPair, 0, len(c.OrderedPartColumnNames))
		for _, partOrderColumnName := range c.OrderedPartColumnNames {
			for _, col := range c.PartColumns {
				if strings.EqualFold(col.columnName, partOrderColumnName) {
					partRefColumnRowKeyIndexes := make([]int, 0, len(col.refColumnNames))
					for _, refColumnName := range col.refColumnNames {
						for rowKeyElementName, index := range c.RowKeyElement.NameIdxMap() {
							if strings.EqualFold(rowKeyElementName, refColumnName) {
								partRefColumnRowKeyIndexes = append(partRefColumnRowKeyIndexes, index)
							}
						}
					}
					pair := NewObColumnIndexesPair(col, partRefColumnRowKeyIndexes)
					relations = append(relations, pair)
				}
			}
		}
		c.OrderedPartRefColumnRowKeyRelations = relations
	}
}

func (c *obPartDescCommon) CommString() string {
	// orderedPartRefColumnRowKeyRelations to string
	var relationsStr string
	relationsStr = relationsStr + "["
	for i, relation := range c.OrderedPartRefColumnRowKeyRelations {
		if i > 0 {
			relationsStr += ", "
		}
		relationsStr += relation.String()
	}
	relationsStr += "]"

	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(c.PartColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += c.PartColumns[i].String()
	}
	partColumnsStr += "]"

	// rowKeyElement to string
	rowKeyElementStr := "nil"
	if c.RowKeyElement != nil {
		rowKeyElementStr = c.RowKeyElement.String()
	}

	return "obPartDescCommon{" +
		"partFuncType:" + strconv.Itoa(int(c.PartFuncType)) + ", " +
		"partExpr:" + c.PartExpr + ", " +
		"orderedPartColumnNames:" + strings.Join(c.OrderedPartColumnNames, ",") + ", " +
		"orderedPartRefColumnRowKeyRelations:" + relationsStr + ", " +
		"partColumns:" + partColumnsStr + ", " +
		"rowKeyElement:" + rowKeyElementStr +
		"}"
}

type obPartDesc interface {
	String() string
	partFuncType() obPartFuncType
	orderedPartColumnNames() []string
	orderedPartRefColumnRowKeyRelations() []*obColumnIndexesPair
	rowKeyElement() *table.ObRowKeyElement
	setRowKeyElement(rowKeyElement *table.ObRowKeyElement)
	setPartColumns(partColumns []*obColumn)
	SetPartNum(partNum int)
	GetPartId(rowKey []interface{}) (int64, error)
}

// evalPartKeyValues calculate the value of the partition key
func evalPartKeyValues(desc obPartDesc, rowKey []interface{}) ([]interface{}, error) {
	if desc == nil {
		return nil, errors.New("part desc is nil")
	}
	if desc.rowKeyElement() == nil {
		return nil, errors.New("rowKey element is nil")
	}
	if len(rowKey) < len(desc.rowKeyElement().NameIdxMap()) {
		return nil, errors.Errorf("rowKey count not match, "+
			"rowKey len:%d, rowKeyElement len:%d", len(rowKey), len(desc.rowKeyElement().NameIdxMap()))
	}
	partRefColumnSize := len(desc.orderedPartRefColumnRowKeyRelations())
	evalValues := make([]interface{}, 0, partRefColumnSize)
	for i := 0; i < partRefColumnSize; i++ {
		relation := desc.orderedPartRefColumnRowKeyRelations()[i]
		evalParams := make([]interface{}, len(relation.indexes))
		for i := 0; i < len(relation.indexes); i++ {
			evalParams[i] = rowKey[relation.indexes[i]]
		}
		val, err := relation.column.EvalValue(evalParams...)
		if err != nil {
			return nil, errors.WithMessagef(err, "eval column value, relations:%s", relation.String())
		}
		evalValues = append(evalValues, val)
	}
	return evalValues, nil
}

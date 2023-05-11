package route

import (
	"errors"
	"strconv"
	"strings"

	"github.com/oceanbase/obkv-table-client-go/log"
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
	column  *obColumn
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
	for i := 0; i < len(c.OrderedPartRefColumnRowKeyRelations); i++ {
		if i > 0 {
			relationsStr += ", "
		}
		relationsStr += c.OrderedPartRefColumnRowKeyRelations[i].String()
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
	setOrderedPartColumnNames(partExpr string)
	orderedPartRefColumnRowKeyRelations() []*obColumnIndexesPair
	rowKeyElement() *table.ObRowKeyElement
	setRowKeyElement(rowKeyElement *table.ObRowKeyElement)
	setPartColumns(partColumns []*obColumn)
	GetPartId(rowKey []interface{}) (int64, error)
}

func evalPartKeyValues(desc obPartDesc, rowKey []interface{}) ([]interface{}, error) {
	if desc == nil {
		log.Warn("part desc is nil")
		return nil, errors.New("part desc is nil")
	}
	if desc.rowKeyElement() == nil {
		log.Warn("rowKey element is nil")
		return nil, errors.New("rowKey element is nil")
	}
	if len(rowKey) < len(desc.rowKeyElement().NameIdxMap()) {
		log.Warn("rowKey count not match",
			log.Int("rowKey len", len(rowKey)),
			log.Int("rowKeyElement len", len(desc.rowKeyElement().NameIdxMap())))
		return nil, errors.New("rowKey count not match")
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
			log.Warn("fail to eval column value", log.String("relations", relation.String()))
			return nil, err
		}
		evalValues = append(evalValues, val)
	}
	return evalValues, nil
}

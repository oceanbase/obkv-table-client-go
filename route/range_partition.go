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

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type obRangePartDesc struct {
	obPartDescCommon
	partSpace                 int
	partNum                   int
	orderedCompareColumns     []*obColumn
	orderedCompareColumnTypes []protocol.ObjType
}

func (d *obRangePartDesc) SetPartNum(partNum int) {
	d.partNum = partNum
}

func newObRangePartDesc() *obRangePartDesc {
	return &obRangePartDesc{}
}

func (d *obRangePartDesc) partFuncType() obPartFuncType {
	return d.PartFuncType
}

func (d *obRangePartDesc) orderedPartColumnNames() []string {
	return d.OrderedPartColumnNames
}

func (d *obRangePartDesc) setOrderedPartColumnNames(partExpr string) {
	// eg:"c1, c2", need to remove ' '
	str := strings.ReplaceAll(partExpr, " ", "")
	d.OrderedPartColumnNames = strings.Split(str, ",")
}

func (d *obRangePartDesc) orderedPartRefColumnRowKeyRelations() []*obColumnIndexesPair {
	return d.OrderedPartRefColumnRowKeyRelations
}
func (d *obRangePartDesc) rowKeyElement() *table.ObRowKeyElement {
	return d.RowKeyElement
}

func (d *obRangePartDesc) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	d.setCommRowKeyElement(rowKeyElement)
}

func (d *obRangePartDesc) setPartColumns(partColumns []*obColumn) {
	d.PartColumns = partColumns
}

func (d *obRangePartDesc) GetPartId(rowKey []interface{}) (int64, error) {
	// todo: impl
	return ObInvalidPartId, errors.New("not implement")
}

func (d *obRangePartDesc) String() string {
	// orderedCompareColumns to string
	var orderedCompareColumnsStr string
	orderedCompareColumnsStr = orderedCompareColumnsStr + "["
	for i := 0; i < len(d.orderedCompareColumns); i++ {
		if i > 0 {
			orderedCompareColumnsStr += ", "
		}
		orderedCompareColumnsStr += d.orderedCompareColumns[i].String()
	}
	orderedCompareColumnsStr += "]"

	// orderedCompareColumnTypes to string
	var orderedCompareColumnTypesStr string
	orderedCompareColumnTypesStr = orderedCompareColumnTypesStr + "["
	for i := 0; i < len(d.orderedCompareColumns); i++ {
		if i > 0 {
			orderedCompareColumnTypesStr += ", "
		}
		orderedCompareColumnTypesStr += d.orderedCompareColumnTypes[i].String()
	}
	orderedCompareColumnTypesStr += "]"

	return "obRangePartDesc{" +
		"comm:" + d.CommString() + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"orderedCompareColumns:" + orderedCompareColumnsStr + ", " +
		"orderedCompareColumnTypes:" + orderedCompareColumnTypesStr +
		"}"
}

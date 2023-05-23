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

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
)

// newObRangePartDesc create a range partition description.
func newObRangePartDesc(
	partSpace int,
	partNum int,
	partFuncType obPartFuncType) *obRangePartDesc {
	return &obRangePartDesc{
		partFuncType: partFuncType,
		partSpace:    partSpace,
		partNum:      partNum,
	}
}

// obRangePartDesc description of the range partition.
type obRangePartDesc struct {
	partFuncType obPartFuncType
	partSpace    int
	partNum      int
	partColumns  []obColumn
}

func (d *obRangePartDesc) PartColumns() []obColumn {
	return d.partColumns
}

func (d *obRangePartDesc) SetPartNum(partNum int) {
	d.partNum = partNum
}

func (d *obRangePartDesc) PartFuncType() obPartFuncType {
	return d.partFuncType
}

func (d *obRangePartDesc) SetPartColumns(partColumns []obColumn) {
	d.partColumns = partColumns
}

// GetPartId get partition id by rowKey.
// Not support range partition now.
func (d *obRangePartDesc) GetPartId(rowKey []*table.Column) (int64, error) {
	return ObInvalidPartId, errors.New("not support range partition now")
}

func (d *obRangePartDesc) String() string {
	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(d.partColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += d.partColumns[i].String()
	}
	partColumnsStr += "]"

	return "obRangePartDesc{" +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partColumns" + partColumnsStr +
		"}"
}

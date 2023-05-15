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

// obHashPartDesc description of the hash partition.
type obHashPartDesc struct {
	*obPartDescCommon
	completeWorks []int64 // all partition id, use in query
	partSpace     int
	partNum       int
}

func (d *obHashPartDesc) SetPartNum(partNum int) {
	d.partNum = partNum
}

func newObHashPartDesc(
	partSpace int,
	partNum int,
	partFuncType obPartFuncType,
	partExpr string) *obHashPartDesc {
	// eg:"c1, c2", need to remove ' '
	str := strings.ReplaceAll(partExpr, " ", "")
	orderedPartColumnNames := strings.Split(str, ",")
	return &obHashPartDesc{
		obPartDescCommon: newObPartDescCommon(partFuncType, partExpr, orderedPartColumnNames),
		partSpace:        partSpace,
		partNum:          partNum,
	}
}

func (d *obHashPartDesc) partFuncType() obPartFuncType {
	return d.PartFuncType
}

func (d *obHashPartDesc) orderedPartColumnNames() []string {
	return d.OrderedPartColumnNames
}

func (d *obHashPartDesc) orderedPartRefColumnRowKeyRelations() []*obColumnIndexesPair {
	return d.OrderedPartRefColumnRowKeyRelations
}

func (d *obHashPartDesc) rowKeyElement() *table.ObRowKeyElement {
	return d.RowKeyElement
}

func (d *obHashPartDesc) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	d.setCommRowKeyElement(rowKeyElement)
}

func (d *obHashPartDesc) setPartColumns(partColumns []*obColumn) {
	d.PartColumns = partColumns
}

// GetPartId get partition id by inner hash function.
func (d *obHashPartDesc) GetPartId(rowKey []interface{}) (int64, error) {
	if len(rowKey) == 0 {
		return ObInvalidPartId, errors.New("rowKey size is 0")
	}
	evalValues, err := evalPartKeyValues(d, rowKey)
	if err != nil {
		return ObInvalidPartId, errors.WithMessagef(err, "eval partition key value, partDesc:%s", d.String())
	}
	longValue, err := parseToNumber(evalValues[0]) // hash part has one param at most
	if err != nil {
		return ObInvalidPartId, errors.WithMessagef(err, "parse long, partDesc:%s", d.String())
	}
	if v, ok := longValue.(int64); !ok {
		return ObInvalidPartId, errors.Errorf("failed to convert to long, value:%T", longValue)
	} else {
		return d.innerHash(v), nil
	}
}

func (d *obHashPartDesc) innerHash(hashVal int64) int64 {
	// abs(hashVal)
	if hashVal < 0 {
		hashVal = -hashVal
	}
	return (int64(d.partSpace) << ObPartIdBitNum) | (hashVal % int64(d.partNum))
}

func (d *obHashPartDesc) String() string {
	// completeWorks to string
	var completeWorksStr string
	completeWorksStr = completeWorksStr + "["
	for i := 0; i < len(d.completeWorks); i++ {
		if i > 0 {
			completeWorksStr += ", "
		}
		completeWorksStr += strconv.Itoa(int(d.completeWorks[i]))
	}
	completeWorksStr += "]"

	return "obHashPartDesc{" +
		"comm:" + d.CommString() + ", " +
		"completeWorks:" + completeWorksStr + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) +
		"}"
}

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

// newObHashPartDesc create a hash partition description.
func newObHashPartDesc(
	partSpace int,
	partNum int,
	partFuncType obPartFuncType) *obHashPartDesc {
	return &obHashPartDesc{
		partFuncType: partFuncType,
		partSpace:    partSpace,
		partNum:      partNum,
	}
}

// obHashPartDesc description of the hash partition.
type obHashPartDesc struct {
	partFuncType  obPartFuncType
	completeWorks []int64 // all partition id, use in query
	partSpace     int
	partNum       int
	partColumns   []obColumn
}

func (d *obHashPartDesc) PartColumns() []obColumn {
	return d.partColumns
}

func (d *obHashPartDesc) PartNum() int {
	return d.partNum
}

func (d *obHashPartDesc) SetPartNum(partNum int) {
	d.partNum = partNum
}

func (d *obHashPartDesc) PartFuncType() obPartFuncType {
	return d.partFuncType
}

func (d *obHashPartDesc) SetPartColumns(partColumns []obColumn) {
	d.partColumns = partColumns
}

// GetPartId get partition id by inner hash function.
func (d *obHashPartDesc) GetPartId(rowKey []*table.Column) (uint64, error) {
	if len(rowKey) == 0 {
		return ObInvalidPartId, errors.New("rowKey size is 0")
	}
	evalValues, err := evalPartKeyValues(d, rowKey)
	if err != nil {
		return ObInvalidPartId, errors.WithMessagef(err, "eval partition key value, partDesc:%s", d.String())
	}
	if len(evalValues) == 0 {
		return ObInvalidPartId, errors.New("invalid eval values length")
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

func (d *obHashPartDesc) GetPartIds(rowKeyPair *table.RangePair) ([]uint64, error) {
	if rowKeyPair == nil {
		return nil, errors.New("rowKeyPair is nil")
	}
	if rowKeyPair.Start() == nil || rowKeyPair.End() == nil {
		return nil, errors.New("startKeys or endKeys in rangePair is nil")
	}
	if rowKeyPair.IsStartEqEnd() {
		// startKey == endKey means that the range is equal to a column
		partId, err := d.GetPartId(rowKeyPair.Start())
		if err != nil {
			return []uint64{ObInvalidPartId}, errors.WithMessagef(err, "get part id, part desc:%s", d.String())
		}
		return []uint64{partId}, nil
	} else {
		// if startKey != endKey, add all partitions to the partition list
		partIds := make([]uint64, 0, d.partNum)
		for i := 0; i < d.partNum; i++ {
			partIds = append(partIds, uint64(i))
		}
		return partIds, nil
	}
}

// innerHash hash method for computing partition id
func (d *obHashPartDesc) innerHash(hashVal int64) uint64 {
	// abs(hashVal)
	if hashVal < 0 {
		hashVal = -hashVal
	}
	return uint64((int64(d.partSpace) << ObPartIdBitNum) | (hashVal % int64(d.partNum)))
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

	return "obHashPartDesc{" +
		"completeWorks:" + completeWorksStr + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partColumns" + partColumnsStr +
		"}"
}

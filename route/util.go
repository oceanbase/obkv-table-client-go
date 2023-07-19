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
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/pkg/errors"
	"strconv"
)

// createInStatement create "(0,1,2...partNum);" string, is used by route model.
func createInStatement(values []uint64) string {
	// Create inStatement "(0,1,2...partNum);".
	var inStatement string
	inStatement += "("
	for i, v := range values {
		if i > 0 {
			inStatement += ", "
		}
		inStatement += strconv.FormatUint(v, 10)
	}
	inStatement += ");"
	return inStatement
}

// murmurHash64A is same as function hash64a() in java jdk.
func murmurHash64A(data []byte, length int, seed int64) int64 {
	const (
		m = 0xc6a4a7935bd1e995
		r = 47
	)
	var k int64
	h := seed ^ int64(uint64(length)*m)

	var um uint64 = m
	var im = int64(um)
	for l := length; l >= 8; l -= 8 {
		k = int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56

		k := k * im
		k ^= int64(uint64(k) >> r)
		k = k * im

		h = h ^ k
		h = h * im
		data = data[8:]
	}

	switch length & 7 {
	case 7:
		h ^= int64(data[6]) << 48
		fallthrough
	case 6:
		h ^= int64(data[5]) << 40
		fallthrough
	case 5:
		h ^= int64(data[4]) << 32
		fallthrough
	case 4:
		h ^= int64(data[3]) << 24
		fallthrough
	case 3:
		h ^= int64(data[2]) << 16
		fallthrough
	case 2:
		h ^= int64(data[1]) << 8
		fallthrough
	case 1:
		h ^= int64(data[0])
		h *= im
	}

	h ^= int64(uint64(h) >> r)
	h *= im
	h ^= int64(uint64(h) >> r)
	return h
}

func parseToNumber(value interface{}) (interface{}, error) {
	if v, ok := value.(string); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		return int64(i), nil
	} else if v, ok := value.(int64); ok {
		return v, nil
	} else if v, ok := value.(int32); ok {
		return int64(v), nil
	} else if v, ok := value.(int16); ok {
		return int64(v), nil
	} else if v, ok := value.(int8); ok {
		return int64(v), nil
	} else {
		return nil, errors.New("invalid type to convert to number")
	}
}

// extractRangePairPartKeyColumns extract part key columns from range pair.
// Notes: if some partition key is missing in startPartColumns or endPartColumns, it will be nil in that position.
func extractRangePairPartKeyColumns(d obPartDesc, rowKeyPair *table.RangePair) ([]*table.Column, []*table.Column, error) {
	if d == nil {
		return nil, nil, errors.New("part desc is nil")
	}
	startPartColumns := make([]*table.Column, 0, len(d.PartColumns()))
	endPartColumns := make([]*table.Column, 0, len(d.PartColumns()))
	for _, column := range d.PartColumns() {
		startPartColumn, err := column.extractColumn(rowKeyPair.Start())
		if err != nil {
			return nil, nil, errors.WithMessagef(err, "extract start part column, column:%s", column.String())
		}
		startPartColumns = append(startPartColumns, startPartColumn)

		endPartColumn, err := column.extractColumn(rowKeyPair.End())
		if err != nil {
			return nil, nil, errors.WithMessagef(err, "extract end part column, column:%s", column.String())
		}
		endPartColumns = append(endPartColumns, endPartColumn)
	}
	return startPartColumns, endPartColumns, nil
}

// isQuerySendAll check if a pair of Partition columns from RangePair need send to all partitions.
// Notes: this function will not check if some partition key is missing in startPartColumns and endPartColumns
func checkQueryPkSendAll(startPartColumns []*table.Column, endPartColumns []*table.Column) bool {
	// If startPartColumns and endPartColumns is nil, it means something wrong o.O
	if startPartColumns == nil || endPartColumns == nil {
		return true
	}
	// If startPartColumns and endPartColumns is empty, it also means something wrong O.o
	if len(startPartColumns) == 0 && len(endPartColumns) == 0 {
		return true
	}
	// If startPartColumns and endPartColumns is not empty, but length not equal, it means send to all partitions.
	if len(startPartColumns) != len(endPartColumns) {
		return true
	}

	// check part columns
	for i := 0; i < len(startPartColumns); i++ {
		// if startPartColumn or endPartColumn is nil, it means send to all partitions.
		if startPartColumns[i] == nil || endPartColumns[i] == nil {
			return true
		}
		// if any part column is extremum, it means send to all partitions.
		if _, ok := startPartColumns[i].Value().(table.Extremum); ok {
			return true
		}
		if _, ok := endPartColumns[i].Value().(table.Extremum); ok {
			return true
		}
		// if startPartColumn and endPartColumn is not equal, it means send to all partitions.
		// Notes: we assume that the startPartColumn and endPartColumn is "ordered" (ordered like obPartDesc.PartColumns),
		//        if you generate key by extractRangePairPartKeyColumns() function, it will be fine.
		if !startPartColumns[i].IsEqual(endPartColumns[i]) {
			return true
		}
	}

	return false
}

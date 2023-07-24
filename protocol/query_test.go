/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObTableQueryEncodeDecode(t *testing.T) {
	util.SetObVersion(4)
	obTableQuery := NewObTableQuery()

	randomLen := rand.Intn(100)
	obNewRanges := make([]*ObNewRange, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		obNewRange := NewObNewRange()
		obNewRange.SetTableId(rand.Uint64())
		obNewRange.SetBorderFlag(ObBorderFlag(rand.Intn(255)))
		randomLen = rand.Intn(100)
		startKey := make([]*ObObject, 0, randomLen)
		endKey := make([]*ObObject, 0, randomLen)
		columns := make([]*table.Column, 0, randomLen)
		for i := 0; i < randomLen; i++ {
			columns = append(columns, table.NewColumn(util.String(10), int64(rand.Intn(10000))))
		}
		for _, column := range columns {
			objMeta, _ := DefaultObjMeta(column.Value())
			startKey = append(startKey, NewObObjectWithParams(objMeta, column.Value()))
			endKey = append(endKey, NewObObjectWithParams(objMeta, column.Value()))
		}
		obNewRange.SetStartKey(startKey)
		obNewRange.SetEndKey(endKey)
		obNewRange.SetFlag(int64(rand.Uint64()))
		obNewRanges = append(obNewRanges, obNewRange)
	}
	obTableQuery.SetKeyRanges(obNewRanges)

	selectColumns := make([]string, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		selectColumns = append(selectColumns, util.String(10))
	}
	obTableQuery.SetSelectColumns(selectColumns)

	obTableQuery.SetFilterString(util.String(rand.Intn(10)))
	obTableQuery.SetLimit(int32(rand.Uint32()))
	obTableQuery.SetOffset(int32(rand.Uint32()))
	obTableQuery.SetScanOrder(ObScanOrder(rand.Intn(255)))
	obTableQuery.SetIndexName(util.String(rand.Intn(10)))
	obTableQuery.SetBatchSize(int32(rand.Uint32()))
	obTableQuery.SetMaxResultSize(int64(rand.Uint64()))
	obTableQuery.SetIsHbaseQuery(true)

	obHTableFilter := NewObHTableFilter()
	obHTableFilter.SetVersion(1)
	obHTableFilter.SetContentLength(0)
	obHTableFilter.SetIsValid(util.ByteToBool(byte(rand.Intn(2))))
	selectColumnQualifierLen := rand.Intn(10)
	selectColumnQualifier := make([][]byte, 0, rand.Intn(selectColumnQualifierLen))
	for i := 0; i < selectColumnQualifierLen; i++ {
		selectColumnQualifier = append(selectColumnQualifier, []byte(util.String(10)))
	}
	obHTableFilter.SetSelectColumnQualifier(selectColumnQualifier)
	obHTableFilter.SetMinStamp(int64(rand.Uint64()))
	obHTableFilter.SetMaxStamp(int64(rand.Uint64()))
	obHTableFilter.SetMaxVersions(int32(rand.Uint32()))
	obHTableFilter.SetLimitPerRowPerCf(int32(rand.Uint32()))
	obHTableFilter.SetOffsetPerRowPerCf(int32(rand.Uint32()))
	obHTableFilter.SetFilterString(util.String(10))
	obTableQuery.SetHTableFilter(obHTableFilter)

	scanRangeColumns := make([]string, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		selectColumns = append(selectColumns, util.String(10))
	}
	obTableQuery.SetScanRangeColumns(scanRangeColumns)

	aggregations := make([]*ObTableAggregation, 0, randomLen)
	obTableQuery.SetAggregations(aggregations)

	payloadLen := obTableQuery.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableQuery.Encode(buffer)

	newObTableQuery := NewObTableQuery()
	newObTableQuery.SetIsHbaseQuery(true)
	newBuffer := bytes.NewBuffer(buf)
	newObTableQuery.Decode(newBuffer)

	assert.EqualValues(t, obTableQuery.KeyRanges(), newObTableQuery.KeyRanges())
	assert.EqualValues(t, obTableQuery.SelectColumns(), newObTableQuery.SelectColumns())
	assert.EqualValues(t, obTableQuery.FilterString(), newObTableQuery.FilterString())
	assert.EqualValues(t, obTableQuery.Limit(), newObTableQuery.Limit())
	assert.EqualValues(t, obTableQuery.Offset(), newObTableQuery.Offset())
	assert.EqualValues(t, obTableQuery.ScanOrder(), newObTableQuery.ScanOrder())
	assert.EqualValues(t, obTableQuery.IndexName(), newObTableQuery.IndexName())
	assert.EqualValues(t, obTableQuery.BatchSize(), newObTableQuery.BatchSize())
	assert.EqualValues(t, obTableQuery.MaxResultSize(), newObTableQuery.MaxResultSize())
	assert.EqualValues(t, obTableQuery.IsHbaseQuery(), newObTableQuery.IsHbaseQuery())
	assert.EqualValues(t, obTableQuery.HTableFilter(), newObTableQuery.HTableFilter())
	assert.EqualValues(t, obTableQuery.ScanRangeColumns(), newObTableQuery.ScanRangeColumns())
	assert.EqualValues(t, obTableQuery.Aggregations(), newObTableQuery.Aggregations())
	assert.EqualValues(t, obTableQuery, newObTableQuery)
}

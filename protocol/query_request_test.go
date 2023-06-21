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

func TestObTableQueryRequestEncodeDecode(t *testing.T) {
	util.SetObVersion(4)
	obTableQueryRequest := NewObTableQueryRequest()
	obTableQueryRequest.SetCredential([]byte(util.String(10)))
	obTableQueryRequest.SetTableName(util.String(10))
	obTableQueryRequest.SetTableId(rand.Uint64())
	obTableQueryRequest.SetPartitionId(rand.Uint64())
	obTableQueryRequest.SetEntityType(ObTableEntityType(rand.Intn(255)))
	obTableQueryRequest.SetConsistencyLevel(ObTableConsistencyLevel(rand.Intn(255)))
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

	aggregations := make([]*ObTableAggregationSingle, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		obTableAggregationSingle := NewObTableAggregationSingle()
		obTableAggregationSingle.SetVersion(1)
		obTableAggregationSingle.SetContentLength(0)
		obTableAggregationSingle.SetAggType(ObTableAggregationType(rand.Intn(255)))
		obTableAggregationSingle.SetAggColumn(util.String(10))
	}
	obTableQuery.SetAggregations(aggregations)
	obTableQueryRequest.SetTableQuery(obTableQuery)

	payloadLen := obTableQueryRequest.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableQueryRequest.Encode(buffer)

	newObTableQueryRequest := NewObTableQueryRequest()
	newObTableQueryRequest.TableQuery().SetIsHbaseQuery(true)
	newBuffer := bytes.NewBuffer(buf)
	newObTableQueryRequest.Decode(newBuffer)

	assert.EqualValues(t, newObTableQueryRequest.Credential(), obTableQueryRequest.Credential())
	assert.EqualValues(t, newObTableQueryRequest.TableName(), obTableQueryRequest.TableName())
	assert.EqualValues(t, newObTableQueryRequest.TableId(), obTableQueryRequest.TableId())
	assert.EqualValues(t, newObTableQueryRequest.PartitionId(), obTableQueryRequest.PartitionId())
	assert.EqualValues(t, newObTableQueryRequest.EntityType(), obTableQueryRequest.EntityType())
	assert.EqualValues(t, newObTableQueryRequest.ConsistencyLevel(), obTableQueryRequest.ConsistencyLevel())
	assert.EqualValues(t, newObTableQueryRequest.TableQuery(), obTableQueryRequest.TableQuery())
	assert.EqualValues(t, newObTableQueryRequest, obTableQueryRequest)
}

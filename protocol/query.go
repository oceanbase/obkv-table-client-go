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
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS WITHOUT WARRANTIES OF ANY KIND
 * EITHER EXPRESS OR IMPLIED INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableQuery struct {
	ObUniVersionHeader
	keyRanges        []*ObNewRange
	selectColumns    []string
	filterString     string
	limit            int32
	offset           int32
	scanOrder        ObScanOrder
	indexName        string
	batchSize        int32
	maxResultSize    int64
	isHbaseQuery     bool
	hTableFilter     *ObHTableFilter
	scanRangeColumns []string
	aggregations     []*ObTableAggregationSingle
}

// NewObTableQueryWithParams creates a new ObTableQuery with parameters.
func NewObTableQueryWithParams(batchSize int32) *ObTableQuery {
	return &ObTableQuery{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		keyRanges:        nil,
		selectColumns:    nil,
		filterString:     "",
		limit:            -1,
		offset:           0,
		scanOrder:        ObScanOrderForward,
		indexName:        "",
		batchSize:        batchSize,
		maxResultSize:    -1,
		isHbaseQuery:     false,
		hTableFilter:     nil,
		scanRangeColumns: nil,
		aggregations:     nil,
	}
}

var hTableFilterDummyBytes = []byte{0x01, 0x00}

func (q *ObTableQuery) KeyRanges() []*ObNewRange {
	return q.keyRanges
}

func (q *ObTableQuery) SetKeyRanges(keyRanges []*ObNewRange) {
	q.keyRanges = keyRanges
}

func (q *ObTableQuery) SelectColumns() []string {
	return q.selectColumns
}

func (q *ObTableQuery) SetSelectColumns(selectColumns []string) {
	q.selectColumns = selectColumns
}

func (q *ObTableQuery) FilterString() string {
	return q.filterString
}

func (q *ObTableQuery) SetFilterString(filterString string) {
	q.filterString = filterString
}

func (q *ObTableQuery) Limit() int32 {
	return q.limit
}

func (q *ObTableQuery) SetLimit(limit int32) {
	q.limit = limit
}

func (q *ObTableQuery) Offset() int32 {
	return q.offset
}

func (q *ObTableQuery) SetOffset(offset int32) {
	q.offset = offset
}

func (q *ObTableQuery) ScanOrder() ObScanOrder {
	return q.scanOrder
}

func (q *ObTableQuery) SetScanOrder(scanOrder ObScanOrder) {
	q.scanOrder = scanOrder
}

func (q *ObTableQuery) IndexName() string {
	return q.indexName
}

func (q *ObTableQuery) SetIndexName(indexName string) {
	q.indexName = indexName
}

func (q *ObTableQuery) BatchSize() int32 {
	return q.batchSize
}

func (q *ObTableQuery) SetBatchSize(batchSize int32) {
	q.batchSize = batchSize
}

func (q *ObTableQuery) MaxResultSize() int64 {
	return q.maxResultSize
}

func (q *ObTableQuery) SetMaxResultSize(maxResultSize int64) {
	q.maxResultSize = maxResultSize
}

func (q *ObTableQuery) HTableFilter() *ObHTableFilter {
	return q.hTableFilter
}

func (q *ObTableQuery) SetHTableFilter(hTableFilter *ObHTableFilter) {
	q.hTableFilter = hTableFilter
}

func (q *ObTableQuery) IsHbaseQuery() bool {
	return q.isHbaseQuery
}

func (q *ObTableQuery) SetIsHbaseQuery(isHbaseQuery bool) {
	q.isHbaseQuery = isHbaseQuery
}

func (q *ObTableQuery) ScanRangeColumns() []string {
	return q.scanRangeColumns
}

func (q *ObTableQuery) SetScanRangeColumns(scanRangeColumns []string) {
	q.scanRangeColumns = scanRangeColumns
}

func (q *ObTableQuery) Aggregations() []*ObTableAggregationSingle {
	return q.aggregations
}

func (q *ObTableQuery) SetAggregations(aggregations []*ObTableAggregationSingle) {
	q.aggregations = aggregations
}

func (q *ObTableQuery) PayloadLen() int {
	return q.PayloadContentLen() + q.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (q *ObTableQuery) PayloadContentLen() int {
	totalLen := 0
	totalLen += util.EncodedLengthByVi64(int64(len(q.keyRanges)))
	for _, obNewRange := range q.keyRanges {
		totalLen += obNewRange.EncodedLength()
	}

	totalLen += util.EncodedLengthByVi64(int64(len(q.selectColumns)))
	for _, column := range q.selectColumns {
		totalLen += util.EncodedLengthByVString(column)
	}

	totalLen += util.EncodedLengthByVString(q.filterString) +
		util.EncodedLengthByVi32(q.limit) +
		util.EncodedLengthByVi32(q.offset) +
		1 + // scanOrder
		util.EncodedLengthByVString(q.indexName) +
		util.EncodedLengthByVi32(q.batchSize) +
		util.EncodedLengthByVi64(q.maxResultSize)

	if q.isHbaseQuery {
		totalLen += q.hTableFilter.PayloadLen()
	} else {
		totalLen += len(hTableFilterDummyBytes)
	}

	totalLen += util.EncodedLengthByVi64(int64(len(q.scanRangeColumns)))
	for _, column := range q.scanRangeColumns {
		util.EncodedLengthByVString(column)
	}

	totalLen += util.EncodedLengthByVi64(int64(len(q.aggregations)))
	for _, tableAggregationSingle := range q.aggregations {
		tableAggregationSingle.PayloadLen()
	}

	q.ObUniVersionHeader.SetContentLength(totalLen)
	return q.ObUniVersionHeader.ContentLength()
}

func (q *ObTableQuery) Encode(buffer *bytes.Buffer) {
	q.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(q.keyRanges)))
	for _, obNewRange := range q.keyRanges {
		obNewRange.Encode(buffer)
	}

	util.EncodeVi64(buffer, int64(len(q.selectColumns)))
	for _, column := range q.selectColumns {
		util.EncodeVString(buffer, column)
	}

	util.EncodeVString(buffer, q.filterString)

	util.EncodeVi32(buffer, q.limit)
	util.EncodeVi32(buffer, q.offset)

	util.PutUint8(buffer, uint8(q.scanOrder))

	util.EncodeVString(buffer, q.indexName)

	util.EncodeVi32(buffer, q.batchSize)

	util.EncodeVi64(buffer, q.maxResultSize)

	if q.isHbaseQuery {
		q.hTableFilter.Encode(buffer)
	} else {
		copy(buffer.Next(len(hTableFilterDummyBytes)), hTableFilterDummyBytes)
	}

	util.EncodeVi64(buffer, int64(len(q.scanRangeColumns)))
	for _, column := range q.scanRangeColumns {
		util.EncodeVString(buffer, column)
	}

	util.EncodeVi64(buffer, int64(len(q.aggregations)))
	for _, tableAggregationSingle := range q.aggregations {
		tableAggregationSingle.Encode(buffer)
	}
}

func (q *ObTableQuery) Decode(buffer *bytes.Buffer) {
}

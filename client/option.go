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

package client

import (
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type ObkvOperationOption interface {
	apply(opts *ObkvOperationOptions)
}

type ObkvOperationOptionFunc func(opts *ObkvOperationOptions)

func NewOperationOptions() *ObkvOperationOptions {
	return &ObkvOperationOptions{
		returnRowKey:         false,
		returnAffectedEntity: false,
		tableFilter:          nil,
	}
}

type ObkvOperationOptions struct {
	returnRowKey         bool
	returnAffectedEntity bool
	tableFilter          filter.ObTableFilter
}

func (f ObkvOperationOptionFunc) apply(opts *ObkvOperationOptions) {
	f(opts)
}

// WithReturnRowKey only work in increment and append operation
func WithReturnRowKey(returnRowKey bool) ObkvOperationOption {
	return ObkvOperationOptionFunc(func(opts *ObkvOperationOptions) {
		opts.returnRowKey = returnRowKey
	})
}

// WithReturnAffectedEntity only work in increment and append operation
func WithReturnAffectedEntity(returnAffectedEntity bool) ObkvOperationOption {
	return ObkvOperationOptionFunc(func(opts *ObkvOperationOptions) {
		opts.returnAffectedEntity = returnAffectedEntity
	})
}

// WithFilter only work in increment append update and delete operation
func WithFilter(tableFilter filter.ObTableFilter) ObkvOperationOption {
	return ObkvOperationOptionFunc(func(opts *ObkvOperationOptions) {
		opts.tableFilter = tableFilter
	})
}

type ObkvQueryOption interface {
	apply(opts *ObkvQueryOptions)
}

type ObkvQueryOptionFunc func(opts *ObkvQueryOptions)

func NewObkvQueryOption() *ObkvQueryOptions {
	return &ObkvQueryOptions{
		queryFilter:   nil,
		hTableFilter:  nil,
		selectColumns: nil,
		indexName:     "",
		batchSize:     -1,
		maxResultSize: -1,
		limit:         -1,
		offset:        0,
		scanOrder:     table.Forward,
		isHbaseQuery:  false,
	}
}

type ObkvQueryOptions struct {
	queryFilter   interface{}
	hTableFilter  interface{}
	selectColumns []string
	indexName     string
	batchSize     int32
	maxResultSize int64
	limit         int32
	offset        int32
	scanOrder     table.ScanOrder
	isHbaseQuery  bool
}

func (f ObkvQueryOptionFunc) apply(opts *ObkvQueryOptions) {
	f(opts)
}

// WithQueryFilter set query filter
func WithQueryFilter(queryFilter interface{}) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.queryFilter = queryFilter
	})
}

// WithHTableFilter set htable filter
func WithHTableFilter(hTableFilter interface{}) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.hTableFilter = hTableFilter
	})
}

// WithSelectColumns set select columns
func WithSelectColumns(selectColumns []string) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.selectColumns = selectColumns
	})
}

// WithIndexName set index name
func WithIndexName(indexName string) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.indexName = indexName
	})
}

// WithBatchSize set batch size
func WithBatchSize(batchSize int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.batchSize = int32(batchSize)
	})
}

// WithMaxResultSize set max result size
func WithMaxResultSize(maxResultSize int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.maxResultSize = int64(maxResultSize)
	})
}

// WithLimit set limit
func WithLimit(limit int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.limit = int32(limit)
	})
}

// WithOffset set offset
func WithOffset(offset int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.offset = int32(offset)
	})
}

// WithScanOrder set scan order
func WithScanOrder(scanOrder table.ScanOrder) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.scanOrder = scanOrder
	})
}

// WithIsHbaseQuery set is hbase query
func WithIsHbaseQuery(isHbaseQuery bool) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.isHbaseQuery = isHbaseQuery
	})
}

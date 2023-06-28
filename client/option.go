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

type OperationOption interface {
	apply(operationOpts *OperationOptions)
}

type OperationOptionFunc func(operationOpts *OperationOptions)

func NewOperationOptions() *OperationOptions {
	return &OperationOptions{
		returnRowKey:         false,
		returnAffectedEntity: false,
		tableFilter:          nil,
	}
}

type OperationOptions struct {
	returnRowKey         bool
	returnAffectedEntity bool
	tableFilter          filter.ObTableFilter
}

func (f OperationOptionFunc) apply(operationOpts *OperationOptions) {
	f(operationOpts)
}

// WithReturnRowKey only work in increment and append operation
func WithReturnRowKey(returnRowKey bool) OperationOption {
	return OperationOptionFunc(func(operationOpts *OperationOptions) {
		operationOpts.returnRowKey = returnRowKey
	})
}

// WithReturnAffectedEntity only work in increment and append operation
func WithReturnAffectedEntity(returnAffectedEntity bool) OperationOption {
	return OperationOptionFunc(func(operationOpts *OperationOptions) {
		operationOpts.returnAffectedEntity = returnAffectedEntity
	})
}

// WithFilter only work in increment append update and delete operation
func WithFilter(tableFilter filter.ObTableFilter) OperationOption {
	return OperationOptionFunc(func(operationOpts *OperationOptions) {
		operationOpts.tableFilter = tableFilter
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

// SetQueryFilter set query filter
func SetQueryFilter(queryFilter interface{}) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.queryFilter = queryFilter
	})
}

// SetHTableFilter set htable filter
func SetHTableFilter(hTableFilter interface{}) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.hTableFilter = hTableFilter
	})
}

// SetSelectColumns set select columns
func SetSelectColumns(selectColumns []string) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.selectColumns = selectColumns
	})
}

// SetIndexName set index name
func SetIndexName(indexName string) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.indexName = indexName
	})
}

// SetBatchSize set batch size
func SetBatchSize(batchSize int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.batchSize = int32(batchSize)
	})
}

// SetMaxResultSize set max result size
func SetMaxResultSize(maxResultSize int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.maxResultSize = int64(maxResultSize)
	})
}

// SetLimit set limit
func SetLimit(limit int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.limit = int32(limit)
	})
}

// SetOffset set offset
func SetOffset(offset int) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.offset = int32(offset)
	})
}

// SetScanOrder set scan order
func SetScanOrder(scanOrder table.ScanOrder) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.scanOrder = scanOrder
	})
}

// SetIsHbaseQuery set is hbase query
func SetIsHbaseQuery(isHbaseQuery bool) ObkvQueryOption {
	return ObkvQueryOptionFunc(func(opts *ObkvQueryOptions) {
		opts.isHbaseQuery = isHbaseQuery
	})
}

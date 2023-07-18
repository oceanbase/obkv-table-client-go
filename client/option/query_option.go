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
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LimitED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package option

import "github.com/oceanbase/obkv-table-client-go/table"

type ObQueryOption interface {
	Apply(opts *ObQueryOptions)
}

type ObQueryOptionFunc func(opts *ObQueryOptions)

func NewObQueryOption() *ObQueryOptions {
	return &ObQueryOptions{
		QueryFilter:   nil,
		HTableFilter:  nil,
		SelectColumns: nil,
		IndexName:     "",
		BatchSize:     -1,
		MaxResultSize: -1,
		Limit:         -1,
		Offset:        0,
		ScanOrder:     table.Forward,
		IsHbaseQuery:  false,
	}
}

type ObQueryOptions struct {
	QueryFilter   interface{}
	HTableFilter  interface{}
	SelectColumns []string
	IndexName     string
	BatchSize     int32
	MaxResultSize int64
	Limit         int32
	Offset        int32
	ScanOrder     table.ScanOrder
	IsHbaseQuery  bool
}

func (f ObQueryOptionFunc) Apply(opts *ObQueryOptions) {
	f(opts)
}

// WithQueryFilter set query filter
func WithQueryFilter(QueryFilter interface{}) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.QueryFilter = QueryFilter
	})
}

// WithHTableFilter set htable filter
func WithHTableFilter(HTableFilter interface{}) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.HTableFilter = HTableFilter
	})
}

// WithSelectColumns set select columns
func WithSelectColumns(SelectColumns []string) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.SelectColumns = SelectColumns
	})
}

// WithIndexName set index name
func WithIndexName(IndexName string) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.IndexName = IndexName
	})
}

// WithBatchSize set batch size
func WithBatchSize(BatchSize int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.BatchSize = int32(BatchSize)
	})
}

// WithMaxResultSize set max result size
func WithMaxResultSize(MaxResultSize int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.MaxResultSize = int64(MaxResultSize)
	})
}

// WithLimit set Limit
func WithLimit(Limit int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.Limit = int32(Limit)
	})
}

// WithOffset set Offset
func WithOffset(Offset int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.Offset = int32(Offset)
	})
}

// WithScanOrder set scan order
func WithScanOrder(ScanOrder table.ScanOrder) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.ScanOrder = ScanOrder
	})
}

// WithIsHbaseQuery set is hbase query
func WithIsHbaseQuery(IsHbaseQuery bool) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.IsHbaseQuery = IsHbaseQuery
	})
}

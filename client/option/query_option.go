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

import (
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/hfilter"
	"github.com/oceanbase/obkv-table-client-go/table"
)

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
		KeyValueMode:  table.DynamicMode,
	}
}

type ObQueryOptions struct {
	QueryFilter   filter.ObTableFilter
	HTableFilter  hfilter.ObHTableFilter
	SelectColumns []string
	IndexName     string
	BatchSize     int32
	MaxResultSize int64
	Limit         int32
	Offset        int32
	ScanOrder     table.ScanOrder
	IsHbaseQuery  bool
	KeyValueMode  table.ObKeyValueMode
}

func (f ObQueryOptionFunc) Apply(opts *ObQueryOptions) {
	f(opts)
}

// WithQueryFilter set query filter
func WithQueryFilter(QueryFilter filter.ObTableFilter) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.QueryFilter = QueryFilter
	})
}

// WithQueryHTableFilter set htable filter
func WithQueryHTableFilter(HTableFilter hfilter.ObHTableFilter) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.HTableFilter = HTableFilter
	})
}

// WithQuerySelectColumns set select columns
func WithQuerySelectColumns(SelectColumns []string) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.SelectColumns = SelectColumns
	})
}

// WithQueryIndexName set index name
func WithQueryIndexName(IndexName string) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.IndexName = IndexName
	})
}

// WithQueryBatchSize set batch size
func WithQueryBatchSize(BatchSize int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.BatchSize = int32(BatchSize)
	})
}

// WithQueryMaxResultSize set max result size
func WithQueryMaxResultSize(MaxResultSize int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.MaxResultSize = int64(MaxResultSize)
	})
}

// WithQueryLimit set Limit
func WithQueryLimit(Limit int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.Limit = int32(Limit)
	})
}

// WithQueryOffset set Offset
func WithQueryOffset(Offset int) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.Offset = int32(Offset)
	})
}

// WithQueryScanOrder set scan order
func WithQueryScanOrder(ScanOrder table.ScanOrder) ObQueryOption {
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

// WithQueryKeyValueMode set key value mode
func WithQueryKeyValueMode(KeyValueMode table.ObKeyValueMode) ObQueryOption {
	return ObQueryOptionFunc(func(opts *ObQueryOptions) {
		opts.KeyValueMode = KeyValueMode
	})
}

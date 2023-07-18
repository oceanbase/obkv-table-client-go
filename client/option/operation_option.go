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

package option

import (
	"github.com/oceanbase/obkv-table-client-go/client/filter"
)

type ObkvOperationOption interface {
	Apply(opts *ObkvOperationOptions)
}

type ObkvOperationOptionFunc func(opts *ObkvOperationOptions)

func NewOperationOptions() *ObkvOperationOptions {
	return &ObkvOperationOptions{
		ReturnRowKey:         false,
		ReturnAffectedEntity: false,
		TableFilter:          nil,
	}
}

type ObkvOperationOptions struct {
	ReturnRowKey         bool
	ReturnAffectedEntity bool
	TableFilter          filter.ObTableFilter
}

func (f ObkvOperationOptionFunc) Apply(opts *ObkvOperationOptions) {
	f(opts)
}

// WithReturnRowKey only work in increment and append operation
func WithReturnRowKey(ReturnRowKey bool) ObkvOperationOption {
	return ObkvOperationOptionFunc(func(opts *ObkvOperationOptions) {
		opts.ReturnRowKey = ReturnRowKey
	})
}

// WithReturnAffectedEntity only work in increment and append operation
func WithReturnAffectedEntity(ReturnAffectedEntity bool) ObkvOperationOption {
	return ObkvOperationOptionFunc(func(opts *ObkvOperationOptions) {
		opts.ReturnAffectedEntity = ReturnAffectedEntity
	})
}

// WithFilter only work in increment append update and delete operation
func WithFilter(TableFilter filter.ObTableFilter) ObkvOperationOption {
	return ObkvOperationOptionFunc(func(opts *ObkvOperationOptions) {
		opts.TableFilter = TableFilter
	})
}

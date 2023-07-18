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

type ObOperationOption interface {
	Apply(opts *ObOperationOptions)
}

type ObOperationOptionFunc func(opts *ObOperationOptions)

func NewOperationOptions() *ObOperationOptions {
	return &ObOperationOptions{
		ReturnRowKey:         false,
		ReturnAffectedEntity: false,
		TableFilter:          nil,
	}
}

type ObOperationOptions struct {
	ReturnRowKey         bool
	ReturnAffectedEntity bool
	TableFilter          filter.ObTableFilter
}

func (f ObOperationOptionFunc) Apply(opts *ObOperationOptions) {
	f(opts)
}

// WithReturnRowKey only work in increment and append operation
func WithReturnRowKey(ReturnRowKey bool) ObOperationOption {
	return ObOperationOptionFunc(func(opts *ObOperationOptions) {
		opts.ReturnRowKey = ReturnRowKey
	})
}

// WithReturnAffectedEntity only work in increment and append operation
func WithReturnAffectedEntity(ReturnAffectedEntity bool) ObOperationOption {
	return ObOperationOptionFunc(func(opts *ObOperationOptions) {
		opts.ReturnAffectedEntity = ReturnAffectedEntity
	})
}

// WithFilter only work in increment append update and delete operation
func WithFilter(TableFilter filter.ObTableFilter) ObOperationOption {
	return ObOperationOptionFunc(func(opts *ObOperationOptions) {
		opts.TableFilter = TableFilter
	})
}

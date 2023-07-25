/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
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
	"github.com/oceanbase/obkv-table-client-go/table"
)

type ObBatchOption interface {
	Apply(opts *ObBatchOptions)
}

type ObBatchOptionFunc func(opts *ObBatchOptions)

func NewObBatchOption() *ObBatchOptions {
	return &ObBatchOptions{
		SamePropertiesNames: false,
		KeyValueMode:        table.DynamicMode,
	}
}

type ObBatchOptions struct {
	SamePropertiesNames bool
	KeyValueMode        table.ObKeyValueMode
}

func (f ObBatchOptionFunc) Apply(opts *ObBatchOptions) {
	f(opts)
}

// WithBatchSamePropertiesNames set samePropertiesNames
func WithBatchSamePropertiesNames(samePropertiesNames bool) ObBatchOption {
	return ObBatchOptionFunc(func(opts *ObBatchOptions) {
		opts.SamePropertiesNames = samePropertiesNames
	})
}

// WithBatchKeyValueMode set keyValueMode
func WithBatchKeyValueMode(keyValueMode table.ObKeyValueMode) ObBatchOption {
	return ObBatchOptionFunc(func(opts *ObBatchOptions) {
		opts.KeyValueMode = keyValueMode
	})
}

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

type ObBatchOption interface {
	Apply(opts *ObBatchOptions)
}

type ObBatchOptionFunc func(opts *ObBatchOptions)

func NewObBatchOption() *ObBatchOptions {
	return &ObBatchOptions{
		ReadOnly:            true,
		SameType:            true,
		SamePropertiesNames: false,
	}
}

type ObBatchOptions struct {
	ReadOnly            bool
	SameType            bool
	SamePropertiesNames bool
}

func (f ObBatchOptionFunc) Apply(opts *ObBatchOptions) {
	f(opts)
}

// WithReadOnly set readonly
func WithReadOnly(readonly bool) ObBatchOption {
	return ObBatchOptionFunc(func(opts *ObBatchOptions) {
		opts.ReadOnly = readonly
	})
}

// WithSameType set sameType
func WithSameType(sameType bool) ObBatchOption {
	return ObBatchOptionFunc(func(opts *ObBatchOptions) {
		opts.SameType = sameType
	})
}

// WithSamePropertiesNames set samePropertiesNames
func WithSamePropertiesNames(samePropertiesNames bool) ObBatchOption {
	return ObBatchOptionFunc(func(opts *ObBatchOptions) {
		opts.SamePropertiesNames = samePropertiesNames
	})
}

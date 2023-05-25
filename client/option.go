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

type ObkvOption interface {
	apply(opts *ObkvOptions)
}
type ObkvOptionFunc func(opts *ObkvOptions)

func NewObkvOption() *ObkvOptions {
	return &ObkvOptions{
		returnRowKey:         false,
		returnAffectedEntity: false,
	}
}

type ObkvOptions struct {
	returnRowKey         bool
	returnAffectedEntity bool
}

func (f ObkvOptionFunc) apply(opts *ObkvOptions) {
	f(opts)
}

// WithReturnRowKey only work in increment and append operation
func WithReturnRowKey(returnRowKey bool) ObkvOption {
	return ObkvOptionFunc(func(opts *ObkvOptions) {
		opts.returnRowKey = returnRowKey
	})
}

// WithReturnAffectedEntity only work in increment and append operation
func WithReturnAffectedEntity(returnAffectedEntity bool) ObkvOption {
	return ObkvOptionFunc(func(opts *ObkvOptions) {
		opts.returnAffectedEntity = returnAffectedEntity
	})
}

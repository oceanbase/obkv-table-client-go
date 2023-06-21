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

type OperationOption interface {
	apply(operationOpts *OperationOptions)
}
type OperationOptionFunc func(operationOpts *OperationOptions)

func NewOperationOptions() *OperationOptions {
	return &OperationOptions{
		returnRowKey:         false,
		returnAffectedEntity: false,
	}
}

type OperationOptions struct {
	returnRowKey         bool
	returnAffectedEntity bool
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

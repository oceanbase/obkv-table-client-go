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

type BatchOperationResult interface {
	// IsEmptySet result is empty or not.
	IsEmptySet() bool
	// GetResults get all results.
	GetResults() []SingleResult
	// Size batch operation size.
	Size() int
	// SuccessIdx indexes of successful operation.
	SuccessIdx() []int
	// ErrorIdx indexes of unsuccessful operation.
	ErrorIdx() []int
}

type obBatchOperationResult struct {
	results []SingleResult
}

func newObBatchOperationResult(results []SingleResult) *obBatchOperationResult {
	return &obBatchOperationResult{results}
}

func (r *obBatchOperationResult) IsEmptySet() bool {
	if len(r.results) == 0 {
		return true
	}
	for _, result := range r.results {
		if !result.IsEmptySet() {
			return false
		}
	}
	return true
}

func (r *obBatchOperationResult) GetResults() []SingleResult {
	return r.results
}

func (r *obBatchOperationResult) Size() int {
	return len(r.results)
}

func (r *obBatchOperationResult) SuccessIdx() []int {
	var successIdx []int
	for i, result := range r.results {
		if result != nil {
			successIdx = append(successIdx, i)
		}
	}
	return successIdx
}

func (r *obBatchOperationResult) ErrorIdx() []int {
	var errorIdx []int
	for i, result := range r.results {
		if result == nil {
			errorIdx = append(errorIdx, i)
		}
	}
	return errorIdx
}

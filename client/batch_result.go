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
	oberror "github.com/oceanbase/obkv-table-client-go/error"
	"github.com/oceanbase/obkv-table-client-go/protocol"
)

type BatchOperationResult interface {
	GetResults() []*protocol.ObTableOperationResponse
	Size() int
	WrongCount() int
	CorrectCount() int
	WrongIndexes() []int
	CorrectIndexes() []int
}

type obBatchOperationResult struct {
	results []*protocol.ObTableOperationResponse
}

func newObBatchOperationResult(results []*protocol.ObTableOperationResponse) *obBatchOperationResult {
	return &obBatchOperationResult{results}
}

func (r *obBatchOperationResult) GetResults() []*protocol.ObTableOperationResponse {
	return r.results
}

func (r *obBatchOperationResult) Size() int {
	return len(r.results)
}

func (r *obBatchOperationResult) WrongCount() int {
	var count = 0
	for _, result := range r.results {
		if oberror.ObErrorCode(result.Header().ErrorNo()) != oberror.ObSuccess {
			count++
		}
	}
	return count
}

func (r *obBatchOperationResult) CorrectCount() int {
	var count = 0
	for _, result := range r.results {
		if oberror.ObErrorCode(result.Header().ErrorNo()) == oberror.ObSuccess {
			count++
		}
	}
	return count
}

func (r *obBatchOperationResult) WrongIndexes() []int {
	var indexes []int
	for i, result := range r.results {
		if oberror.ObErrorCode(result.Header().ErrorNo()) != oberror.ObSuccess {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func (r *obBatchOperationResult) CorrectIndexes() []int {
	var indexes = make([]int, 0, len(r.results))
	for i, result := range r.results {
		if oberror.ObErrorCode(result.Header().ErrorNo()) == oberror.ObSuccess {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObBatchOperationResult_Size(t *testing.T) {
	result := newObSingleResult(1, nil)
	emptyResult := make([]SingleResult, 0)
	batchResult := newObBatchOperationResult(emptyResult)
	assert.EqualValues(t, 0, batchResult.Size())
	batchResult = newObBatchOperationResult([]SingleResult{result})
	assert.EqualValues(t, 1, batchResult.Size())
}

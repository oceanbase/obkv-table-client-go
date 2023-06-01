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

package protocol

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObTableBatchOperationResponseEncodeDecode(t *testing.T) {
	randomLen := rand.Intn(10)
	obTableOperations := make([]*ObTableOperationResponse, 0, randomLen)

	for i := 0; i < randomLen; i++ {
		obTableOperations = append(obTableOperations, NewObTableOperationResponse())
	}

	obTableBatchOperationResponse := NewObTableBatchOperationResponse()
	obTableBatchOperationResponse.SetVersion(1)
	obTableBatchOperationResponse.SetContentLength(0)
	obTableBatchOperationResponse.SetObTableOperationResponses(obTableOperations)

	payloadLen := obTableBatchOperationResponse.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableBatchOperationResponse.Encode(buffer)

	newObTableBatchOperationResponse := NewObTableBatchOperationResponse()
	newBuffer := bytes.NewBuffer(buf)
	newObTableBatchOperationResponse.Decode(newBuffer)

	assert.EqualValues(t, obTableBatchOperationResponse.Version(), newObTableBatchOperationResponse.Version())
	assert.EqualValues(t, obTableBatchOperationResponse.ContentLength(), newObTableBatchOperationResponse.ContentLength())
	assert.EqualValues(t, len(obTableBatchOperationResponse.ObTableOperationResponses()), len(newObTableBatchOperationResponse.ObTableOperationResponses()))
}

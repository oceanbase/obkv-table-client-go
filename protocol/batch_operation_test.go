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

	"github.com/oceanbase/obkv-table-client-go/table"
)

func TestObTableBatchOperationEncodeDecode(t *testing.T) {
	obTableBatchOperation := NewObTableBatchOperation()

	obTableBatchOperation.SetVersion(1)
	obTableBatchOperation.SetContentLength(0)
	obTableBatchOperation.SetReadOnly(true)
	obTableBatchOperation.SetSamePropertiesNames(false)

	randomLen := rand.Intn(10)
	obTableOperations := make([]*ObTableOperation, 0, randomLen)
	obTableBatchOperation.SetObTableOperations(obTableOperations)

	for i := 0; i < randomLen; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
		mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
		tableOperation, _ := NewObTableOperationWithParams(ObTableOperationType(rand.Intn(8)), rowKey, mutateColumns)
		obTableBatchOperation.AppendObTableOperation(tableOperation)
	}

	payloadLen := obTableBatchOperation.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableBatchOperation.Encode(buffer)

	newObTableBatchOperation := NewObTableBatchOperation()
	newBuffer := bytes.NewBuffer(buf)
	newObTableBatchOperation.Decode(newBuffer)

	assert.EqualValues(t, obTableBatchOperation.Version(), newObTableBatchOperation.Version())
	assert.EqualValues(t, obTableBatchOperation.ContentLength(), newObTableBatchOperation.ContentLength())
	assert.EqualValues(t, obTableBatchOperation.ReadOnly(), newObTableBatchOperation.ReadOnly())
	assert.EqualValues(t, obTableBatchOperation.SamePropertiesNames(), newObTableBatchOperation.SamePropertiesNames())
	assert.EqualValues(t, obTableBatchOperation.ObTableOperations(), newObTableBatchOperation.ObTableOperations())
	assert.EqualValues(t, obTableBatchOperation, newObTableBatchOperation)
}

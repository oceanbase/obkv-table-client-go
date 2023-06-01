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
	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObTableBatchOperationRequestEncodeDecode(t *testing.T) {
	obTableBatchOperation := NewObTableBatchOperation()

	randomLen := rand.Intn(10)
	obTableOperations := make([]*ObTableOperation, 0, randomLen)
	obTableBatchOperation.SetObTableOperations(obTableOperations)

	for i := 0; i < randomLen; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
		columns := []*table.Column{table.NewColumn("c2", int64(1))}
		tableOperation, _ := NewObTableOperationWithParams(ObTableOperationType(rand.Intn(8)), rowKey, columns)
		obTableBatchOperation.AppendObTableOperation(tableOperation)
	}

	batchOperationRequest := NewObTableBatchOperationRequest()
	batchOperationRequest.SetVersion(1)
	batchOperationRequest.SetContentLength(0)
	batchOperationRequest.SetCredential([]byte(util.String(rand.Intn(20))))
	batchOperationRequest.SetTableName(util.String(rand.Intn(10)))
	batchOperationRequest.SetTableId(rand.Uint64())
	batchOperationRequest.SetObTableEntityType(ObTableEntityType(rand.Intn(3)))
	batchOperationRequest.SetObTableBatchOperation(obTableBatchOperation)
	batchOperationRequest.SetObTableConsistencyLevel(ObTableConsistencyLevel(rand.Intn(2)))
	batchOperationRequest.SetReturnRowKey(util.ByteToBool(byte(rand.Intn(2))))
	batchOperationRequest.SetReturnAffectedEntity(util.ByteToBool(byte(rand.Intn(2))))
	batchOperationRequest.SetReturnAffectedRows(util.ByteToBool(byte(rand.Intn(2))))
	batchOperationRequest.SetPartitionId(rand.Uint64())
	batchOperationRequest.SetAtomicOperation(util.ByteToBool(byte(rand.Intn(2))))

	payloadLen := batchOperationRequest.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	batchOperationRequest.Encode(buffer)

	newObTableBatchOperation := NewObTableBatchOperationRequest()
	newBuffer := bytes.NewBuffer(buf)
	newObTableBatchOperation.Decode(newBuffer)

	assert.EqualValues(t, batchOperationRequest.Credential(), newObTableBatchOperation.Credential())
	assert.EqualValues(t, batchOperationRequest.TableName(), newObTableBatchOperation.TableName())
	assert.EqualValues(t, batchOperationRequest.TableId(), newObTableBatchOperation.TableId())
	assert.EqualValues(t, batchOperationRequest.ObTableEntityType(), newObTableBatchOperation.ObTableEntityType())
	assert.EqualValues(t, batchOperationRequest.ObTableBatchOperation(), newObTableBatchOperation.ObTableBatchOperation())
	assert.EqualValues(t, batchOperationRequest.ObTableConsistencyLevel(), newObTableBatchOperation.ObTableConsistencyLevel())
	assert.EqualValues(t, batchOperationRequest.ReturnRowKey(), newObTableBatchOperation.ReturnRowKey())
	assert.EqualValues(t, batchOperationRequest.ReturnAffectedEntity(), newObTableBatchOperation.ReturnAffectedEntity())
	assert.EqualValues(t, batchOperationRequest.ReturnAffectedRows(), newObTableBatchOperation.ReturnAffectedRows())
	assert.EqualValues(t, batchOperationRequest.PartitionId(), newObTableBatchOperation.PartitionId())
	assert.EqualValues(t, batchOperationRequest.AtomicOperation(), newObTableBatchOperation.AtomicOperation())
	assert.EqualValues(t, len(batchOperationRequest.String()), len(newObTableBatchOperation.String()))
	assert.EqualValues(t, batchOperationRequest, newObTableBatchOperation)
}

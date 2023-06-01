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

func TestObTableOperationRequestEncodeDecode(t *testing.T) {
	obTableOperationRequest := NewObTableOperationRequest()

	obTableOperation := NewObTableOperation()

	obTableEntity := NewObTableEntityWithParams(0, 0)
	obTableEntity.SetVersion(1)
	obTableEntity.SetContentLength(0)
	randomLen := rand.Intn(100)
	rowKey := make([]*table.Column, 0, randomLen)
	columns := make([]*table.Column, 0, randomLen)

	for i := 0; i < randomLen; i++ {
		rowKey = append(rowKey, table.NewColumn(util.String(10), int64(rand.Intn(10000))))
		columns = append(columns, table.NewColumn(util.String(10), int64(rand.Intn(10000))))
	}

	for _, column := range rowKey {
		objMeta, _ := DefaultObjMeta(column.Value())

		object := NewObObjectWithParams(objMeta, column.Value())

		obTableEntity.AppendRowKeyElement(object)
	}

	for _, column := range columns {
		objMeta, _ := DefaultObjMeta(column.Value())

		object := NewObObjectWithParams(objMeta, column.Value())

		obTableEntity.SetProperty(column.Name(), object)
	}

	obTableOperation.SetOpType(ObTableOperationType(rand.Intn(8)))
	obTableOperation.SetEntity(obTableEntity)

	obTableOperationRequest.SetCredential([]byte(util.String(10)))
	obTableOperationRequest.SetTableName(util.String(10))
	obTableOperationRequest.SetTableId(rand.Uint64())
	obTableOperationRequest.SetPartitionId(rand.Uint64())
	obTableOperationRequest.SetEntityType(ObTableEntityType(rand.Intn(3)))
	obTableOperationRequest.SetTableOperation(obTableOperation)
	obTableOperationRequest.SetConsistencyLevel(ObTableConsistencyLevel(rand.Intn(2)))
	obTableOperationRequest.SetReturnRowKey(util.ByteToBool(byte(rand.Intn(2))))
	obTableOperationRequest.SetReturnAffectedEntity(util.ByteToBool(byte(rand.Intn(2))))
	obTableOperationRequest.SetReturnAffectedRows(util.ByteToBool(byte(rand.Intn(2))))

	payloadLen := obTableOperationRequest.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableOperationRequest.Encode(buffer)

	newObTableOperationRequest := NewObTableOperationRequest()
	newObTableOperation := NewObTableOperation()
	newObTableOperationRequest.SetTableOperation(newObTableOperation)

	newBuffer := bytes.NewBuffer(buf)
	newObTableOperationRequest.Decode(newBuffer)

	assert.EqualValues(t, obTableOperationRequest.Credential(), newObTableOperationRequest.Credential())
	assert.EqualValues(t, obTableOperationRequest.TableName(), newObTableOperationRequest.TableName())
	assert.EqualValues(t, obTableOperationRequest.TableId(), newObTableOperationRequest.TableId())
	assert.EqualValues(t, obTableOperationRequest.PartitionId(), newObTableOperationRequest.PartitionId())
	assert.EqualValues(t, obTableOperationRequest.EntityType(), newObTableOperationRequest.EntityType())
	assert.EqualValues(t, obTableOperationRequest.TableOperation(), newObTableOperationRequest.TableOperation())
	assert.EqualValues(t, obTableOperationRequest.ConsistencyLevel(), newObTableOperationRequest.ConsistencyLevel())
	assert.EqualValues(t, obTableOperationRequest.ReturnRowKey(), newObTableOperationRequest.ReturnRowKey())
	assert.EqualValues(t, obTableOperationRequest.ReturnAffectedEntity(), newObTableOperationRequest.ReturnAffectedEntity())
	assert.EqualValues(t, obTableOperationRequest.ReturnAffectedRows(), newObTableOperationRequest.ReturnAffectedRows())
	assert.EqualValues(t, len(obTableOperationRequest.String()), len(newObTableOperationRequest.String()))
	assert.EqualValues(t, obTableOperationRequest, newObTableOperationRequest)
}

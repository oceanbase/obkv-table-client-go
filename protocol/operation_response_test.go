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

func TestObTableOperationResponseEncodeDecode(t *testing.T) {
	obTableOperationResponse := NewObTableOperationResponse()
	obTableOperationResponse.SetVersion(1)
	obTableOperationResponse.SetContentLength(0)
	obTableResponse := NewObTableResponse()
	obTableResponse.SetErrorNo(int32(rand.Uint32()))
	obTableResponse.SetSqlState([]byte(util.String(20)))
	obTableResponse.SetMsg([]byte(util.String(20)))

	obTableOperationResponse.SetHeader(obTableResponse)
	obTableOperationResponse.SetOperationType(ObTableOperationType(rand.Intn(8)))

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

	obTableOperationResponse.SetEntity(obTableEntity)
	obTableOperationResponse.SetAffectedRows(int64(rand.Uint64()))

	payloadLen := obTableOperationResponse.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableOperationResponse.Encode(buffer)

	newObTableOperationResponse := NewObTableOperationResponse()

	newBuffer := bytes.NewBuffer(buf)
	newObTableOperationResponse.Decode(newBuffer)

	assert.EqualValues(t, obTableOperationResponse.Header(), newObTableOperationResponse.Header())
	assert.EqualValues(t, obTableOperationResponse.OperationType(), newObTableOperationResponse.OperationType())
	assert.EqualValues(t, obTableOperationResponse.Entity(), newObTableOperationResponse.Entity())
	assert.EqualValues(t, obTableOperationResponse.AffectedRows(), newObTableOperationResponse.AffectedRows())
	assert.EqualValues(t, obTableOperationResponse, newObTableOperationResponse)
}

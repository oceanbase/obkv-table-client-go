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

func TestObTableOperationEncodeDecode(t *testing.T) {
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

	payloadLen := obTableOperation.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableOperation.Encode(buffer)

	newObTableOperation := NewObTableOperation()
	newBuffer := bytes.NewBuffer(buf)
	newObTableOperation.Decode(newBuffer)

	assert.EqualValues(t, obTableOperation.OpType(), newObTableOperation.OpType())
	assert.EqualValues(t, obTableOperation.Entity(), newObTableOperation.Entity())
	assert.EqualValues(t, len(obTableOperation.String()), len(newObTableOperation.String()))
	assert.EqualValues(t, obTableOperation, newObTableOperation)
}

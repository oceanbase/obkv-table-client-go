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

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

func TestObBatchExecutor_String(t *testing.T) {
	executor := &obBatchExecutor{}
	assert.Equal(t, "obBatchExecutor{tableName:, rowKeyName:[], isAtomic:false}", executor.String())

	executor = newObBatchExecutor("test", nil)
	assert.Equal(t, "obBatchExecutor{tableName:test, rowKeyName:[], isAtomic:true}", executor.String())

	executor.SetIsAtomic(false)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddInsertOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)
	assert.Equal(t, "obBatchExecutor{tableName:test, rowKeyName:[c1], isAtomic:false}", executor.String())
}

func TestObBatchExecutor_AddInsertOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddInsertOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)

	err = executor.AddInsertOp(nil, mutateColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddInsertOp(rowKey, nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddUpdateOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddUpdateOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)

	err = executor.AddUpdateOp(nil, mutateColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddUpdateOp(rowKey, nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddInsertOrUpdateOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddInsertOrUpdateOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)

	err = executor.AddInsertOrUpdateOp(nil, mutateColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddInsertOrUpdateOp(rowKey, nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddReplaceOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddReplaceOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)

	err = executor.AddReplaceOp(nil, mutateColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddReplaceOp(rowKey, nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddIncrementOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddIncrementOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)

	err = executor.AddIncrementOp(nil, mutateColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddIncrementOp(rowKey, nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddAppendOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	err := executor.AddAppendOp(rowKey, mutateColumns)
	assert.Equal(t, nil, err)

	err = executor.AddAppendOp(nil, mutateColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddAppendOp(rowKey, nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddDeleteOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	err := executor.AddDeleteOp(rowKey)
	assert.Equal(t, nil, err)

	err = executor.AddDeleteOp(nil)
	assert.NotEqual(t, nil, err)
}

func TestObBatchExecutor_AddGetOp(t *testing.T) {
	executor := newObBatchExecutor("test", nil)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	getColumns := []string{"c1", "c2"}
	err := executor.AddGetOp(rowKey, getColumns)
	assert.Equal(t, nil, err)

	err = executor.AddGetOp(nil, getColumns)
	assert.NotEqual(t, nil, err)
	err = executor.AddGetOp(rowKey, nil)
	assert.Equal(t, nil, err)
}

func TestObPartOp_String(t *testing.T) {
	op := &obPartOp{}
	assert.Equal(t, "obPartOp{tableParam:nil, ops:[]}", op.String())

	op = newPartOp(nil)
	assert.Equal(t, "obPartOp{tableParam:nil, ops:[]}", op.String())

	tableParam := NewObTableParam(&ObTable{}, testTableIdV3, testPartIdV3)
	op = newPartOp(tableParam)
	assert.EqualValues(t, "obPartOp{tableParam:ObTableParam{table:ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}, tableId:1099511677791, partitionId:0}, ops:[]}", op.String())

	singleOp := newSingleOp(0, nil)
	op.addOperation(singleOp)
	assert.Equal(t, "obPartOp{tableParam:ObTableParam{table:ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}, tableId:1099511677791, partitionId:0}, ops:[obSingleOp{indexOfBatch:0, op:nil}]}", op.String())
}

func TestObSingleOp_String(t *testing.T) {
	op := &obSingleOp{}
	assert.Equal(t, "obSingleOp{indexOfBatch:0, op:nil}", op.String())

	op = newSingleOp(0, nil)
	assert.Equal(t, "obSingleOp{indexOfBatch:0, op:nil}", op.String())

	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	req, err := protocol.NewObTableOperationWithParams(protocol.ObTableOperationGet, rowKey, mutateColumns)
	assert.Equal(t, nil, err)
	op = newSingleOp(0, req)
	assert.Equal(t, "obSingleOp{indexOfBatch:0, op:ObTableOperation{ObUniVersionHeader:ObUniVersionHeader{version:1, contentLength:0}, opType:0, entity:ObTableEntity{ObUniVersionHeader:ObUniVersionHeader{version:1, contentLength:0}, rowKey:[ObObject{meta:ObObjectMeta{objType:ObObjType{type:ObInt64Type}, collationLevel:5, collationType:63, scale:-1}, value:1}], properties:{m[c2]=ObObject{meta:ObObjectMeta{objType:ObObjType{type:ObInt64Type}, collationLevel:5, collationType:63, scale:-1}, value:1}}}}}", op.String())
}

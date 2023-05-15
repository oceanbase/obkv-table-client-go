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
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableOperationRequest struct {
	*UniVersionHeader
	credential           []byte
	tableName            string
	tableId              uint64
	partitionId          int64
	entityType           TableEntityType
	tableOperation       *TableOperation
	consistencyLevel     TableConsistencyLevel
	returnRowKey         bool
	returnAffectedEntity bool
	returnAffectedRows   bool
}

func NewTableOperationRequest(
	tableName string,
	tableId uint64,
	partitionId int64,
	operationType TableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	timeout time.Duration,
	flag uint16) (*TableOperationRequest, error) {
	tableOperation, err := NewTableOperation(operationType, rowKey, columns)
	if err != nil {
		return nil, errors.WithMessage(err, "create table operation")
	}
	uniVersionHeader := NewUniVersionHeader()
	uniVersionHeader.SetFlag(flag)
	uniVersionHeader.SetTimeout(timeout)

	return &TableOperationRequest{
		UniVersionHeader:     uniVersionHeader,
		credential:           nil, // when execute set
		tableName:            tableName,
		tableId:              tableId,
		partitionId:          partitionId,
		entityType:           Dynamic,
		tableOperation:       tableOperation,
		consistencyLevel:     Strong,
		returnRowKey:         false,
		returnAffectedEntity: false,
		returnAffectedRows:   true,
	}, nil
}

func (r *TableOperationRequest) TableName() string {
	return r.tableName
}

func (r *TableOperationRequest) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *TableOperationRequest) TableId() uint64 {
	return r.tableId
}

func (r *TableOperationRequest) SetTableId(tableId uint64) {
	r.tableId = tableId
}

func (r *TableOperationRequest) PartitionId() int64 {
	return r.partitionId
}

func (r *TableOperationRequest) SetPartitionId(partitionId int64) {
	r.partitionId = partitionId
}

func (r *TableOperationRequest) EntityType() TableEntityType {
	return r.entityType
}

func (r *TableOperationRequest) SetEntityType(entityType TableEntityType) {
	r.entityType = entityType
}

func (r *TableOperationRequest) TableOperation() *TableOperation {
	return r.tableOperation
}

func (r *TableOperationRequest) SetTableOperation(tableOperation *TableOperation) {
	r.tableOperation = tableOperation
}

func (r *TableOperationRequest) ConsistencyLevel() TableConsistencyLevel {
	return r.consistencyLevel
}

func (r *TableOperationRequest) SetConsistencyLevel(consistencyLevel TableConsistencyLevel) {
	r.consistencyLevel = consistencyLevel
}

func (r *TableOperationRequest) ReturnRowKey() bool {
	return r.returnRowKey
}

func (r *TableOperationRequest) SetReturnRowKey(returnRowKey bool) {
	r.returnRowKey = returnRowKey
}

func (r *TableOperationRequest) ReturnAffectedEntity() bool {
	return r.returnAffectedEntity
}

func (r *TableOperationRequest) SetReturnAffectedEntity(returnAffectedEntity bool) {
	r.returnAffectedEntity = returnAffectedEntity
}

func (r *TableOperationRequest) ReturnAffectedRows() bool {
	return r.returnAffectedRows
}

func (r *TableOperationRequest) SetReturnAffectedRows(returnAffectedRows bool) {
	r.returnAffectedRows = returnAffectedRows
}

func (r *TableOperationRequest) PCode() TablePacketCode {
	return TableApiExecute
}

func (r *TableOperationRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.UniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *TableOperationRequest) PayloadContentLen() int {
	totalLen := 0
	if util.ObVersion() >= 4 {
		totalLen =
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				8 + // todo partitionId
				5 + // entityType consistencyLevel returnRowKey returnAffectedEntity returnAffectedRows
				r.tableOperation.PayloadLen()
	} else {
		totalLen =
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				util.EncodedLengthByVi64(r.partitionId) + // todo partitionId
				5 + // entityType consistencyLevel returnRowKey returnAffectedEntity returnAffectedRows
				r.tableOperation.PayloadLen()
	}

	r.UniVersionHeader.SetContentLength(totalLen)
	return r.UniVersionHeader.ContentLength()
}

func (r *TableOperationRequest) Credential() []byte {
	return r.credential
}

func (r *TableOperationRequest) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *TableOperationRequest) Encode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Encode(buffer)

	util.EncodeBytesString(buffer, r.credential)

	util.EncodeVString(buffer, r.tableName)

	util.EncodeVi64(buffer, int64(r.tableId))

	if util.ObVersion() >= 4 {
		util.PutUint64(buffer, uint64(r.partitionId))
	} else {
		util.EncodeVi64(buffer, r.partitionId)
	}

	util.PutUint8(buffer, uint8(r.entityType))

	r.tableOperation.Encode(buffer)

	util.PutUint8(buffer, uint8(r.consistencyLevel))

	util.PutUint8(buffer, util.BoolToByte(r.returnRowKey))

	util.PutUint8(buffer, util.BoolToByte(r.returnAffectedEntity))

	util.PutUint8(buffer, util.BoolToByte(r.returnAffectedRows))
}

func (r *TableOperationRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *TableOperationRequest) String() string {
	// todo: impl
	return "TableOperationRequest{" +
		"}"
}

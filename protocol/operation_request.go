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

type ObTableOperationRequest struct {
	*ObUniVersionHeader
	*ObPayloadBase
	credential           []byte
	tableName            string
	tableId              uint64
	partitionId          uint64
	entityType           ObTableEntityType
	tableOperation       *ObTableOperation
	consistencyLevel     ObTableConsistencyLevel
	returnRowKey         bool
	returnAffectedEntity bool
	returnAffectedRows   bool
}

func NewObTableOperationRequest(
	tableName string,
	tableId uint64,
	partitionId uint64,
	tableOperationType ObTableOperationType,
	rowKey []*table.Column,
	columns []*table.Column,
	timeout time.Duration,
	flag uint16) (*ObTableOperationRequest, error) {
	tableOperation, err := NewObTableOperation(tableOperationType, rowKey, columns)
	if err != nil {
		return nil, errors.WithMessage(err, "create table operation")
	}
	uniVersionHeader := NewObUniVersionHeader()
	obPayloadBase := NewObPayloadBase()
	obPayloadBase.SetFlag(flag)
	obPayloadBase.SetTimeout(timeout)

	return &ObTableOperationRequest{
		ObUniVersionHeader:   uniVersionHeader,
		ObPayloadBase:        obPayloadBase,
		credential:           nil, // when execute set
		tableName:            tableName,
		tableId:              tableId,
		partitionId:          partitionId,
		entityType:           ObTableEntityTypeDynamic,
		tableOperation:       tableOperation,
		consistencyLevel:     ObTableConsistencyLevelStrong,
		returnRowKey:         false,
		returnAffectedEntity: false,
		returnAffectedRows:   true,
	}, nil
}

func (r *ObTableOperationRequest) TableName() string {
	return r.tableName
}

func (r *ObTableOperationRequest) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *ObTableOperationRequest) TableId() uint64 {
	return r.tableId
}

func (r *ObTableOperationRequest) SetTableId(tableId uint64) {
	r.tableId = tableId
}

func (r *ObTableOperationRequest) PartitionId() uint64 {
	return r.partitionId
}

func (r *ObTableOperationRequest) SetPartitionId(partitionId uint64) {
	r.partitionId = partitionId
}

func (r *ObTableOperationRequest) EntityType() ObTableEntityType {
	return r.entityType
}

func (r *ObTableOperationRequest) SetEntityType(entityType ObTableEntityType) {
	r.entityType = entityType
}

func (r *ObTableOperationRequest) TableOperation() *ObTableOperation {
	return r.tableOperation
}

func (r *ObTableOperationRequest) SetTableOperation(tableOperation *ObTableOperation) {
	r.tableOperation = tableOperation
}

func (r *ObTableOperationRequest) ConsistencyLevel() ObTableConsistencyLevel {
	return r.consistencyLevel
}

func (r *ObTableOperationRequest) SetConsistencyLevel(consistencyLevel ObTableConsistencyLevel) {
	r.consistencyLevel = consistencyLevel
}

func (r *ObTableOperationRequest) ReturnRowKey() bool {
	return r.returnRowKey
}

func (r *ObTableOperationRequest) SetReturnRowKey(returnRowKey bool) {
	r.returnRowKey = returnRowKey
}

func (r *ObTableOperationRequest) ReturnAffectedEntity() bool {
	return r.returnAffectedEntity
}

func (r *ObTableOperationRequest) SetReturnAffectedEntity(returnAffectedEntity bool) {
	r.returnAffectedEntity = returnAffectedEntity
}

func (r *ObTableOperationRequest) ReturnAffectedRows() bool {
	return r.returnAffectedRows
}

func (r *ObTableOperationRequest) SetReturnAffectedRows(returnAffectedRows bool) {
	r.returnAffectedRows = returnAffectedRows
}

func (r *ObTableOperationRequest) PCode() ObTablePacketCode {
	return ObTableApiExecute
}

func (r *ObTableOperationRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableOperationRequest) PayloadContentLen() int {
	totalLen := 0
	if util.ObVersion() >= 4 {
		totalLen =
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				8 + // partitionId
				5 + // obTableEntityType obTableConsistencyLevel returnRowKey returnAffectedEntity returnAffectedRows
				r.tableOperation.PayloadLen()
	} else {
		totalLen =
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				util.EncodedLengthByVi64(int64(r.partitionId)) + // partitionId
				5 + // obTableEntityType obTableConsistencyLevel returnRowKey returnAffectedEntity returnAffectedRows
				r.tableOperation.PayloadLen()
	}

	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableOperationRequest) Credential() []byte {
	return r.credential
}

func (r *ObTableOperationRequest) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *ObTableOperationRequest) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.EncodeBytesString(buffer, r.credential)

	util.EncodeVString(buffer, r.tableName)

	util.EncodeVi64(buffer, int64(r.tableId))

	if util.ObVersion() >= 4 {
		util.PutUint64(buffer, r.partitionId)
	} else {
		util.EncodeVi64(buffer, int64(r.partitionId))
	}

	util.PutUint8(buffer, uint8(r.entityType))

	r.tableOperation.Encode(buffer)

	util.PutUint8(buffer, uint8(r.consistencyLevel))

	util.PutUint8(buffer, util.BoolToByte(r.returnRowKey))

	util.PutUint8(buffer, util.BoolToByte(r.returnAffectedEntity))

	util.PutUint8(buffer, util.BoolToByte(r.returnAffectedRows))
}

func (r *ObTableOperationRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *ObTableOperationRequest) String() string {
	return ""
}

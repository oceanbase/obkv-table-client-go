package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableBatchOperationRequest struct {
	*UniVersionHeader
	credential           []byte
	tableName            string
	tableId              uint64
	entityType           TableEntityType
	tableBatchOperation  *TableBatchOperation
	consistencyLevel     TableConsistencyLevel
	returnRowKey         bool
	returnAffectedEntity bool
	returnAffectedRows   bool
	partitionId          int64 // todo batch request partitionId different
	atomicOperation      bool
}

func (r *TableBatchOperationRequest) Credential() []byte {
	return r.credential
}

func (r *TableBatchOperationRequest) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *TableBatchOperationRequest) TableName() string {
	return r.tableName
}

func (r *TableBatchOperationRequest) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *TableBatchOperationRequest) TableId() uint64 {
	return r.tableId
}

func (r *TableBatchOperationRequest) SetTableId(tableId uint64) {
	r.tableId = tableId
}

func (r *TableBatchOperationRequest) EntityType() TableEntityType {
	return r.entityType
}

func (r *TableBatchOperationRequest) SetEntityType(entityType TableEntityType) {
	r.entityType = entityType
}

func (r *TableBatchOperationRequest) TableBatchOperation() *TableBatchOperation {
	return r.tableBatchOperation
}

func (r *TableBatchOperationRequest) SetTableBatchOperation(tableBatchOperation *TableBatchOperation) {
	r.tableBatchOperation = tableBatchOperation
}

func (r *TableBatchOperationRequest) ConsistencyLevel() TableConsistencyLevel {
	return r.consistencyLevel
}

func (r *TableBatchOperationRequest) SetConsistencyLevel(consistencyLevel TableConsistencyLevel) {
	r.consistencyLevel = consistencyLevel
}

func (r *TableBatchOperationRequest) ReturnRowKey() bool {
	return r.returnRowKey
}

func (r *TableBatchOperationRequest) SetReturnRowKey(returnRowKey bool) {
	r.returnRowKey = returnRowKey
}

func (r *TableBatchOperationRequest) ReturnAffectedEntity() bool {
	return r.returnAffectedEntity
}

func (r *TableBatchOperationRequest) SetReturnAffectedEntity(returnAffectedEntity bool) {
	r.returnAffectedEntity = returnAffectedEntity
}

func (r *TableBatchOperationRequest) ReturnAffectedRows() bool {
	return r.returnAffectedRows
}

func (r *TableBatchOperationRequest) SetReturnAffectedRows(returnAffectedRows bool) {
	r.returnAffectedRows = returnAffectedRows
}

func (r *TableBatchOperationRequest) PartitionId() int64 {
	return r.partitionId
}

func (r *TableBatchOperationRequest) SetPartitionId(partitionId int64) {
	r.partitionId = partitionId
}

func (r *TableBatchOperationRequest) AtomicOperation() bool {
	return r.atomicOperation
}

func (r *TableBatchOperationRequest) SetAtomicOperation(atomicOperation bool) {
	r.atomicOperation = atomicOperation
}

func (r *TableBatchOperationRequest) PCode() TablePacketCode {
	return TableApiBatchExecute
}

func (r *TableBatchOperationRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.UniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *TableBatchOperationRequest) PayloadContentLen() int {
	totalLen := 0
	if globalVersion >= 4 { // todo version
		totalLen =
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				6 + // entityType consistencyLevel returnRowKey returnAffectedEntity returnAffectedRows atomicOperation
				8 + // todo partitionId
				r.tableBatchOperation.PayloadLen()
	} else {
		totalLen =
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				6 + // entityType consistencyLevel returnRowKey returnAffectedEntity returnAffectedRows atomicOperation
				util.EncodedLengthByVi64(r.partitionId) + // todo partitionId\
				r.tableBatchOperation.PayloadLen()
	}

	r.UniVersionHeader.SetContentLength(totalLen)
	return r.UniVersionHeader.ContentLength()
}

func (r *TableBatchOperationRequest) Encode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Encode(buffer)

	util.EncodeBytesString(buffer, r.credential)

	util.EncodeVString(buffer, r.tableName)

	util.EncodeVi64(buffer, int64(r.tableId))

	util.PutUint8(buffer, uint8(r.entityType))

	r.tableBatchOperation.Encode(buffer)

	util.PutUint8(buffer, uint8(r.consistencyLevel))

	util.PutUint8(buffer, util.BoolToByte(r.returnRowKey))

	util.PutUint8(buffer, util.BoolToByte(r.returnAffectedEntity))

	util.PutUint8(buffer, util.BoolToByte(r.returnAffectedRows))

	if globalVersion >= 4 { // todo version
		util.PutUint64(buffer, uint64(r.partitionId))
	} else {
		util.EncodeVi64(buffer, r.partitionId)
	}

	util.PutUint8(buffer, util.BoolToByte(r.atomicOperation))
}

func (r *TableBatchOperationRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableOperationRequest struct {
	*UniVersionHeader
	credential           []byte
	tableName            string
	tableId              int64
	partitionId          int64
	entityType           TableEntityType
	tableOperation       *TableOperation
	consistencyLevel     TableConsistencyLevel
	returnRowKey         bool
	returnAffectedEntity bool
	returnAffectedRows   bool
}

type TableEntityType uint8

const (
	Dynamic TableEntityType = iota
	KV
	HKV
)

type TableConsistencyLevel uint8

const (
	Strong TableConsistencyLevel = iota
	Eventual
)

func (r *TableOperationRequest) TableName() string {
	return r.tableName
}

func (r *TableOperationRequest) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *TableOperationRequest) TableId() int64 {
	return r.tableId
}

func (r *TableOperationRequest) SetTableId(tableId int64) {
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

func (r *TableOperationRequest) PayloadLen() int64 {
	// TODO implement me
	panic("implement me")
}

func (r *TableOperationRequest) PayloadContentLen() int64 {
	// TODO implement me
	panic("implement me")
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

	util.EncodeVi64(buffer, r.tableId)

	util.EncodeVi64(buffer, r.partitionId)

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

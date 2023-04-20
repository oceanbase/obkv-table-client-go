package protocol

import (
	"bytes"
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

type TableEntityType int8

const (
	Dynamic TableEntityType = iota
	KV
	HKV
)

type TableConsistencyLevel int8

const (
	Strong TableConsistencyLevel = iota
	Eventual
)

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

func (r *TableOperationRequest) SessionId() uint64 {
	return 0
}

func (r *TableOperationRequest) SetSessionId(sessionId uint64) {
	return
}

func (r *TableOperationRequest) Credential() []byte {
	return r.credential
}

func (r *TableOperationRequest) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *TableOperationRequest) Encode(buffer *bytes.Buffer) {
	// var index = 0
	//
	// payloadLength := r.PayloadLen()
	// requestBuf := make([]byte, payloadLength)
	//
	// headerLen := r.UniVersionHeader.UniVersionHeaderLen()
	// r.UniVersionHeader.Encode(requestBuf[:headerLen])
	// index += headerLen
	//
	// needLength := util.NeedLengthByBytesString(r.credential)
	// util.EncodeBytesString(requestBuf[index:index+needLength], r.credential)
	// index += needLength
	//
	// needLength = util.NeedLengthByVString(r.tableName)
	// util.EncodeVString(requestBuf[index:index+needLength], r.tableName)
	// index += needLength
	//
	// needLength = util.NeedLengthByVi64(r.tableId)
	// util.EncodeVi64(requestBuf[index:index+needLength], r.tableId)
	// index += needLength
	//
	// needLength = util.NeedLengthByVi64(r.partitionId)
	// util.EncodeVi64(requestBuf[index:index+needLength], r.partitionId)
	// index += needLength
	//
	// util.PutUint8(requestBuf[index:index+1], uint8(r.entityType))
	// index++

	// tableOperationPayloadLen := r.tableOperation.PayloadLen()
	// r.tableOperation.Encode()
}

func (r *TableOperationRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

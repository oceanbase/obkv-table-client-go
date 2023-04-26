package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableBatchOperation struct {
	*UniVersionHeader
	tableOperations     []*TableOperation
	readOnly            bool
	sameType            bool
	samePropertiesNames bool
}

func (o *TableBatchOperation) TableOperations() []*TableOperation {
	return o.tableOperations
}

func (o *TableBatchOperation) SetTableOperations(tableOperations []*TableOperation) {
	o.tableOperations = tableOperations
}

func (o *TableBatchOperation) AppendTableOperation(tableOperation *TableOperation) {
	o.tableOperations = append(o.tableOperations, tableOperation)
}

func (o *TableBatchOperation) ReadOnly() bool {
	return o.readOnly
}

func (o *TableBatchOperation) SetReadOnly(readOnly bool) {
	o.readOnly = readOnly
}

func (o *TableBatchOperation) SameType() bool {
	return o.sameType
}

func (o *TableBatchOperation) SetSameType(sameType bool) {
	o.sameType = sameType
}

func (o *TableBatchOperation) SamePropertiesNames() bool {
	return o.samePropertiesNames
}

func (o *TableBatchOperation) SetSamePropertiesNames(samePropertiesNames bool) {
	o.samePropertiesNames = samePropertiesNames
}

func (o *TableBatchOperation) PayloadLen() int {
	return o.PayloadContentLen() + o.UniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (o *TableBatchOperation) PayloadContentLen() int {
	totalLen := 0

	totalLen += util.EncodedLengthByVi64(int64(len(o.tableOperations)))

	for _, tableOperation := range o.tableOperations {
		totalLen += tableOperation.PayloadLen()
	}

	totalLen += 3 // readOnly sameType samePropertiesNames

	o.UniVersionHeader.SetContentLength(totalLen)
	return o.UniVersionHeader.ContentLength()
}

func (o *TableBatchOperation) Encode(buffer *bytes.Buffer) {
	o.UniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(o.tableOperations)))

	for _, tableOperation := range o.tableOperations {
		tableOperation.Encode(buffer)
	}

	util.PutUint8(buffer, util.BoolToByte(o.readOnly))

	util.PutUint8(buffer, util.BoolToByte(o.sameType))

	util.PutUint8(buffer, util.BoolToByte(o.samePropertiesNames))
}

func (o *TableBatchOperation) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

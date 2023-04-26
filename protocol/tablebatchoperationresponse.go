package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableBatchOperationResponse struct {
	*UniVersionHeader
	tableOperationResponses []*TableOperationResponse
}

func (r *TableBatchOperationResponse) TableOperationResponses() []*TableOperationResponse {
	return r.tableOperationResponses
}

func (r *TableBatchOperationResponse) SetTableOperationResponses(tableOperationResponses []*TableOperationResponse) {
	r.tableOperationResponses = tableOperationResponses
}

func (r *TableBatchOperationResponse) AppendTableOperationResponse(tableOperationResponse *TableOperationResponse) {
	r.tableOperationResponses = append(r.tableOperationResponses, tableOperationResponse)
}

func (r *TableBatchOperationResponse) PCode() TablePacketCode {
	return TableApiBatchExecute
}

func (r *TableBatchOperationResponse) PayloadLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *TableBatchOperationResponse) PayloadContentLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *TableBatchOperationResponse) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *TableBatchOperationResponse) Decode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Decode(buffer)

	responsesLen := util.DecodeVi64(buffer)

	var i int64
	for i = 0; i < responsesLen; i++ {
		tableOperationResponse := NewTableOperationResponse()
		tableOperationResponse.Decode(buffer)
		r.tableOperationResponses = append(r.tableOperationResponses, tableOperationResponse)
	}
}

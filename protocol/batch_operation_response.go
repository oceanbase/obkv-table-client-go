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

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableBatchOperationResponse struct {
	*UniVersionHeader
	tableOperationResponses []*TableOperationResponse
}

func NewTableBatchOperationResponse() *TableBatchOperationResponse {
	return &TableBatchOperationResponse{
		UniVersionHeader:        NewUniVersionHeader(),
		tableOperationResponses: nil,
	}
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

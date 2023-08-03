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

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableBatchOperationResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	obTableOperationResponses []*ObTableOperationResponse
}

func NewObTableBatchOperationResponse() *ObTableBatchOperationResponse {
	return &ObTableBatchOperationResponse{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		ObPayloadBase: ObPayloadBase{
			uniqueId:  0,
			sequence:  0,
			tenantId:  1,
			sessionId: 0,
			flag:      0,
			timeout:   10 * 1000 * time.Millisecond,
		},
		obTableOperationResponses: nil,
	}
}

func (r *ObTableBatchOperationResponse) ObTableOperationResponses() []*ObTableOperationResponse {
	return r.obTableOperationResponses
}

func (r *ObTableBatchOperationResponse) SetObTableOperationResponses(obTableOperationResponses []*ObTableOperationResponse) {
	r.obTableOperationResponses = obTableOperationResponses
}

func (r *ObTableBatchOperationResponse) AppendObTableOperationResponse(obTableOperationResponse *ObTableOperationResponse) {
	r.obTableOperationResponses = append(r.obTableOperationResponses, obTableOperationResponse)
}

func (r *ObTableBatchOperationResponse) PCode() ObTablePacketCode {
	return ObTableApiBatchExecute
}

func (r *ObTableBatchOperationResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableBatchOperationResponse) PayloadContentLen() int {
	totalLen := util.EncodedLengthByVi64(int64(len(r.obTableOperationResponses)))

	for _, response := range r.obTableOperationResponses {
		totalLen += response.PayloadLen()
	}

	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableBatchOperationResponse) Credential() []byte {
	return nil
}

func (r *ObTableBatchOperationResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableBatchOperationResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(r.obTableOperationResponses)))

	for _, response := range r.obTableOperationResponses {
		response.Encode(buffer)
	}
}

func (r *ObTableBatchOperationResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	responsesLen := util.DecodeVi64(buffer)

	r.obTableOperationResponses = make([]*ObTableOperationResponse, 0, responsesLen)

	var i int64
	for i = 0; i < responsesLen; i++ {
		obTableOperationResponse := NewObTableOperationResponse()
		obTableOperationResponse.Decode(buffer)
		r.obTableOperationResponses = append(r.obTableOperationResponses, obTableOperationResponse)
	}
}

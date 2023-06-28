/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
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
	"time"
)

type ObTableAsyncQueryRequest struct {
	ObUniVersionHeader
	ObPayloadBase
	tableQueryRequest *ObTableQueryRequest
	querySessionId    int64
	queryType         ObQueryOperationType
}

// NewObTableAsyncQueryRequestWithParams creates a new ObTableAsyncQueryRequest.
func NewObTableAsyncQueryRequestWithParams(
	tableQueryRequest *ObTableQueryRequest,
	timeout time.Duration,
	flag uint16) *ObTableAsyncQueryRequest {
	return &ObTableAsyncQueryRequest{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		ObPayloadBase: ObPayloadBase{
			uniqueId:  0,
			sequence:  0,
			tenantId:  1,
			sessionId: 0,
			flag:      flag,
			timeout:   timeout,
		},
		tableQueryRequest: tableQueryRequest,
		querySessionId:    0,
		queryType:         QueryStart,
	}
}

func (r *ObTableAsyncQueryRequest) PCode() ObTablePacketCode {
	return ObTableApiExecuteAsyncQuery
}

func (r *ObTableAsyncQueryRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableAsyncQueryRequest) PayloadContentLen() int {
	totalLen := r.tableQueryRequest.PayloadLen() + util.EncodedLengthByVi64(r.querySessionId) + 1
	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableAsyncQueryRequest) Credential() []byte {
	return r.tableQueryRequest.Credential()
}

func (r *ObTableAsyncQueryRequest) SetCredential(credential []byte) {
	r.tableQueryRequest.SetCredential(credential)
}

// SetQuerySessionId sets the query session id.
func (r *ObTableAsyncQueryRequest) SetQuerySessionId(querySessionId int64) {
	r.querySessionId = querySessionId
}

// SetQueryType sets the query type.
func (r *ObTableAsyncQueryRequest) SetQueryType(queryType ObQueryOperationType) {
	r.queryType = queryType
}

func (r *ObTableAsyncQueryRequest) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	r.tableQueryRequest.Encode(buffer)

	util.EncodeVi64(buffer, r.querySessionId)

	util.PutUint8(buffer, uint8(r.queryType))
}

func (r *ObTableAsyncQueryRequest) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.tableQueryRequest.Decode(buffer)

	r.querySessionId = util.DecodeVi64(buffer)

	r.queryType = ObQueryOperationType(util.Uint8(buffer))
}

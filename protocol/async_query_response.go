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
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableAsyncQueryResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	querySessionId       int64
	isEnd                bool
	obTableQueryResponse *ObTableQueryResponse
}

// NewObTableAsyncQueryResponse creates a new ObTableAsyncQueryResponse.
func NewObTableAsyncQueryResponse() *ObTableAsyncQueryResponse {
	return &ObTableAsyncQueryResponse{
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
		querySessionId:       0,
		isEnd:                false,
		obTableQueryResponse: NewObTableQueryResponse(),
	}
}

func (r *ObTableAsyncQueryResponse) QuerySessionId() int64 {
	return r.querySessionId
}

func (r *ObTableAsyncQueryResponse) SetQuerySessionId(querySessionId int64) {
	r.querySessionId = querySessionId
}

func (r *ObTableAsyncQueryResponse) IsEnd() bool {
	return r.isEnd
}

func (r *ObTableAsyncQueryResponse) SetIsEnd(isEnd bool) {
	r.isEnd = isEnd
}

func (r *ObTableAsyncQueryResponse) PropertiesNames() []string {
	return r.obTableQueryResponse.PropertiesNames()
}

func (r *ObTableAsyncQueryResponse) ResultRowCount() int64 {
	return r.obTableQueryResponse.RowCount()
}

func (r *ObTableAsyncQueryResponse) PropertiesRows() [][]*ObObject {
	return r.obTableQueryResponse.PropertiesRows()
}

func (r *ObTableAsyncQueryResponse) PCode() ObTablePacketCode {
	return ObTableApiExecuteAsyncQuery
}

func (r *ObTableAsyncQueryResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableAsyncQueryResponse) PayloadContentLen() int {
	totalLen := r.obTableQueryResponse.PayloadLen() + util.EncodedLengthByVi64(r.querySessionId) + 1
	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableAsyncQueryResponse) Credential() []byte {
	return nil
}

func (r *ObTableAsyncQueryResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableAsyncQueryResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	r.obTableQueryResponse.Encode(buffer)

	util.PutUint8(buffer, util.BoolToByte(r.isEnd))

	util.EncodeVi64(buffer, r.querySessionId)
}

func (r *ObTableAsyncQueryResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.obTableQueryResponse.Decode(buffer)

	r.isEnd = util.ByteToBool(util.Uint8(buffer))

	r.querySessionId = util.DecodeVi64(buffer)
}

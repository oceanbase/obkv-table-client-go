/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
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

type ObTableQueryAndMutateResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	affectedRows   int64
	affectedEntity *ObTableQueryResponse
}

func NewObTableQueryAndMutateResponse() *ObTableQueryAndMutateResponse {
	return &ObTableQueryAndMutateResponse{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		ObPayloadBase: ObPayloadBase{
			uniqueId:  0,
			sequence:  0,
			tenantId:  1,
			sessionId: 0,
			flag:      7,
			timeout:   10 * 1000 * time.Millisecond,
		},
		affectedRows:   0,
		affectedEntity: NewObTableQueryResponse(),
	}
}

func (r *ObTableQueryAndMutateResponse) AffectedRows() int64 {
	return r.affectedRows
}

func (r *ObTableQueryAndMutateResponse) SetAffectedRows(affectedRows int64) {
	r.affectedRows = affectedRows
}

func (r *ObTableQueryAndMutateResponse) AffectedEntity() *ObTableQueryResponse {
	return r.affectedEntity
}

func (r *ObTableQueryAndMutateResponse) SetAffectedEntity(affectedEntity *ObTableQueryResponse) {
	r.affectedEntity = affectedEntity
}

func (r *ObTableQueryAndMutateResponse) PCode() ObTablePacketCode {
	return ObTableApiQueryAndMute
}

func (r *ObTableQueryAndMutateResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableQueryAndMutateResponse) PayloadContentLen() int {
	totalLen := 0
	totalLen +=
		util.EncodedLengthByVi64(r.affectedRows) +
			r.affectedEntity.PayloadLen()

	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableQueryAndMutateResponse) Credential() []byte {
	return nil
}

func (r *ObTableQueryAndMutateResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableQueryAndMutateResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, r.affectedRows)

	r.affectedEntity.Encode(buffer)
}

func (r *ObTableQueryAndMutateResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.affectedRows = util.DecodeVi64(buffer)

	r.affectedEntity.Decode(buffer)
}

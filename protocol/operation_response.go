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

const (
	IsInsertUpDoInsertMask = 1 << 0
	IsInsertUpDoPutMask    = 1 << 1
)

type ObTableOperationResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	header        *ObTableResponse
	operationType ObTableOperationType
	entity        *ObTableEntity
	affectedRows  int64
	flags         uint64
}

func NewObTableOperationResponse() *ObTableOperationResponse {
	return &ObTableOperationResponse{
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
		header:        NewObTableResponse(),
		operationType: ObTableOperationGet,
		entity:        NewObTableEntity(),
		affectedRows:  0,
		flags:         0,
	}
}

func (r *ObTableOperationResponse) Header() *ObTableResponse {
	return r.header
}

func (r *ObTableOperationResponse) SetHeader(header *ObTableResponse) {
	r.header = header
}

func (r *ObTableOperationResponse) OperationType() ObTableOperationType {
	return r.operationType
}

func (r *ObTableOperationResponse) SetOperationType(operationType ObTableOperationType) {
	r.operationType = operationType
}

func (r *ObTableOperationResponse) Entity() *ObTableEntity {
	return r.entity
}

func (r *ObTableOperationResponse) SetEntity(entity *ObTableEntity) {
	r.entity = entity
}

func (r *ObTableOperationResponse) AffectedRows() int64 {
	return r.affectedRows
}

func (r *ObTableOperationResponse) SetAffectedRows(affectedRows int64) {
	r.affectedRows = affectedRows
}

func (r *ObTableOperationResponse) Flags() uint64 {
	return r.flags
}

func (r *ObTableOperationResponse) PCode() ObTablePacketCode {
	if r.operationType == ObTableOperationRedis {
		return ObTableApiRedis
	}
	return ObTableApiExecute
}

func (r *ObTableOperationResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableOperationResponse) PayloadContentLen() int {
	totalLen := r.header.PayloadLen() +
		1 +
		r.entity.PayloadLen() +
		util.EncodedLengthByVi64(r.affectedRows) +
		util.EncodedLengthByVi64(int64(r.flags))

	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableOperationResponse) Credential() []byte {
	return nil
}

func (r *ObTableOperationResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableOperationResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	r.header.Encode(buffer)

	util.PutUint8(buffer, uint8(r.operationType))

	r.entity.Encode(buffer)

	util.EncodeVi64(buffer, r.affectedRows)
	util.EncodeVi64(buffer, int64(r.flags))
}

func (r *ObTableOperationResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.header.Decode(buffer)

	r.operationType = ObTableOperationType(util.Uint8(buffer))

	r.entity.Decode(buffer)

	r.affectedRows = util.DecodeVi64(buffer)

	r.flags = uint64(util.DecodeVi64(buffer))
}

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

type ObTableOperationResponse struct {
	*ObUniVersionHeader
	header        *ObTableResponse
	operationType ObTableOperationType
	entity        *ObTableEntity
	affectedRows  int64
}

func NewObTableOperationResponse() *ObTableOperationResponse {
	return &ObTableOperationResponse{
		ObUniVersionHeader: NewObUniVersionHeader(),
		header:             NewObTableResponse(),
		operationType:      ObTableOperationGet,
		entity:             NewObTableEntity(),
		affectedRows:       0,
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

func (r *ObTableOperationResponse) PCode() ObTablePacketCode {
	return ObTableApiExecute
}

func (r *ObTableOperationResponse) PayloadLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *ObTableOperationResponse) PayloadContentLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *ObTableOperationResponse) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *ObTableOperationResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.header.Decode(buffer)

	r.operationType = ObTableOperationType(util.Uint8(buffer))

	r.entity.Decode(buffer)

	r.affectedRows = util.DecodeVi64(buffer)
}

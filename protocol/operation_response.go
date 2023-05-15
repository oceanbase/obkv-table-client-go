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

type TableOperationResponse struct {
	*UniVersionHeader
	header        *TableResponse
	operationType TableOperationType
	entity        *TableEntity
	affectedRows  int64
}

func NewTableOperationResponse() *TableOperationResponse {
	return &TableOperationResponse{
		UniVersionHeader: NewUniVersionHeader(),
		header:           NewTableResponse(),
		operationType:    Get,
		entity:           NewTableEntity(),
		affectedRows:     0,
	}
}

func (r *TableOperationResponse) Header() *TableResponse {
	return r.header
}

func (r *TableOperationResponse) SetHeader(header *TableResponse) {
	r.header = header
}

func (r *TableOperationResponse) OperationType() TableOperationType {
	return r.operationType
}

func (r *TableOperationResponse) SetOperationType(operationType TableOperationType) {
	r.operationType = operationType
}

func (r *TableOperationResponse) Entity() *TableEntity {
	return r.entity
}

func (r *TableOperationResponse) SetEntity(entity *TableEntity) {
	r.entity = entity
}

func (r *TableOperationResponse) AffectedRows() int64 {
	return r.affectedRows
}

func (r *TableOperationResponse) SetAffectedRows(affectedRows int64) {
	r.affectedRows = affectedRows
}

func (r *TableOperationResponse) PCode() TablePacketCode {
	return TableApiExecute
}

func (r *TableOperationResponse) PayloadLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *TableOperationResponse) PayloadContentLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *TableOperationResponse) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *TableOperationResponse) Decode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Decode(buffer)

	r.header.Decode(buffer)

	r.operationType = TableOperationType(util.Uint8(buffer))

	r.entity.Decode(buffer)

	r.affectedRows = util.DecodeVi64(buffer)
}

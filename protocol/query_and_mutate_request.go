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

type ObTableQueryAndMutateRequest struct {
	ObUniVersionHeader
	ObPayloadBase
	credential          []byte
	tableName           string
	tableId             uint64
	partitionId         uint64
	entityType          ObTableEntityType
	tableQueryAndMutate *ObTableQueryAndMutate
}

func NewObTableQueryAndMutateRequest() *ObTableQueryAndMutateRequest {
	return &ObTableQueryAndMutateRequest{
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
		credential:          nil,
		tableName:           "",
		tableId:             0,
		partitionId:         0,
		entityType:          0,
		tableQueryAndMutate: NewObTableQueryAndMutate(),
	}
}

func (r *ObTableQueryAndMutateRequest) TableName() string {
	return r.tableName
}

func (r *ObTableQueryAndMutateRequest) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *ObTableQueryAndMutateRequest) TableId() uint64 {
	return r.tableId
}

func (r *ObTableQueryAndMutateRequest) SetTableId(tableId uint64) {
	r.tableId = tableId
}

func (r *ObTableQueryAndMutateRequest) PartitionId() uint64 {
	return r.partitionId
}

func (r *ObTableQueryAndMutateRequest) SetPartitionId(partitionId uint64) {
	r.partitionId = partitionId
}

func (r *ObTableQueryAndMutateRequest) EntityType() ObTableEntityType {
	return r.entityType
}

func (r *ObTableQueryAndMutateRequest) SetEntityType(entityType ObTableEntityType) {
	r.entityType = entityType
}

func (r *ObTableQueryAndMutateRequest) TableQueryAndMutate() *ObTableQueryAndMutate {
	return r.tableQueryAndMutate
}

func (r *ObTableQueryAndMutateRequest) SetTableQueryAndMutate(tableQueryAndMutate *ObTableQueryAndMutate) {
	r.tableQueryAndMutate = tableQueryAndMutate
}

func (r *ObTableQueryAndMutateRequest) PCode() ObTablePacketCode {
	return ObTableApiQueryAndMute
}

func (r *ObTableQueryAndMutateRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableQueryAndMutateRequest) PayloadContentLen() int {
	totalLen := 0
	if util.ObVersion() >= 4 {
		totalLen +=
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				8 + // partitionId
				1 + // entityType
				r.tableQueryAndMutate.PayloadLen()
	} else {
		totalLen +=
			util.EncodedLengthByBytesString(r.credential) +
				util.EncodedLengthByVString(r.tableName) +
				util.EncodedLengthByVi64(int64(r.tableId)) +
				util.EncodedLengthByVi64(int64(r.partitionId)) + // partitionId
				1 + // entityType
				r.tableQueryAndMutate.PayloadLen()
	}
	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableQueryAndMutateRequest) Credential() []byte {
	return r.credential
}

func (r *ObTableQueryAndMutateRequest) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *ObTableQueryAndMutateRequest) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.EncodeBytesString(buffer, r.credential)

	util.EncodeVString(buffer, r.tableName)

	util.EncodeVi64(buffer, int64(r.tableId))

	if util.ObVersion() >= 4 {
		util.PutUint64(buffer, r.partitionId)
	} else {
		util.EncodeVi64(buffer, int64(r.partitionId))
	}

	util.PutUint8(buffer, uint8(r.entityType))

	r.tableQueryAndMutate.Encode(buffer)
}

func (r *ObTableQueryAndMutateRequest) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.credential = util.DecodeBytesString(buffer)

	r.tableName = util.DecodeVString(buffer)

	r.tableId = uint64(util.DecodeVi64(buffer))

	if util.ObVersion() >= 4 {
		r.partitionId = util.Uint64(buffer)
	} else {
		r.partitionId = uint64(util.DecodeVi64(buffer))
	}

	r.entityType = ObTableEntityType(util.Uint8(buffer))

	r.tableQueryAndMutate.Decode(buffer)
}

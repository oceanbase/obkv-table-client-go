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
	"strconv"
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableMoveResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	replicaInfo *ObTableMoveReplicaInfo
	reserved    uint64
}

func (r *ObTableMoveResponse) ReplicaInfo() *ObTableMoveReplicaInfo {
	return r.replicaInfo
}

func (r *ObTableMoveResponse) SetReplicaInfo(replicaInfo *ObTableMoveReplicaInfo) {
	r.replicaInfo = replicaInfo
}

func (r *ObTableMoveResponse) Reserved() uint64 {
	return r.reserved
}

func (r *ObTableMoveResponse) SetReserved(reserved uint64) {
	r.reserved = reserved
}

func NewObTableMoveResponse() *ObTableMoveResponse {
	return &ObTableMoveResponse{
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
		replicaInfo: NewObTableMoveReplicaInfo(),
		reserved:    0,
	}
}

func (r *ObTableMoveResponse) Valid() bool {
	return r.replicaInfo != nil && r.replicaInfo.tableId != 0
}

func (r *ObTableMoveResponse) PCode() ObTablePacketCode {
	return ObTableApiMove
}

func (r *ObTableMoveResponse) Credential() []byte {
	return nil
}

func (r *ObTableMoveResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableMoveResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableMoveResponse) PayloadContentLen() int {
	totalLen := r.replicaInfo.PayloadLen()

	totalLen += util.EncodedLengthByVi64(int64(r.reserved))

	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableMoveResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	r.replicaInfo.Encode(buffer)

	util.EncodeVi64(buffer, int64(r.reserved))
}

func (r *ObTableMoveResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.replicaInfo.Decode(buffer)

	r.reserved = uint64(util.DecodeVi64(buffer))
}

func (r *ObTableMoveResponse) String() string {
	var ObUniVersionHeaderStr = "nil"
	if r.ObUniVersionHeader != (ObUniVersionHeader{}) {
		ObUniVersionHeaderStr = r.ObUniVersionHeader.String()
	}

	return "ObAddr{" +
		"ObUniVersionHeader:" + ObUniVersionHeaderStr + ", " +
		"replicaInfo:" + r.replicaInfo.String() + ", " +
		"reserved:" + strconv.Itoa(int(r.reserved)) + ", " +
		"}"
}

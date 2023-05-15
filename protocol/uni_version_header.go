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

// ObUniVersionHeader ...
type ObUniVersionHeader struct {
	version       int64
	contentLength int

	flag      uint16
	tenantId  uint64
	sessionId uint64
	timeout   time.Duration

	uniqueId uint64 // rpc header traceId0
	sequence uint64 // rpc header traceId1
}

func NewObUniVersionHeader() *ObUniVersionHeader {
	return &ObUniVersionHeader{
		version:       1,
		contentLength: 0,
		flag:          7,
		tenantId:      1,
		sessionId:     0,
		timeout:       10 * 1000 * time.Millisecond,
		uniqueId:      0,
		sequence:      0,
	}
}

func (h *ObUniVersionHeader) UniVersionHeaderLen() int {
	return util.EncodedLengthByVi64(h.version) + util.EncodedLengthByVi64(int64(h.contentLength))
}

func (h *ObUniVersionHeader) ContentLength() int {
	return h.contentLength
}

func (h *ObUniVersionHeader) SetContentLength(contentLength int) {
	h.contentLength = contentLength
}

func (h *ObUniVersionHeader) PCode() ObTablePacketCode {
	return ObTableApiErrorPacket
}

func (h *ObUniVersionHeader) UniqueId() uint64 {
	return h.uniqueId
}

func (h *ObUniVersionHeader) SetUniqueId(uniqueId uint64) {
	h.uniqueId = uniqueId
}

func (h *ObUniVersionHeader) Sequence() uint64 {
	return h.sequence
}

func (h *ObUniVersionHeader) SetSequence(sequence uint64) {
	h.sequence = sequence
}

func (h *ObUniVersionHeader) TenantId() uint64 {
	return h.tenantId
}

func (h *ObUniVersionHeader) SetTenantId(tenantId uint64) {
	h.tenantId = tenantId
}

func (h *ObUniVersionHeader) SessionId() uint64 {
	return h.sessionId
}

func (h *ObUniVersionHeader) SetSessionId(sessionId uint64) {
	h.sessionId = sessionId
}

func (h *ObUniVersionHeader) Flag() uint16 {
	return h.flag
}

func (h *ObUniVersionHeader) SetFlag(flag uint16) {
	h.flag = flag
}

func (h *ObUniVersionHeader) Version() int64 {
	return h.version
}

func (h *ObUniVersionHeader) SetVersion(version int64) {
	h.version = version
}

func (h *ObUniVersionHeader) Timeout() time.Duration {
	return h.timeout
}

func (h *ObUniVersionHeader) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *ObUniVersionHeader) Credential() []byte {
	return nil
}

func (h *ObUniVersionHeader) SetCredential(credential []byte) {
	return
}

func (h *ObUniVersionHeader) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, h.version)
	util.EncodeVi64(buffer, int64(h.contentLength)) // payloadLen
}

func (h *ObUniVersionHeader) Decode(buffer *bytes.Buffer) {
	h.version = util.DecodeVi64(buffer)
	h.contentLength = int(util.DecodeVi64(buffer)) // contentLength useless right now
}

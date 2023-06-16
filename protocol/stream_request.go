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
)

type ObTableStreamRequest struct {
	ObUniVersionHeader
	ObPayloadBase
	sessionId uint64
	flag      uint16
}

func (r *ObTableStreamRequest) SessionId() uint64 {
	return r.sessionId
}

func (r *ObTableStreamRequest) SetSessionId(sessionId uint64) {
	r.sessionId = sessionId
}

func (r *ObTableStreamRequest) Flag() uint16 {
	return r.flag
}

func (r *ObTableStreamRequest) SetFlag(flag uint16) {
	r.flag = flag
}

func (r *ObTableStreamRequest) PCode() ObTablePacketCode {
	return ObTableApiExecuteQuery
}

func (r *ObTableStreamRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableStreamRequest) PayloadContentLen() int {
	r.ObUniVersionHeader.SetContentLength(0)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableStreamRequest) Credential() []byte {
	return nil
}

func (r *ObTableStreamRequest) SetCredential(credential []byte) {
	return
}

func (r *ObTableStreamRequest) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)
}

func (r *ObTableStreamRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

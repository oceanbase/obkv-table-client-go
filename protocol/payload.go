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
	"time"
)

// ObPayload ...
type ObPayload interface {
	ProtoEncoder
	ProtoDecoder

	PCode() ObTablePacketCode

	PayloadLen() int

	PayloadContentLen() int

	Credential() []byte
	SetCredential(credential []byte)

	Version() int64
	SetVersion(version int64)

	UniqueId() uint64
	SetUniqueId(uniqueId uint64)

	Sequence() uint64
	SetSequence(sequence uint64)

	Flag() uint16
	SetFlag(flag uint16)

	TenantId() uint64
	SetTenantId(tenantId uint64)

	SessionId() uint64
	SetSessionId(sessionId uint64)

	Timeout() time.Duration
	SetTimeout(timeout time.Duration)

	MoveResponse() *ObTableMoveResponse
	SetMoveResponse(response *ObTableMoveResponse)
}

package protocol

import (
	"time"
)

// Payload ...
type Payload interface {
	ProtoEncoder
	ProtoDecoder

	PCode() TablePacketCode

	PayloadLen() int

	PayloadContentLen() int

	UniqueId() uint64
	SetUniqueId(uniqueId uint64)

	Sequence() uint64
	SetSequence(sequence uint64)

	TenantId() uint64
	SetTenantId(tenantId uint64)

	SessionId() uint64
	SetSessionId(sessionId uint64)

	Flag() uint16
	SetFlag(flag uint16)

	Version() int64
	SetVersion(version int64)

	Timeout() time.Duration
	SetTimeout(timeout time.Duration)

	Credential() []byte
	SetCredential(credential []byte)
}

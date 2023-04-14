package protocol

import (
	"time"
)

// Payload ...
type Payload interface {
	ProtoEncoder
	ProtoDecoder

	PCode() TablePacketCode

	PayloadLen() int64

	PayloadContentLen() int64

	TenantId() uint64
	SetTenantId(tenantId uint64)

	SessionId() uint64
	SetSessionId(sessionId uint64)

	Flag() uint16
	SetFlag(flag uint16)

	Version() int64
	SetVersion(version int64)

	Timeout() time.Duration
	SetTimeout(duration time.Duration)

	Credential() string
	SetCredential(credential string)
}
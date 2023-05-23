package protocol

import (
	"time"
)

// ObPayloadBase payload base
type ObPayloadBase struct {
	uniqueId uint64 // rpc header traceId0
	sequence uint64 // rpc header traceId1

	flag      uint16
	tenantId  uint64
	sessionId uint64

	timeout time.Duration
}

func NewObPayloadBase() *ObPayloadBase {
	return &ObPayloadBase{
		uniqueId:  0,
		sequence:  0,
		tenantId:  1,
		sessionId: 0,
		flag:      7,
		timeout:   10 * 1000 * time.Millisecond,
	}
}

func (p *ObPayloadBase) UniqueId() uint64 {
	return p.uniqueId
}

func (p *ObPayloadBase) SetUniqueId(uniqueId uint64) {
	p.uniqueId = uniqueId
}

func (p *ObPayloadBase) Sequence() uint64 {
	return p.sequence
}

func (p *ObPayloadBase) SetSequence(sequence uint64) {
	p.sequence = sequence
}

func (p *ObPayloadBase) Flag() uint16 {
	return p.flag
}

func (p *ObPayloadBase) SetFlag(flag uint16) {
	p.flag = flag
}

func (p *ObPayloadBase) TenantId() uint64 {
	return p.tenantId
}

func (p *ObPayloadBase) SetTenantId(tenantId uint64) {
	p.tenantId = tenantId
}

func (p *ObPayloadBase) SessionId() uint64 {
	return p.sessionId
}

func (p *ObPayloadBase) SetSessionId(sessionId uint64) {
	p.sessionId = sessionId
}

func (p *ObPayloadBase) Timeout() time.Duration {
	return p.timeout
}

func (p *ObPayloadBase) SetTimeout(timeout time.Duration) {
	p.timeout = timeout
}

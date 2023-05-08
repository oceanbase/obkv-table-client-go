package protocol

import (
	"bytes"
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

// UniVersionHeader ...
type UniVersionHeader struct {
	version       int64
	contentLength int

	flag      uint16
	tenantId  uint64
	sessionId uint64
	timeout   time.Duration
}

func NewUniVersionHeader() *UniVersionHeader {
	return &UniVersionHeader{
		version:       1,
		contentLength: 0,
		flag:          7,
		tenantId:      1,
		sessionId:     0,
		timeout:       10 * 1000 * time.Millisecond,
	}
}

func (h *UniVersionHeader) UniVersionHeaderLen() int {
	return util.EncodedLengthByVi64(h.version) + util.EncodedLengthByVi64(int64(h.contentLength))
}

func (h *UniVersionHeader) ContentLength() int {
	return h.contentLength
}

func (h *UniVersionHeader) SetContentLength(contentLength int) {
	h.contentLength = contentLength
}

func (h *UniVersionHeader) PCode() TablePacketCode {
	return TableApiErrorPacket
}

func (h *UniVersionHeader) TenantId() uint64 {
	return h.tenantId
}

func (h *UniVersionHeader) SetTenantId(tenantId uint64) {
	h.tenantId = tenantId
}

func (h *UniVersionHeader) SessionId() uint64 {
	return h.sessionId
}

func (h *UniVersionHeader) SetSessionId(sessionId uint64) {
	h.sessionId = sessionId
}

func (h *UniVersionHeader) Flag() uint16 {
	return h.flag
}

func (h *UniVersionHeader) SetFlag(flag uint16) {
	h.flag = flag
}

func (h *UniVersionHeader) Version() int64 {
	return h.version
}

func (h *UniVersionHeader) SetVersion(version int64) {
	h.version = version
}

func (h *UniVersionHeader) Timeout() time.Duration {
	return h.timeout
}

func (h *UniVersionHeader) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *UniVersionHeader) Credential() []byte {
	return nil
}

func (h *UniVersionHeader) SetCredential(credential []byte) {
	return
}

func (h *UniVersionHeader) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, h.version)
	util.EncodeVi64(buffer, int64(h.contentLength)) // payloadLen
}

func (h *UniVersionHeader) Decode(buffer *bytes.Buffer) {
	h.version = util.DecodeVi64(buffer)
	h.contentLength = int(util.DecodeVi64(buffer)) // contentLength useless right now
}

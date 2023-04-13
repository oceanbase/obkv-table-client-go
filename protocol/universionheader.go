package protocol

import (
	"bytes"
	"sync"
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

// UniVersionHeader ...
type UniVersionHeader struct {
	version       int64
	contentLength int64

	calculateLengthOnce sync.Once

	flag      uint16
	channelId uint32
	uniqueId  uint32
	sequence  uint32
	tenantId  uint64

	timeout time.Duration
}

func NewUniVersionHeader() *UniVersionHeader {
	return &UniVersionHeader{
		version:             1,
		contentLength:       0,
		calculateLengthOnce: sync.Once{},
		flag:                7,
		channelId:           0,
		uniqueId:            0,
		sequence:            0,
		tenantId:            1,
		timeout:             10 * 1000 * time.Millisecond,
	}
}

func (h *UniVersionHeader) UniVersionHeaderLen() int {
	return util.NeedLengthByVi64(h.version) + util.NeedLengthByVi64(h.contentLength)
}

func (h *UniVersionHeader) Version() int64 {
	return h.version
}

func (h *UniVersionHeader) SetVersion(version int64) {
	h.version = version
}

func (h *UniVersionHeader) ContentLength() int64 {
	return h.contentLength
}

func (h *UniVersionHeader) SetContentLength(contentLength int64) {
	h.contentLength = contentLength
}

func (h *UniVersionHeader) Flag() uint16 {
	return h.flag
}

func (h *UniVersionHeader) SetFlag(flag uint16) {
	h.flag = flag
}

func (h *UniVersionHeader) ChannelId() uint32 {
	return h.channelId
}

func (h *UniVersionHeader) SetChannelId(channelId uint32) {
	h.channelId = channelId
}

func (h *UniVersionHeader) UniqueId() uint32 {
	return h.uniqueId
}

func (h *UniVersionHeader) SetUniqueId(uniqueId uint32) {
	h.uniqueId = uniqueId
}

func (h *UniVersionHeader) Sequence() uint32 {
	return h.sequence
}

func (h *UniVersionHeader) SetSequence(sequence uint32) {
	h.sequence = sequence
}

func (h *UniVersionHeader) TenantId() uint64 {
	return h.tenantId
}

func (h *UniVersionHeader) SetTenantId(tenantId uint64) {
	h.tenantId = tenantId
}

func (h *UniVersionHeader) Timeout() time.Duration {
	return h.timeout
}

func (h *UniVersionHeader) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *UniVersionHeader) Encode(buf []byte) {
	versionLen := util.NeedLengthByVi64(h.version)
	util.EncodeVi64(buf[:versionLen], h.version)
	util.EncodeVi64(buf[versionLen:], h.contentLength)
}

func (h *UniVersionHeader) Decode(buffer *bytes.Buffer) {
	h.version = util.DecodeVi64(buffer)
	h.contentLength = util.DecodeVi64(buffer) // payloadLen
}

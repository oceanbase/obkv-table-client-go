package protocol

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"
)

const (
	EzHeaderLength       = 16
	version        uint8 = 1
)

var (
	MagicHeaderFlag = []uint8{version, 0xDB, 0xDB, 0xCE}
	Reserved        = []uint8{0, 0, 0, 0}
)

// EzHeader ...
type EzHeader struct {
	contentLen uint32 // ContentLen = RpcHeader + Payload
	channelId  uint32
}

func (h *EzHeader) ContentLen() uint32 {
	return h.contentLen
}

func (h *EzHeader) SetContentLen(contentLen uint32) {
	h.contentLen = contentLen
}

func (h *EzHeader) ChannelId() uint32 {
	return h.channelId
}

func (h *EzHeader) SetChannelId(channelId uint32) {
	h.channelId = channelId
}

func NewEzHeader() *EzHeader {
	return &EzHeader{}
}

func (h *EzHeader) Encode() []byte {
	ezHeaderBuf := make([]byte, EzHeaderLength)
	copy(ezHeaderBuf[0:4], MagicHeaderFlag)
	binary.BigEndian.PutUint32(ezHeaderBuf[4:8], h.contentLen)
	binary.BigEndian.PutUint32(ezHeaderBuf[8:12], h.channelId)
	copy(ezHeaderBuf[12:16], Reserved)
	return ezHeaderBuf
}

func (h *EzHeader) Decode(buffer *bytes.Buffer) error {
	if ok := bytes.Equal(MagicHeaderFlag, buffer.Next(4)); !ok {
		return errors.New("magic header flag not match")
	}
	h.contentLen = binary.BigEndian.Uint32(buffer.Next(4))
	h.channelId = binary.BigEndian.Uint32(buffer.Next(4))
	_ = buffer.Next(4) // reserved
	return nil
}

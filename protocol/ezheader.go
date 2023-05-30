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
	"encoding/binary"

	"github.com/pkg/errors"
)

const (
	EzHeaderLength       = 16
	version        uint8 = 1
)

var (
	magicHeaderFlag = []uint8{version, 0xDB, 0xDB, 0xCE}
	reserved        = []uint8{0, 0, 0, 0}

	ErrHeaderNotMatch = errors.New("magic header flag not match")
)

// EzHeader ...
type EzHeader struct {
	contentLen uint32 // ContentLen = ObRpcHeader + ObPayload
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

func (h *EzHeader) Encode(ezHeaderBuf []byte) {
	copy(ezHeaderBuf[:4], magicHeaderFlag)
	binary.BigEndian.PutUint32(ezHeaderBuf[4:8], h.contentLen)
	binary.BigEndian.PutUint32(ezHeaderBuf[8:12], h.channelId)
	copy(ezHeaderBuf[12:16], reserved)
}

func (h *EzHeader) Decode(ezHeaderBuf []byte) error {
	if ok := bytes.Equal(magicHeaderFlag, ezHeaderBuf[0:4]); !ok {
		return ErrHeaderNotMatch
	}
	h.contentLen = binary.BigEndian.Uint32(ezHeaderBuf[4:8])
	h.channelId = binary.BigEndian.Uint32(ezHeaderBuf[8:12])
	_ = ezHeaderBuf[12:16] // reserved
	return nil
}

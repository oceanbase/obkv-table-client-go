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

	"github.com/oceanbase/obkv-table-client-go/util"
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

func NewEzHeader() *EzHeader {
	return &EzHeader{}
}

func (h *EzHeader) Encode(ezHeaderBuf []byte) {
	copy(ezHeaderBuf[:4], MagicHeaderFlag)
	binary.BigEndian.PutUint32(ezHeaderBuf[4:8], h.contentLen)
	binary.BigEndian.PutUint32(ezHeaderBuf[8:12], h.channelId)
	copy(ezHeaderBuf[12:16], Reserved)
}

func (h *EzHeader) Decode(buffer *bytes.Buffer) error {
	if ok := bytes.Equal(MagicHeaderFlag, buffer.Next(4)); !ok {
		return errors.New("magic header flag not match")
	}
	h.contentLen = util.Uint32(buffer)
	h.channelId = util.Uint32(buffer)
	_ = buffer.Next(4) // reserved
	return nil
}

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

	"github.com/oceanbase/obkv-table-client-go/util"
)

// ObUniVersionHeader ...
type ObUniVersionHeader struct {
	version       int64
	contentLength int
}

func NewObUniVersionHeader() *ObUniVersionHeader {
	return &ObUniVersionHeader{
		version:       1,
		contentLength: 0,
	}
}

func (h *ObUniVersionHeader) Version() int64 {
	return h.version
}

func (h *ObUniVersionHeader) SetVersion(version int64) {
	h.version = version
}

func (h *ObUniVersionHeader) ContentLength() int {
	return h.contentLength
}

func (h *ObUniVersionHeader) SetContentLength(contentLength int) {
	h.contentLength = contentLength
}

func (h *ObUniVersionHeader) UniVersionHeaderLen() int {
	return util.EncodedLengthByVi64(h.version) + util.EncodedLengthByVi64(int64(h.contentLength))
}

func (h *ObUniVersionHeader) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, h.version)
	util.EncodeVi64(buffer, int64(h.contentLength)) // payloadLen
}

func (h *ObUniVersionHeader) Decode(buffer *bytes.Buffer) {
	h.version = util.DecodeVi64(buffer)
	h.contentLength = int(util.DecodeVi64(buffer)) // contentLength useless right now
}

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

type ObRpcResponseWarningMsg struct {
	ObUniVersionHeader
	msg       []byte
	timestamp int64
	logLevel  int32
	lineNo    int32
	code      int32
}

func NewObRpcResponseWarningMsg() *ObRpcResponseWarningMsg {
	return &ObRpcResponseWarningMsg{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		msg:       nil,
		timestamp: 0,
		logLevel:  0,
		lineNo:    0,
		code:      0,
	}
}

func (m *ObRpcResponseWarningMsg) Msg() []byte {
	return m.msg
}

func (m *ObRpcResponseWarningMsg) SetMsg(msg []byte) {
	m.msg = msg
}

func (m *ObRpcResponseWarningMsg) Timestamp() int64 {
	return m.timestamp
}

func (m *ObRpcResponseWarningMsg) SetTimestamp(timestamp int64) {
	m.timestamp = timestamp
}

func (m *ObRpcResponseWarningMsg) LogLevel() int32 {
	return m.logLevel
}

func (m *ObRpcResponseWarningMsg) SetLogLevel(logLevel int32) {
	m.logLevel = logLevel
}

func (m *ObRpcResponseWarningMsg) LineNo() int32 {
	return m.lineNo
}

func (m *ObRpcResponseWarningMsg) SetLineNo(lineNo int32) {
	m.lineNo = lineNo
}

func (m *ObRpcResponseWarningMsg) Code() int32 {
	return m.code
}

func (m *ObRpcResponseWarningMsg) SetCode(code int32) {
	m.code = code
}

func (m *ObRpcResponseWarningMsg) PayloadLen() int {
	return m.PayloadContentLen() + m.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (m *ObRpcResponseWarningMsg) PayloadContentLen() int {
	totalLen := util.EncodedLengthByBytes(m.msg) +
		util.EncodedLengthByVi64(m.timestamp) +
		util.EncodedLengthByVi32(m.logLevel) +
		util.EncodedLengthByVi32(m.lineNo) +
		util.EncodedLengthByVi32(m.code)

	m.ObUniVersionHeader.SetContentLength(totalLen)
	return m.ObUniVersionHeader.ContentLength()
}

func (m *ObRpcResponseWarningMsg) Encode(buffer *bytes.Buffer) {
	m.ObUniVersionHeader.Encode(buffer)

	util.EncodeBytes(buffer, m.msg)
	util.EncodeVi64(buffer, m.timestamp)
	util.EncodeVi32(buffer, m.logLevel)
	util.EncodeVi32(buffer, m.lineNo)
	util.EncodeVi32(buffer, m.code)
}

func (m *ObRpcResponseWarningMsg) Decode(buffer *bytes.Buffer) {
	m.ObUniVersionHeader.Decode(buffer)

	m.msg = util.DecodeBytes(buffer)
	m.timestamp = util.DecodeVi64(buffer)
	m.logLevel = util.DecodeVi32(buffer)
	m.lineNo = util.DecodeVi32(buffer)
	m.code = util.DecodeVi32(buffer)
}

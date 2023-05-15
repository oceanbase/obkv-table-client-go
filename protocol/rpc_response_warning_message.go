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
	*ObUniVersionHeader
	msg       []byte
	timestamp int64
	logLevel  int32
	lineNo    int32
	code      int32
}

func NewObRpcResponseWarningMsg() *ObRpcResponseWarningMsg {
	return &ObRpcResponseWarningMsg{
		ObUniVersionHeader: NewObUniVersionHeader(),
		msg:                nil,
		timestamp:          0,
		logLevel:           0,
		lineNo:             0,
		code:               0,
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

func (m *ObRpcResponseWarningMsg) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (m *ObRpcResponseWarningMsg) Decode(buffer *bytes.Buffer) {
	m.ObUniVersionHeader.Decode(buffer)

	m.msg = util.DecodeBytes(buffer)
	m.timestamp = util.DecodeVi64(buffer)
	m.logLevel = util.DecodeVi32(buffer)
	m.lineNo = util.DecodeVi32(buffer)
	m.code = util.DecodeVi32(buffer)
}

package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type RpcResponseWarningMsg struct {
	*UniVersionHeader
	msg       []byte
	timestamp int64
	logLevel  int32
	lineNo    int32
	code      int32
}

func NewRpcResponseWarningMsg() *RpcResponseWarningMsg {
	return &RpcResponseWarningMsg{
		UniVersionHeader: NewUniVersionHeader(),
		msg:              nil,
		timestamp:        0,
		logLevel:         0,
		lineNo:           0,
		code:             0,
	}
}

func (m *RpcResponseWarningMsg) Msg() []byte {
	return m.msg
}

func (m *RpcResponseWarningMsg) SetMsg(msg []byte) {
	m.msg = msg
}

func (m *RpcResponseWarningMsg) Timestamp() int64 {
	return m.timestamp
}

func (m *RpcResponseWarningMsg) SetTimestamp(timestamp int64) {
	m.timestamp = timestamp
}

func (m *RpcResponseWarningMsg) LogLevel() int32 {
	return m.logLevel
}

func (m *RpcResponseWarningMsg) SetLogLevel(logLevel int32) {
	m.logLevel = logLevel
}

func (m *RpcResponseWarningMsg) LineNo() int32 {
	return m.lineNo
}

func (m *RpcResponseWarningMsg) SetLineNo(lineNo int32) {
	m.lineNo = lineNo
}

func (m *RpcResponseWarningMsg) Code() int32 {
	return m.code
}

func (m *RpcResponseWarningMsg) SetCode(code int32) {
	m.code = code
}

func (m *RpcResponseWarningMsg) Encode() []byte {
	// TODO implement me
	panic("implement me")
}

func (m *RpcResponseWarningMsg) Decode(buffer *bytes.Buffer) {
	m.UniVersionHeader.Decode(buffer)

	m.msg = util.DecodeBytes(buffer)
	m.timestamp = util.DecodeVi64(buffer)
	m.logLevel = util.DecodeVi32(buffer)
	m.lineNo = util.DecodeVi32(buffer)
	m.code = util.DecodeVi32(buffer)
}

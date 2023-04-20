package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type RpcResponseCode struct {
	*UniVersionHeader
	code        ResponseCode
	msg         []byte
	warningMsgs []*RpcResponseWarningMsg
}

func NewRpcResponseCode() *RpcResponseCode {
	return &RpcResponseCode{
		UniVersionHeader: NewUniVersionHeader(),
		code:             0,
		msg:              nil,
		warningMsgs:      nil,
	}
}

func (c *RpcResponseCode) Code() ResponseCode {
	return c.code
}

func (c *RpcResponseCode) SetCode(code ResponseCode) {
	c.code = code
}

func (c *RpcResponseCode) Msg() []byte {
	return c.msg
}

func (c *RpcResponseCode) SetMsg(msg []byte) {
	c.msg = msg
}

func (c *RpcResponseCode) WarningMsgs() []*RpcResponseWarningMsg {
	return c.warningMsgs
}

func (c *RpcResponseCode) SetWarningMsgs(warningMsgs []*RpcResponseWarningMsg) {
	c.warningMsgs = warningMsgs
}

func (c *RpcResponseCode) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (c *RpcResponseCode) Decode(buffer *bytes.Buffer) {
	c.UniVersionHeader.Decode(buffer)

	c.code = ResponseCode(util.DecodeVi32(buffer))
	c.msg = util.DecodeBytes(buffer)

	waringMsgsLen := util.DecodeVi32(buffer)
	if waringMsgsLen > 0 {
		rpcResponseWarningMsg := NewRpcResponseWarningMsg()
		rpcResponseWarningMsg.Decode(buffer)
		c.warningMsgs = append(c.warningMsgs, rpcResponseWarningMsg)
	}
}

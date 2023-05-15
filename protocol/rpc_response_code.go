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

	oberror "github.com/oceanbase/obkv-table-client-go/error"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type RpcResponseCode struct {
	*UniVersionHeader
	code        oberror.ObErrorCode
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

func (c *RpcResponseCode) Code() oberror.ObErrorCode {
	return c.code
}

func (c *RpcResponseCode) SetCode(code oberror.ObErrorCode) {
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

	c.code = oberror.ObErrorCode(util.DecodeVi32(buffer))
	c.msg = util.DecodeBytes(buffer)

	waringMsgsLen := int(util.DecodeVi32(buffer))

	for i := 0; i < waringMsgsLen; i++ {
		rpcResponseWarningMsg := NewRpcResponseWarningMsg()
		rpcResponseWarningMsg.Decode(buffer)
		c.warningMsgs = append(c.warningMsgs, rpcResponseWarningMsg)
	}
}

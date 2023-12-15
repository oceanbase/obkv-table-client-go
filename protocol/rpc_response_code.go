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

type ObRpcResponseCode struct {
	ObUniVersionHeader
	code        oberror.ObErrorCode
	msg         []byte
	warningMsgs []*ObRpcResponseWarningMsg
}

func NewObRpcResponseCode() *ObRpcResponseCode {
	return &ObRpcResponseCode{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		code:        0,
		msg:         nil,
		warningMsgs: nil,
	}
}

func (c *ObRpcResponseCode) Code() oberror.ObErrorCode {
	return c.code
}

func (c *ObRpcResponseCode) SetCode(code oberror.ObErrorCode) {
	c.code = code
}

func (c *ObRpcResponseCode) Msg() string {
	if len(c.msg) > 0 {
		return string(c.msg[:len(c.msg)-1])
	}
	return string(c.msg)
}

func (c *ObRpcResponseCode) SetMsg(msg []byte) {
	c.msg = msg
}

func (c *ObRpcResponseCode) WarningMsgs() []*ObRpcResponseWarningMsg {
	return c.warningMsgs
}

func (c *ObRpcResponseCode) SetWarningMsgs(warningMsgs []*ObRpcResponseWarningMsg) {
	c.warningMsgs = warningMsgs
}

func (c *ObRpcResponseCode) PayloadLen() int {
	return c.PayloadContentLen() + c.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (c *ObRpcResponseCode) PayloadContentLen() int {
	totalLen := util.EncodedLengthByVi32(int32(c.code)) +
		util.EncodedLengthByBytes(c.msg)

	totalLen += util.EncodedLengthByVi32(int32(len(c.warningMsgs)))
	for _, msg := range c.warningMsgs {
		totalLen += msg.PayloadLen()
	}

	c.ObUniVersionHeader.SetContentLength(totalLen)
	return c.ObUniVersionHeader.ContentLength()
}

func (c *ObRpcResponseCode) Encode(buffer *bytes.Buffer) {
	c.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi32(buffer, int32(c.code))
	util.EncodeBytes(buffer, c.msg)

	util.EncodeVi32(buffer, int32(len(c.warningMsgs)))

	for _, warningMsg := range c.warningMsgs {
		warningMsg.Encode(buffer)
	}
}

func (c *ObRpcResponseCode) Decode(buffer *bytes.Buffer) {
	c.ObUniVersionHeader.Decode(buffer)

	c.code = oberror.ObErrorCode(util.DecodeVi32(buffer))
	c.msg = util.DecodeBytes(buffer)

	waringMsgsLen := int(util.DecodeVi32(buffer))

	c.warningMsgs = make([]*ObRpcResponseWarningMsg, 0, waringMsgsLen)

	for i := 0; i < waringMsgsLen; i++ {
		rpcResponseWarningMsg := NewObRpcResponseWarningMsg()
		rpcResponseWarningMsg.Decode(buffer)
		c.warningMsgs = append(c.warningMsgs, rpcResponseWarningMsg)
	}
}

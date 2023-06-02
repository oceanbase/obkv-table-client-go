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
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObRpcResponseWarningMsgEncodeDecode(t *testing.T) {
	obRpcResponseWarningMsg := NewObRpcResponseWarningMsg()
	obRpcResponseWarningMsg.SetVersion(1)
	obRpcResponseWarningMsg.SetContentLength(0)
	obRpcResponseWarningMsg.SetMsg([]byte(util.String(20)))
	obRpcResponseWarningMsg.SetTimestamp(int64(rand.Uint64()))
	obRpcResponseWarningMsg.SetLogLevel(int32(rand.Uint32()))
	obRpcResponseWarningMsg.SetLineNo(int32(rand.Uint32()))
	obRpcResponseWarningMsg.SetCode(int32(rand.Uint32()))

	payloadLen := obRpcResponseWarningMsg.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obRpcResponseWarningMsg.Encode(buffer)

	newObRpcResponseWarningMsg := NewObRpcResponseWarningMsg()

	newBuffer := bytes.NewBuffer(buf)
	newObRpcResponseWarningMsg.Decode(newBuffer)
	assert.EqualValues(t, obRpcResponseWarningMsg.Version(), newObRpcResponseWarningMsg.Version())
	assert.EqualValues(t, obRpcResponseWarningMsg.ContentLength(), newObRpcResponseWarningMsg.ContentLength())
	assert.EqualValues(t, obRpcResponseWarningMsg.Timestamp(), newObRpcResponseWarningMsg.Timestamp())
	assert.EqualValues(t, obRpcResponseWarningMsg.LogLevel(), newObRpcResponseWarningMsg.LogLevel())
	assert.EqualValues(t, obRpcResponseWarningMsg.LineNo(), newObRpcResponseWarningMsg.LineNo())
	assert.EqualValues(t, obRpcResponseWarningMsg.Code(), newObRpcResponseWarningMsg.Code())
	assert.EqualValues(t, obRpcResponseWarningMsg, newObRpcResponseWarningMsg)
}

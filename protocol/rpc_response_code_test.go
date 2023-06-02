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

	oberror "github.com/oceanbase/obkv-table-client-go/error"
	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObRpcResponseCodeEncodeDecode(t *testing.T) {
	obRpcResponseCode := NewObRpcResponseCode()
	obRpcResponseCode.SetVersion(1)
	obRpcResponseCode.SetContentLength(0)
	obRpcResponseCode.SetCode(oberror.ObErrorCode(rand.Uint32()))
	obRpcResponseCode.SetMsg([]byte(util.String(20)))

	randomLen := rand.Intn(20)
	obRpcResponseWarningMsgs := make([]*ObRpcResponseWarningMsg, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		obRpcResponseWarningMsg := NewObRpcResponseWarningMsg()
		obRpcResponseWarningMsg.SetVersion(1)
		obRpcResponseWarningMsg.SetContentLength(0)
		obRpcResponseWarningMsg.SetMsg([]byte(util.String(20)))
		obRpcResponseWarningMsg.SetTimestamp(int64(rand.Uint64()))
		obRpcResponseWarningMsg.SetLogLevel(int32(rand.Uint32()))
		obRpcResponseWarningMsg.SetLineNo(int32(rand.Uint32()))
		obRpcResponseWarningMsg.SetCode(int32(rand.Uint32()))
		obRpcResponseWarningMsgs = append(obRpcResponseWarningMsgs, obRpcResponseWarningMsg)
	}
	obRpcResponseCode.SetWarningMsgs(obRpcResponseWarningMsgs)

	payloadLen := obRpcResponseCode.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obRpcResponseCode.Encode(buffer)

	newObRpcResponseCode := NewObRpcResponseCode()

	newBuffer := bytes.NewBuffer(buf)
	newObRpcResponseCode.Decode(newBuffer)

	assert.EqualValues(t, obRpcResponseCode.Version(), newObRpcResponseCode.Version())
	assert.EqualValues(t, obRpcResponseCode.ContentLength(), newObRpcResponseCode.ContentLength())
	assert.EqualValues(t, obRpcResponseCode.Code(), newObRpcResponseCode.Code())
	assert.EqualValues(t, obRpcResponseCode.Msg(), newObRpcResponseCode.Msg())
	assert.EqualValues(t, obRpcResponseCode.WarningMsgs(), newObRpcResponseCode.WarningMsgs())
	assert.EqualValues(t, obRpcResponseCode, newObRpcResponseCode)
}

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

func TestObLoginRequestEncodeDecode(t *testing.T) {
	obLoginRequest := NewObLoginRequest(util.String(10), util.String(10), util.String(10), util.String(10))
	obLoginRequest.SetAuthMethod(uint8(rand.Intn(100)))
	obLoginRequest.SetClientType(uint8(rand.Intn(100)))
	obLoginRequest.SetClientVersion(uint8(rand.Intn(100)))
	obLoginRequest.SetReversed1(uint8(rand.Intn(100)))
	obLoginRequest.SetClientCapabilities(int32(rand.Intn(100)))
	obLoginRequest.SetMaxPacketSize(int32(rand.Intn(100)))
	obLoginRequest.SetReversed2(int32(rand.Intn(100)))
	obLoginRequest.SetReversed3(int64(rand.Intn(100)))
	obLoginRequest.SetTenantName(util.String(10))
	obLoginRequest.SetUserName(util.String(10))
	obLoginRequest.SetPassSecret(util.String(10))
	obLoginRequest.SetPassScramble(util.String(10))
	obLoginRequest.SetDatabaseName(util.String(10))
	obLoginRequest.SetTtlUs(int64(rand.Intn(10)))

	payloadLen := obLoginRequest.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obLoginRequest.Encode(buffer)

	newObLoginRequest := NewObLoginRequest(util.String(10), util.String(10), util.String(10), util.String(10))
	newBuffer := bytes.NewBuffer(buf)
	newObLoginRequest.Decode(newBuffer)

	assert.EqualValues(t, obLoginRequest.AuthMethod(), newObLoginRequest.AuthMethod())
	assert.EqualValues(t, obLoginRequest.ClientType(), newObLoginRequest.ClientType())
	assert.EqualValues(t, obLoginRequest.ClientVersion(), newObLoginRequest.ClientVersion())
	assert.EqualValues(t, obLoginRequest.Reversed1(), newObLoginRequest.Reversed1())
	assert.EqualValues(t, obLoginRequest.ClientCapabilities(), newObLoginRequest.ClientCapabilities())
	assert.EqualValues(t, obLoginRequest.MaxPacketSize(), newObLoginRequest.MaxPacketSize())
	assert.EqualValues(t, obLoginRequest.Reversed2(), newObLoginRequest.Reversed2())
	assert.EqualValues(t, obLoginRequest.Reversed3(), newObLoginRequest.Reversed3())
	assert.EqualValues(t, obLoginRequest.TenantName(), newObLoginRequest.TenantName())
	assert.EqualValues(t, obLoginRequest.UserName(), newObLoginRequest.UserName())
	assert.EqualValues(t, obLoginRequest.PassSecret(), newObLoginRequest.PassSecret())
	assert.EqualValues(t, obLoginRequest.PassScramble(), newObLoginRequest.PassScramble())
	assert.EqualValues(t, obLoginRequest.DatabaseName(), newObLoginRequest.DatabaseName())
	assert.EqualValues(t, obLoginRequest.TtlUs(), newObLoginRequest.TtlUs())
	assert.EqualValues(t, obLoginRequest, newObLoginRequest)
}

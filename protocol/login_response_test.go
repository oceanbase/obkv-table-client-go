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

func TestObLoginResponseEncodeDecode(t *testing.T) {
	obLoginResponse := NewObLoginResponse()
	obLoginResponse.SetServerCapabilities(int32(rand.Intn(100)))
	obLoginResponse.SetReserved1(int32(rand.Intn(100)))
	obLoginResponse.SetReserved2(int64(rand.Intn(100)))
	obLoginResponse.SetServerVersion(util.String(10))
	obLoginResponse.SetCredential([]byte(util.String(10)))
	obLoginResponse.SetUserId(int64(rand.Intn(100)))
	obLoginResponse.SetDatabaseId(int64(rand.Intn(100)))

	payloadLen := obLoginResponse.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obLoginResponse.Encode(buffer)

	newObLoginResponse := NewObLoginResponse()
	newBuffer := bytes.NewBuffer(buf)
	newObLoginResponse.Decode(newBuffer)

	assert.EqualValues(t, obLoginResponse.ServerCapabilities(), newObLoginResponse.ServerCapabilities())
	assert.EqualValues(t, obLoginResponse.Reserved1(), newObLoginResponse.Reserved1())
	assert.EqualValues(t, obLoginResponse.Reserved2(), newObLoginResponse.Reserved2())
	assert.EqualValues(t, obLoginResponse.ServerVersion(), newObLoginResponse.ServerVersion())
	assert.EqualValues(t, obLoginResponse.Credential(), newObLoginResponse.Credential())
	assert.EqualValues(t, obLoginResponse.UserId(), newObLoginResponse.UserId())
	assert.EqualValues(t, obLoginResponse.DatabaseId(), newObLoginResponse.DatabaseId())
	assert.EqualValues(t, obLoginResponse, obLoginResponse)
}

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

func TestObTableResponseEncodeDecode(t *testing.T) {
	obTableResponse := NewObTableResponse()
	obTableResponse.SetVersion(1)
	obTableResponse.SetContentLength(0)
	obTableResponse.SetErrorNo(int32(rand.Uint32()))
	obTableResponse.SetSqlState([]byte(util.String(20)))
	obTableResponse.SetMsg([]byte(util.String(20)))

	payloadLen := obTableResponse.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableResponse.Encode(buffer)

	newObTableResponse := NewObTableResponse()

	newBuffer := bytes.NewBuffer(buf)
	newObTableResponse.Decode(newBuffer)

	assert.EqualValues(t, obTableResponse.Version(), newObTableResponse.Version())
	assert.EqualValues(t, obTableResponse.ContentLength(), newObTableResponse.ContentLength())
	assert.EqualValues(t, obTableResponse.ErrorNo(), newObTableResponse.ErrorNo())
	assert.EqualValues(t, obTableResponse.SqlState(), newObTableResponse.SqlState())
	assert.EqualValues(t, obTableResponse.Msg(), newObTableResponse.Msg())
	assert.EqualValues(t, obTableResponse, newObTableResponse)
}

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

package util

import (
	"bytes"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceToString(t *testing.T) {
	assert.Equal(t, "true", InterfaceToString(true))
	assert.Equal(t, "false", InterfaceToString(false))
	assert.Equal(t, "1", InterfaceToString(int(1)))
	assert.Equal(t, "1", InterfaceToString(int8(1)))
	assert.Equal(t, "1", InterfaceToString(int16(1)))
	assert.Equal(t, "1", InterfaceToString(int32(1)))
	assert.Equal(t, "1", InterfaceToString(int64(1)))
	assert.Equal(t, "1", InterfaceToString(uint(1)))
	assert.Equal(t, "1", InterfaceToString(uint8(1)))
	assert.Equal(t, "1", InterfaceToString(uint16(1)))
	assert.Equal(t, "1", InterfaceToString(uint32(1)))
	assert.Equal(t, "1", InterfaceToString(uint32(1)))
	assert.Equal(t, "1", InterfaceToString(uint64(1)))
	assert.Equal(t, "1.1", InterfaceToString(float32(1.1)))
	assert.Equal(t, "1.1", InterfaceToString(1.1))
	assert.Equal(t, "abc", InterfaceToString("abc"))

	type myTestStruct struct {
		str string
	}
	v := myTestStruct{"abc"}
	assert.Equal(t, InterfaceToString(v), "{abc}")
}

func TestStringArrayToString(t *testing.T) {
	strArr := []string{"hello", "test", "world"}
	assert.Equal(t, StringArrayToString(strArr), "[hello, test, world]")
	assert.Equal(t, StringArrayToString([]string{}), "[]")
}

func TestInterfacesToString(t *testing.T) {
	s := []interface{}{1, 2, 3, 4, "abc", 1.1}
	assert.Equal(t, "[1, 2, 3, 4, abc, 1.1]", InterfacesToString(s))
}

func TestSkipBytes(t *testing.T) {
	buf := &bytes.Buffer{}
	SkipBytes(buf, 0)
	buf = &bytes.Buffer{}
	SkipBytes(buf, 1)
}

func TestStringToBytes(t *testing.T) {
	assert.EqualValues(t, []byte(nil), StringToBytes(""))
	b := StringToBytes("abc")
	assert.EqualValues(t, []byte{0x61, 0x62, 0x63}, b)
}

func TestBytesToString(t *testing.T) {
	assert.Equal(t, "", BytesToString([]byte{}))
	assert.Equal(t, "abc", BytesToString([]byte{0x61, 0x62, 0x63}))
}

func TestBoolToByte(t *testing.T) {
	assert.EqualValues(t, 1, BoolToByte(true))
	assert.EqualValues(t, 0, BoolToByte(false))
}

func TestByteToBool(t *testing.T) {
	assert.EqualValues(t, true, ByteToBool(1))
	assert.EqualValues(t, false, ByteToBool(0))
}

func TestConvertIpToUint32(t *testing.T) {
	assert.EqualValues(t, uint32(0x7f000001), ConvertIpToUint32(net.IP{127, 0, 0, 1}))
	assert.EqualValues(t, 0, ConvertIpToUint32(net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func TestConvertUint32ToIp(t *testing.T) {
	assert.EqualValues(t, net.IP{127, 0, 0, 1}, ConvertUint32ToIp(uint32(0x7f000001)))
	assert.EqualValues(t, net.IP{0, 0, 0, 0}, ConvertUint32ToIp(0))
}

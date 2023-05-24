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
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceToString(t *testing.T) {
	assert.Equal(t, InterfaceToString("string"), "string")
	assert.Equal(t, InterfaceToString(int(1)), "1")
	assert.Equal(t, InterfaceToString(true), "true")
	assert.Equal(t, InterfaceToString(false), "false")
	assert.Equal(t, InterfaceToString(float32(3.14)), "3.14")
	assert.Equal(t, InterfaceToString(3.14), "3.14")
	assert.Equal(t, InterfaceToString(complex64(3.14)), "(3.140000+0.000000i)")
	assert.Equal(t, InterfaceToString(complex128(3.14)), "(3.140000+0.000000i)")
	assert.Equal(t, InterfaceToString(byte(1)), "\x01")
	assert.Equal(t, InterfaceToString(errors.New("error")), "error")
	assert.Equal(t, InterfaceToString(nil), "<nil>")
	type myTestStruct struct {
	}
	v := myTestStruct{}
	assert.Equal(t, InterfaceToString(v), "{}")
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

func TestConvertIpToUint32(t *testing.T) {
	assert.EqualValues(t, uint32(0x7f000001), ConvertIpToUint32(net.IP{127, 0, 0, 1}))
	assert.EqualValues(t, 0, ConvertIpToUint32(net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

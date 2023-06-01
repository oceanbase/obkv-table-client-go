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
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"unsafe"
)

func InterfaceToString(i interface{}) string {
	switch t := i.(type) {
	case bool:
		return strconv.FormatBool(t)
	case int:
		return strconv.Itoa(t)
	case int8:
		return strconv.FormatInt(int64(t), 10)
	case int16:
		return strconv.FormatInt(int64(t), 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case int64:
		return strconv.FormatInt(t, 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint8:
		return strconv.FormatUint(uint64(t), 10)
	case uint16:
		return strconv.FormatUint(uint64(t), 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		return t
	default:
		return fmt.Sprintf("%v", t)
	}
}

func InterfacesToString(values []interface{}) string {
	var str string
	str = str + "["
	for i := 0; i < len(values); i++ {
		if i > 0 {
			str += ", "
		}
		str += InterfaceToString(values[i])
	}
	str += "]"
	return str
}

func StringArrayToString(strArr []string) string {
	var str string
	str = str + "["
	for i := 0; i < len(strArr); i++ {
		if i > 0 {
			str += ", "
		}
		str += strArr[i]
	}
	str += "]"
	return str
}

func SkipBytes(buffer *bytes.Buffer, skipLen int) {
	if skipLen > 0 {
		buffer.Next(skipLen)
	}
}

func StringToBytes(str string) []byte {
	if str == "" {
		return nil
	}
	return *(*[]byte)(unsafe.Pointer(&str))
}

func BytesToString(bys []byte) string {
	if len(bys) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&bys))
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func ByteToBool(b byte) bool {
	if b == 0 {
		return false
	}
	return true
}

func ConvertIpToUint32(ip net.IP) uint32 {
	if len(ip) == net.IPv6len {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func ConvertUint32ToIp(num uint32) net.IP {
	ip := make(net.IP, net.IPv4len)
	binary.BigEndian.PutUint32(ip, num)
	return ip
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

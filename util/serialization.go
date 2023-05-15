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
	"math"
)

const (
	obMaxV1b uint64 = 1<<7 - 1
	obMaxV2b uint64 = 1<<14 - 1
	obMaxV3b uint64 = 1<<21 - 1
	obMaxV4b uint64 = 1<<28 - 1
	obMaxV5b uint64 = 1<<35 - 1
	obMaxV6b uint64 = 1<<42 - 1
	obMaxV7b uint64 = 1<<49 - 1
	obMaxV8b uint64 = 1<<56 - 1
	obMaxV9b uint64 = 1<<63 - 1

	end byte = 0
)

var obMax = []uint64{
	obMaxV1b,
	obMaxV2b,
	obMaxV3b,
	obMaxV4b,
	obMaxV5b,
	obMaxV6b,
	obMaxV7b,
	obMaxV8b,
	obMaxV9b,
}

func EncodedLengthByVi32(value int32) (encodedLength int) {
	if value < 0 {
		return 5
	}
	var unsignedValue = uint64(value)
	for _, max := range obMax {
		encodedLength++
		if unsignedValue <= max {
			break
		}
	}
	return encodedLength
}

func EncodedLengthByVi64(value int64) (encodedLength int) {
	if value < 0 {
		return 10
	}
	var unsignedValue = uint64(value)
	for _, max := range obMax {
		encodedLength++
		if unsignedValue <= max {
			break
		}
	}
	return encodedLength
}

func EncodedLengthByVf32(value float32) (encodedLength int) {
	return EncodedLengthByVi32(int32(math.Float32bits(value)))
}

func EncodedLengthByVf64(value float64) (encodedLength int) {
	return EncodedLengthByVi64(int64(math.Float64bits(value)))
}

func EncodedLengthByVString(str string) (encodedLength int) {
	return EncodedLengthByVi64(int64(len(str))) + len(str) + 1
}

func EncodedLengthByBytesString(bys []byte) (encodedLength int) {
	return EncodedLengthByVi64(int64(len(bys))) + len(bys) + 1
}

func EncodedLengthByBytes(bys []byte) (encodedLength int) {
	return EncodedLengthByVi64(int64(len(bys))) + len(bys)
}

func EncodeVi32(buffer *bytes.Buffer, value int32) {
	var buf = buffer.Next(EncodedLengthByVi32(value))
	var unsignedValue = uint32(value)
	var index = 0
	for uint64(unsignedValue) > obMaxV1b {
		buf[index] = byte(unsignedValue | 0x80)
		index++
		unsignedValue = unsignedValue >> 7
	}
	buf[index] = byte(unsignedValue & 0x7f)
}

func EncodeVi64(buffer *bytes.Buffer, value int64) {
	var buf = buffer.Next(EncodedLengthByVi64(value))
	var unsignedValue = uint64(value)
	var index = 0
	for unsignedValue > obMaxV1b {
		buf[index] = byte(unsignedValue | 0x80)
		index++
		unsignedValue = unsignedValue >> 7
	}
	buf[index] = byte(unsignedValue & 0x7f)
}

func EncodeVf32(buffer *bytes.Buffer, value float32) {
	EncodeVi32(buffer, int32(math.Float32bits(value)))
}

func EncodeVf64(buffer *bytes.Buffer, value float64) {
	EncodeVi64(buffer, int64(math.Float64bits(value)))
}

func EncodeVString(buffer *bytes.Buffer, str string) {
	// encode string header
	var strLen = len(str)
	EncodeVi64(buffer, int64(strLen))
	// copy string body
	var buf = buffer.Next(strLen + 1)
	var strBytes = StringToBytes(str)
	copy(buf[:strLen], strBytes)
	// copy string end
	buf[strLen] = end
}

func EncodeBytesString(buffer *bytes.Buffer, bys []byte) {
	// encode bytes header
	var bytesLen = len(bys)
	EncodeVi64(buffer, int64(bytesLen))
	// copy bytes body
	var buf = buffer.Next(bytesLen + 1)
	copy(buf[:bytesLen], bys)
	// copy bytes end
	buf[bytesLen] = end
}

func EncodeBytes(buffer *bytes.Buffer, bys []byte) {
	// encode bytes header
	var bytesLen = len(bys)
	EncodeVi64(buffer, int64(bytesLen))
	// copy bytes body
	var buf = buffer.Next(bytesLen)
	copy(buf[:bytesLen], bys)
}

func DecodeVi32(buffer *bytes.Buffer) int32 {
	var ret uint32 = 0
	var shift uint32 = 0
	var index = 0
	buf := buffer.Bytes()
	for index = range buf {
		ret |= (uint32(buf[index]) & 0x7f) << shift
		shift += 7
		if buf[index]&0x80 == 0 {
			break
		}
	}
	buffer.Next(index + 1)
	return int32(ret)
}

func DecodeVi64(buffer *bytes.Buffer) int64 {
	var ret uint64 = 0
	var shift uint64 = 0
	var index = 0
	buf := buffer.Bytes()
	for index = range buf {
		ret |= (uint64(buf[index]) & 0x7f) << shift
		shift += 7
		if buf[index]&0x80 == 0 {
			break
		}
	}
	buffer.Next(index + 1)
	return int64(ret)
}

func DecodeVf32(buffer *bytes.Buffer) float32 {
	return math.Float32frombits(uint32(DecodeVi32(buffer)))
}

func DecodeVf64(buffer *bytes.Buffer) float64 {
	return math.Float64frombits(uint64(DecodeVi64(buffer)))
}

func DecodeVString(buffer *bytes.Buffer) string {
	// decode string header
	strLen := DecodeVi32(buffer)
	// copy str body
	strBodyBuf := make([]byte, strLen)
	copy(strBodyBuf, buffer.Next(int(strLen)))
	// skip the end byte '0'
	SkipBytes(buffer, 1)
	return BytesToString(strBodyBuf)
}

func DecodeBytesString(buffer *bytes.Buffer) []byte {
	// decode bytes header
	bytesLen := DecodeVi32(buffer)
	// copy bytes body
	bytesBodyBuf := make([]byte, bytesLen)
	copy(bytesBodyBuf, buffer.Next(int(bytesLen)))
	// skip the end byte '0'
	SkipBytes(buffer, 1)
	return bytesBodyBuf
}

func DecodeBytes(buffer *bytes.Buffer) []byte {
	// decode bytes header
	bytesLen := DecodeVi32(buffer)
	bytesBodyBuf := make([]byte, bytesLen)
	// copy bytes body
	copy(bytesBodyBuf, buffer.Next(int(bytesLen)))
	return bytesBodyBuf
}

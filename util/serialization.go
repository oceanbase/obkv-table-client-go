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

func NeedLengthByVi32(value int32) (needLength int) {
	if value < 0 {
		return 5
	}
	var unsignedValue = uint64(value)
	for _, max := range obMax {
		needLength++
		if unsignedValue <= max {
			break
		}
	}
	return needLength
}

func NeedLengthByVi64(value int64) (needLength int) {
	if value < 0 {
		return 10
	}
	var unsignedValue = uint64(value)
	for _, max := range obMax {
		needLength++
		if unsignedValue <= max {
			break
		}
	}
	return needLength
}

func NeedLengthByVf32(value float32) (needLength int) {
	return NeedLengthByVi32(int32(math.Float32bits(value)))
}

func NeedLengthByVf64(value float64) (needLength int) {
	return NeedLengthByVi64(int64(math.Float64bits(value)))
}

func NeedLengthByVString(str string) (needLength int) {
	return NeedLengthByVi64(int64(len(str))) + len(str) + 1
}

func NeedLengthByBytesString(bys []byte) (needLength int) {
	return NeedLengthByVi64(int64(len(bys))) + len(bys) + 1
}

func NeedLengthByBytes(bys []byte) (needLength int) {
	return NeedLengthByVi64(int64(len(bys))) + len(bys)
}

func EncodeVi32(buf []byte, value int32) {
	var unsignedValue = uint32(value)
	var index = 0
	for uint64(unsignedValue) > obMaxV1b {
		buf[index] = byte(unsignedValue | 0x80)
		index++
		unsignedValue = unsignedValue >> 7
	}
	buf[index] = byte(unsignedValue & 0x7f)
}

func EncodeVi64(buf []byte, value int64) {
	var unsignedValue = uint64(value)
	var index = 0
	for unsignedValue > obMaxV1b {
		buf[index] = byte(unsignedValue | 0x80)
		index++
		unsignedValue = unsignedValue >> 7
	}
	buf[index] = byte(unsignedValue & 0x7f)
}

func EncodeVf32(buf []byte, value float32) {
	EncodeVi32(buf, int32(math.Float32bits(value)))
}

func EncodeVf64(buf []byte, value float64) {
	EncodeVi64(buf, int64(math.Float64bits(value)))
}

func EncodeVString(buf []byte, str string) {
	// encode string header
	var strLen = len(str)
	var strHeadNeedLength = NeedLengthByVi64(int64(strLen))
	EncodeVi64(buf[:strHeadNeedLength], int64(strLen))
	// copy string body
	var endIndex = len(buf) - 1
	var strBytes = StringToBytes(str)
	copy(buf[strHeadNeedLength:endIndex], strBytes)
	// copy string end
	buf[endIndex] = end
}

func EncodeBytesString(buf []byte, bys []byte) {
	// encode bytes header
	var bytesLen = len(bys)
	var bytesHeadNeedLength = NeedLengthByVi64(int64(bytesLen))
	EncodeVi64(buf[:bytesHeadNeedLength], int64(bytesLen))
	// copy bytes body
	var endIndex = len(buf) - 1
	copy(buf[bytesHeadNeedLength:endIndex], bys)
	// copy bytes end
	buf[endIndex] = end
}

func EncodeBytes(buf []byte, bys []byte) {
	// encode bytes header
	var bytesLen = len(bys)
	var bytesHeadNeedLength = NeedLengthByVi64(int64(bytesLen))
	EncodeVi64(buf[:bytesHeadNeedLength], int64(bytesLen))
	// copy bytes body
	copy(buf[bytesHeadNeedLength:], bys)
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

package util

import (
	"bytes"
	"encoding/binary"
)

var endian = binary.BigEndian

func PutUint8(buffer *bytes.Buffer, v uint8) {
	buffer.Next(1)[0] = v
}

func PutUint16(buffer *bytes.Buffer, v uint16) {
	endian.PutUint16(buffer.Next(2), v)
}

func PutUint32(buffer *bytes.Buffer, v uint32) {
	endian.PutUint32(buffer.Next(4), v)
}

func PutUint64(buffer *bytes.Buffer, v uint64) {
	endian.PutUint64(buffer.Next(8), v)
}

func Uint8(buffer *bytes.Buffer) uint8 {
	return buffer.Next(1)[0]
}

func Uint16(buffer *bytes.Buffer) uint16 {
	return endian.Uint16(buffer.Next(2))
}

func Uint32(buffer *bytes.Buffer) uint32 {
	return endian.Uint32(buffer.Next(4))
}

func Uint64(buffer *bytes.Buffer) uint64 {
	return endian.Uint64(buffer.Next(8))
}

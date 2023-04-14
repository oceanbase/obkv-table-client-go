package util

import (
	"bytes"
	"unsafe"
)

func SkipBytes(buffer *bytes.Buffer, skipLen int) {
	if skipLen > 0 {
		buffer.Next(skipLen)
	}
}

func StringToBytes(str string) []byte {
	if str == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func BytesToString(bys []byte) string {
	if len(bys) == 0 {
		return ""
	}
	// unsafeString converts a []byte to a string with no allocation.
	// The caller must not modify b while the result string is in use.
	return unsafe.String(unsafe.SliceData(bys), len(bys))
}

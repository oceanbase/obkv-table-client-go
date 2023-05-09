package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"unsafe"
)

func InterfaceToString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case complex64:
		return fmt.Sprintf("(%f+%fi)", real(v), imag(v))
	case complex128:
		return fmt.Sprintf("(%f+%fi)", real(v), imag(v))
	case byte:
		return string(v)
	case rune:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	case uintptr, unsafe.Pointer:
		return fmt.Sprintf("%p", v)
	default:
		return fmt.Sprintf("%v", v)
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

package util

func Uint8(b []byte) uint8 {
	_ = b[0]
	return b[0]
}

func PutUint8(b []byte, v uint8) {
	_ = b[0]
	b[0] = v
}

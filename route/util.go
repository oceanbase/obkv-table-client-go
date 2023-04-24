package route

import (
	"strconv"
)

func CreateInStatement(values []int) string {
	// Create inStatement "(0,1,2...partNum);".
	var inStatement string
	inStatement += "("
	for i, v := range values {
		if i > 0 {
			inStatement += ", "
		}
		inStatement += strconv.Itoa(v)
	}
	inStatement += ");"
	return inStatement
}

func MurmurHash64A(data []byte, length int, seed int64) int64 {
	const (
		m = 0xc6a4a7935bd1e995
		r = 47
	)
	var k int64
	h := seed ^ int64(uint64(length)*m)

	var ubigm uint64 = m
	var ibigm = int64(ubigm)
	for l := length; l >= 8; l -= 8 {
		k = int64(int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56)

		k := k * ibigm
		k ^= int64(uint64(k) >> r)
		k = k * ibigm

		h = h ^ k
		h = h * ibigm
		data = data[8:]
	}

	switch length & 7 {
	case 7:
		h ^= int64(data[6]) << 48
		fallthrough
	case 6:
		h ^= int64(data[5]) << 40
		fallthrough
	case 5:
		h ^= int64(data[4]) << 32
		fallthrough
	case 4:
		h ^= int64(data[3]) << 24
		fallthrough
	case 3:
		h ^= int64(data[2]) << 16
		fallthrough
	case 2:
		h ^= int64(data[1]) << 8
		fallthrough
	case 1:
		h ^= int64(data[0])
		h *= ibigm
	}

	h ^= int64(uint64(h) >> r)
	h *= ibigm
	h ^= int64(uint64(h) >> r)
	return h
}

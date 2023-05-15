package route

import (
	"errors"
	"strconv"
)

func createInStatement(values []int) string {
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

func murmurHash64A(data []byte, length int, seed int64) int64 {
	const (
		m = 0xc6a4a7935bd1e995
		r = 47
	)
	var k int64
	h := seed ^ int64(uint64(length)*m)

	var um uint64 = m
	var im = int64(um)
	for l := length; l >= 8; l -= 8 {
		k = int64(int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56)

		k := k * im
		k ^= int64(uint64(k) >> r)
		k = k * im

		h = h ^ k
		h = h * im
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
		h *= im
	}

	h ^= int64(uint64(h) >> r)
	h *= im
	h ^= int64(uint64(h) >> r)
	return h
}

func parseToNumber(value interface{}) (interface{}, error) {
	if v, ok := value.(string); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		return int64(i), nil
	} else if v, ok := value.(int); ok { // todo  not support int
		return int64(v), nil
	} else if v, ok := value.(int64); ok {
		return v, nil
	} else if v, ok := value.(int32); ok {
		return int64(v), nil
	} else if v, ok := value.(int16); ok {
		return int64(v), nil
	} else if v, ok := value.(int8); ok {
		return int64(v), nil
	} else {
		return nil, errors.New("invalid type to convert to number")
	}
}

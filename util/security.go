package util

import (
	"crypto/sha1"
	"time"
)

var bys = []byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o',
	'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'z', 'x', 'c', 'v', 'b', 'n', 'm', 'Q', 'W',
	'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'Z', 'X',
	'C', 'V', 'B', 'N', 'M',
}

const (
	multiplier     int64 = 0x5DEECE66D
	addend         int64 = 0xB
	mask           int64 = (1 << 48) - 1
	integerMask    int64 = (1 << 33) - 1
	seedUniquifier int64 = 8682522807148012
)

var seed = ((seedUniquifier + time.Now().UnixNano()) ^ multiplier) & mask

func GetPasswordScramble(passwordScrambleLen int) string {
	scrambleBys := make([]byte, passwordScrambleLen)
	for i := range scrambleBys {
		scrambleBys[i] = randomByte()
	}

	return BytesToString(scrambleBys)
}

func ScramblePassword(passWord string, seed string) string {
	if passWord == "" {
		return ""
	}
	hash := sha1.New()

	hash.Write(StringToBytes(passWord))
	pass1 := hash.Sum(nil)
	hash.Reset()

	hash.Write(pass1)
	pass2 := hash.Sum(nil)
	hash.Reset()

	hash.Write(StringToBytes(seed))
	hash.Write(pass2)
	pass3 := hash.Sum(nil)
	for i := range pass3 {
		pass3[i] = pass3[i] ^ pass1[i]
	}

	return BytesToString(pass3)
}

func randomByte() byte {
	var ran = int((random() & integerMask) >> 16)
	return bys[ran%len(bys)]
}

func random() int64 {
	var oldSeed = seed
	var nextSeed int64 = 0
	for {
		nextSeed = (oldSeed*multiplier + addend) & mask
		if oldSeed != nextSeed {
			break
		}
	}
	seed = nextSeed
	return nextSeed
}

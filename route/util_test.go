package route

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtil_createInStatement(t *testing.T) {
	inStr := createInStatement(nil)
	assert.Equal(t, "();", inStr)
	inStr = createInStatement([]int{1})
	assert.Equal(t, "(1);", inStr)
	inStr = createInStatement([]int{1, 2})
	assert.Equal(t, "(1, 2);", inStr)
}

func TestUtil_murmurHash64A(t *testing.T) {
	result := murmurHash64A([]byte{1}, len([]byte{1}), int64(0))
	assert.Equal(t, int64(-5720937396023583481), result)
	result = murmurHash64A([]byte{1}, len([]byte{1}), int64(1))
	assert.Equal(t, int64(6351753276682545529), result)
	result = murmurHash64A([]byte{1, 2, 3}, len([]byte{1, 2, 3}), int64(123456789))
	assert.Equal(t, int64(-4356950700900923028), result)
}

package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceToString(t *testing.T) {
	assert.Equal(t, InterfaceToString("string"), "string")
	assert.Equal(t, InterfaceToString(int(1)), "1")
	assert.Equal(t, InterfaceToString(true), "true")
	assert.Equal(t, InterfaceToString(false), "false")
	assert.Equal(t, InterfaceToString(float32(3.14)), "3.14")
	assert.Equal(t, InterfaceToString(3.14), "3.14")
	assert.Equal(t, InterfaceToString(complex64(3.14)), "(3.140000+0.000000i)")
	assert.Equal(t, InterfaceToString(complex128(3.14)), "(3.140000+0.000000i)")
	assert.Equal(t, InterfaceToString(byte(1)), "\x01")
	assert.Equal(t, InterfaceToString(errors.New("error")), "error")
	assert.Equal(t, InterfaceToString(nil), "<nil>")
	type myTestStruct struct {
	}
	v := myTestStruct{}
	assert.Equal(t, InterfaceToString(v), "{0}")
}

func TestStringArrayToString(t *testing.T) {
	strArr := []string{"hello", "test", "world"}
	assert.Equal(t, StringArrayToString(strArr), "[hello, test, world]")
	assert.Equal(t, StringArrayToString([]string{}), "[]")
}

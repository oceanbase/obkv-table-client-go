package protocol

import "github.com/oceanbase/obkv-table-client-go/util"

type ObBytesString struct {
	bytesVal []byte
}

func (v *ObBytesString) BytesVal() []byte {
	return v.bytesVal
}

func (v *ObBytesString) String() string {
	return "ObBytesString{" +
		"bytesVal:" + string(v.bytesVal) +
		"}"
}

func newObBytesString(bytesVal []byte) *ObBytesString {
	return &ObBytesString{bytesVal: bytesVal}
}

func newObBytesStringFromString(str string) *ObBytesString {
	return &ObBytesString{util.StringToBytes(str)}
}

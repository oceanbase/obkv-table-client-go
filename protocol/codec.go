package protocol

import (
	"bytes"
)

type ProtoEncoder interface {
	Encode() []byte
}

type ProtoDecoder interface {
	Decode(*bytes.Buffer)
}

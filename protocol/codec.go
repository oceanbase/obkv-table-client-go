package protocol

import (
	"bytes"
)

type ProtoEncoder interface {
	Encode(*bytes.Buffer)
}

type ProtoDecoder interface {
	Decode(*bytes.Buffer)
}

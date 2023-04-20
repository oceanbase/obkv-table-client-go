package protocol

import (
	"bytes"
)

type Object struct {
	meta  *ObjectMeta
	value interface{}
}

func (o *Object) Meta() *ObjectMeta {
	return o.meta
}

func (o *Object) SetMeta(meta *ObjectMeta) {
	o.meta = meta
}

func (o *Object) Value() interface{} {
	return o.value
}

func (o *Object) SetValue(value interface{}) {
	o.value = value
}

func (o *Object) Encode(buffer *bytes.Buffer) {
	o.meta.Encode(buffer)
	o.meta.ObjType().Encode(buffer, o.value)
}

func (o *Object) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (o *Object) EncodedLength() int {
	return o.meta.EncodedLength() + o.meta.ObjType().EncodedLength(o.value)
}

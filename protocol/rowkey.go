package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type RowKey struct {
	keys []*Object
}

func NewRowKey() *RowKey {
	return &RowKey{
		keys: nil,
	}
}

func (k *RowKey) Keys() []*Object {
	return k.keys
}

func (k *RowKey) SetKeys(keys []*Object) {
	k.keys = keys
}

func (k *RowKey) AppendKey(key *Object) {
	k.keys = append(k.keys, key)
}

func (k *RowKey) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, int64(len(k.keys)))

	for _, key := range k.keys {
		key.Encode(buffer)
	}
}

func (k *RowKey) Decode(buffer *bytes.Buffer) {
	keysLen := util.DecodeVi64(buffer)

	var i int64
	for i = 0; i < keysLen; i++ {
		key := NewObject()
		key.Decode(buffer)
		k.keys = append(k.keys, key)
	}
}

func (k *RowKey) EncodedLength() int {
	totalLen := util.EncodedLengthByVi64(int64(len(k.keys)))

	for _, key := range k.keys {
		totalLen += key.EncodedLength()
	}

	return totalLen
}

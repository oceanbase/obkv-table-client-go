package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type RowKey struct {
	keys []*Object
}

func (k *RowKey) Keys() []*Object {
	return k.keys
}

func (k *RowKey) SetKeys(keys []*Object) {
	k.keys = keys
}

func (k *RowKey) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, int64(len(k.keys)))

	for _, key := range k.keys {
		key.Encode(buffer)
	}
}

func (k *RowKey) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (k *RowKey) EncodedLength() int {
	totalLen := util.EncodedLengthByVi64(int64(len(k.keys)))

	for _, key := range k.keys {
		totalLen += key.EncodedLength()
	}

	return totalLen
}

package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableEntity struct {
	*UniVersionHeader
	rowKey     *RowKey
	properties map[string]*Object
}

func (e *TableEntity) RowKey() *RowKey {
	return e.rowKey
}

func (e *TableEntity) SetRowKey(rowKey *RowKey) {
	e.rowKey = rowKey
}

func (e *TableEntity) Properties() map[string]*Object {
	return e.properties
}

func (e *TableEntity) SetProperties(properties map[string]*Object) {
	e.properties = properties
}

func (e *TableEntity) PayloadLen() int {
	return e.PayloadContentLen() + e.UniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (e *TableEntity) PayloadContentLen() int {
	totalLen := e.rowKey.EncodedLength()

	totalLen += util.EncodedLengthByVi64(int64(len(e.properties)))

	for key, value := range e.properties {
		totalLen += util.EncodedLengthByVString(key)
		totalLen += value.EncodedLength()
	}

	e.UniVersionHeader.SetContentLength(totalLen)
	return e.UniVersionHeader.ContentLength()
}

func (e *TableEntity) Encode(buffer *bytes.Buffer) {
	e.UniVersionHeader.Encode(buffer)

	e.rowKey.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(e.properties)))

	for key, value := range e.properties {
		util.EncodeVString(buffer, key)
		value.Encode(buffer)
	}
}

func (e *TableEntity) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

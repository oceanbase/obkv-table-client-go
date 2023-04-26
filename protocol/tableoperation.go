package protocol

import (
	"bytes"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableOperation struct {
	*UniVersionHeader
	opType TableOperationType
	entity *TableEntity
}

func NewTableOperation(
	operationType TableOperationType,
	rowKeys []interface{},
	columnNames []string,
	properties []interface{}) (*TableOperation, error) {
	tableEntity := NewTableEntity()

	// add rowKey
	for _, rowKey := range rowKeys {
		objMeta, err := DefaultObjMeta(rowKey)
		if err != nil {
			return nil, errors.Wrap(err, "create obj meta by row key")
		}

		object := NewObject()
		object.SetMeta(objMeta)
		object.SetValue(rowKey)

		tableEntity.RowKey().AppendKey(object)
	}

	// add column
	for i, columnName := range columnNames {
		var value interface{} = nil
		if properties != nil {
			value = properties[i]
		}
		objMeta, err := DefaultObjMeta(value)
		if err != nil {
			return nil, errors.Wrap(err, "create obj meta by column")
		}

		object := NewObject()
		object.SetMeta(objMeta)
		object.SetValue(value)

		tableEntity.SetProperty(columnName, object)
	}

	return &TableOperation{
		UniVersionHeader: NewUniVersionHeader(),
		opType:           operationType,
		entity:           tableEntity,
	}, nil
}

type TableOperationType uint8

// todo add prefix
const (
	Get TableOperationType = iota
	Insert
	Del
	Update
	InsertOrUpdate
	Replace
	Increment
	Append
)

func (o *TableOperation) OpType() TableOperationType {
	return o.opType
}

func (o *TableOperation) SetOpType(opType TableOperationType) {
	o.opType = opType
}

func (o *TableOperation) Entity() *TableEntity {
	return o.entity
}

func (o *TableOperation) SetEntity(entity *TableEntity) {
	o.entity = entity
}

func (o *TableOperation) PayloadLen() int {
	return o.PayloadContentLen() + o.UniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (o *TableOperation) PayloadContentLen() int {
	totalLen := 1 + // opType
		o.entity.PayloadLen()

	o.UniVersionHeader.SetContentLength(totalLen)
	return o.UniVersionHeader.ContentLength()
}

func (o *TableOperation) Encode(buffer *bytes.Buffer) {
	o.UniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, uint8(o.opType))

	o.entity.Encode(buffer)
}

func (o *TableOperation) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

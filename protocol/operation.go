/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableOperation struct {
	*UniVersionHeader
	opType TableOperationType
	entity *TableEntity
}

func NewTableOperation(
	operationType TableOperationType,
	rowKey []*table.Column,
	columns []*table.Column) (*TableOperation, error) {
	tableEntity := NewTableEntity()

	// add rowKey
	for _, column := range rowKey {
		objMeta, err := DefaultObjMeta(column.Value())
		if err != nil {
			return nil, errors.Wrap(err, "create obj meta by row key")
		}

		object := NewObject()
		object.SetMeta(objMeta)
		object.SetValue(column.Value())

		tableEntity.RowKey().AppendKey(object)
	}

	// add column
	for _, column := range columns {
		objMeta, err := DefaultObjMeta(column.Value())
		if err != nil {
			return nil, errors.Wrap(err, "create obj meta by column")
		}

		object := NewObject()
		object.SetMeta(objMeta)
		object.SetValue(column.Value())

		tableEntity.SetProperty(column.Name(), object)
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

func (o *TableOperation) String() string {
	// todo:impl
	return "TableOperation{" +
		"}"
}

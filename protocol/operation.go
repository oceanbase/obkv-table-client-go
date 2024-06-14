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
	"strconv"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableOperation struct {
	ObUniVersionHeader
	opType ObTableOperationType
	entity *ObTableEntity
}

func NewObTableOperation() *ObTableOperation {
	return &ObTableOperation{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		opType: 0,
		entity: NewObTableEntity(),
	}
}

func NewObTableOperationWithParams(
	tableOperationType ObTableOperationType,
	rowKey []*table.Column,
	columns []*table.Column) (*ObTableOperation, error) {
	tableEntity := NewObTableEntityWithParams(len(rowKey), len(columns))

	// add rowKey
	for _, column := range rowKey {
		objMeta, err := DefaultObjMeta(column.Value())
		if err != nil {
			return nil, errors.WithMessage(err, "create obj meta by row key")
		}

		object := NewObObjectWithParams(objMeta, column.Value())

		tableEntity.AppendRowKeyElement(object)
	}

	// add column
	for _, column := range columns {
		objMeta, err := DefaultObjMeta(column.Value())
		if err != nil {
			return nil, errors.WithMessage(err, "create obj meta by column")
		}

		object := NewObObjectWithParams(objMeta, column.Value())

		tableEntity.SetProperty(column.Name(), object)
	}

	return &ObTableOperation{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		opType: tableOperationType,
		entity: tableEntity,
	}, nil
}

type ObTableOperationType uint8

const (
	ObTableOperationGet ObTableOperationType = iota
	ObTableOperationInsert
	ObTableOperationDel
	ObTableOperationUpdate
	ObTableOperationInsertOrUpdate
	ObTableOperationReplace
	ObTableOperationIncrement
	ObTableOperationAppend
	ObTableOperationRedis = 13
)

func (o *ObTableOperation) OpType() ObTableOperationType {
	return o.opType
}

func (o *ObTableOperation) SetOpType(opType ObTableOperationType) {
	o.opType = opType
}

func (o *ObTableOperation) Entity() *ObTableEntity {
	return o.entity
}

func (o *ObTableOperation) SetEntity(entity *ObTableEntity) {
	o.entity = entity
}

func (o *ObTableOperation) PayloadLen() int {
	return o.PayloadContentLen() + o.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (o *ObTableOperation) PayloadContentLen() int {
	totalLen := 1 + // opType
		o.entity.PayloadLen()

	o.ObUniVersionHeader.SetContentLength(totalLen)
	return o.ObUniVersionHeader.ContentLength()
}

func (o *ObTableOperation) Encode(buffer *bytes.Buffer) {
	o.ObUniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, uint8(o.opType))

	o.entity.Encode(buffer)
}

func (o *ObTableOperation) Decode(buffer *bytes.Buffer) {
	o.ObUniVersionHeader.Decode(buffer)

	o.opType = ObTableOperationType(util.Uint8(buffer))

	o.entity.Decode(buffer)
}

func (o *ObTableOperation) String() string {
	var ObUniVersionHeaderStr = "nil"
	if o.ObUniVersionHeader != (ObUniVersionHeader{}) {
		ObUniVersionHeaderStr = o.ObUniVersionHeader.String()
	}

	var entityStr = "nil"
	if o.entity != nil {
		entityStr = o.entity.String()
	}
	return "ObTableOperation{" +
		"ObUniVersionHeader:" + ObUniVersionHeaderStr + ", " +
		"opType:" + strconv.Itoa(int(o.opType)) + ", " +
		"entity:" + entityStr +
		"}"
}

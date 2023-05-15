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

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableEntityType uint8

const (
	Dynamic TableEntityType = iota
	KV
	HKV
)

type TableEntity struct {
	*UniVersionHeader
	rowKey     *RowKey
	properties map[string]*Object
}

func NewTableEntity() *TableEntity {
	return &TableEntity{
		UniVersionHeader: NewUniVersionHeader(),
		rowKey:           NewRowKey(),
		properties:       make(map[string]*Object),
	}
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

func (e *TableEntity) GetProperty(name string) *Object {
	return e.properties[name]
}

func (e *TableEntity) SetProperty(name string, property *Object) {
	e.properties[name] = property
}

func (e *TableEntity) DelProperty(name string) {
	delete(e.properties, name)
}

// GetSimpleProperties todo optimize
func (e *TableEntity) GetSimpleProperties() map[string]interface{} {
	m := make(map[string]interface{}, len(e.properties))
	for k, v := range e.properties {
		m[k] = v.value
	}
	return m
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
	e.UniVersionHeader.Decode(buffer)

	e.rowKey.Decode(buffer)

	propertiesLen := util.DecodeVi64(buffer)

	var i int64
	for i = 0; i < propertiesLen; i++ {
		name := util.DecodeVString(buffer)

		property := NewObject()
		property.Decode(buffer)

		e.properties[name] = property
	}
}

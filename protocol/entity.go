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

type ObTableEntityType uint8

const (
	ObTableEntityTypeDynamic ObTableEntityType = iota
	ObTableEntityTypeKV
	ObTableEntityTypeHKV
)

type ObTableEntity struct {
	*ObUniVersionHeader
	rowKey     *RowKey
	properties map[string]*ObObject
}

func NewObTableEntity() *ObTableEntity {
	return &ObTableEntity{
		ObUniVersionHeader: NewObUniVersionHeader(),
		rowKey:             NewRowKey(),
		properties:         make(map[string]*ObObject),
	}
}

func (e *ObTableEntity) RowKey() *RowKey {
	return e.rowKey
}

func (e *ObTableEntity) SetRowKey(rowKey *RowKey) {
	e.rowKey = rowKey
}

func (e *ObTableEntity) Properties() map[string]*ObObject {
	return e.properties
}

func (e *ObTableEntity) SetProperties(properties map[string]*ObObject) {
	e.properties = properties
}

func (e *ObTableEntity) GetProperty(name string) *ObObject {
	return e.properties[name]
}

func (e *ObTableEntity) SetProperty(name string, property *ObObject) {
	e.properties[name] = property
}

func (e *ObTableEntity) DelProperty(name string) {
	delete(e.properties, name)
}

// GetSimpleProperties todo optimize
func (e *ObTableEntity) GetSimpleProperties() map[string]interface{} {
	m := make(map[string]interface{}, len(e.properties))
	for k, v := range e.properties {
		m[k] = v.value
	}
	return m
}

func (e *ObTableEntity) PayloadLen() int {
	return e.PayloadContentLen() + e.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (e *ObTableEntity) PayloadContentLen() int {
	totalLen := e.rowKey.EncodedLength()

	totalLen += util.EncodedLengthByVi64(int64(len(e.properties)))

	for key, value := range e.properties {
		totalLen += util.EncodedLengthByVString(key)
		totalLen += value.EncodedLength()
	}

	e.ObUniVersionHeader.SetContentLength(totalLen)
	return e.ObUniVersionHeader.ContentLength()
}

func (e *ObTableEntity) Encode(buffer *bytes.Buffer) {
	e.ObUniVersionHeader.Encode(buffer)

	e.rowKey.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(e.properties)))

	for key, value := range e.properties {
		util.EncodeVString(buffer, key)
		value.Encode(buffer)
	}
}

func (e *ObTableEntity) Decode(buffer *bytes.Buffer) {
	e.ObUniVersionHeader.Decode(buffer)

	e.rowKey.Decode(buffer)

	propertiesLen := util.DecodeVi64(buffer)

	var i int64
	for i = 0; i < propertiesLen; i++ {
		name := util.DecodeVString(buffer)

		property := NewObObject()
		property.Decode(buffer)

		e.properties[name] = property
	}
}

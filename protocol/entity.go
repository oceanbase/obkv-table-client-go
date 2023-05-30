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
	ObUniVersionHeader
	rowKey     []*ObObject
	properties map[string]*ObObject
}

func NewObTableEntity() *ObTableEntity {
	return &ObTableEntity{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		rowKey:     nil,
		properties: nil,
	}
}

func NewObTableEntityWithParams(rowKeyLen int, propertiesLen int) *ObTableEntity {
	return &ObTableEntity{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		rowKey:     make([]*ObObject, 0, rowKeyLen),
		properties: make(map[string]*ObObject, propertiesLen),
	}
}

func (e *ObTableEntity) RowKey() []*ObObject {
	return e.rowKey
}

func (e *ObTableEntity) SetRowKey(rowKey []*ObObject) {
	e.rowKey = rowKey
}

func (e *ObTableEntity) AppendRowKeyElement(object *ObObject) {
	e.rowKey = append(e.rowKey, object)
}

func (e *ObTableEntity) GetRowKeyValue() []interface{} {
	rowKey := make([]interface{}, 0, len(e.rowKey))
	for idx := range e.rowKey {
		rowKey = append(rowKey, e.rowKey[idx].value)
	}
	return rowKey
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

func (e *ObTableEntity) PayloadLen() int {
	return e.PayloadContentLen() + e.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (e *ObTableEntity) PayloadContentLen() int {
	totalLen := util.EncodedLengthByVi64(int64(len(e.rowKey)))

	for _, key := range e.rowKey {
		totalLen += key.EncodedLength()
	}

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

	util.EncodeVi64(buffer, int64(len(e.rowKey)))

	for _, key := range e.rowKey {
		key.Encode(buffer)
	}

	util.EncodeVi64(buffer, int64(len(e.properties)))

	for key, value := range e.properties {
		util.EncodeVString(buffer, key)
		value.Encode(buffer)
	}
}

func (e *ObTableEntity) Decode(buffer *bytes.Buffer) {
	e.ObUniVersionHeader.Decode(buffer)

	rowKeyLen := util.DecodeVi64(buffer)

	e.rowKey = make([]*ObObject, 0, rowKeyLen)

	var i int64
	for i = 0; i < rowKeyLen; i++ {
		key := NewObObject()
		key.Decode(buffer)
		e.rowKey = append(e.rowKey, key)
	}

	propertiesLen := util.DecodeVi64(buffer)

	e.properties = make(map[string]*ObObject, propertiesLen)

	for i = 0; i < propertiesLen; i++ {
		name := util.DecodeVString(buffer)

		property := NewObObject()
		property.Decode(buffer)

		e.properties[name] = property
	}
}

func (e *ObTableEntity) String() string {
	var ObUniVersionHeaderStr = "nil"
	if e.ObUniVersionHeader != (ObUniVersionHeader{}) {
		ObUniVersionHeaderStr = e.ObUniVersionHeader.String()
	}

	var rowKeyStr = "nil"
	if e.rowKey != nil {
		var keysStr string
		keysStr = keysStr + "["
		for i := 0; i < len(e.rowKey); i++ {
			if i > 0 {
				keysStr += ", "
			}
			if e.rowKey[i] != nil {
				keysStr += e.rowKey[i].String()
			} else {
				keysStr += "nil"
			}
		}
		keysStr += "]"
		rowKeyStr = "rowKey:" + keysStr
	}

	var propertiesStr = "properties:{"
	var i = 0
	for k, v := range e.properties {
		if i > 0 {
			propertiesStr += ", "
		}
		i++

		objStr := "nil"
		if v != nil {
			objStr = v.String()
		}
		propertiesStr += "m[" + k + "]=" + objStr
	}
	propertiesStr += "}"

	return "ObTableEntity{" +
		"ObUniVersionHeader:" + ObUniVersionHeaderStr + ", " +
		"rowKey:" + rowKeyStr + ", " +
		"propertiesStr:" + propertiesStr +
		"}"
}

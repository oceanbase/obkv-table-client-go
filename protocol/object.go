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
)

type Object struct {
	meta  *ObjectMeta
	value interface{}
}

func NewObject() *Object {
	return &Object{
		meta:  NewObjectMeta(),
		value: nil,
	}
}

func NewObjectWithParams(meta *ObjectMeta, value interface{}) *Object {
	return &Object{meta: meta, value: value}
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
	o.meta.Decode(buffer)
	o.value = o.meta.ObjType().Decode(buffer, o.meta.CsType())
}

func (o *Object) EncodedLength() int {
	return o.meta.EncodedLength() + o.meta.ObjType().EncodedLength(o.value)
}

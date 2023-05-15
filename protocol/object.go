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

type ObObject struct {
	meta  *ObObjectMeta
	value interface{}
}

func NewObObject() *ObObject {
	return &ObObject{
		meta:  NewObjectMeta(),
		value: nil,
	}
}

func NewObjectWithParams(meta *ObObjectMeta, value interface{}) *ObObject {
	return &ObObject{meta: meta, value: value}
}

func (o *ObObject) Meta() *ObObjectMeta {
	return o.meta
}

func (o *ObObject) SetMeta(meta *ObObjectMeta) {
	o.meta = meta
}

func (o *ObObject) Value() interface{} {
	return o.value
}

func (o *ObObject) SetValue(value interface{}) {
	o.value = value
}

func (o *ObObject) Encode(buffer *bytes.Buffer) {
	o.meta.Encode(buffer)
	o.meta.ObjType().Encode(buffer, o.value)
}

func (o *ObObject) Decode(buffer *bytes.Buffer) {
	o.meta.Decode(buffer)
	o.value = o.meta.ObjType().Decode(buffer, o.meta.CollationType())
}

func (o *ObObject) EncodedLength() int {
	return o.meta.EncodedLength() + o.meta.ObjType().EncodedLength(o.value)
}

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

type ObObject struct {
	meta  ObObjectMeta
	value interface{}
}

func NewObObject() *ObObject {
	return &ObObject{
		meta:  ObObjectMeta{},
		value: nil,
	}
}

func NewObObjectWithParams(meta ObObjectMeta, value interface{}) *ObObject {
	return &ObObject{
		meta:  meta,
		value: value,
	}
}

func (o *ObObject) Meta() ObObjectMeta {
	return o.meta
}

func (o *ObObject) SetMeta(meta ObObjectMeta) {
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

func GetMin() *ObObject {
	objType, _ := NewObjType(ObObjTypeExtendTypeValue)
	// -3 = -1 -2 = MaxUint64 - 2
	return NewObObjectWithParams(objType.DefaultObjMeta(), int64(-3))
}

func GetMax() *ObObject {
	objType, _ := NewObjType(ObObjTypeExtendTypeValue)
	// -2 = -1 -1 = MaxUint64 - 1
	return NewObObjectWithParams(objType.DefaultObjMeta(), int64(-2))
}

func (o *ObObject) String() string {
	var metaStr = "nil"
	if o.meta != (ObObjectMeta{}) {
		metaStr = o.meta.String()
	}
	return "ObObject{" +
		"meta:" + metaStr + ", " +
		"value:" + util.InterfaceToString(o.value) +
		"}"
}

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

package route

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

func TestObKeyPartDesc_String(t *testing.T) {
	desc := &obKeyPartDesc{}
	assert.Equal(t, "obKeyPartDesc{partSpace:0, partNum:0, partColumns[]}", desc.String())
	desc = newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2)
	assert.Equal(t, "obKeyPartDesc{partSpace:0, partNum:10, partColumns[]}", desc.String())

	desc = newObKeyPartDesc(0, 10, partFuncTypeHash)
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", objType, protocol.ObCollationTypeBinary)
	desc.SetPartColumns([]obColumn{col})
	assert.Equal(t, "obKeyPartDesc{partSpace:0, partNum:10, partColumns[obSimpleColumn{columnName:c1, objType:ObObjType{type:ObInt64Type}, collationType:63}]}", desc.String())
}

func TestObKeyPartDesc_GetPartId(t *testing.T) {
	desc := newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2)
	partId, err := desc.GetPartId([]*table.Column{})
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)

	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", objType, protocol.ObCollationTypeBinary)
	desc.SetPartColumns([]obColumn{col})
	partId, err = desc.GetPartId([]*table.Column{table.NewColumn("c1", int64(1))})
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 6, partId)
}

func TestObKeyPartDesc_intToInt64(t *testing.T) {
	v, _ := intToInt64(true)
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(false)
	assert.EqualValues(t, 0, v)
	v, _ = intToInt64(int8(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(uint8(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(int16(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(uint16(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(1)
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(uint(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(int32(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(uint32(1))
	assert.EqualValues(t, 1, v)
	v, _ = intToInt64(int64(1))
	assert.EqualValues(t, 1, v)
	_, err := intToInt64("abc")
	assert.NotEqual(t, nil, err)
}

func TestObKeyPartDesc_toHashCode(t *testing.T) {
	// int64
	desc := newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2)
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", objType, protocol.ObCollationTypeBinary)
	desc.SetPartColumns([]obColumn{col})
	var hashCode int64 = 0
	hashCode, err := desc.toHashCode(int64(123), col, hashCode, partFuncTypeKeyImplV2)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, -1234695888563024189, hashCode)
	// varchar ObCollationTypeUtf8mb4GeneralCi
	desc = newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2)
	objType, _ = protocol.NewObjType(protocol.ObObjTypeVarcharTypeValue)
	col = newObSimpleColumn("c1", objType, protocol.ObCollationTypeUtf8mb4GeneralCi)
	desc.SetPartColumns([]obColumn{col})
	hashCode = 0
	hashCode, err = desc.toHashCode("abc", col, hashCode, partFuncTypeKeyImplV2)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, -925032385642742853, hashCode)
	// varchar ObCollationTypeUtf8mb4Bin
	desc = newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2)
	objType, _ = protocol.NewObjType(protocol.ObObjTypeVarcharTypeValue)
	col = newObSimpleColumn("c1", objType, protocol.ObCollationTypeUtf8mb4Bin)
	desc.SetPartColumns([]obColumn{col})
	hashCode = 0
	hashCode, err = desc.toHashCode("abc", col, hashCode, partFuncTypeKeyImplV2)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, -7148968302806999301, hashCode)
	// varchar ObCollationTypeInvalid
	desc = newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2)
	objType, _ = protocol.NewObjType(protocol.ObObjTypeVarcharTypeValue)
	col = newObSimpleColumn("c1", objType, protocol.ObCollationTypeInvalid)
	desc.SetPartColumns([]obColumn{col})
	hashCode = 0
	_, err = desc.toHashCode("abc", col, hashCode, partFuncTypeKeyImplV2)
	assert.NotEqual(t, nil, err)
}

func TestObKeyPartDesc_longHash(t *testing.T) {
	var hashCode int64 = 0
	desc := obKeyPartDesc{}
	assert.EqualValues(t, -8089716718896805586, desc.longHash(1, hashCode))
	hashCode = 0
	assert.EqualValues(t, 4559785591630536017, desc.longHash(12, hashCode))
	hashCode = 0
	assert.EqualValues(t, -8599128181304237022, desc.longHash(12345, hashCode))
}

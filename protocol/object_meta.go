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
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObObjectMeta struct {
	objType        ObObjType
	collationLevel ObCollationLevel
	collationType  ObCollationType
	scale          byte
}

func NewObObjectMeta() *ObObjectMeta {
	return &ObObjectMeta{objType: nil, collationLevel: 0, collationType: 0, scale: 0}
}

func NewObObjectMetaWithParams(objType ObObjType, obCollationLevel ObCollationLevel, obCollationType ObCollationType, scale byte) *ObObjectMeta {
	return &ObObjectMeta{objType: objType, collationLevel: obCollationLevel, collationType: obCollationType, scale: scale}
}

func (m *ObObjectMeta) ObjType() ObObjType {
	return m.objType
}

func (m *ObObjectMeta) SetObjType(objType ObObjType) {
	m.objType = objType
}

func (m *ObObjectMeta) CollationLevel() ObCollationLevel {
	return m.collationLevel
}

func (m *ObObjectMeta) SetCollationLevel(collationLevel ObCollationLevel) {
	m.collationLevel = collationLevel
}

func (m *ObObjectMeta) CollationType() ObCollationType {
	return m.collationType
}

func (m *ObObjectMeta) SetCollationType(collationType ObCollationType) {
	m.collationType = collationType
}

func (m *ObObjectMeta) Scale() byte {
	return m.scale
}

func (m *ObObjectMeta) SetScale(scale byte) {
	m.scale = scale
}

func (m *ObObjectMeta) String() string {
	var objTypeStr = "nil"
	if m.objType != nil {
		objTypeStr = m.objType.String()
	}
	return "ObObjectMeta{" +
		"objType:" + objTypeStr + ", " +
		"collationLevel:" + strconv.Itoa(int(m.collationLevel)) + ", " +
		"collationType:" + strconv.Itoa(int(m.collationType)) + ", " +
		"scale:" + strconv.Itoa(int(m.scale)) +
		"}"
}

func (m *ObObjectMeta) Encode(buffer *bytes.Buffer) {
	util.PutUint8(buffer, uint8(m.objType.Value()))
	util.PutUint8(buffer, uint8(m.collationLevel))
	util.PutUint8(buffer, uint8(m.collationType))
	util.PutUint8(buffer, m.scale)
}

func (m *ObObjectMeta) Decode(buffer *bytes.Buffer) {
	m.objType = ObObjTypeValue(util.Uint8(buffer)).ValueOf()
	m.collationLevel = ObCollationLevel(util.Uint8(buffer))
	m.collationType = ObCollationType(util.Uint8(buffer))
	m.scale = util.Uint8(buffer)
}

func (m *ObObjectMeta) EncodedLength() int {
	return 4 // objType collationLevel collationType scale
}

type ObObjType interface {
	Encode(buffer *bytes.Buffer, value interface{})
	Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{}
	EncodedLength(value interface{}) int
	DefaultObjMeta() *ObObjectMeta
	CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error)
	Value() ObObjTypeValue
	String() string
}

func DefaultObjMeta(value interface{}) (*ObObjectMeta, error) {
	if value == nil {
		return ObObjTypes[ObObjTypeNullTypeValue].DefaultObjMeta(), nil
	}
	switch value.(type) {
	case bool:
		return ObObjTypes[ObObjTypeTinyIntTypeValue].DefaultObjMeta(), nil
	case int8:
		return ObObjTypes[ObObjTypeTinyIntTypeValue].DefaultObjMeta(), nil
	case uint8:
		return ObObjTypes[ObObjTypeUTinyIntTypeValue].DefaultObjMeta(), nil
	case int16:
		return ObObjTypes[ObObjTypeSmallIntTypeValue].DefaultObjMeta(), nil
	case uint16:
		return ObObjTypes[ObObjTypeUSmallIntTypeValue].DefaultObjMeta(), nil
	case int32:
		return ObObjTypes[ObObjTypeInt32TypeValue].DefaultObjMeta(), nil
	case uint32:
		return ObObjTypes[ObObjTypeUInt32TypeValue].DefaultObjMeta(), nil
	case int64:
		return ObObjTypes[ObObjTypeInt64TypeValue].DefaultObjMeta(), nil
	case uint64:
		return ObObjTypes[ObObjTypeUInt64TypeValue].DefaultObjMeta(), nil
	case float32:
		return ObObjTypes[ObObjTypeFloatTypeValue].DefaultObjMeta(), nil
	case float64:
		return ObObjTypes[ObObjTypeDoubleTypeValue].DefaultObjMeta(), nil
	case string:
		return ObObjTypes[ObObjTypeVarcharTypeValue].DefaultObjMeta(), nil
	case []byte:
		return ObObjTypes[ObObjTypeVarcharTypeValue].DefaultObjMeta(), nil
	case time.Duration:
		return ObObjTypes[ObObjTypeDateTimeTypeValue].DefaultObjMeta(), nil
	default:
		return nil, errors.Errorf("not match objmeta, value: %v", value)
	}
}

func NewObjType(value ObObjTypeValue) (ObObjType, error) {
	switch value {
	case ObObjTypeNullTypeValue:
		return &ObNullType{value: value}, nil
	case ObObjTypeTinyIntTypeValue:
		return &ObTinyIntType{value: value}, nil
	case ObObjTypeSmallIntTypeValue:
		return &ObSmallIntType{value: value}, nil
	case ObObjTypeMediumIntTypeValue:
		return &ObMediumIntType{value: value}, nil
	case ObObjTypeInt32TypeValue:
		return &ObInt32Type{value: value}, nil
	case ObObjTypeInt64TypeValue:
		return &ObInt64Type{value: value}, nil
	case ObObjTypeUTinyIntTypeValue:
		return &ObUTinyIntType{value: value}, nil
	case ObObjTypeUSmallIntTypeValue:
		return &ObUSmallIntType{value: value}, nil
	case ObObjTypeUMediumIntTypeValue:
		return &ObUMediumIntType{value: value}, nil
	case ObObjTypeUInt32TypeValue:
		return &ObUInt32Type{value: value}, nil
	case ObObjTypeUInt64TypeValue:
		return &ObUInt64Type{value: value}, nil
	case ObObjTypeFloatTypeValue:
		return &ObFloatType{value: value}, nil
	case ObObjTypeDoubleTypeValue:
		return &ObDoubleType{value: value}, nil
	case ObObjTypeUFloatTypeValue:
		return &ObUFloatType{value: value}, nil
	case ObObjTypeUDoubleTypeValue:
		return &ObUDoubleType{value: value}, nil
	case ObObjTypeNumberTypeValue:
		return &ObNumberType{value: value}, nil
	case ObObjTypeUNumberTypeValue:
		return &ObUNumberType{value: value}, nil
	case ObObjTypeDateTimeTypeValue:
		return &ObDateTimeType{value: value}, nil
	case ObObjTypeTimestampTypeValue:
		return &ObTimestampType{value: value}, nil
	case ObObjTypeDateTypeValue:
		return &ObDateType{value: value}, nil
	case ObObjTypeTimeTypeValue:
		return &ObTimeType{value: value}, nil
	case ObObjTypeYearTypeValue:
		return &ObYearType{value: value}, nil
	case ObObjTypeVarcharTypeValue:
		return &ObVarcharType{value: value}, nil
	case ObObjTypeCharTypeValue:
		return &ObCharType{value: value}, nil
	case ObObjTypeHexStringTypeValue:
		return &ObHexStringType{value: value}, nil
	case ObObjTypeExtendTypeValue:
		return &ObExtendType{value: value}, nil
	case ObObjTypeUnknownTypeValue:
		return &ObUnknownType{value: value}, nil
	case ObObjTypeTinyTextTypeValue:
		return &ObTinyTextType{value: value}, nil
	case ObObjTypeTextTypeValue:
		return &ObTextType{value: value}, nil
	case ObObjTypeMediumTextTypeValue:
		return &ObMediumTextType{value: value}, nil
	case ObObjTypeLongTextTypeValue:
		return &ObLongTextType{value: value}, nil
	case ObObjTypeBitTypeValue:
		return &ObBitType{value: value}, nil
	default:
		return nil, errors.Errorf("not match objtype, value: %d", value)
	}
}

type ObObjTypeValue uint8

var ObObjTypes = []ObObjType{
	ObObjTypeNullTypeValue:       &ObNullType{value: ObObjTypeNullTypeValue},
	ObObjTypeTinyIntTypeValue:    &ObTinyIntType{value: ObObjTypeTinyIntTypeValue},
	ObObjTypeSmallIntTypeValue:   &ObSmallIntType{value: ObObjTypeSmallIntTypeValue},
	ObObjTypeMediumIntTypeValue:  &ObMediumIntType{value: ObObjTypeMediumIntTypeValue},
	ObObjTypeInt32TypeValue:      &ObInt32Type{value: ObObjTypeInt32TypeValue},
	ObObjTypeInt64TypeValue:      &ObInt64Type{value: ObObjTypeInt64TypeValue},
	ObObjTypeUTinyIntTypeValue:   &ObUTinyIntType{value: ObObjTypeUTinyIntTypeValue},
	ObObjTypeUSmallIntTypeValue:  &ObUSmallIntType{value: ObObjTypeUSmallIntTypeValue},
	ObObjTypeUMediumIntTypeValue: &ObUMediumIntType{value: ObObjTypeUMediumIntTypeValue},
	ObObjTypeUInt32TypeValue:     &ObUInt32Type{value: ObObjTypeUInt32TypeValue},
	ObObjTypeUInt64TypeValue:     &ObUInt64Type{value: ObObjTypeUInt64TypeValue},
	ObObjTypeFloatTypeValue:      &ObFloatType{value: ObObjTypeFloatTypeValue},
	ObObjTypeDoubleTypeValue:     &ObDoubleType{value: ObObjTypeDoubleTypeValue},
	ObObjTypeUFloatTypeValue:     &ObUFloatType{value: ObObjTypeUFloatTypeValue},
	ObObjTypeUDoubleTypeValue:    &ObUDoubleType{value: ObObjTypeUDoubleTypeValue},
	ObObjTypeNumberTypeValue:     &ObNumberType{value: ObObjTypeNumberTypeValue},
	ObObjTypeUNumberTypeValue:    &ObUNumberType{value: ObObjTypeUNumberTypeValue},
	ObObjTypeDateTimeTypeValue:   &ObDateTimeType{value: ObObjTypeDateTimeTypeValue},
	ObObjTypeTimestampTypeValue:  &ObTimestampType{value: ObObjTypeTimestampTypeValue},
	ObObjTypeDateTypeValue:       &ObDateType{value: ObObjTypeDateTypeValue},
	ObObjTypeTimeTypeValue:       &ObTimeType{value: ObObjTypeTimeTypeValue},
	ObObjTypeYearTypeValue:       &ObYearType{value: ObObjTypeYearTypeValue},
	ObObjTypeVarcharTypeValue:    &ObVarcharType{value: ObObjTypeVarcharTypeValue},
	ObObjTypeCharTypeValue:       &ObCharType{value: ObObjTypeCharTypeValue},
	ObObjTypeHexStringTypeValue:  &ObHexStringType{value: ObObjTypeHexStringTypeValue},
	ObObjTypeExtendTypeValue:     &ObExtendType{value: ObObjTypeExtendTypeValue},
	ObObjTypeUnknownTypeValue:    &ObUnknownType{value: ObObjTypeUnknownTypeValue},
	ObObjTypeTinyTextTypeValue:   &ObTinyTextType{value: ObObjTypeTinyTextTypeValue},
	ObObjTypeTextTypeValue:       &ObTextType{value: ObObjTypeTextTypeValue},
	ObObjTypeMediumTextTypeValue: &ObMediumTextType{value: ObObjTypeMediumTextTypeValue},
	ObObjTypeLongTextTypeValue:   &ObLongTextType{value: ObObjTypeLongTextTypeValue},
	ObObjTypeBitTypeValue:        &ObBitType{value: ObObjTypeBitTypeValue},
}

func (v ObObjTypeValue) ValueOf() ObObjType {
	return ObObjTypes[v]
}

const (
	ObObjTypeNullTypeValue ObObjTypeValue = iota
	ObObjTypeTinyIntTypeValue
	ObObjTypeSmallIntTypeValue
	ObObjTypeMediumIntTypeValue
	ObObjTypeInt32TypeValue
	ObObjTypeInt64TypeValue
	ObObjTypeUTinyIntTypeValue
	ObObjTypeUSmallIntTypeValue
	ObObjTypeUMediumIntTypeValue
	ObObjTypeUInt32TypeValue
	ObObjTypeUInt64TypeValue
	ObObjTypeFloatTypeValue
	ObObjTypeDoubleTypeValue
	ObObjTypeUFloatTypeValue
	ObObjTypeUDoubleTypeValue
	ObObjTypeNumberTypeValue
	ObObjTypeUNumberTypeValue
	ObObjTypeDateTimeTypeValue
	ObObjTypeTimestampTypeValue
	ObObjTypeDateTypeValue
	ObObjTypeTimeTypeValue
	ObObjTypeYearTypeValue
	ObObjTypeVarcharTypeValue
	ObObjTypeCharTypeValue
	ObObjTypeHexStringTypeValue
	ObObjTypeExtendTypeValue
	ObObjTypeUnknownTypeValue
	ObObjTypeTinyTextTypeValue
	ObObjTypeTextTypeValue
	ObObjTypeMediumTextTypeValue
	ObObjTypeLongTextTypeValue
	ObObjTypeBitTypeValue
)

type ObCollationLevel uint8

const (
	ObCollationLevelExplicit ObCollationLevel = iota
	ObCollationLevelNone
	ObCollationLevelImplicit
	ObCollationLevelSysConst
	ObCollationLevelCoercible
	ObCollationLevelNumeric
	ObCollationLevelIgnorable
	CollationLevelInvalid ObCollationLevel = 127
)

func (l ObCollationLevel) String() string {
	switch l {
	case ObCollationLevelExplicit:
		return "ObCollationLevelExplicit"
	case ObCollationLevelNone:
		return "ObCollationLevelNone"
	case ObCollationLevelImplicit:
		return "ObCollationLevelImplicit"
	case ObCollationLevelSysConst:
		return "ObCollationLevelSysConst"
	case ObCollationLevelCoercible:
		return "ObCollationLevelCoercible"
	case ObCollationLevelNumeric:
		return "ObCollationLevelNumeric"
	case ObCollationLevelIgnorable:
		return "ObCollationLevelIgnorable"
	default:
		return "ObCollationLevelInvalid"
	}
}

type ObCollationType uint8

const (
	ObCollationTypeInvalid          ObCollationType = 0
	ObCollationTypeUtf8mb4GeneralCi ObCollationType = 45
	ObCollationTypeUtf8mb4Bin       ObCollationType = 46
	ObCollationTypeBinary           ObCollationType = 63
	ObCollationTypeCollationFree    ObCollationType = 100
	ObCollationTypeMax              ObCollationType = 101
)

func (t ObCollationType) String() string {
	switch t {
	case ObCollationTypeUtf8mb4GeneralCi:
		return "ObCollationTypeUtf8mb4GeneralCi"
	case ObCollationTypeUtf8mb4Bin:
		return "ObCollationTypeUtf8mb4Bin"
	case ObCollationTypeBinary:
		return "ObCollationTypeBinary"
	case ObCollationTypeCollationFree:
		return "ObCollationTypeCollationFree"
	case ObCollationTypeMax:
		return "ObCollationTypeMax"
	default:
		return "ObCollationTypeInvalid"
	}
}

type ObNullType struct {
	value ObObjTypeValue
}

func (t *ObNullType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObNullType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObNullType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObNullType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelIgnorable, ObCollationTypeBinary, 10)
}

func (t *ObNullType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("null type can not parse to comparable value")
}

func (t *ObNullType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObNullType) String() string {
	return "ObObjType{" +
		"type:" + "ObNullType" +
		"}"
}

type ObTinyIntType struct {
	value ObObjTypeValue
}

func (t *ObTinyIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	switch v := value.(type) {
	case bool:
		if v {
			util.PutUint8(buffer, 1)
		} else {
			util.PutUint8(buffer, 0)
		}
	case int8:
		util.PutUint8(buffer, uint8(v))
	}
}

func (t *ObTinyIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.Uint8(buffer)
}

func (t *ObTinyIntType) EncodedLength(value interface{}) int {
	return 1
}

func (t *ObTinyIntType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObTinyIntType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case int8:
		return v, nil
	default:
		return nil, errors.Errorf("tiny int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObTinyIntType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObTinyIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObTinyIntType" +
		"}"
}

type ObSmallIntType struct {
	value ObObjTypeValue
}

func (t *ObSmallIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(int16)))
}

func (t *ObSmallIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *ObSmallIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(int16)))
}

func (t *ObSmallIntType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObSmallIntType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(int16); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("small int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObSmallIntType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObSmallIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObSmallIntType" +
		"}"
}

type ObMediumIntType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObMediumIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, value.(int32))
}

func (t *ObMediumIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *ObMediumIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(value.(int32))
}

func (t *ObMediumIntType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObMediumIntType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("medium int type is not support")
}

func (t *ObMediumIntType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObMediumIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObMediumIntType" +
		"}"
}

type ObInt32Type struct {
	value ObObjTypeValue
}

func (t *ObInt32Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, value.(int32))
}

func (t *ObInt32Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *ObInt32Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(value.(int32))
}

func (t *ObInt32Type) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObInt32Type) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(int32); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("int32 type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObInt32Type) Value() ObObjTypeValue {
	return t.value
}

func (t *ObInt32Type) String() string {
	return "ObObjType{" +
		"type:" + "ObInt32Type" +
		"}"
}

type ObInt64Type struct {
	value ObObjTypeValue
}

func (t *ObInt64Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t *ObInt64Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObInt64Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t *ObInt64Type) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObInt64Type) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(int64); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("int64 type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObInt64Type) Value() ObObjTypeValue {
	return t.value
}

func (t *ObInt64Type) String() string {
	return "ObObjType{" +
		"type:" + "ObInt64Type" +
		"}"
}

type ObUTinyIntType struct {
	value ObObjTypeValue
}

func (t *ObUTinyIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.PutUint8(buffer, value.(uint8))
}

func (t *ObUTinyIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.Uint8(buffer)
}

func (t *ObUTinyIntType) EncodedLength(value interface{}) int {
	return 1
}

func (t *ObUTinyIntType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUTinyIntType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(uint8); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uTiny int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObUTinyIntType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUTinyIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObUTinyIntType" +
		"}"
}

type ObUSmallIntType struct {
	value ObObjTypeValue
}

func (t *ObUSmallIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint16)))
}

func (t *ObUSmallIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *ObUSmallIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint16)))
}

func (t *ObUSmallIntType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUSmallIntType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(uint16); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uSmall int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObUSmallIntType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUSmallIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObUSmallIntType" +
		"}"
}

type ObUMediumIntType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObUMediumIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint32)))
}

func (t *ObUMediumIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *ObUMediumIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint32)))
}

func (t *ObUMediumIntType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUMediumIntType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("uMedium int type is not support")
}

func (t *ObUMediumIntType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUMediumIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObUMediumIntType" +
		"}"
}

type ObUInt32Type struct {
	value ObObjTypeValue
}

func (t *ObUInt32Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint32)))
}

func (t *ObUInt32Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *ObUInt32Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint32)))
}

func (t *ObUInt32Type) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUInt32Type) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(uint32); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uInt int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObUInt32Type) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUInt32Type) String() string {
	return "ObObjType{" +
		"type:" + "ObUInt32Type" +
		"}"
}

type ObUInt64Type struct {
	value ObObjTypeValue
}

func (t *ObUInt64Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(uint64)))
}

func (t *ObUInt64Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObUInt64Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(uint64)))
}

func (t *ObUInt64Type) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUInt64Type) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(uint64); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uInt64 type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObUInt64Type) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUInt64Type) String() string {
	return "ObObjType{" +
		"type:" + "ObUInt64Type" +
		"}"
}

type ObFloatType struct {
	value ObObjTypeValue
}

func (t *ObFloatType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVf32(buffer, value.(float32))
}

func (t *ObFloatType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVf32(buffer)
}

func (t *ObFloatType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVf32(value.(float32))
}

func (t *ObFloatType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObFloatType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(float32); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("float type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObFloatType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObFloatType) String() string {
	return "ObObjType{" +
		"type:" + "ObFloatType" +
		"}"
}

type ObDoubleType struct {
	value ObObjTypeValue
}

func (t *ObDoubleType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVf64(buffer, value.(float64))
}

func (t *ObDoubleType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVf64(buffer)
}

func (t *ObDoubleType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVf64(value.(float64))
}

func (t *ObDoubleType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObDoubleType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(float64); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("double type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObDoubleType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObDoubleType) String() string {
	return "ObObjType{" +
		"type:" + "ObDoubleType" +
		"}"
}

type ObUFloatType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObUFloatType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObUFloatType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObUFloatType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObUFloatType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUFloatType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("uFloat type is not support")
}

func (t *ObUFloatType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUFloatType) String() string {
	return "ObObjType{" +
		"type:" + "ObUFloatType" +
		"}"
}

type ObUDoubleType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObUDoubleType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObUDoubleType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObUDoubleType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObUDoubleType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUDoubleType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("uDouble type is not support")
}

func (t *ObUDoubleType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUDoubleType) String() string {
	return "ObObjType{" +
		"type:" + "ObUDoubleType" +
		"}"
}

type ObNumberType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObNumberType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObNumberType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObNumberType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObNumberType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObNumberType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("number type is not support")
}

func (t *ObNumberType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObNumberType) String() string {
	return "ObObjType{" +
		"type:" + "ObNumberType" +
		"}"
}

type ObUNumberType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObUNumberType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObUNumberType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObUNumberType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObUNumberType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUNumberType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("uNumber type is not support")
}

func (t *ObUNumberType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUNumberType) String() string {
	return "ObObjType{" +
		"type:" + "ObUNumberType" +
		"}"
}

type ObDateTimeType struct {
	value ObObjTypeValue
}

func (t *ObDateTimeType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration)))
}

func (t *ObDateTimeType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObDateTimeType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *ObDateTimeType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObDateTimeType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("date time type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObDateTimeType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObDateTimeType) String() string {
	return "ObObjType{" +
		"type:" + "ObDateTimeType" +
		"}"
}

type ObTimestampType struct {
	value ObObjTypeValue
}

func (t *ObTimestampType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration)))
}

func (t *ObTimestampType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObTimestampType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *ObTimestampType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObTimestampType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("time stamp type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObTimestampType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObTimestampType) String() string {
	return "ObObjType{" +
		"type:" + "ObTimestampType" +
		"}"
}

type ObDateType struct {
	value ObObjTypeValue
}

func (t *ObDateType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration)))
}

func (t *ObDateType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObDateType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *ObDateType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObDateType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("data type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObDateType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObDateType) String() string {
	return "ObObjType{" +
		"type:" + "ObDateType" +
		"}"
}

type ObTimeType struct {
	value ObObjTypeValue
}

func (t *ObTimeType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration)))
}

func (t *ObTimeType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObTimeType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *ObTimeType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObTimeType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("time type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *ObTimeType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObTimeType) String() string {
	return "ObObjType{" +
		"type:" + "ObTimeType" +
		"}"
}

type ObYearType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObYearType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.PutUint8(buffer, value.(uint8))
}

func (t *ObYearType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.Uint8(buffer)
}

func (t *ObYearType) EncodedLength(value interface{}) int {
	return 1
}

func (t *ObYearType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObYearType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("year type is not support")
}

func (t *ObYearType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObYearType) String() string {
	return "ObObjType{" +
		"type:" + "ObYearType" +
		"}"
}

type ObVarcharType struct {
	value ObObjTypeValue
}

func (t *ObVarcharType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *ObVarcharType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t *ObVarcharType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *ObVarcharType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelExplicit, ObCollationTypeUtf8mb4GeneralCi, 10)
}

func (t *ObVarcharType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if ObCollationTypeBinary == obCollationType {
		if v, ok := value.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("varchar type parse to comparable failed, not match value, value: %v", v)
		}
	} else {
		if v, ok := value.(string); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("varchar type parse to comparable failed, not match value, value: %v", v)
		}
	}
}

func (t *ObVarcharType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObVarcharType) String() string {
	return "ObObjType{" +
		"type:" + "ObVarcharType" +
		"}"
}

type ObCharType struct {
	value ObObjTypeValue
}

func (t *ObCharType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *ObCharType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t *ObCharType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *ObCharType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelExplicit, ObCollationTypeUtf8mb4GeneralCi, 10)
}

func (t *ObCharType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if ObCollationTypeBinary == obCollationType {
		if v, ok := value.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("char type parse to comparable failed, not match value, value: %v", v)
		}
	} else {
		if v, ok := value.(string); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("char type parse to comparable failed, not match value, value: %v", v)
		}
	}
}

func (t *ObCharType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObCharType) String() string {
	return "ObObjType{" +
		"type:" + "ObCharType" +
		"}"
}

type ObHexStringType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObHexStringType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObHexStringType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObHexStringType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObHexStringType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObHexStringType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("hex string type is not support")
}

func (t *ObHexStringType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObHexStringType) String() string {
	return "ObObjType{" +
		"type:" + "ObHexStringType" +
		"}"
}

type ObExtendType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObExtendType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t *ObExtendType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObExtendType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t *ObExtendType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObExtendType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("extend type is not support")
}

func (t *ObExtendType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObExtendType) String() string {
	return "ObObjType{" +
		"type:" + "ObExtendType" +
		"}"
}

type ObUnknownType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObUnknownType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *ObUnknownType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t *ObUnknownType) EncodedLength(value interface{}) int {
	return 0
}

func (t *ObUnknownType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObUnknownType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("unknown type is not support")
}

func (t *ObUnknownType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObUnknownType) String() string {
	return "ObObjType{" +
		"type:" + "ObUnknownType" +
		"}"
}

type ObTinyTextType struct {
	value ObObjTypeValue
}

func (t *ObTinyTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *ObTinyTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t *ObTinyTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *ObTinyTextType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObTinyTextType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if ObCollationTypeBinary == obCollationType {
		if v, ok := value.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("tiny text type parse to comparable failed, not match value, value: %v", v)
		}
	} else {
		if v, ok := value.(string); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("tiny text type parse to comparable failed, not match value, value: %v", v)
		}
	}
}

func (t *ObTinyTextType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObTinyTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObTinyTextType" +
		"}"
}

type ObTextType struct {
	value ObObjTypeValue
}

func (t *ObTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *ObTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t *ObTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *ObTextType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObTextType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if ObCollationTypeBinary == obCollationType {
		if v, ok := value.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("text type parse to comparable failed, not match value, value: %v", v)
		}
	} else {
		if v, ok := value.(string); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("text type parse to comparable failed, not match value, value: %v", v)
		}
	}
}

func (t *ObTextType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObTextType" +
		"}"
}

type ObMediumTextType struct {
	value ObObjTypeValue
}

func (t *ObMediumTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *ObMediumTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t *ObMediumTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *ObMediumTextType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObMediumTextType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if ObCollationTypeBinary == obCollationType {
		if v, ok := value.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("medium text type parse to comparable failed, not match value, value: %v", v)
		}
	} else {
		if v, ok := value.(string); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("medium text type parse to comparable failed, not match value, value: %v", v)
		}
	}
}

func (t *ObMediumTextType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObMediumTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObMediumTextType" +
		"}"
}

type ObLongTextType struct {
	value ObObjTypeValue
}

func (t *ObLongTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *ObLongTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t *ObLongTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *ObLongTextType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObLongTextType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	if ObCollationTypeBinary == obCollationType {
		if v, ok := value.([]byte); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("long text type parse to comparable failed, not match value, value: %v", v)
		}
	} else {
		if v, ok := value.(string); ok {
			return v, nil
		} else {
			return nil, errors.Errorf("long text type parse to comparable failed, not match value, value: %v", v)
		}
	}
}

func (t *ObLongTextType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObLongTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObLongTextType" +
		"}"
}

type ObBitType struct { // TODO not support
	value ObObjTypeValue
}

func (t *ObBitType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t *ObBitType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ObBitType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t *ObBitType) DefaultObjMeta() *ObObjectMeta {
	return NewObObjectMetaWithParams(t, ObCollationLevelNumeric, ObCollationTypeBinary, 10)
}

func (t *ObBitType) CheckTypeForValue(value interface{}, obCollationType ObCollationType) (interface{}, error) {
	return nil, errors.New("bit type is not support")
}

func (t *ObBitType) Value() ObObjTypeValue {
	return t.value
}

func (t *ObBitType) String() string {
	return "ObObjType{" +
		"type:" + "ObBitType" +
		"}"
}

func EncodeText(buffer *bytes.Buffer, value interface{}) {
	switch v := value.(type) {
	case string:
		util.EncodeVString(buffer, v)
	case []byte:
		util.EncodeBytesString(buffer, v)
	default:
		// do nothing
		return
	}
}

func DecodeText(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	if ObCollationTypeBinary == obCollationType {
		return util.DecodeBytesString(buffer)
	} else {
		return util.DecodeVString(buffer)
	}
}

func EncodedLengthByText(value interface{}) int {
	switch v := value.(type) {
	case string:
		return util.EncodedLengthByVString(v)
	case []byte:
		return util.EncodedLengthByBytesString(v)
	default:
		// do nothing
		return 0
	}
}

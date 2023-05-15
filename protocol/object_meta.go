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
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObjectMeta struct {
	objType ObjType
	csLevel CollationLevel
	csType  CollationType
	scale   byte
}

func NewObjectMeta() *ObjectMeta {
	return &ObjectMeta{objType: nil, csLevel: 0, csType: 0, scale: 0}
}

func NewObjectMetaWithParams(objType ObjType, csLevel CollationLevel, csType CollationType, scale byte) *ObjectMeta {
	return &ObjectMeta{objType: objType, csLevel: csLevel, csType: csType, scale: scale}
}

func (m *ObjectMeta) ObjType() ObjType {
	return m.objType
}

func (m *ObjectMeta) SetObjType(objType ObjType) {
	m.objType = objType
}

func (m *ObjectMeta) CsLevel() CollationLevel {
	return m.csLevel
}

func (m *ObjectMeta) SetCsLevel(csLevel CollationLevel) {
	m.csLevel = csLevel
}

func (m *ObjectMeta) CsType() CollationType {
	return m.csType
}

func (m *ObjectMeta) SetCsType(csType CollationType) {
	m.csType = csType
}

func (m *ObjectMeta) Scale() byte {
	return m.scale
}

func (m *ObjectMeta) SetScale(scale byte) {
	m.scale = scale
}

func (m *ObjectMeta) Encode(buffer *bytes.Buffer) {
	util.PutUint8(buffer, uint8(m.objType.Value()))
	util.PutUint8(buffer, uint8(m.csLevel))
	util.PutUint8(buffer, uint8(m.csType))
	util.PutUint8(buffer, m.scale)
}

func (m *ObjectMeta) Decode(buffer *bytes.Buffer) {
	m.objType = ObjTypeValue(util.Uint8(buffer)).ValueOf()
	m.csLevel = CollationLevel(util.Uint8(buffer))
	m.csType = CollationType(util.Uint8(buffer))
	m.scale = util.Uint8(buffer)
}

func (m *ObjectMeta) EncodedLength() int {
	return 4 // objType csLevel csType scale
}

type ObjType interface {
	Encode(buffer *bytes.Buffer, value interface{})
	Decode(buffer *bytes.Buffer, csType CollationType) interface{}
	EncodedLength(value interface{}) int
	DefaultObjMeta() *ObjectMeta
	CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error)
	Value() ObjTypeValue
	String() string
}

func DefaultObjMeta(value interface{}) (*ObjectMeta, error) {
	if value == nil {
		return ObjTypes[ObjTypeNullTypeValue].DefaultObjMeta(), nil
	}
	switch value.(type) {
	case bool:
		return ObjTypes[ObjTypeTinyIntTypeValue].DefaultObjMeta(), nil
	case int8:
		return ObjTypes[ObjTypeTinyIntTypeValue].DefaultObjMeta(), nil
	case uint8:
		return ObjTypes[ObjTypeUTinyIntTypeValue].DefaultObjMeta(), nil
	case int16:
		return ObjTypes[ObjTypeSmallIntTypeValue].DefaultObjMeta(), nil
	case uint16:
		return ObjTypes[ObjTypeUSmallIntTypeValue].DefaultObjMeta(), nil
	case int32:
		return ObjTypes[ObjTypeInt32TypeValue].DefaultObjMeta(), nil
	case uint32:
		return ObjTypes[ObjTypeUInt32TypeValue].DefaultObjMeta(), nil
	case int64:
		return ObjTypes[ObjTypeInt64TypeValue].DefaultObjMeta(), nil
	case uint64:
		return ObjTypes[ObjTypeUInt64TypeValue].DefaultObjMeta(), nil
	case float32:
		return ObjTypes[ObjTypeFloatTypeValue].DefaultObjMeta(), nil
	case float64:
		return ObjTypes[ObjTypeDoubleTypeValue].DefaultObjMeta(), nil
	case string:
		return ObjTypes[ObjTypeVarcharTypeValue].DefaultObjMeta(), nil
	case []byte:
		return ObjTypes[ObjTypeVarcharTypeValue].DefaultObjMeta(), nil
	case time.Duration:
		return ObjTypes[ObjTypeDateTimeTypeValue].DefaultObjMeta(), nil
	default:
		return nil, errors.Errorf("not match objmeta, value: %v", value)
	}
}

func NewObjType(value ObjTypeValue) (ObjType, error) {
	switch value {
	case ObjTypeNullTypeValue:
		return &NullType{value: value}, nil
	case ObjTypeTinyIntTypeValue:
		return &TinyIntType{value: value}, nil
	case ObjTypeSmallIntTypeValue:
		return &SmallIntType{value: value}, nil
	case ObjTypeMediumIntTypeValue:
		return &MediumIntType{value: value}, nil
	case ObjTypeInt32TypeValue:
		return &Int32Type{value: value}, nil
	case ObjTypeInt64TypeValue:
		return &Int64Type{value: value}, nil
	case ObjTypeUTinyIntTypeValue:
		return &UTinyIntType{value: value}, nil
	case ObjTypeUSmallIntTypeValue:
		return &USmallIntType{value: value}, nil
	case ObjTypeUMediumIntTypeValue:
		return &UMediumIntType{value: value}, nil
	case ObjTypeUInt32TypeValue:
		return &UInt32Type{value: value}, nil
	case ObjTypeUInt64TypeValue:
		return &UInt64Type{value: value}, nil
	case ObjTypeFloatTypeValue:
		return &FloatType{value: value}, nil
	case ObjTypeDoubleTypeValue:
		return &DoubleType{value: value}, nil
	case ObjTypeUFloatTypeValue:
		return &UFloatType{value: value}, nil
	case ObjTypeUDoubleTypeValue:
		return &UDoubleType{value: value}, nil
	case ObjTypeNumberTypeValue:
		return &NumberType{value: value}, nil
	case ObjTypeUNumberTypeValue:
		return &UNumberType{value: value}, nil
	case ObjTypeDateTimeTypeValue:
		return &DateTimeType{value: value}, nil
	case ObjTypeTimestampTypeValue:
		return &TimestampType{value: value}, nil
	case ObjTypeDateTypeValue:
		return &DateType{value: value}, nil
	case ObjTypeTimeTypeValue:
		return &TimeType{value: value}, nil
	case ObjTypeYearTypeValue:
		return &YearType{value: value}, nil
	case ObjTypeVarcharTypeValue:
		return &VarcharType{value: value}, nil
	case ObjTypeCharTypeValue:
		return &CharType{value: value}, nil
	case ObjTypeHexStringTypeValue:
		return &HexStringType{value: value}, nil
	case ObjTypeExtendTypeValue:
		return &ExtendType{value: value}, nil
	case ObjTypeUnknownTypeValue:
		return &UnknownType{value: value}, nil
	case ObjTypeTinyTextTypeValue:
		return &TinyTextType{value: value}, nil
	case ObjTypeTextTypeValue:
		return &TextType{value: value}, nil
	case ObjTypeMediumTextTypeValue:
		return &MediumTextType{value: value}, nil
	case ObjTypeLongTextTypeValue:
		return &LongTextType{value: value}, nil
	case ObjTypeBitTypeValue:
		return &BitType{value: value}, nil
	default:
		return nil, errors.Errorf("not match objtype, value: %d", value)
	}
}

type ObjTypeValue uint8

var ObjTypes = []ObjType{
	ObjTypeNullTypeValue:       &NullType{value: ObjTypeNullTypeValue},
	ObjTypeTinyIntTypeValue:    &TinyIntType{value: ObjTypeTinyIntTypeValue},
	ObjTypeSmallIntTypeValue:   &SmallIntType{value: ObjTypeSmallIntTypeValue},
	ObjTypeMediumIntTypeValue:  &MediumIntType{value: ObjTypeMediumIntTypeValue},
	ObjTypeInt32TypeValue:      &Int32Type{value: ObjTypeInt32TypeValue},
	ObjTypeInt64TypeValue:      &Int64Type{value: ObjTypeInt64TypeValue},
	ObjTypeUTinyIntTypeValue:   &UTinyIntType{value: ObjTypeUTinyIntTypeValue},
	ObjTypeUSmallIntTypeValue:  &USmallIntType{value: ObjTypeUSmallIntTypeValue},
	ObjTypeUMediumIntTypeValue: &UMediumIntType{value: ObjTypeUMediumIntTypeValue},
	ObjTypeUInt32TypeValue:     &UInt32Type{value: ObjTypeUInt32TypeValue},
	ObjTypeUInt64TypeValue:     &UInt64Type{value: ObjTypeUInt64TypeValue},
	ObjTypeFloatTypeValue:      &FloatType{value: ObjTypeFloatTypeValue},
	ObjTypeDoubleTypeValue:     &DoubleType{value: ObjTypeDoubleTypeValue},
	ObjTypeUFloatTypeValue:     &UFloatType{value: ObjTypeUFloatTypeValue},
	ObjTypeUDoubleTypeValue:    &UDoubleType{value: ObjTypeUDoubleTypeValue},
	ObjTypeNumberTypeValue:     &NumberType{value: ObjTypeNumberTypeValue},
	ObjTypeUNumberTypeValue:    &UNumberType{value: ObjTypeUNumberTypeValue},
	ObjTypeDateTimeTypeValue:   &DateTimeType{value: ObjTypeDateTimeTypeValue},
	ObjTypeTimestampTypeValue:  &TimestampType{value: ObjTypeTimestampTypeValue},
	ObjTypeDateTypeValue:       &DateType{value: ObjTypeDateTypeValue},
	ObjTypeTimeTypeValue:       &TimeType{value: ObjTypeTimeTypeValue},
	ObjTypeYearTypeValue:       &YearType{value: ObjTypeYearTypeValue},
	ObjTypeVarcharTypeValue:    &VarcharType{value: ObjTypeVarcharTypeValue},
	ObjTypeCharTypeValue:       &CharType{value: ObjTypeCharTypeValue},
	ObjTypeHexStringTypeValue:  &HexStringType{value: ObjTypeHexStringTypeValue},
	ObjTypeExtendTypeValue:     &ExtendType{value: ObjTypeExtendTypeValue},
	ObjTypeUnknownTypeValue:    &UnknownType{value: ObjTypeUnknownTypeValue},
	ObjTypeTinyTextTypeValue:   &TinyTextType{value: ObjTypeTinyTextTypeValue},
	ObjTypeTextTypeValue:       &TextType{value: ObjTypeTextTypeValue},
	ObjTypeMediumTextTypeValue: &MediumTextType{value: ObjTypeMediumTextTypeValue},
	ObjTypeLongTextTypeValue:   &LongTextType{value: ObjTypeLongTextTypeValue},
	ObjTypeBitTypeValue:        &BitType{value: ObjTypeBitTypeValue},
}

func (v ObjTypeValue) ValueOf() ObjType {
	return ObjTypes[v]
}

const (
	ObjTypeNullTypeValue ObjTypeValue = iota
	ObjTypeTinyIntTypeValue
	ObjTypeSmallIntTypeValue
	ObjTypeMediumIntTypeValue
	ObjTypeInt32TypeValue
	ObjTypeInt64TypeValue
	ObjTypeUTinyIntTypeValue
	ObjTypeUSmallIntTypeValue
	ObjTypeUMediumIntTypeValue
	ObjTypeUInt32TypeValue
	ObjTypeUInt64TypeValue
	ObjTypeFloatTypeValue
	ObjTypeDoubleTypeValue
	ObjTypeUFloatTypeValue
	ObjTypeUDoubleTypeValue
	ObjTypeNumberTypeValue
	ObjTypeUNumberTypeValue
	ObjTypeDateTimeTypeValue
	ObjTypeTimestampTypeValue
	ObjTypeDateTypeValue
	ObjTypeTimeTypeValue
	ObjTypeYearTypeValue
	ObjTypeVarcharTypeValue
	ObjTypeCharTypeValue
	ObjTypeHexStringTypeValue
	ObjTypeExtendTypeValue
	ObjTypeUnknownTypeValue
	ObjTypeTinyTextTypeValue
	ObjTypeTextTypeValue
	ObjTypeMediumTextTypeValue
	ObjTypeLongTextTypeValue
	ObjTypeBitTypeValue
)

type CollationLevel uint8

const (
	CollationLevelExplicit CollationLevel = iota
	CollationLevelNone
	CollationLevelImplicit
	CollationLevelSysConst
	CollationLevelCoercible
	CollationLevelNumeric
	CollationLevelIgnorable
	CollationLevelInvalid CollationLevel = 127
)

func (l CollationLevel) String() string {
	switch l {
	case CollationLevelExplicit:
		return "CollationLevelExplicit"
	case CollationLevelNone:
		return "CollationLevelNone"
	case CollationLevelImplicit:
		return "CollationLevelImplicit"
	case CollationLevelSysConst:
		return "CollationLevelSysConst"
	case CollationLevelCoercible:
		return "CollationLevelCoercible"
	case CollationLevelNumeric:
		return "CollationLevelNumeric"
	case CollationLevelIgnorable:
		return "CollationLevelIgnorable"
	default:
		return "CollationLevelInvalid"
	}
}

type CollationType uint8

const (
	CollationTypeInvalid          CollationType = 0
	CollationTypeUtf8mb4GeneralCi CollationType = 45
	CollationTypeUtf8mb4Bin       CollationType = 46
	CollationTypeBinary           CollationType = 63
	CollationTypeCollationFree    CollationType = 100
	CollationTypeMax              CollationType = 101
)

func (t CollationType) String() string {
	switch t {
	case CollationTypeUtf8mb4GeneralCi:
		return "CollationTypeUtf8mb4GeneralCi"
	case CollationTypeUtf8mb4Bin:
		return "CollationTypeUtf8mb4Bin"
	case CollationTypeBinary:
		return "CollationTypeBinary"
	case CollationTypeCollationFree:
		return "CollationTypeCollationFree"
	case CollationTypeMax:
		return "CollationTypeMax"
	default:
		return "CollationTypeInvalid"
	}
}

type NullType struct {
	value ObjTypeValue
}

func (t *NullType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *NullType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *NullType) EncodedLength(value interface{}) int {
	return 0
}

func (t *NullType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelIgnorable, CollationTypeBinary, 10)
}

func (t *NullType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("null type can not parse to comparable value")
}

func (t *NullType) Value() ObjTypeValue {
	return t.value
}

func (t *NullType) String() string {
	return "ObjType{" +
		"type:" + "NullType" +
		"}"
}

type TinyIntType struct {
	value ObjTypeValue
}

func (t *TinyIntType) Encode(buffer *bytes.Buffer, value interface{}) {
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

func (t *TinyIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.Uint8(buffer)
}

func (t *TinyIntType) EncodedLength(value interface{}) int {
	return 1
}

func (t *TinyIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *TinyIntType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case int8:
		return v, nil
	default:
		return nil, errors.Errorf("tiny int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *TinyIntType) Value() ObjTypeValue {
	return t.value
}

func (t *TinyIntType) String() string {
	return "ObjType{" +
		"type:" + "TinyIntType" +
		"}"
}

type SmallIntType struct {
	value ObjTypeValue
}

func (t *SmallIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(int16)))
}

func (t *SmallIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *SmallIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(int16)))
}

func (t *SmallIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *SmallIntType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(int16); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("small int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *SmallIntType) Value() ObjTypeValue {
	return t.value
}

func (t *SmallIntType) String() string {
	return "ObjType{" +
		"type:" + "SmallIntType" +
		"}"
}

type MediumIntType struct { // TODO not support
	value ObjTypeValue
}

func (t *MediumIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, value.(int32))
}

func (t *MediumIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *MediumIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(value.(int32))
}

func (t *MediumIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *MediumIntType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("medium int type is not support")
}

func (t *MediumIntType) Value() ObjTypeValue {
	return t.value
}

func (t *MediumIntType) String() string {
	return "ObjType{" +
		"type:" + "MediumIntType" +
		"}"
}

type Int32Type struct {
	value ObjTypeValue
}

func (t *Int32Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, value.(int32))
}

func (t *Int32Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *Int32Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(value.(int32))
}

func (t *Int32Type) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *Int32Type) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(int32); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("int32 type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *Int32Type) Value() ObjTypeValue {
	return t.value
}

func (t *Int32Type) String() string {
	return "ObjType{" +
		"type:" + "Int32Type" +
		"}"
}

type Int64Type struct {
	value ObjTypeValue
}

func (t *Int64Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t *Int64Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *Int64Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t *Int64Type) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *Int64Type) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(int64); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("int64 type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *Int64Type) Value() ObjTypeValue {
	return t.value
}

func (t *Int64Type) String() string {
	return "ObjType{" +
		"type:" + "Int64Type" +
		"}"
}

type UTinyIntType struct {
	value ObjTypeValue
}

func (t *UTinyIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.PutUint8(buffer, value.(uint8))
}

func (t *UTinyIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.Uint8(buffer)
}

func (t *UTinyIntType) EncodedLength(value interface{}) int {
	return 1
}

func (t *UTinyIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UTinyIntType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(uint8); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uTiny int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *UTinyIntType) Value() ObjTypeValue {
	return t.value
}

func (t *UTinyIntType) String() string {
	return "ObjType{" +
		"type:" + "TinyIntType" +
		"}"
}

type USmallIntType struct {
	value ObjTypeValue
}

func (t *USmallIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint16)))
}

func (t *USmallIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *USmallIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint16)))
}

func (t *USmallIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *USmallIntType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(uint16); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uSmall int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *USmallIntType) Value() ObjTypeValue {
	return t.value
}

func (t *USmallIntType) String() string {
	return "ObjType{" +
		"type:" + "USmallIntType" +
		"}"
}

type UMediumIntType struct { // TODO not support
	value ObjTypeValue
}

func (t *UMediumIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint32)))
}

func (t *UMediumIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *UMediumIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint32)))
}

func (t *UMediumIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UMediumIntType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("uMedium int type is not support")
}

func (t *UMediumIntType) Value() ObjTypeValue {
	return t.value
}

func (t *UMediumIntType) String() string {
	return "ObjType{" +
		"type:" + "UMediumIntType" +
		"}"
}

type UInt32Type struct {
	value ObjTypeValue
}

func (t *UInt32Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint32)))
}

func (t *UInt32Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t *UInt32Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint32)))
}

func (t *UInt32Type) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UInt32Type) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(uint32); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uInt int type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *UInt32Type) Value() ObjTypeValue {
	return t.value
}

func (t *UInt32Type) String() string {
	return "ObjType{" +
		"type:" + "UInt32Type" +
		"}"
}

type UInt64Type struct {
	value ObjTypeValue
}

func (t *UInt64Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(uint64)))
}

func (t *UInt64Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *UInt64Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(uint64)))
}

func (t *UInt64Type) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UInt64Type) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(uint64); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("uInt64 type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *UInt64Type) Value() ObjTypeValue {
	return t.value
}

func (t *UInt64Type) String() string {
	return "ObjType{" +
		"type:" + "UInt64Type" +
		"}"
}

type FloatType struct {
	value ObjTypeValue
}

func (t *FloatType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVf32(buffer, value.(float32))
}

func (t *FloatType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVf32(buffer)
}

func (t *FloatType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVf32(value.(float32))
}

func (t *FloatType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *FloatType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(float32); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("float type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *FloatType) Value() ObjTypeValue {
	return t.value
}

func (t *FloatType) String() string {
	return "ObjType{" +
		"type:" + "FloatType" +
		"}"
}

type DoubleType struct {
	value ObjTypeValue
}

func (t *DoubleType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVf64(buffer, value.(float64))
}

func (t *DoubleType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVf64(buffer)
}

func (t *DoubleType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVf64(value.(float64))
}

func (t *DoubleType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *DoubleType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(float64); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("double type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *DoubleType) Value() ObjTypeValue {
	return t.value
}

func (t *DoubleType) String() string {
	return "ObjType{" +
		"type:" + "DoubleType" +
		"}"
}

type UFloatType struct { // TODO not support
	value ObjTypeValue
}

func (t *UFloatType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *UFloatType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *UFloatType) EncodedLength(value interface{}) int {
	return 0
}

func (t *UFloatType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UFloatType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("uFloat type is not support")
}

func (t *UFloatType) Value() ObjTypeValue {
	return t.value
}

func (t *UFloatType) String() string {
	return "ObjType{" +
		"type:" + "UFloatType" +
		"}"
}

type UDoubleType struct { // TODO not support
	value ObjTypeValue
}

func (t *UDoubleType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *UDoubleType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *UDoubleType) EncodedLength(value interface{}) int {
	return 0
}

func (t *UDoubleType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UDoubleType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("uDouble type is not support")
}

func (t *UDoubleType) Value() ObjTypeValue {
	return t.value
}

func (t *UDoubleType) String() string {
	return "ObjType{" +
		"type:" + "UDoubleType" +
		"}"
}

type NumberType struct { // TODO not support
	value ObjTypeValue
}

func (t *NumberType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *NumberType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *NumberType) EncodedLength(value interface{}) int {
	return 0
}

func (t *NumberType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *NumberType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("number type is not support")
}

func (t *NumberType) Value() ObjTypeValue {
	return t.value
}

func (t *NumberType) String() string {
	return "ObjType{" +
		"type:" + "NumberType" +
		"}"
}

type UNumberType struct { // TODO not support
	value ObjTypeValue
}

func (t *UNumberType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *UNumberType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *UNumberType) EncodedLength(value interface{}) int {
	return 0
}

func (t *UNumberType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UNumberType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("uNumber type is not support")
}

func (t *UNumberType) Value() ObjTypeValue {
	return t.value
}

func (t *UNumberType) String() string {
	return "ObjType{" +
		"type:" + "UNumberType" +
		"}"
}

type DateTimeType struct {
	value ObjTypeValue
}

func (t *DateTimeType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration))) // todo
}

func (t *DateTimeType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *DateTimeType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *DateTimeType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *DateTimeType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("date time type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *DateTimeType) Value() ObjTypeValue {
	return t.value
}

func (t *DateTimeType) String() string {
	return "ObjType{" +
		"type:" + "DateTimeType" +
		"}"
}

type TimestampType struct {
	value ObjTypeValue
}

func (t *TimestampType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration))) // todo
}

func (t *TimestampType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *TimestampType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *TimestampType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *TimestampType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("time stamp type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *TimestampType) Value() ObjTypeValue {
	return t.value
}

func (t *TimestampType) String() string {
	return "ObjType{" +
		"type:" + "TimestampType" +
		"}"
}

type DateType struct {
	value ObjTypeValue
}

func (t *DateType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration))) // todo
}

func (t *DateType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *DateType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *DateType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *DateType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("data type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *DateType) Value() ObjTypeValue {
	return t.value
}

func (t *DateType) String() string {
	return "ObjType{" +
		"type:" + "DateType" +
		"}"
}

type TimeType struct {
	value ObjTypeValue
}

func (t *TimeType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration))) // todo
}

func (t *TimeType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *TimeType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t *TimeType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *TimeType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if v, ok := value.(time.Duration); ok {
		return v, nil
	} else {
		return nil, errors.Errorf("time type parse to comparable failed, not match value, value: %v", v)
	}
}

func (t *TimeType) Value() ObjTypeValue {
	return t.value
}

func (t *TimeType) String() string {
	return "ObjType{" +
		"type:" + "TimeType" +
		"}"
}

type YearType struct { // TODO not support
	value ObjTypeValue
}

func (t *YearType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.PutUint8(buffer, value.(uint8))
}

func (t *YearType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.Uint8(buffer)
}

func (t *YearType) EncodedLength(value interface{}) int {
	return 1
}

func (t *YearType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *YearType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("year type is not support")
}

func (t *YearType) Value() ObjTypeValue {
	return t.value
}

func (t *YearType) String() string {
	return "ObjType{" +
		"type:" + "YearType" +
		"}"
}

type VarcharType struct {
	value ObjTypeValue
}

func (t *VarcharType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *VarcharType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return DecodeText(buffer, csType)
}

func (t *VarcharType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *VarcharType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelExplicit, CollationTypeUtf8mb4GeneralCi, 10)
}

func (t *VarcharType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if CollationTypeBinary == csType {
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

func (t *VarcharType) Value() ObjTypeValue {
	return t.value
}

func (t *VarcharType) String() string {
	return "ObjType{" +
		"type:" + "VarcharType" +
		"}"
}

type CharType struct {
	value ObjTypeValue
}

func (t *CharType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *CharType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return DecodeText(buffer, csType)
}

func (t *CharType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *CharType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelExplicit, CollationTypeUtf8mb4GeneralCi, 10)
}

func (t *CharType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if CollationTypeBinary == csType {
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

func (t *CharType) Value() ObjTypeValue {
	return t.value
}

func (t *CharType) String() string {
	return "ObjType{" +
		"type:" + "CharType" +
		"}"
}

type HexStringType struct { // TODO not support
	value ObjTypeValue
}

func (t *HexStringType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *HexStringType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *HexStringType) EncodedLength(value interface{}) int {
	return 0
}

func (t *HexStringType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *HexStringType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("hex string type is not support")
}

func (t *HexStringType) Value() ObjTypeValue {
	return t.value
}

func (t *HexStringType) String() string {
	return "ObjType{" +
		"type:" + "HexStringType" +
		"}"
}

type ExtendType struct { // TODO not support
	value ObjTypeValue
}

func (t *ExtendType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t *ExtendType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *ExtendType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t *ExtendType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *ExtendType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("extend type is not support")
}

func (t *ExtendType) Value() ObjTypeValue {
	return t.value
}

func (t *ExtendType) String() string {
	return "ObjType{" +
		"type:" + "ExtendType" +
		"}"
}

type UnknownType struct { // TODO not support
	value ObjTypeValue
}

func (t *UnknownType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t *UnknownType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *UnknownType) EncodedLength(value interface{}) int {
	return 0
}

func (t *UnknownType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *UnknownType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("unknown type is not support")
}

func (t *UnknownType) Value() ObjTypeValue {
	return t.value
}

func (t *UnknownType) String() string {
	return "ObjType{" +
		"type:" + "UnknownType" +
		"}"
}

type TinyTextType struct {
	value ObjTypeValue
}

func (t *TinyTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *TinyTextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return DecodeText(buffer, csType)
}

func (t *TinyTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *TinyTextType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *TinyTextType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if CollationTypeBinary == csType {
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

func (t *TinyTextType) Value() ObjTypeValue {
	return t.value
}

func (t *TinyTextType) String() string {
	return "ObjType{" +
		"type:" + "TinyTextType" +
		"}"
}

type TextType struct {
	value ObjTypeValue
}

func (t *TextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *TextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return DecodeText(buffer, csType)
}

func (t *TextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *TextType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *TextType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if CollationTypeBinary == csType {
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

func (t *TextType) Value() ObjTypeValue {
	return t.value
}

func (t *TextType) String() string {
	return "ObjType{" +
		"type:" + "TextType" +
		"}"
}

type MediumTextType struct {
	value ObjTypeValue
}

func (t *MediumTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *MediumTextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return DecodeText(buffer, csType)
}

func (t *MediumTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *MediumTextType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *MediumTextType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if CollationTypeBinary == csType {
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

func (t *MediumTextType) Value() ObjTypeValue {
	return t.value
}

func (t *MediumTextType) String() string {
	return "ObjType{" +
		"type:" + "MediumTextType" +
		"}"
}

type LongTextType struct {
	value ObjTypeValue
}

func (t *LongTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t *LongTextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return DecodeText(buffer, csType)
}

func (t *LongTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t *LongTextType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *LongTextType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	if CollationTypeBinary == csType {
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

func (t *LongTextType) Value() ObjTypeValue {
	return t.value
}

func (t *LongTextType) String() string {
	return "ObjType{" +
		"type:" + "LongTextType" +
		"}"
}

type BitType struct { // TODO not support
	value ObjTypeValue
}

func (t *BitType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t *BitType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t *BitType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t *BitType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMetaWithParams(t, CollationLevelNumeric, CollationTypeBinary, 10)
}

func (t *BitType) CheckTypeForValue(value interface{}, csType CollationType) (interface{}, error) {
	return nil, errors.New("bit type is not support")
}

func (t *BitType) Value() ObjTypeValue {
	return t.value
}

func (t *BitType) String() string {
	return "ObjType{" +
		"type:" + "BitType" +
		"}"
}

func EncodeText(buffer *bytes.Buffer, value interface{}) {
	switch v := value.(type) {
	case string:
		util.EncodeVString(buffer, v)
	case []byte:
		util.EncodeBytesString(buffer, v)
	default:
		// todo do nothing
		return
	}
}

func DecodeText(buffer *bytes.Buffer, csType CollationType) interface{} {
	if CollationTypeBinary == csType {
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
		// todo do nothing
		return 0
	}
}

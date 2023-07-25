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

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObObjectMeta struct {
	objType        ObObjType
	collationLevel ObCollationLevel
	collationType  ObCollationType
	scale          int8
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

func (m *ObObjectMeta) Scale() int8 {
	return m.scale
}

func (m *ObObjectMeta) SetScale(scale int8) {
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
	util.PutUint8(buffer, uint8(m.scale))
}

func (m *ObObjectMeta) Decode(buffer *bytes.Buffer) {
	m.objType = ObObjTypeValue(util.Uint8(buffer)).ValueOf()
	m.collationLevel = ObCollationLevel(util.Uint8(buffer))
	m.collationType = ObCollationType(util.Uint8(buffer))
	m.scale = int8(util.Uint8(buffer))
}

func (m *ObObjectMeta) EncodedLength() int {
	return 4 // objType collationLevel collationType scale
}

type ObObjType interface {
	Encode(buffer *bytes.Buffer, value interface{})
	Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{}
	EncodedLength(value interface{}) int
	DefaultObjMeta() ObObjectMeta
	Value() ObObjTypeValue
	String() string
}

func DefaultObjMeta(value interface{}) (ObObjectMeta, error) {
	if value == nil {
		return ObObjTypeNull.DefaultObjMeta(), nil
	}
	switch value.(type) {
	case bool:
		return ObObjTypeTinyInt.DefaultObjMeta(), nil
	case int8:
		return ObObjTypeTinyInt.DefaultObjMeta(), nil
	case uint8:
		return ObObjTypeUTinyInt.DefaultObjMeta(), nil
	case int16:
		return ObObjTypeSmallInt.DefaultObjMeta(), nil
	case uint16:
		return ObObjTypeUSmallInt.DefaultObjMeta(), nil
	case int32:
		return ObObjTypeInt32.DefaultObjMeta(), nil
	case uint32:
		return ObObjTypeUInt32.DefaultObjMeta(), nil
	case int64:
		return ObObjTypeInt64.DefaultObjMeta(), nil
	case uint64:
		return ObObjTypeUInt64.DefaultObjMeta(), nil
	case float32:
		return ObObjTypeFloat.DefaultObjMeta(), nil
	case float64:
		return ObObjTypeDouble.DefaultObjMeta(), nil
	case string:
		return ObObjTypeVarchar.DefaultObjMeta(), nil
	case []byte:
		return ObObjTypeVarchar.DefaultObjMeta(), nil
	case table.DateTime:
		return ObObjTypeDateTime.DefaultObjMeta(), nil
	case table.TimeStamp:
		return ObObjTypeTimestamp.DefaultObjMeta(), nil
	case table.Extremum:
		return ObObjTypeExtend.DefaultObjMeta(), nil
	default:
		return ObObjectMeta{}, errors.Errorf("not match objmeta, value: %v", value)
	}
}

func NewObjType(value ObObjTypeValue) (ObObjType, error) {
	switch value {
	case ObObjTypeNullTypeValue:
		return ObNullType(ObObjTypeNullTypeValue), nil
	case ObObjTypeTinyIntTypeValue:
		return ObTinyIntType(ObObjTypeTinyIntTypeValue), nil
	case ObObjTypeSmallIntTypeValue:
		return ObSmallIntType(ObObjTypeSmallIntTypeValue), nil
	case ObObjTypeMediumIntTypeValue:
		return ObMediumIntType(ObObjTypeMediumIntTypeValue), nil
	case ObObjTypeInt32TypeValue:
		return ObInt32Type(ObObjTypeInt32TypeValue), nil
	case ObObjTypeInt64TypeValue:
		return ObInt64Type(ObObjTypeInt64TypeValue), nil
	case ObObjTypeUTinyIntTypeValue:
		return ObUTinyIntType(ObObjTypeUTinyIntTypeValue), nil
	case ObObjTypeUSmallIntTypeValue:
		return ObUSmallIntType(ObObjTypeUSmallIntTypeValue), nil
	case ObObjTypeUMediumIntTypeValue:
		return ObUMediumIntType(ObObjTypeUMediumIntTypeValue), nil
	case ObObjTypeUInt32TypeValue:
		return ObUInt32Type(ObObjTypeUInt32TypeValue), nil
	case ObObjTypeUInt64TypeValue:
		return ObUInt64Type(ObObjTypeUInt64TypeValue), nil
	case ObObjTypeFloatTypeValue:
		return ObFloatType(ObObjTypeFloatTypeValue), nil
	case ObObjTypeDoubleTypeValue:
		return ObDoubleType(ObObjTypeDoubleTypeValue), nil
	case ObObjTypeUFloatTypeValue:
		return ObUFloatType(ObObjTypeUFloatTypeValue), nil
	case ObObjTypeUDoubleTypeValue:
		return ObUDoubleType(ObObjTypeUDoubleTypeValue), nil
	case ObObjTypeNumberTypeValue:
		return ObNumberType(ObObjTypeNumberTypeValue), nil
	case ObObjTypeUNumberTypeValue:
		return ObUNumberType(ObObjTypeUNumberTypeValue), nil
	case ObObjTypeDateTimeTypeValue:
		return ObDateTimeType(ObObjTypeDateTimeTypeValue), nil
	case ObObjTypeTimestampTypeValue:
		return ObTimestampType(ObObjTypeTimestampTypeValue), nil
	case ObObjTypeDateTypeValue:
		return ObDateType(ObObjTypeDateTypeValue), nil
	case ObObjTypeTimeTypeValue:
		return ObTimeType(ObObjTypeTimeTypeValue), nil
	case ObObjTypeYearTypeValue:
		return ObYearType(ObObjTypeYearTypeValue), nil
	case ObObjTypeVarcharTypeValue:
		return ObVarcharType(ObObjTypeVarcharTypeValue), nil
	case ObObjTypeCharTypeValue:
		return ObCharType(ObObjTypeCharTypeValue), nil
	case ObObjTypeHexStringTypeValue:
		return ObHexStringType(ObObjTypeHexStringTypeValue), nil
	case ObObjTypeExtendTypeValue:
		return ObExtendType(ObObjTypeExtendTypeValue), nil
	case ObObjTypeUnknownTypeValue:
		return ObUnknownType(ObObjTypeUnknownTypeValue), nil
	case ObObjTypeTinyTextTypeValue:
		return ObTinyTextType(ObObjTypeTinyTextTypeValue), nil
	case ObObjTypeTextTypeValue:
		return ObTextType(ObObjTypeTextTypeValue), nil
	case ObObjTypeMediumTextTypeValue:
		return ObMediumTextType(ObObjTypeMediumTextTypeValue), nil
	case ObObjTypeLongTextTypeValue:
		return ObLongTextType(ObObjTypeLongTextTypeValue), nil
	case ObObjTypeBitTypeValue:
		return ObBitType(ObObjTypeBitTypeValue), nil
	default:
		return nil, errors.Errorf("not match objtype, value: %d", value)
	}
}

type ObObjTypeValue uint8

var (
	ObObjTypeNull       = ObNullType(ObObjTypeNullTypeValue)
	ObObjTypeTinyInt    = ObTinyIntType(ObObjTypeTinyIntTypeValue)
	ObObjTypeSmallInt   = ObSmallIntType(ObObjTypeSmallIntTypeValue)
	ObObjTypeMediumInt  = ObMediumIntType(ObObjTypeMediumIntTypeValue)
	ObObjTypeInt32      = ObInt32Type(ObObjTypeInt32TypeValue)
	ObObjTypeInt64      = ObInt64Type(ObObjTypeInt64TypeValue)
	ObObjTypeUTinyInt   = ObUTinyIntType(ObObjTypeUTinyIntTypeValue)
	ObObjTypeUSmallInt  = ObUSmallIntType(ObObjTypeUSmallIntTypeValue)
	ObObjTypeUMediumInt = ObUMediumIntType(ObObjTypeUMediumIntTypeValue)
	ObObjTypeUInt32     = ObUInt32Type(ObObjTypeUInt32TypeValue)
	ObObjTypeUInt64     = ObUInt64Type(ObObjTypeUInt64TypeValue)
	ObObjTypeFloat      = ObFloatType(ObObjTypeFloatTypeValue)
	ObObjTypeDouble     = ObDoubleType(ObObjTypeDoubleTypeValue)
	ObObjTypeUFloat     = ObUFloatType(ObObjTypeUFloatTypeValue)
	ObObjTypeUDouble    = ObUDoubleType(ObObjTypeUDoubleTypeValue)
	ObObjTypeNumber     = ObNumberType(ObObjTypeNumberTypeValue)
	ObObjTypeUNumber    = ObUNumberType(ObObjTypeUNumberTypeValue)
	ObObjTypeDateTime   = ObDateTimeType(ObObjTypeDateTimeTypeValue)
	ObObjTypeTimestamp  = ObTimestampType(ObObjTypeTimestampTypeValue)
	ObObjTypeDate       = ObDateType(ObObjTypeDateTypeValue)
	ObObjTypeTime       = ObTimeType(ObObjTypeTimeTypeValue)
	ObObjTypeYear       = ObYearType(ObObjTypeYearTypeValue)
	ObObjTypeVarchar    = ObVarcharType(ObObjTypeVarcharTypeValue)
	ObObjTypeChar       = ObCharType(ObObjTypeCharTypeValue)
	ObObjTypeHexString  = ObHexStringType(ObObjTypeHexStringTypeValue)
	ObObjTypeExtend     = ObExtendType(ObObjTypeExtendTypeValue)
	ObObjTypeUnknown    = ObUnknownType(ObObjTypeUnknownTypeValue)
	ObObjTypeTinyText   = ObTinyTextType(ObObjTypeTinyTextTypeValue)
	ObObjTypeText       = ObTextType(ObObjTypeTextTypeValue)
	ObObjTypeMediumText = ObMediumTextType(ObObjTypeMediumTextTypeValue)
	ObObjTypeLongText   = ObLongTextType(ObObjTypeLongTextTypeValue)
	ObObjTypeBit        = ObBitType(ObObjTypeBitTypeValue)
)

var ObObjTypes = []ObObjType{
	ObObjTypeNullTypeValue:       ObObjTypeNull,
	ObObjTypeTinyIntTypeValue:    ObObjTypeTinyInt,
	ObObjTypeSmallIntTypeValue:   ObObjTypeSmallInt,
	ObObjTypeMediumIntTypeValue:  ObObjTypeMediumInt,
	ObObjTypeInt32TypeValue:      ObObjTypeInt32,
	ObObjTypeInt64TypeValue:      ObObjTypeInt64,
	ObObjTypeUTinyIntTypeValue:   ObObjTypeUTinyInt,
	ObObjTypeUSmallIntTypeValue:  ObObjTypeUSmallInt,
	ObObjTypeUMediumIntTypeValue: ObObjTypeUMediumInt,
	ObObjTypeUInt32TypeValue:     ObObjTypeUInt32,
	ObObjTypeUInt64TypeValue:     ObObjTypeUInt64,
	ObObjTypeFloatTypeValue:      ObObjTypeFloat,
	ObObjTypeDoubleTypeValue:     ObObjTypeDouble,
	ObObjTypeUFloatTypeValue:     ObObjTypeUFloat,
	ObObjTypeUDoubleTypeValue:    ObObjTypeUDouble,
	ObObjTypeNumberTypeValue:     ObObjTypeNumber,
	ObObjTypeUNumberTypeValue:    ObObjTypeUNumber,
	ObObjTypeDateTimeTypeValue:   ObObjTypeDateTime,
	ObObjTypeTimestampTypeValue:  ObObjTypeTimestamp,
	ObObjTypeDateTypeValue:       ObObjTypeDate,
	ObObjTypeTimeTypeValue:       ObObjTypeTime,
	ObObjTypeYearTypeValue:       ObObjTypeYear,
	ObObjTypeVarcharTypeValue:    ObObjTypeVarchar,
	ObObjTypeCharTypeValue:       ObObjTypeChar,
	ObObjTypeHexStringTypeValue:  ObObjTypeHexString,
	ObObjTypeExtendTypeValue:     ObObjTypeExtend,
	ObObjTypeUnknownTypeValue:    ObObjTypeUnknown,
	ObObjTypeTinyTextTypeValue:   ObObjTypeTinyText,
	ObObjTypeTextTypeValue:       ObObjTypeText,
	ObObjTypeMediumTextTypeValue: ObObjTypeMediumText,
	ObObjTypeLongTextTypeValue:   ObObjTypeLongText,
	ObObjTypeBitTypeValue:        ObObjTypeBit,
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

type ObNullType ObObjTypeValue

func (t ObNullType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObNullType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObNullType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObNullType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelIgnorable, ObCollationTypeBinary, -1}
}

func (t ObNullType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObNullType) String() string {
	return "ObObjType{" +
		"type:" + "ObNullType" +
		"}"
}

type ObTinyIntType ObObjTypeValue

func (t ObTinyIntType) Encode(buffer *bytes.Buffer, value interface{}) {
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

func (t ObTinyIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return int8(util.Uint8(buffer))
}

func (t ObTinyIntType) EncodedLength(value interface{}) int {
	return 1
}

func (t ObTinyIntType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObTinyIntType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObTinyIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObTinyIntType" +
		"}"
}

type ObSmallIntType ObObjTypeValue

func (t ObSmallIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(int16)))
}

func (t ObSmallIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return int16(util.DecodeVi32(buffer))
}

func (t ObSmallIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(int16)))
}

func (t ObSmallIntType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObSmallIntType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObSmallIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObSmallIntType" +
		"}"
}

type ObMediumIntType ObObjTypeValue // TODO not support

func (t ObMediumIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, value.(int32))
}

func (t ObMediumIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t ObMediumIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(value.(int32))
}

func (t ObMediumIntType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObMediumIntType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObMediumIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObMediumIntType" +
		"}"
}

type ObInt32Type ObObjTypeValue

func (t ObInt32Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, value.(int32))
}

func (t ObInt32Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t ObInt32Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(value.(int32))
}

func (t ObInt32Type) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObInt32Type) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObInt32Type) String() string {
	return "ObObjType{" +
		"type:" + "ObInt32Type" +
		"}"
}

type ObInt64Type ObObjTypeValue

func (t ObInt64Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t ObInt64Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t ObInt64Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t ObInt64Type) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObInt64Type) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObInt64Type) String() string {
	return "ObObjType{" +
		"type:" + "ObInt64Type" +
		"}"
}

type ObUTinyIntType ObObjTypeValue

func (t ObUTinyIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.PutUint8(buffer, value.(uint8))
}

func (t ObUTinyIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.Uint8(buffer)
}

func (t ObUTinyIntType) EncodedLength(value interface{}) int {
	return 1
}

func (t ObUTinyIntType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUTinyIntType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUTinyIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObUTinyIntType" +
		"}"
}

type ObUSmallIntType ObObjTypeValue

func (t ObUSmallIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint16)))
}

func (t ObUSmallIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return uint16(util.DecodeVi32(buffer))
}

func (t ObUSmallIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint16)))
}

func (t ObUSmallIntType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUSmallIntType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUSmallIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObUSmallIntType" +
		"}"
}

type ObUMediumIntType ObObjTypeValue // TODO not support

func (t ObUMediumIntType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint32)))
}

func (t ObUMediumIntType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi32(buffer)
}

func (t ObUMediumIntType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint32)))
}

func (t ObUMediumIntType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUMediumIntType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUMediumIntType) String() string {
	return "ObObjType{" +
		"type:" + "ObUMediumIntType" +
		"}"
}

type ObUInt32Type ObObjTypeValue

func (t ObUInt32Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi32(buffer, int32(value.(uint32)))
}

func (t ObUInt32Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return uint32(util.DecodeVi32(buffer))
}

func (t ObUInt32Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi32(int32(value.(uint32)))
}

func (t ObUInt32Type) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUInt32Type) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUInt32Type) String() string {
	return "ObObjType{" +
		"type:" + "ObUInt32Type" +
		"}"
}

type ObUInt64Type ObObjTypeValue

func (t ObUInt64Type) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(uint64)))
}

func (t ObUInt64Type) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return uint64(util.DecodeVi64(buffer))
}

func (t ObUInt64Type) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(uint64)))
}

func (t ObUInt64Type) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUInt64Type) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUInt64Type) String() string {
	return "ObObjType{" +
		"type:" + "ObUInt64Type" +
		"}"
}

type ObFloatType ObObjTypeValue

func (t ObFloatType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVf32(buffer, value.(float32))
}

func (t ObFloatType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVf32(buffer)
}

func (t ObFloatType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVf32(value.(float32))
}

func (t ObFloatType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObFloatType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObFloatType) String() string {
	return "ObObjType{" +
		"type:" + "ObFloatType" +
		"}"
}

type ObDoubleType ObObjTypeValue

func (t ObDoubleType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVf64(buffer, value.(float64))
}

func (t ObDoubleType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVf64(buffer)
}

func (t ObDoubleType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVf64(value.(float64))
}

func (t ObDoubleType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObDoubleType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObDoubleType) String() string {
	return "ObObjType{" +
		"type:" + "ObDoubleType" +
		"}"
}

type ObUFloatType ObObjTypeValue // TODO not support

func (t ObUFloatType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObUFloatType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObUFloatType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObUFloatType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUFloatType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUFloatType) String() string {
	return "ObObjType{" +
		"type:" + "ObUFloatType" +
		"}"
}

type ObUDoubleType ObObjTypeValue // TODO not support

func (t ObUDoubleType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObUDoubleType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObUDoubleType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObUDoubleType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUDoubleType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUDoubleType) String() string {
	return "ObObjType{" +
		"type:" + "ObUDoubleType" +
		"}"
}

type ObNumberType ObObjTypeValue // TODO not support

func (t ObNumberType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObNumberType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObNumberType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObNumberType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObNumberType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObNumberType) String() string {
	return "ObObjType{" +
		"type:" + "ObNumberType" +
		"}"
}

type ObUNumberType ObObjTypeValue // TODO not support

func (t ObUNumberType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObUNumberType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObUNumberType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObUNumberType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUNumberType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUNumberType) String() string {
	return "ObObjType{" +
		"type:" + "ObUNumberType" +
		"}"
}

type ObDateTimeType ObObjTypeValue

func (t ObDateTimeType) Encode(buffer *bytes.Buffer, value interface{}) {
	v := time.Time(value.(table.DateTime))
	util.EncodeVi64(buffer, v.UnixMicro()) // store UTC
}

func (t ObDateTimeType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return time.UnixMicro(util.DecodeVi64(buffer)).In(time.UTC) // show UTC
}

func (t ObDateTimeType) EncodedLength(value interface{}) int {
	v := time.Time(value.(table.DateTime))
	return util.EncodedLengthByVi64(v.UnixMicro())
}

func (t ObDateTimeType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObDateTimeType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObDateTimeType) String() string {
	return "ObObjType{" +
		"type:" + "ObDateTimeType" +
		"}"
}

type ObTimestampType ObObjTypeValue

func (t ObTimestampType) Encode(buffer *bytes.Buffer, value interface{}) {
	v := time.Time(value.(table.TimeStamp))
	util.EncodeVi64(buffer, v.UnixMicro()) // store UTC
}

func (t ObTimestampType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return time.UnixMicro(util.DecodeVi64(buffer)) // show local
}

func (t ObTimestampType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(time.Time(value.(table.TimeStamp)).UnixMicro())
}

func (t ObTimestampType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObTimestampType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObTimestampType) String() string {
	return "ObObjType{" +
		"type:" + "ObTimestampType" +
		"}"
}

type ObDateType ObObjTypeValue // TODO not support

func (t ObDateType) Encode(buffer *bytes.Buffer, value interface{}) {
	// TODO not support
}

func (t ObDateType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
	// TODO not support
}

func (t ObDateType) EncodedLength(value interface{}) int {
	return -1
	// TODO not support
}

func (t ObDateType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObDateType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObDateType) String() string {
	return "ObObjType{" +
		"type:" + "ObDateType" +
		"}"
}

type ObTimeType ObObjTypeValue

func (t ObTimeType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(time.Duration)))
}

func (t ObTimeType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t ObTimeType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(time.Duration)))
}

func (t ObTimeType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObTimeType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObTimeType) String() string {
	return "ObObjType{" +
		"type:" + "ObTimeType" +
		"}"
}

type ObYearType ObObjTypeValue // TODO not support

func (t ObYearType) Encode(buffer *bytes.Buffer, value interface{}) {
	// TODO not support
}

func (t ObYearType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	// TODO not support
	return nil
}

func (t ObYearType) EncodedLength(value interface{}) int {
	// TODO not support
	return -1
}

func (t ObYearType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObYearType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObYearType) String() string {
	return "ObObjType{" +
		"type:" + "ObYearType" +
		"}"
}

type ObVarcharType ObObjTypeValue

func (t ObVarcharType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t ObVarcharType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t ObVarcharType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t ObVarcharType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelExplicit, ObCollationTypeUtf8mb4GeneralCi, -1}
}

func (t ObVarcharType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObVarcharType) String() string {
	return "ObObjType{" +
		"type:" + "ObVarcharType" +
		"}"
}

type ObCharType ObObjTypeValue

func (t ObCharType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t ObCharType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t ObCharType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t ObCharType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelExplicit, ObCollationTypeUtf8mb4GeneralCi, -1}
}

func (t ObCharType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObCharType) String() string {
	return "ObObjType{" +
		"type:" + "ObCharType" +
		"}"
}

type ObHexStringType ObObjTypeValue // TODO not support

func (t ObHexStringType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObHexStringType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObHexStringType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObHexStringType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObHexStringType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObHexStringType) String() string {
	return "ObObjType{" +
		"type:" + "ObHexStringType" +
		"}"
}

type ObExtendType ObObjTypeValue // TODO: Only Extremum use ExtendType now

func (t ObExtendType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, int64(value.(table.Extremum)))
}

func (t ObExtendType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t ObExtendType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(int64(value.(table.Extremum)))
}

func (t ObExtendType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObExtendType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObExtendType) String() string {
	return "ObObjType{" +
		"type:" + "ObExtendType" +
		"}"
}

type ObUnknownType ObObjTypeValue // TODO not support

func (t ObUnknownType) Encode(buffer *bytes.Buffer, value interface{}) {
	return
}

func (t ObUnknownType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return nil
}

func (t ObUnknownType) EncodedLength(value interface{}) int {
	return 0
}

func (t ObUnknownType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObUnknownType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObUnknownType) String() string {
	return "ObObjType{" +
		"type:" + "ObUnknownType" +
		"}"
}

type ObTinyTextType ObObjTypeValue

func (t ObTinyTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t ObTinyTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t ObTinyTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t ObTinyTextType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObTinyTextType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObTinyTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObTinyTextType" +
		"}"
}

type ObTextType ObObjTypeValue

func (t ObTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t ObTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t ObTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t ObTextType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObTextType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObTextType" +
		"}"
}

type ObMediumTextType ObObjTypeValue

func (t ObMediumTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t ObMediumTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t ObMediumTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t ObMediumTextType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObMediumTextType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObMediumTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObMediumTextType" +
		"}"
}

type ObLongTextType ObObjTypeValue

func (t ObLongTextType) Encode(buffer *bytes.Buffer, value interface{}) {
	EncodeText(buffer, value)
}

func (t ObLongTextType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return DecodeText(buffer, obCollationType)
}

func (t ObLongTextType) EncodedLength(value interface{}) int {
	return EncodedLengthByText(value)
}

func (t ObLongTextType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObLongTextType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObLongTextType) String() string {
	return "ObObjType{" +
		"type:" + "ObLongTextType" +
		"}"
}

type ObBitType ObObjTypeValue // TODO not support

func (t ObBitType) Encode(buffer *bytes.Buffer, value interface{}) {
	util.EncodeVi64(buffer, value.(int64))
}

func (t ObBitType) Decode(buffer *bytes.Buffer, obCollationType ObCollationType) interface{} {
	return util.DecodeVi64(buffer)
}

func (t ObBitType) EncodedLength(value interface{}) int {
	return util.EncodedLengthByVi64(value.(int64))
}

func (t ObBitType) DefaultObjMeta() ObObjectMeta {
	return ObObjectMeta{t, ObCollationLevelNumeric, ObCollationTypeBinary, -1}
}

func (t ObBitType) Value() ObObjTypeValue {
	return ObObjTypeValue(t)
}

func (t ObBitType) String() string {
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
	default: // do nothing
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
	default: // do nothing
		return 0
	}
}

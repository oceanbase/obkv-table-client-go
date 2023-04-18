package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObjectMeta struct {
	objType ObjType
	csLevel CollationLevel
	csType  CollationType
	scale   byte
}

func NewObjectMeta(objType ObjType, csLevel CollationLevel, csType CollationType, scale byte) *ObjectMeta {
	return &ObjectMeta{objType: objType, csLevel: csLevel, csType: csType, scale: scale}
}

type ObjType interface {
	Encode(buf []byte, value interface{})
	Decode(buffer *bytes.Buffer, csType CollationType) interface{}
	EncodedSize() int
	DefaultObjMeta() *ObjectMeta
	Value() ObjTypeValue
}

func NewObjType(value ObjTypeValue) ObjType {
	switch value {
	case ObjTypeNullTypeValue:
		return &NullType{
			value: value,
		}
	case ObjTypeTinyIntTypeValue:
		return &TinyIntType{
			value: value,
		}
	case ObjTypeSmallIntTypeValue:
		return &SmallIntType{
			value: value,
		}
	case ObjTypeMediumIntTypeValue:
		return &MediumIntType{
			value: value,
		}
	case ObjTypeInt32TypeValue:
		return &Int32Type{
			value: value,
		}
	case ObjTypeInt64TypeValue:
		return &Int64Type{
			value: value,
		}
	case ObjTypeUTinyIntTypeValue:
		return &UTinyIntType{
			value: value,
		}
	case ObjTypeUSmallIntTypeValue:
		return &USmallIntType{
			value: value,
		}
	case ObjTypeUMediumIntTypeValue:
		return &UMediumIntType{
			value: value,
		}
	case ObjTypeUInt32TypeValue:
		return &UInt32Type{
			value: value,
		}
	case ObjTypeUInt64TypeValue:
		return &UInt64Type{
			value: value,
		}
	case ObjTypeFloatTypeValue:
		return &FloatType{
			value: value,
		}
	case ObjTypeDoubleTypeValue:
		return &DoubleType{
			value: value,
		}
	case ObjTypeUFloatTypeValue:
		return &UFloatType{
			value: value,
		}
	case ObjTypeUDoubleTypeValue:
		return &UDoubleType{
			value: value,
		}
	case ObjTypeNumberTypeValue:
		return &NumberType{
			value: value,
		}
	case ObjTypeUNumberTypeValue:
		return &UNumberType{
			value: value,
		}
	case ObjTypeDateTimeTypeValue:
		return &DateTimeType{
			value: value,
		}
	case ObjTypeTimestampTypeValue:
		return &TimestampType{
			value: value,
		}
	case ObjTypeDateTypeValue:
		return &DateType{
			value: value,
		}
	case ObjTypeTimeTypeValue:
		return &TimeType{
			value: value,
		}
	case ObjTypeYearTypeValue:
		return &YearType{
			value: value,
		}
	case ObjTypeVarcharTypeValue:
		return &VarcharType{
			value: value,
		}
	case ObjTypeCharTypeValue:
		return &CharType{
			value: value,
		}
	case ObjTypeHexStringTypeValue:
		return &HexStringType{
			value: value,
		}
	case ObjTypeExtendTypeValue:
		return &ExtendType{
			value: value,
		}
	case ObjTypeUnknownTypeValue:
		return &UnknownType{
			value: value,
		}
	case ObjTypeTinyTextTypeValue:
		return &TinyTextType{
			value: value,
		}
	case ObjTypeTextTypeValue:
		return &TextType{
			value: value,
		}
	case ObjTypeMediumTextTypeValue:
		return &MediumTextType{
			value: value,
		}
	case ObjTypeLongTextTypeValue:
		return &LongTextType{
			value: value,
		}
	case ObjTypeBitTypeValue:
		return &BitType{
			value: value,
		}
	}
	return nil
}

type ObjTypeValue int32

var ObjTypes = []ObjType{
	ObjTypeNullTypeValue:       NewObjType(ObjTypeNullTypeValue),
	ObjTypeTinyIntTypeValue:    NewObjType(ObjTypeSmallIntTypeValue),
	ObjTypeSmallIntTypeValue:   NewObjType(ObjTypeSmallIntTypeValue),
	ObjTypeMediumIntTypeValue:  NewObjType(ObjTypeMediumIntTypeValue),
	ObjTypeInt32TypeValue:      NewObjType(ObjTypeInt32TypeValue),
	ObjTypeInt64TypeValue:      NewObjType(ObjTypeInt64TypeValue),
	ObjTypeUTinyIntTypeValue:   NewObjType(ObjTypeUTinyIntTypeValue),
	ObjTypeUSmallIntTypeValue:  NewObjType(ObjTypeUSmallIntTypeValue),
	ObjTypeUMediumIntTypeValue: NewObjType(ObjTypeUMediumIntTypeValue),
	ObjTypeUInt32TypeValue:     NewObjType(ObjTypeUInt32TypeValue),
	ObjTypeUInt64TypeValue:     NewObjType(ObjTypeUInt64TypeValue),
	ObjTypeFloatTypeValue:      NewObjType(ObjTypeFloatTypeValue),
	ObjTypeDoubleTypeValue:     NewObjType(ObjTypeDoubleTypeValue),
	ObjTypeUFloatTypeValue:     NewObjType(ObjTypeUFloatTypeValue),
	ObjTypeUDoubleTypeValue:    NewObjType(ObjTypeUDoubleTypeValue),
	ObjTypeNumberTypeValue:     NewObjType(ObjTypeNumberTypeValue),
	ObjTypeUNumberTypeValue:    NewObjType(ObjTypeUNumberTypeValue),
	ObjTypeDateTimeTypeValue:   NewObjType(ObjTypeDateTimeTypeValue),
	ObjTypeTimestampTypeValue:  NewObjType(ObjTypeTimestampTypeValue),
	ObjTypeDateTypeValue:       NewObjType(ObjTypeDateTypeValue),
	ObjTypeTimeTypeValue:       NewObjType(ObjTypeTimeTypeValue),
	ObjTypeYearTypeValue:       NewObjType(ObjTypeYearTypeValue),
	ObjTypeVarcharTypeValue:    NewObjType(ObjTypeVarcharTypeValue),
	ObjTypeCharTypeValue:       NewObjType(ObjTypeCharTypeValue),
	ObjTypeHexStringTypeValue:  NewObjType(ObjTypeHexStringTypeValue),
	ObjTypeExtendTypeValue:     NewObjType(ObjTypeExtendTypeValue),
	ObjTypeUnknownTypeValue:    NewObjType(ObjTypeUnknownTypeValue),
	ObjTypeTinyTextTypeValue:   NewObjType(ObjTypeTinyTextTypeValue),
	ObjTypeTextTypeValue:       NewObjType(ObjTypeTextTypeValue),
	ObjTypeMediumTextTypeValue: NewObjType(ObjTypeMediumTextTypeValue),
	ObjTypeLongTextTypeValue:   NewObjType(ObjTypeLongTextTypeValue),
	ObjTypeBitTypeValue:        NewObjType(ObjTypeBitTypeValue),
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

type CollationLevel int32

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

type CollationType int32

const (
	CollationTypeInvalid          CollationType = 0
	CollationTypeUtf8mb4GeneralCi CollationType = 45
	CollationTypeUtf8mb4Bin       CollationType = 46
	CollationTypeBinary           CollationType = 63
	CollationTypeCollationFree    CollationType = 100
	CollationTypeMax              CollationType = 101
)

type NullType struct {
	value ObjTypeValue
}

func (t *NullType) Encode(buf []byte, value interface{}) {
	return
}

func (t *NullType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return nil
}

func (t *NullType) EncodedSize() int {
	return 0
}

func (t *NullType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMeta(t, csLevelIgnorable, csTypeBinary, 10)
}

func (t *NullType) Value() ObjTypeValue {
	return t.value
}

type TinyIntType struct {
	value ObjTypeValue
}

func (t *TinyIntType) Encode(buf []byte, value interface{}) {
	switch v := value.(type) {
	case bool:
		if v {
			util.PutUint8(buf, 1)
		} else {
			util.PutUint8(buf, 2)
		}
	case byte:
		util.PutUint8(buf, v)
	}
}

func (t *TinyIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	return util.Uint8(buffer.Next(1))
}

func (t *TinyIntType) EncodedSize() int {
	return 1
}

func (t *TinyIntType) DefaultObjMeta() *ObjectMeta {
	return NewObjectMeta(t, csLevelNumeric, csTypeBinary, 10)
}

func (t *TinyIntType) Value() ObjTypeValue {
	return t.value
}

type SmallIntType struct {
	value ObjTypeValue
}

func (t *SmallIntType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *SmallIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *SmallIntType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *SmallIntType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *SmallIntType) Value() ObjTypeValue {
	return t.value
}

type MediumIntType struct {
	value ObjTypeValue
}

func (t *MediumIntType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *MediumIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *MediumIntType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *MediumIntType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *MediumIntType) Value() ObjTypeValue {
	return t.value
}

type Int32Type struct {
	value ObjTypeValue
}

func (t *Int32Type) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *Int32Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *Int32Type) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *Int32Type) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *Int32Type) Value() ObjTypeValue {
	return t.value
}

type Int64Type struct {
	value ObjTypeValue
}

func (t *Int64Type) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *Int64Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *Int64Type) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *Int64Type) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *Int64Type) Value() ObjTypeValue {
	return t.value
}

type UTinyIntType struct {
	value ObjTypeValue
}

func (t *UTinyIntType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UTinyIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UTinyIntType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UTinyIntType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UTinyIntType) Value() ObjTypeValue {
	return t.value
}

type USmallIntType struct {
	value ObjTypeValue
}

func (t *USmallIntType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *USmallIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *USmallIntType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *USmallIntType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *USmallIntType) Value() ObjTypeValue {
	return t.value
}

type UMediumIntType struct {
	value ObjTypeValue
}

func (t *UMediumIntType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UMediumIntType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UMediumIntType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UMediumIntType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UMediumIntType) Value() ObjTypeValue {
	return t.value
}

type UInt32Type struct {
	value ObjTypeValue
}

func (t *UInt32Type) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UInt32Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UInt32Type) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UInt32Type) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UInt32Type) Value() ObjTypeValue {
	return t.value
}

type UInt64Type struct {
	value ObjTypeValue
}

func (t *UInt64Type) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UInt64Type) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UInt64Type) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UInt64Type) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UInt64Type) Value() ObjTypeValue {
	return t.value
}

type FloatType struct {
	value ObjTypeValue
}

func (t *FloatType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *FloatType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *FloatType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *FloatType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *FloatType) Value() ObjTypeValue {
	return t.value
}

type DoubleType struct {
	value ObjTypeValue
}

func (t *DoubleType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *DoubleType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *DoubleType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *DoubleType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *DoubleType) Value() ObjTypeValue {
	return t.value
}

type UFloatType struct {
	value ObjTypeValue
}

func (t *UFloatType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UFloatType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UFloatType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UFloatType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UFloatType) Value() ObjTypeValue {
	return t.value
}

type UDoubleType struct {
	value ObjTypeValue
}

func (t *UDoubleType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UDoubleType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UDoubleType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UDoubleType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UDoubleType) Value() ObjTypeValue {
	return t.value
}

type NumberType struct {
	value ObjTypeValue
}

func (t *NumberType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *NumberType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *NumberType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *NumberType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *NumberType) Value() ObjTypeValue {
	return t.value
}

type UNumberType struct {
	value ObjTypeValue
}

func (t *UNumberType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UNumberType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UNumberType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UNumberType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UNumberType) Value() ObjTypeValue {
	return t.value
}

type DateTimeType struct {
	value ObjTypeValue
}

func (t *DateTimeType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *DateTimeType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *DateTimeType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *DateTimeType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *DateTimeType) Value() ObjTypeValue {
	return t.value
}

type TimestampType struct {
	value ObjTypeValue
}

func (t *TimestampType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *TimestampType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *TimestampType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *TimestampType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *TimestampType) Value() ObjTypeValue {
	return t.value
}

type DateType struct {
	value ObjTypeValue
}

func (t *DateType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *DateType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *DateType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *DateType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *DateType) Value() ObjTypeValue {
	return t.value
}

type TimeType struct {
	value ObjTypeValue
}

func (t *TimeType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *TimeType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *TimeType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *TimeType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *TimeType) Value() ObjTypeValue {
	return t.value
}

type YearType struct {
	value ObjTypeValue
}

func (t *YearType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *YearType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *YearType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *YearType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *YearType) Value() ObjTypeValue {
	return t.value
}

type VarcharType struct {
	value ObjTypeValue
}

func (t *VarcharType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *VarcharType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *VarcharType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *VarcharType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *VarcharType) Value() ObjTypeValue {
	return t.value
}

type CharType struct {
	value ObjTypeValue
}

func (t *CharType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *CharType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *CharType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *CharType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *CharType) Value() ObjTypeValue {
	return t.value
}

type HexStringType struct {
	value ObjTypeValue
}

func (t *HexStringType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *HexStringType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *HexStringType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *HexStringType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *HexStringType) Value() ObjTypeValue {
	return t.value
}

type ExtendType struct {
	value ObjTypeValue
}

func (t *ExtendType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *ExtendType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *ExtendType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *ExtendType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *ExtendType) Value() ObjTypeValue {
	return t.value
}

type UnknownType struct {
	value ObjTypeValue
}

func (t *UnknownType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *UnknownType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *UnknownType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *UnknownType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *UnknownType) Value() ObjTypeValue {
	return t.value
}

type TinyTextType struct {
	value ObjTypeValue
}

func (t *TinyTextType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *TinyTextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *TinyTextType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *TinyTextType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *TinyTextType) Value() ObjTypeValue {
	return t.value
}

type TextType struct {
	value ObjTypeValue
}

func (t *TextType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *TextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *TextType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *TextType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *TextType) Value() ObjTypeValue {
	return t.value
}

type MediumTextType struct {
	value ObjTypeValue
}

func (t *MediumTextType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *MediumTextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *MediumTextType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *MediumTextType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *MediumTextType) Value() ObjTypeValue {
	return t.value
}

type LongTextType struct {
	value ObjTypeValue
}

func (t *LongTextType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *LongTextType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *LongTextType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *LongTextType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *LongTextType) Value() ObjTypeValue {
	return t.value
}

type BitType struct {
	value ObjTypeValue
}

func (t *BitType) Encode(buf []byte, value interface{}) {
	// TODO implement me
	panic("implement me")
}

func (t *BitType) Decode(buffer *bytes.Buffer, csType CollationType) interface{} {
	// TODO implement me
	panic("implement me")
}

func (t *BitType) EncodedSize() int {
	// TODO implement me
	panic("implement me")
}

func (t *BitType) DefaultObjMeta() *ObjectMeta {
	// TODO implement me
	panic("implement me")
}

func (t *BitType) Value() ObjTypeValue {
	return t.value
}

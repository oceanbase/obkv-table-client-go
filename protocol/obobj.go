package protocol

import (
	"bytes"
	"errors"
	"github.com/oceanbase/obkv-table-client-go/log"
	"strconv"
)

type ObObjMeta struct {
	objType ObObjType
	csLevel ObCollationLevel
	csType  ObCollationType
	scale   byte // When ObBitType, scale storages the length of bit
}

func (m *ObObjMeta) encode() []byte {
	//todo:imp
	return []byte{}
}

func (m *ObObjMeta) decode(buf bytes.Buffer) interface{} {
	//todo:imp
	return nil
}

func (m *ObObjMeta) getEncodeSize() int {
	return 4
}

func (m *ObObjMeta) ToString() string {
	return "ObObjMeta{" +
		"objType:" + m.objType.ToString() + ", " +
		"csLevel:" + m.csLevel.ToString() + ", " +
		"csType:" + m.csType.ToString() + ", " +
		"scale:" + strconv.Itoa(int(m.scale)) +
		"}"
}

const (
	ObNullTypeValue       = 0
	ObTinyIntTypeValue    = 1
	ObSmallIntTypeValue   = 2
	ObMediumIntTypeValue  = 3
	ObInt32TypeValue      = 4
	ObInt64TypeValue      = 5
	ObUTinyIntTypeValue   = 6
	ObUSmallIntTypeValue  = 7
	ObUMediumIntTypeValue = 8
	ObUInt32TypeValue     = 9
	ObUInt64TypeValue     = 10
	ObFloatTypeValue      = 11
	ObDoubleTypeValue     = 12
	ObUFloatTypeValue     = 13
	ObUDoubleTypeValue    = 14
	ObNumberTypeValue     = 15
	ObUNumberTypeValue    = 16
	ObDateTimeTypeValue   = 17
	ObTimestampTypeValue  = 18
	ObDateTypeValue       = 19
	ObTimeTypeValue       = 20
	ObYearTypeValue       = 21
	ObVarcharTypeValue    = 22
	ObCharTypeValue       = 23
	ObHexStringTypeValue  = 24
	ObExtendTypeValue     = 25
	ObUnknownTypeValue    = 26
	ObTinyTextTypeValue   = 27
	ObTextTypeValue       = 28
	ObMediumTextTypeValue = 29
	ObLongTextTypeValue   = 30
	ObBitTypeValue        = 31
)

type ObObjType interface {
	ToString() string
	encode(obj interface{}) ([]byte, error)
	decode(buf bytes.Buffer, csType ObCollationType) interface{}
	getEncodeSize(obj interface{}) int
	getDefaultObjMeta() ObObjMeta
	parseToComparable() (interface{}, error)
	getValue() int
}

func NewObObjType(value int) (ObObjType, error) {
	switch value {
	case ObNullTypeValue:
		return ObNullType{ObNullTypeValue}, nil
	case ObTinyIntTypeValue:
		return ObTinyIntType{ObTinyIntTypeValue}, nil
	case ObSmallIntTypeValue:
		return ObSmallIntType{ObSmallIntTypeValue}, nil
	case ObMediumIntTypeValue:
		return ObMediumIntType{ObMediumIntTypeValue}, nil
	case ObInt32TypeValue:
		return ObInt32Type{ObInt32TypeValue}, nil
	case ObInt64TypeValue:
		return ObInt64Type{ObInt64TypeValue}, nil
	case ObUTinyIntTypeValue:
		return ObUTinyIntType{ObUTinyIntTypeValue}, nil
	case ObUSmallIntTypeValue:
		return ObUSmallIntType{ObUSmallIntTypeValue}, nil
	case ObUMediumIntTypeValue:
		return ObUMediumIntType{ObUMediumIntTypeValue}, nil
	case ObUInt32TypeValue:
		return ObUInt32Type{ObUInt32TypeValue}, nil
	case ObUInt64TypeValue:
		return ObUInt64Type{ObUInt64TypeValue}, nil
	case ObFloatTypeValue:
		return ObFloatType{ObFloatTypeValue}, nil
	case ObDoubleTypeValue:
		return ObDoubleType{ObDoubleTypeValue}, nil
	case ObUFloatTypeValue:
		return ObUFloatType{ObUFloatTypeValue}, nil
	case ObUDoubleTypeValue:
		return ObUDoubleType{ObUDoubleTypeValue}, nil
	case ObNumberTypeValue:
		return ObNumberType{ObNumberTypeValue}, nil
	case ObUNumberTypeValue:
		return ObUNumberType{ObUNumberTypeValue}, nil
	case ObDateTimeTypeValue:
		return ObDateTimeType{ObDateTimeTypeValue}, nil
	case ObTimestampTypeValue:
		return ObTimestampType{ObTimestampTypeValue}, nil
	case ObDateTypeValue:
		return ObDateType{ObDateTypeValue}, nil
	case ObTimeTypeValue:
		return ObTimeType{ObTimeTypeValue}, nil
	case ObYearTypeValue:
		return ObYearType{ObYearTypeValue}, nil
	case ObVarcharTypeValue:
		return ObVarcharType{ObVarcharTypeValue}, nil
	case ObCharTypeValue:
		return ObCharType{ObCharTypeValue}, nil
	case ObHexStringTypeValue:
		return ObHexStringType{ObHexStringTypeValue}, nil
	case ObExtendTypeValue:
		return ObExtendType{ObExtendTypeValue}, nil
	case ObUnknownTypeValue:
		return ObUnknownType{ObUnknownTypeValue}, nil
	case ObTinyTextTypeValue:
		return ObTinyTextType{ObTinyTextTypeValue}, nil
	case ObTextTypeValue:
		return ObTextType{ObTextTypeValue}, nil
	case ObMediumTextTypeValue:
		return ObMediumTextType{ObMediumTextTypeValue}, nil
	case ObLongTextTypeValue:
		return ObLongTextType{ObLongTextTypeValue}, nil
	case ObBitTypeValue:
		return ObBitType{ObBitTypeValue}, nil
	default:
		log.Warn("invalid object type value", log.Int("value", value))
		err := errors.New("invalid object type value")
		return nil, err
	}
}

type ObNullType struct {
	value int // ObNullTypeValue
}

func (t ObNullType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObNullType" +
		"}"
}

func (t ObNullType) encode(obj interface{}) ([]byte, error) {
	return []byte{0}, nil
}

func (t ObNullType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	return nil
}

func (t ObNullType) getEncodeSize(obj interface{}) int {
	return 0
}

func (t ObNullType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelIgnorable},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObNullType) parseToComparable() (interface{}, error) {
	return nil, nil
}

func (t ObNullType) getValue() int {
	return t.value
}

type ObTinyIntType struct {
	value int // ObTinyIntTypeValue
}

func (t ObTinyIntType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObTinyIntType" +
		"}"
}

func (t ObTinyIntType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTinyIntType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObTinyIntType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObTinyIntType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObTinyIntType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTinyIntType) getValue() int {
	return t.value
}

type ObSmallIntType struct {
	value int // ObSmallIntTypeValue
}

func (t ObSmallIntType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObSmallIntType" +
		"}"
}

func (t ObSmallIntType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObSmallIntType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObSmallIntType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObSmallIntType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObSmallIntType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObSmallIntType) getValue() int {
	return t.value
}

type ObMediumIntType struct {
	value int // ObMediumIntTypeValue
}

func (t ObMediumIntType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObMediumIntType" +
		"}"
}

func (t ObMediumIntType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObMediumIntType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObMediumIntType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObMediumIntType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObMediumIntType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObMediumIntType) getValue() int {
	return t.value
}

type ObInt32Type struct {
	value int // ObInt32TypeValue
}

func (t ObInt32Type) ToString() string {
	return "ObObjType{" +
		"type:" + "ObInt32Type" +
		"}"
}

func (t ObInt32Type) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObInt32Type) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObInt32Type) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObInt32Type) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObInt32Type) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObInt32Type) getValue() int {
	return t.value
}

type ObInt64Type struct {
	value int // ObInt64TypeValue
}

func (t ObInt64Type) ToString() string {
	return "ObObjType{" +
		"type:" + "ObInt64Type" +
		"}"
}

func (t ObInt64Type) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObInt64Type) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObInt64Type) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObInt64Type) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObInt64Type) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObInt64Type) getValue() int {
	return t.value
}

type ObUTinyIntType struct {
	value int // ObUTinyIntTypeValue
}

func (t ObUTinyIntType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUTinyIntType" +
		"}"
}

func (t ObUTinyIntType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUTinyIntType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUTinyIntType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUTinyIntType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUTinyIntType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUTinyIntType) getValue() int {
	return t.value
}

type ObUSmallIntType struct {
	value int // ObUSmallIntTypeValue
}

func (t ObUSmallIntType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUSmallIntType" +
		"}"
}

func (t ObUSmallIntType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUSmallIntType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUSmallIntType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUSmallIntType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUSmallIntType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUSmallIntType) getValue() int {
	return t.value
}

type ObUMediumIntType struct {
	value int // ObUMediumIntTypeValue
}

func (t ObUMediumIntType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUMediumIntType" +
		"}"
}

func (t ObUMediumIntType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUMediumIntType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUMediumIntType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUMediumIntType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUMediumIntType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUMediumIntType) getValue() int {
	return t.value
}

type ObUInt32Type struct {
	value int // ObUInt32TypeValue
}

func (t ObUInt32Type) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUInt32Type" +
		"}"
}

func (t ObUInt32Type) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUInt32Type) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUInt32Type) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUInt32Type) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUInt32Type) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUInt32Type) getValue() int {
	return t.value
}

type ObUInt64Type struct {
	value int // ObUInt64TypeValue
}

func (t ObUInt64Type) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUInt64Type" +
		"}"
}

func (t ObUInt64Type) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUInt64Type) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUInt64Type) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUInt64Type) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUInt64Type) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUInt64Type) getValue() int {
	return t.value
}

type ObFloatType struct {
	value int // ObFloatTypeValue
}

func (t ObFloatType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObFloatType" +
		"}"
}

func (t ObFloatType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObFloatType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObFloatType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObFloatType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObFloatType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObFloatType) getValue() int {
	return t.value
}

type ObDoubleType struct {
	value int // ObDoubleTypeValue
}

func (t ObDoubleType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObDoubleType" +
		"}"
}

func (t ObDoubleType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObDoubleType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObDoubleType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObDoubleType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObDoubleType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObDoubleType) getValue() int {
	return t.value
}

type ObUFloatType struct {
	value int // ObUFloatTypeValue
}

func (t ObUFloatType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUFloatType" +
		"}"
}

func (t ObUFloatType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUFloatType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUFloatType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUFloatType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUFloatType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUFloatType) getValue() int {
	return t.value
}

type ObUDoubleType struct {
	value int // ObUDoubleTypeValue
}

func (t ObUDoubleType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUDoubleType" +
		"}"
}

func (t ObUDoubleType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUDoubleType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUDoubleType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUDoubleType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUDoubleType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUDoubleType) getValue() int {
	return t.value
}

type ObNumberType struct {
	value int // ObNumberTypeValue
}

func (t ObNumberType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObNumberType" +
		"}"
}

func (t ObNumberType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObNumberType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObNumberType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObNumberType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObNumberType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObNumberType) getValue() int {
	return t.value
}

type ObUNumberType struct {
	value int // ObUNumberTypeValue
}

func (t ObUNumberType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUNumberType" +
		"}"
}

func (t ObUNumberType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUNumberType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUNumberType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUNumberType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObUNumberType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUNumberType) getValue() int {
	return t.value
}

type ObDateTimeType struct {
	value int // ObDateTimeTypeValue
}

func (t ObDateTimeType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObDateTimeType" +
		"}"
}

func (t ObDateTimeType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObDateTimeType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObDateTimeType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObDateTimeType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObDateTimeType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObDateTimeType) getValue() int {
	return t.value
}

type ObTimestampType struct {
	value int // ObTimestampTypeValue
}

func (t ObTimestampType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObTimestampType" +
		"}"
}

func (t ObTimestampType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTimestampType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObTimestampType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObTimestampType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObTimestampType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTimestampType) getValue() int {
	return t.value
}

type ObDateType struct {
	value int // ObDateTypeValue
}

func (t ObDateType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObDateType" +
		"}"
}

func (t ObDateType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObDateType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObDateType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObDateType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObDateType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObDateType) getValue() int {
	return t.value
}

type ObTimeType struct {
	value int // ObTimeTypeValue
}

func (t ObTimeType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObTimeType" +
		"}"
}

func (t ObTimeType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTimeType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObTimeType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObTimeType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObTimeType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTimeType) getValue() int {
	return t.value
}

type ObYearType struct {
	value int // ObYearTypeValue
}

func (t ObYearType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObYearType" +
		"}"
}

func (t ObYearType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObYearType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObYearType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObYearType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObYearType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObYearType) getValue() int {
	return t.value
}

type ObVarcharType struct {
	value int // ObVarcharTypeValue
}

func (t ObVarcharType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObVarcharType" +
		"}"
}

func (t ObVarcharType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObVarcharType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObVarcharType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObVarcharType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelExplicit},
		ObCollationType{csTypeUtf8mb4GeneralCi},
		byte(10),
	}
}

func (t ObVarcharType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObVarcharType) getValue() int {
	return t.value
}

type ObCharType struct {
	value int // ObCharTypeValue
}

func (t ObCharType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObCharType" +
		"}"
}

func (t ObCharType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObCharType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObCharType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObCharType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelExplicit},
		ObCollationType{csTypeUtf8mb4GeneralCi},
		byte(10),
	}
}

func (t ObCharType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObCharType) getValue() int {
	return t.value
}

type ObHexStringType struct {
	value int // ObHexStringTypeValue
}

func (t ObHexStringType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObHexStringType" +
		"}"
}

func (t ObHexStringType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObHexStringType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObHexStringType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObHexStringType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelInvalid},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObHexStringType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObHexStringType) getValue() int {
	return t.value
}

type ObExtendType struct {
	value int // ObExtendTypeValue
}

func (t ObExtendType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObExtendType" +
		"}"
}

func (t ObExtendType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObExtendType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObExtendType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObExtendType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelInvalid},
		ObCollationType{csTypeInvalid},
		byte(10),
	}
}

func (t ObExtendType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObExtendType) getValue() int {
	return t.value
}

type ObUnknownType struct {
	value int // ObUnknownTypeValue
}

func (t ObUnknownType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObUnknownType" +
		"}"
}

func (t ObUnknownType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUnknownType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObUnknownType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObUnknownType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelInvalid},
		ObCollationType{csTypeInvalid},
		byte(10),
	}
}

func (t ObUnknownType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObUnknownType) getValue() int {
	return t.value
}

type ObTinyTextType struct {
	value int // ObTinyTextTypeValue
}

func (t ObTinyTextType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObTinyTextType" +
		"}"
}

func (t ObTinyTextType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTinyTextType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObTinyTextType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObTinyTextType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelImplicit},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObTinyTextType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTinyTextType) getValue() int {
	return t.value
}

type ObTextType struct {
	value int // ObTextTypeValue
}

func (t ObTextType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObTextType" +
		"}"
}

func (t ObTextType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTextType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObTextType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObTextType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelImplicit},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObTextType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObTextType) getValue() int {
	return t.value
}

type ObMediumTextType struct {
	value int // ObMediumTextTypeValue
}

func (t ObMediumTextType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObMediumTextType" +
		"}"
}

func (t ObMediumTextType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObMediumTextType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObMediumTextType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObMediumTextType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelImplicit},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObMediumTextType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObMediumTextType) getValue() int {
	return t.value
}

type ObLongTextType struct {
	value int // ObLongTextTypeValue
}

func (t ObLongTextType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObLongTextType" +
		"}"
}

func (t ObLongTextType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObLongTextType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObLongTextType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObLongTextType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelImplicit},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObLongTextType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObLongTextType) getValue() int {
	return t.value
}

type ObBitType struct {
	value int // ObBitTypeValue
}

func (t ObBitType) ToString() string {
	return "ObObjType{" +
		"type:" + "ObBitType" +
		"}"
}

func (t ObBitType) encode(obj interface{}) ([]byte, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObBitType) decode(buf bytes.Buffer, csType ObCollationType) interface{} {
	// todo:imp
	return nil
}

func (t ObBitType) getEncodeSize(obj interface{}) int {
	// todo:imp
	return 0
}

func (t ObBitType) getDefaultObjMeta() ObObjMeta {
	return ObObjMeta{t,
		ObCollationLevel{csLevelNumeric},
		ObCollationType{csTypeBinary},
		byte(10),
	}
}

func (t ObBitType) parseToComparable() (interface{}, error) {
	// todo:imp
	return nil, errors.New("not implement")
}

func (t ObBitType) getValue() int {
	return t.value
}

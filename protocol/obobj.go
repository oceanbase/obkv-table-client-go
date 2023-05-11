package protocol

import (
	"bytes"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/util"
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

func (m *ObObjMeta) String() string {
	// objType to string
	var objTypeStr string
	if m.objType != nil {
		objTypeStr = m.objType.String()
	} else {
		objTypeStr = "nil"
	}
	return "ObObjMeta{" +
		"objType:" + objTypeStr + ", " +
		"csLevel:" + m.csLevel.String() + ", " +
		"csType:" + m.csType.String() + ", " +
		"scale:" + strconv.Itoa(int(m.scale)) +
		"}"
}

func ParseToLong(obj interface{}) (interface{}, error) {
	if v, ok := obj.(string); ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("failed to convert string to int", log.String("obj", v))
			return nil, err
		}
		return int64(i), nil
	} else if v, ok := obj.(int64); ok {
		return v, nil
	} else if v, ok := obj.(int); ok {
		return int64(v), nil
	} else if v, ok := obj.(int32); ok {
		return int64(v), nil
	} else if v, ok := obj.(int16); ok {
		return int64(v), nil
	} else if v, ok := obj.(int8); ok {
		return int64(v), nil
	} else {
		log.Warn("invalid type to convert to long", log.String("obj", util.InterfaceToString(obj)))
		return nil, errors.New("invalid type to convert to long")
	}
}

func parseTimestamp(obj interface{}) (time.Time, error) {
	if v, ok := obj.(time.Time); ok {
		return v, nil
	} else if v, ok := obj.(string); ok {
		return time.Parse("2006-01-02 15:04:01", v) // UTC
	} else {
		log.Warn("unexpected type to timestamp")
		return time.Time{}, errors.New("unexpected type to timestamp")
	}
}

func parseTextToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	if csType.value == CsTypeBinary {
		if v, ok := obj.([]byte); ok {
			return string(v), nil
		} else if v, ok := obj.(string); ok {
			return v, nil
		} else {
			log.Warn("unexpected type")
			return nil, errors.New("unexpected obj type")
		}
	} else {
		if v, ok := obj.(string); ok {
			return v, nil
		} else if _, ok := obj.([]byte); ok {
			// todo:impl
			// return Serialization.decodeVString(((ObBytesString) object).bytes);
			return nil, errors.New("need impl Serialization.decodeVString")
		} else {
			log.Warn("unexpected type")
			return nil, errors.New("unexpected obj type")
		}
	}
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
	String() string
	encode(obj interface{}) ([]byte, error)
	decode(buf bytes.Buffer, csType ObCollationType) interface{}
	getEncodeSize(obj interface{}) int
	getDefaultObjMeta() ObObjMeta
	ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error)
	GetValue() int
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

func (t ObNullType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObNullType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, nil
}

func (t ObNullType) GetValue() int {
	return t.value
}

type ObTinyIntType struct {
	value int // ObTinyIntTypeValue
}

func (t ObTinyIntType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObTinyIntType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		if v >= math.MinInt8 && v <= math.MaxInt8 {
			return byte(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObTinyIntType) GetValue() int {
	return t.value
}

type ObSmallIntType struct {
	value int // ObSmallIntTypeValue
}

func (t ObSmallIntType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObSmallIntType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		if v >= math.MinInt16 && v <= math.MaxInt16 {
			return int16(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObSmallIntType) GetValue() int {
	return t.value
}

type ObMediumIntType struct {
	value int // ObMediumIntTypeValue
}

func (t ObMediumIntType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObMediumIntType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObMediumIntType is not supported")
}

func (t ObMediumIntType) GetValue() int {
	return t.value
}

type ObInt32Type struct {
	value int // ObInt32TypeValue
}

func (t ObInt32Type) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObInt32Type) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return int32(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObInt32Type) GetValue() int {
	return t.value
}

type ObInt64Type struct {
	value int // ObInt64TypeValue
}

func (t ObInt64Type) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObInt64Type) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return ParseToLong(obj)
}

func (t ObInt64Type) GetValue() int {
	return t.value
}

type ObUTinyIntType struct {
	value int // ObUTinyIntTypeValue
}

func (t ObUTinyIntType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUTinyIntType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		if v >= 0 && v <= math.MaxUint8 {
			return int16(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObUTinyIntType) GetValue() int {
	return t.value
}

type ObUSmallIntType struct {
	value int // ObUSmallIntTypeValue
}

func (t ObUSmallIntType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUSmallIntType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		if v >= 0 && v <= math.MaxUint16 {
			return int32(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObUSmallIntType) GetValue() int {
	return t.value
}

type ObUMediumIntType struct {
	value int // ObUMediumIntTypeValue
}

func (t ObUMediumIntType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUMediumIntType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		const Uint24Max = (1 << 24) - 1
		if v >= 0 && v <= Uint24Max {
			return int32(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObUMediumIntType) GetValue() int {
	return t.value
}

type ObUInt32Type struct {
	value int // ObUInt32TypeValue
}

func (t ObUInt32Type) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUInt32Type) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	value, err := ParseToLong(obj)
	if err != nil {
		log.Warn("failed to parse to long or null")
		return nil, err
	}
	if v, ok := value.(int64); ok {
		if v >= 0 && v <= math.MaxUint32 {
			return int64(v), nil
		} else {
			log.Warn("value out of range", log.Int64("value", v))
			return nil, err
		}
	} else {
		log.Warn("failed to convert value to int64")
		return nil, errors.New("failed to convert value to int64")
	}
}

func (t ObUInt32Type) GetValue() int {
	return t.value
}

type ObUInt64Type struct {
	value int // ObUInt64TypeValue
}

func (t ObUInt64Type) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUInt64Type) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return ParseToLong(obj)
}

func (t ObUInt64Type) GetValue() int {
	return t.value
}

type ObFloatType struct {
	value int // ObFloatTypeValue
}

func (t ObFloatType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObFloatType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	if v, ok := obj.(float32); ok {
		return v, nil
	} else if v, ok := obj.(string); ok {
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			log.Warn("failed to convert string to float32", log.String("obj", v))
			return nil, err
		}
		return float32(f), nil
	} else {
		log.Warn("unexpected type")
		return nil, errors.New("unexpected type")
	}
}

func (t ObFloatType) GetValue() int {
	return t.value
}

type ObDoubleType struct {
	value int // ObDoubleTypeValue
}

func (t ObDoubleType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObDoubleType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	if v, ok := obj.(float64); ok {
		return v, nil
	} else if v, ok := obj.(string); ok {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Warn("failed to convert string to float64", log.String("obj", v))
			return nil, err
		}
		return f, nil
	} else {
		log.Warn("unexpected type")
		return nil, errors.New("unexpected type")
	}
}

func (t ObDoubleType) GetValue() int {
	return t.value
}

type ObUFloatType struct {
	value int // ObUFloatTypeValue
}

func (t ObUFloatType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUFloatType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObUFloatType is not supported")
}

func (t ObUFloatType) GetValue() int {
	return t.value
}

type ObUDoubleType struct {
	value int // ObUDoubleTypeValue
}

func (t ObUDoubleType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUDoubleType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObUDoubleType is not supported")
}

func (t ObUDoubleType) GetValue() int {
	return t.value
}

type ObNumberType struct {
	value int // ObNumberTypeValue
}

func (t ObNumberType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObNumberType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObNumberType is not supported")
}

func (t ObNumberType) GetValue() int {
	return t.value
}

type ObUNumberType struct {
	value int // ObUNumberTypeValue
}

func (t ObUNumberType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObUNumberType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObUNumberType is not supported")
}

func (t ObUNumberType) GetValue() int {
	return t.value
}

type ObDateTimeType struct {
	value int // ObDateTimeTypeValue
}

func (t ObDateTimeType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObDateTimeType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTimestamp(obj)
}

func (t ObDateTimeType) GetValue() int {
	return t.value
}

type ObTimestampType struct {
	value int // ObTimestampTypeValue
}

func (t ObTimestampType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObTimestampType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTimestamp(obj)
}

func (t ObTimestampType) GetValue() int {
	return t.value
}

type ObDateType struct {
	value int // ObDateTypeValue
}

func (t ObDateType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObDateType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	if v, ok := obj.(time.Time); ok {
		return v, nil
	} else if v, ok := obj.(string); ok {
		return time.ParseInLocation("YYYY-MM-DD", v, time.Local)
	} else {
		log.Warn("unexpected type")
		return time.Time{}, errors.New("unexpected type")
	}
}

func (t ObDateType) GetValue() int {
	return t.value
}

type ObTimeType struct {
	value int // ObTimeTypeValue
}

func (t ObTimeType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObTimeType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObTimeType is not supported")
}

func (t ObTimeType) GetValue() int {
	return t.value
}

type ObYearType struct {
	value int // ObYearTypeValue
}

func (t ObYearType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObYearType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObYearType is not supported")
}

func (t ObYearType) GetValue() int {
	return t.value
}

type ObVarcharType struct {
	value int // ObVarcharTypeValue
}

func (t ObVarcharType) String() string {
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
		ObCollationType{CsTypeUtf8mb4GeneralCi},
		byte(10),
	}
}

func (t ObVarcharType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTextToComparable(obj, csType)
}

func (t ObVarcharType) GetValue() int {
	return t.value
}

type ObCharType struct {
	value int // ObCharTypeValue
}

func (t ObCharType) String() string {
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
		ObCollationType{CsTypeUtf8mb4GeneralCi},
		byte(10),
	}
}

func (t ObCharType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTextToComparable(obj, csType)
}

func (t ObCharType) GetValue() int {
	return t.value
}

type ObHexStringType struct {
	value int // ObHexStringTypeValue
}

func (t ObHexStringType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObHexStringType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObHexStringType is not supported")
}

func (t ObHexStringType) GetValue() int {
	return t.value
}

type ObExtendType struct {
	value int // ObExtendTypeValue
}

func (t ObExtendType) String() string {
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
		ObCollationType{CsTypeInvalid},
		byte(10),
	}
}

func (t ObExtendType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObExtendType is not supported")
}

func (t ObExtendType) GetValue() int {
	return t.value
}

type ObUnknownType struct {
	value int // ObUnknownTypeValue
}

func (t ObUnknownType) String() string {
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
		ObCollationType{CsTypeInvalid},
		byte(10),
	}
}

func (t ObUnknownType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObUnknownType is not supported")
}

func (t ObUnknownType) GetValue() int {
	return t.value
}

type ObTinyTextType struct {
	value int // ObTinyTextTypeValue
}

func (t ObTinyTextType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObTinyTextType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTextToComparable(obj, csType)
}

func (t ObTinyTextType) GetValue() int {
	return t.value
}

type ObTextType struct {
	value int // ObTextTypeValue
}

func (t ObTextType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObTextType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTextToComparable(obj, csType)
}

func (t ObTextType) GetValue() int {
	return t.value
}

type ObMediumTextType struct {
	value int // ObMediumTextTypeValue
}

func (t ObMediumTextType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObMediumTextType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTextToComparable(obj, csType)
}

func (t ObMediumTextType) GetValue() int {
	return t.value
}

type ObLongTextType struct {
	value int // ObLongTextTypeValue
}

func (t ObLongTextType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObLongTextType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return parseTextToComparable(obj, csType)
}

func (t ObLongTextType) GetValue() int {
	return t.value
}

type ObBitType struct {
	value int // ObBitTypeValue
}

func (t ObBitType) String() string {
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
		ObCollationType{CsTypeBinary},
		byte(10),
	}
}

func (t ObBitType) ParseToComparable(obj interface{}, csType ObCollationType) (interface{}, error) {
	return nil, errors.New("ObBitType is not supported")
}

func (t ObBitType) GetValue() int {
	return t.value
}

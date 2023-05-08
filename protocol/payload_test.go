package protocol

import (
	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestObCollationType_ToString(t *testing.T) {
	coll := ObCollationType{}
	assert.Equal(t, "ObCollationType{collationType:CsTypeInvalid}", coll.String())
	coll = NewObCollationType(CsTypeInvalid)
	assert.Equal(t, "ObCollationType{collationType:CsTypeInvalid}", coll.String())
	coll = NewObCollationType(CsTypeUtf8mb4GeneralCi)
	assert.Equal(t, "ObCollationType{collationType:CsTypeUtf8mb4GeneralCi}", coll.String())
	coll = NewObCollationType(CsTypeUtf8mb4Bin)
	assert.Equal(t, "ObCollationType{collationType:CsTypeUtf8mb4Bin}", coll.String())
	coll = NewObCollationType(CsTypeBinary)
	assert.Equal(t, "ObCollationType{collationType:CsTypeBinary}", coll.String())
	coll = NewObCollationType(CsTypeCollationFree)
	assert.Equal(t, "ObCollationType{collationType:CsTypeCollationFree}", coll.String())
	coll = NewObCollationType(CsTypeMax)
	assert.Equal(t, "ObCollationType{collationType:CsTypeMax}", coll.String())
}

func TestObCollationLevel_ToString(t *testing.T) {
	level := ObCollationLevel{}
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelExplicit}", level.String())
	level = newObCollationLevel(csLevelExplicit)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelExplicit}", level.String())
	level = newObCollationLevel(csLevelNone)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelNone}", level.String())
	level = newObCollationLevel(csLevelImplicit)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelImplicit}", level.String())
	level = newObCollationLevel(csLevelSysConst)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelSysConst}", level.String())
	level = newObCollationLevel(csLevelCoercible)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelCoercible}", level.String())
	level = newObCollationLevel(csLevelNumeric)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelNumeric}", level.String())
	level = newObCollationLevel(csLevelIgnorable)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelIgnorable}", level.String())
	level = newObCollationLevel(csLevelInvalid)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelInvalid}", level.String())
}

func TestObColumn_ToString(t *testing.T) {
	column := &route.ObColumn{}
	assert.Equal(t, "ObColumn{"+
		"columnName:, "+
		"index:0, "+
		"objType:nil, "+
		"collationType:ObCollationType{collationType:CsTypeInvalid}, "+
		"refColumnNames:[], "+
		"isGenColumn:false, "+
		"columnExpress:nil}",
		column.String(),
	)
	objType, _ := NewObObjType(1)
	collType := NewObCollationType(63)
	column = route.NewObSimpleColumn("testColumnName", 0, objType, collType)
	assert.Equal(t, "ObColumn{"+
		"columnName:testColumnName, "+
		"index:0, "+
		"objType:ObObjType{type:ObTinyIntType}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[testColumnName], "+
		"isGenColumn:false, "+
		"columnExpress:nil}",
		column.String(),
	)
}

func TestObObjMeta_ToString(t *testing.T) {
	meta := ObObjMeta{}
	assert.Equal(t, "ObObjMeta{"+
		"objType:nil, "+
		"csLevel:ObCollationLevel{collationLevel:csLevelExplicit}, "+
		"csType:ObCollationType{collationType:CsTypeInvalid}, "+
		"scale:0}",
		meta.String(),
	)
}

func TestObVString_ToString(t *testing.T) {
	v := ObVString{}
	assert.Equal(t, "ObVString{stringVal:, bytesVal:, encodeBytes:}", v.String())
	str := "test"
	v = ObVString{str, []byte(str), []byte(str)}
	assert.Equal(t, "ObVString{stringVal:test, bytesVal:test, encodeBytes:test}", v.String())
}

func TestObBytesString_ToString(t *testing.T) {
	v := ObBytesString{}
	assert.Equal(t, "ObBytesString{bytesVal:}", v.String())
	str := "test"
	v = ObBytesString{[]byte(str)}
	assert.Equal(t, "ObBytesString{bytesVal:test}", v.String())
}

func TestParse(t *testing.T) {
	// test ParseToLong
	{
		// string
		str := string("")
		_, err := ParseToLong(str)
		assert.NotEqual(t, err, nil)
		str = string("0")
		res, err := ParseToLong(str)
		assert.Equal(t, err, nil)
		assert.Equal(t, res, int64(0))
		// ObVString
		vstr := ObVString{}
		_, err = ParseToLong(vstr)
		assert.NotEqual(t, err, nil)
		vstr = ObVString{stringVal: "0"}
		res, err = ParseToLong(vstr)
		assert.Equal(t, err, nil)
		assert.Equal(t, res, int64(0))
		// int64
		i64 := int64(666)
		res, err = ParseToLong(i64)
		assert.Equal(t, err, nil)
		assert.Equal(t, res, int64(666))
		// int32
		i32 := int32(666)
		res, err = ParseToLong(i32)
		assert.Equal(t, err, nil)
		assert.Equal(t, res, int64(666))
		// int16
		i16 := int16(666)
		res, err = ParseToLong(i16)
		assert.Equal(t, err, nil)
		assert.Equal(t, res, int64(666))
		// int8
		i8 := int8(10)
		res, err = ParseToLong(i8)
		assert.Equal(t, err, nil)
		assert.Equal(t, res, int64(10))
		// other
		_, err = ParseToLong(nil) // nil
		assert.NotEqual(t, err, nil)
		_, err = ParseToLong(uint64(10)) // uint64
		assert.NotEqual(t, err, nil)
	}

	// test parseTimestamp
	{
		ti := time.Now()
		_, err := parseTimestamp(ti)
		assert.Equal(t, err, nil)
		_, err = parseTimestamp("2020-11-11 11:01:02")
		assert.Equal(t, err, nil)
		_, err = parseTimestamp(123)
		assert.NotEqual(t, err, nil)
	}

	// test parseTextToComparable
	{
		{
			// cs type = CsTypeBinary
			csBin := ObCollationType{CsTypeBinary}
			bStr := ObBytesString{[]byte{0, 1, 2}}
			res, err := parseTextToComparable(bStr, csBin)
			assert.Equal(t, err, nil)
			assert.Equal(t, res, bStr)
			bArr := []byte{0, 1, 2}
			res, err = parseTextToComparable(bArr, csBin)
			assert.Equal(t, err, nil)
			assert.Equal(t, res, bStr)
			vStr := ObVString{bytesVal: bArr}
			res, err = parseTextToComparable(vStr, csBin)
			assert.Equal(t, err, nil)
			assert.Equal(t, res, bStr)
		}
		{
			// cs type = CsTypeUtf8mb4GeneralCi
			csBin := ObCollationType{CsTypeUtf8mb4GeneralCi}
			bStr := ObBytesString{[]byte{0, 1, 2}}
			_, err := parseTextToComparable(bStr, csBin)
			assert.NotEqual(t, err, nil) // todo
			bArr := []byte{0, 1, 2}
			_, err = parseTextToComparable(bArr, csBin)
			assert.NotEqual(t, err, nil) // todo
			vStr := ObVString{stringVal: "abc"}
			_, err = parseTextToComparable(vStr, csBin)
			assert.Equal(t, err, nil)
		}
	}
}

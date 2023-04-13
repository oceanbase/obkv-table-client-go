package protocol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObCollationType_ToString(t *testing.T) {
	coll := ObCollationType{}
	assert.Equal(t, "ObCollationType{collationType:csTypeInvalid}", coll.ToString())
	coll = NewObCollationType(csTypeInvalid)
	assert.Equal(t, "ObCollationType{collationType:csTypeInvalid}", coll.ToString())
	coll = NewObCollationType(csTypeUtf8mb4GeneralCi)
	assert.Equal(t, "ObCollationType{collationType:csTypeUtf8mb4GeneralCi}", coll.ToString())
	coll = NewObCollationType(csTypeUtf8mb4Bin)
	assert.Equal(t, "ObCollationType{collationType:csTypeUtf8mb4Bin}", coll.ToString())
	coll = NewObCollationType(csTypeBinary)
	assert.Equal(t, "ObCollationType{collationType:csTypeBinary}", coll.ToString())
	coll = NewObCollationType(csTypeCollationFree)
	assert.Equal(t, "ObCollationType{collationType:csTypeCollationFree}", coll.ToString())
	coll = NewObCollationType(csTypeMax)
	assert.Equal(t, "ObCollationType{collationType:csTypeMax}", coll.ToString())
}

func TestObCollationLevel_ToString(t *testing.T) {
	level := ObCollationLevel{}
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelExplicit}", level.ToString())
	level = newObCollationLevel(csLevelExplicit)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelExplicit}", level.ToString())
	level = newObCollationLevel(csLevelNone)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelNone}", level.ToString())
	level = newObCollationLevel(csLevelImplicit)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelImplicit}", level.ToString())
	level = newObCollationLevel(csLevelSysConst)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelSysConst}", level.ToString())
	level = newObCollationLevel(csLevelCoercible)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelCoercible}", level.ToString())
	level = newObCollationLevel(csLevelNumeric)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelNumeric}", level.ToString())
	level = newObCollationLevel(csLevelIgnorable)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelIgnorable}", level.ToString())
	level = newObCollationLevel(csLevelInvalid)
	assert.Equal(t, "ObCollationLevel{collationLevel:csLevelInvalid}", level.ToString())
}

func TestObColumn_ToString(t *testing.T) {
	column := &ObColumn{}
	assert.Equal(t, "ObColumn{"+
		"columnName:, "+
		"index:0, "+
		"objType:nil, "+
		"collationType:ObCollationType{collationType:csTypeInvalid}, "+
		"refColumnNames:[], "+
		"isGenColumn:false, "+
		"columnExpress:nil}",
		column.ToString(),
	)
	objType, _ := NewObObjType(1)
	collType := NewObCollationType(63)
	column = NewObSimpleColumn("testColumnName", 0, objType, collType)
	assert.Equal(t, "ObColumn{"+
		"columnName:testColumnName, "+
		"index:0, "+
		"objType:ObObjType{type:ObTinyIntType}, "+
		"collationType:ObCollationType{collationType:csTypeBinary}, "+
		"refColumnNames:[testColumnName], "+
		"isGenColumn:false, "+
		"columnExpress:nil}",
		column.ToString(),
	)
}

func TestObObjMeta_ToString(t *testing.T) {
	meta := ObObjMeta{}
	assert.Equal(t, "ObObjMeta{"+
		"objType:nil, "+
		"csLevel:ObCollationLevel{collationLevel:csLevelExplicit}, "+
		"csType:ObCollationType{collationType:csTypeInvalid}, "+
		"scale:0}",
		meta.ToString(),
	)
}

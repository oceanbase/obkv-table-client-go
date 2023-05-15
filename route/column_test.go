package route

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/protocol"
)

func TestObColumn_String(t *testing.T) {
	col := &obColumn{}
	assert.Equal(t, "obColumn{columnName:, index:0, objType:nil, collationType:ObCollationType{collationType:CsTypeInvalid}, refColumnNames:[], isGenColumn:false, columnExpress:nil}", col.String())
	objType, _ := protocol.NewObjType(protocol.ObjTypeTinyIntTypeValue)
	col = newObSimpleColumn("c1", 1, objType, protocol.CollationTypeUtf8mb4GeneralCi)
	assert.Equal(t, "obColumn{columnName:c1, index:1, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeUtf8mb4GeneralCi}, refColumnNames:[c1], isGenColumn:false, columnExpress:nil}", col.String())
}

// todo: test after obobj type refactoring
func TestObColumn_EvalValue(t *testing.T) {
	objType, _ := protocol.NewObjType(protocol.ObjTypeVarcharTypeValue)
	col := newObSimpleColumn("c1", 1, objType, protocol.CollationTypeUtf8mb4GeneralCi)
	_, err := col.EvalValue(0)
	assert.NotEqual(t, nil, err)
}

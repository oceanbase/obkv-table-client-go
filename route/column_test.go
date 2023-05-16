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
)

func TestObColumn_String(t *testing.T) {
	col := &obColumn{}
	assert.Equal(t, "obColumn{columnName:, index:0, objType:nil, collationType:ObCollationTypeInvalid, refColumnNames:[], isGenColumn:false, columnExpress:nil}", col.String())
	objType, _ := protocol.NewObjType(protocol.ObObjTypeTinyIntTypeValue)
	col = newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeUtf8mb4GeneralCi)
	assert.Equal(t, "obColumn{columnName:c1, index:1, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationTypeUtf8mb4GeneralCi, refColumnNames:[c1], isGenColumn:false, columnExpress:nil}", col.String())
}

func TestObColumn_EvalValue(t *testing.T) {
	// simple column
	objType, _ := protocol.NewObjType(protocol.ObObjTypeVarcharTypeValue)
	col := newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeUtf8mb4GeneralCi)
	value, err := col.EvalValue("hello")
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "hello", value)
}

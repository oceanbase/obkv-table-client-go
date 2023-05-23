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
	"github.com/oceanbase/obkv-table-client-go/table"
)

func TestObGeneratedColumn_String(t *testing.T) {
	col := &obGeneratedColumn{}
	assert.Equal(t, "obGeneratedColumn{columnName:, objType:nil, collationType:0, refColumnNames:[], columnExpress:nil}", col.String())
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col = newObGeneratedColumn("c1", objType, protocol.ObCollationTypeBinary, nil, nil)
	assert.Equal(t, "obGeneratedColumn{columnName:c1, objType:ObObjType{type:ObInt64Type}, collationType:63, refColumnNames:[], columnExpress:nil}", col.String())

	col = newObGeneratedColumn("c1", objType, protocol.ObCollationTypeBinary, []string{"c1"}, nil)
	assert.Equal(t, "obGeneratedColumn{columnName:c1, objType:ObObjType{type:ObInt64Type}, collationType:63, refColumnNames:[c1], columnExpress:nil}", col.String())
}

func TestObGeneratedColumn_eval(t *testing.T) {
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObGeneratedColumn("c1", objType, protocol.ObCollationTypeBinary, nil, nil)
	_, err := col.eval([]*table.Column{table.NewColumn("c1", int64(1))})
	assert.NotEqual(t, nil, err)
}

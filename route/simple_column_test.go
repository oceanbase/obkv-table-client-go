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

func TestObSimpleColumn_String(t *testing.T) {
	col := &obSimpleColumn{}
	assert.Equal(t, "obSimpleColumn{columnName:, objType:nil, collationType:0}", col.String())
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col = newObSimpleColumn("c1", objType, protocol.ObCollationTypeBinary)
	assert.Equal(t, "obSimpleColumn{columnName:c1, objType:ObObjType{type:ObInt64Type}, collationType:63}", col.String())
}

func TestObSimpleColumn_eval(t *testing.T) {
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", objType, protocol.ObCollationTypeBinary)
	value, err := col.eval([]*table.Column{table.NewColumn("c1", int64(1))})
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, value)
	value, err = col.eval([]*table.Column{table.NewColumn("C1", int64(1))})
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, value)

	value, err = col.eval([]*table.Column{})
	assert.NotEqual(t, nil, err)
	value, err = col.eval([]*table.Column{
		table.NewColumn("C1", int64(1)),
		table.NewColumn("C2", int64(2)),
	})
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, value)
}

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

func TestObRangePartDesc_String(t *testing.T) {
	desc := &obRangePartDesc{}
	assert.Equal(t, "obRangePartDesc{partSpace:0, partNum:0, partColumns[]}", desc.String())
	desc = newObRangePartDesc(0, 10, partFuncTypeRange)
	assert.Equal(t, "obRangePartDesc{partSpace:0, partNum:10, partColumns[]}", desc.String())

	desc = newObRangePartDesc(0, 10, partFuncTypeHash)
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", objType, protocol.ObCollationTypeBinary)
	desc.SetPartColumns([]obColumn{col})
	assert.Equal(t, "obRangePartDesc{partSpace:0, partNum:10, partColumns[obSimpleColumn{columnName:c1, objType:ObObjType{type:ObInt64Type}, collationType:63}]}", desc.String())
}

func TestObRangePartDesc_GetPartId(t *testing.T) {
	desc := &obRangePartDesc{}
	partId, err := desc.GetPartId([]*table.Column{})
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)
}

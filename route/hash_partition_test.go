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

func TestObHashPartDesc_String(t *testing.T) {
	desc := &obHashPartDesc{}
	assert.Equal(t, "obHashPartDesc{comm:nil, completeWorks:[], partSpace:0, partNum:0}", desc.String())
	desc = newObHashPartDesc(0, 10, partFuncTypeHash, "c1")
	assert.Equal(t, "obHashPartDesc{comm:obPartDescCommon{partFuncType:0, partExpr:c1, orderedPartColumnNames:c1, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:nil}, completeWorks:[], partSpace:0, partNum:10}", desc.String())
}

func TestObHashPartDesc_GetPartId(t *testing.T) {
	desc := newObHashPartDesc(0, 10, partFuncTypeHash, "c1")
	partId, err := desc.GetPartId([]interface{}{1})
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)

	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeBinary)
	desc.PartColumns = []*obColumn{col}
	nameIdxMap := make(map[string]int)
	nameIdxMap["c1"] = 0
	rowkeyElement := table.NewObRowKeyElement(nameIdxMap)
	desc.setCommRowKeyElement(rowkeyElement)
	partId, err = desc.GetPartId([]interface{}{int64(1)})
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, partId)
}

func TestObHashPartDesc_innerHash(t *testing.T) {
	hashDesc := obHashPartDesc{partSpace: 0, partNum: 10}
	hashVal := hashDesc.innerHash(0)
	assert.Equal(t, int64(0), hashVal)
	hashDesc = obHashPartDesc{partSpace: 0, partNum: 10}
	hashVal = hashDesc.innerHash(1)
	assert.Equal(t, int64(1), hashVal)
	hashVal = hashDesc.innerHash(11)
	assert.Equal(t, int64(1), hashVal)
	hashVal = hashDesc.innerHash(-1)
	assert.Equal(t, int64(1), hashVal)
}

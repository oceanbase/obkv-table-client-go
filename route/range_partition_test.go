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
)

func TestObRangePartDesc_String(t *testing.T) {
	desc := &obRangePartDesc{}
	assert.Equal(t, "obRangePartDesc{comm:nil, partSpace:0, partNum:0, orderedCompareColumns:[], orderedCompareColumnTypes:[]}", desc.String())
	desc = newObRangePartDesc(0, 10, partFuncTypeRange, "c1")
	assert.Equal(t, "obRangePartDesc{comm:obPartDescCommon{partFuncType:3, partExpr:c1, orderedPartColumnNames:c1, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:nil}, partSpace:0, partNum:10, orderedCompareColumns:[], orderedCompareColumnTypes:[]}", desc.String())
}

func TestObRangePartDesc_GetPartId(t *testing.T) {
	desc := &obRangePartDesc{}
	_, err := desc.GetPartId([]interface{}{1})
	assert.NotEqual(t, nil, err)

}

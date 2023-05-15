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

func TestPartFuncType_isRangePart(t *testing.T) {
	assert.True(t, isRangePart(partFuncTypeRange))
	assert.True(t, isRangePart(partFuncTypeRangeCol))
	assert.False(t, isRangePart(partFuncTypeUnknown))
	assert.False(t, isRangePart(partFuncTypeHash))
	assert.False(t, isRangePart(partFuncTypeKey))
	assert.False(t, isRangePart(partFuncTypeKeyImpl))
	assert.False(t, isRangePart(partFuncTypeList))
	assert.False(t, isRangePart(partFuncTypeKeyV2))
	assert.False(t, isRangePart(partFuncTypeListCol))
	assert.False(t, isRangePart(partFuncTypeHashV2))
	assert.False(t, isRangePart(partFuncTypeKeyV3))
	assert.False(t, isRangePart(partFuncTypeKeyImplV2))
}

func TestPartFuncType_isKeyPart(t *testing.T) {
	assert.True(t, isKeyPart(partFuncTypeKey))
	assert.True(t, isKeyPart(partFuncTypeKeyImpl))
	assert.True(t, isKeyPart(partFuncTypeKeyV2))
	assert.True(t, isKeyPart(partFuncTypeKeyV3))
	assert.True(t, isKeyPart(partFuncTypeKeyImplV2))
	assert.False(t, isKeyPart(partFuncTypeUnknown))
	assert.False(t, isKeyPart(partFuncTypeHash))
	assert.False(t, isKeyPart(partFuncTypeList))
	assert.False(t, isKeyPart(partFuncTypeListCol))
	assert.False(t, isKeyPart(partFuncTypeHashV2))
	assert.False(t, isKeyPart(partFuncTypeRange))
	assert.False(t, isKeyPart(partFuncTypeRangeCol))
}

func TestPartFuncType_isHashPart(t *testing.T) {
	assert.True(t, isHashPart(partFuncTypeHash))
	assert.True(t, isHashPart(partFuncTypeHashV2))
	assert.False(t, isHashPart(partFuncTypeUnknown))
	assert.False(t, isHashPart(partFuncTypeKey))
	assert.False(t, isHashPart(partFuncTypeKeyImpl))
	assert.False(t, isHashPart(partFuncTypeList))
	assert.False(t, isHashPart(partFuncTypeKeyV2))
	assert.False(t, isHashPart(partFuncTypeListCol))
	assert.False(t, isHashPart(partFuncTypeKeyV3))
	assert.False(t, isHashPart(partFuncTypeKeyImplV2))
	assert.False(t, isHashPart(partFuncTypeRange))
	assert.False(t, isHashPart(partFuncTypeRangeCol))
}

func TestPartFuncType_isListPart(t *testing.T) {
	assert.True(t, isListPart(partFuncTypeList))
	assert.True(t, isListPart(partFuncTypeListCol))
	assert.False(t, isListPart(partFuncTypeUnknown))
	assert.False(t, isListPart(partFuncTypeHash))
	assert.False(t, isListPart(partFuncTypeKey))
	assert.False(t, isListPart(partFuncTypeKeyImpl))
	assert.False(t, isListPart(partFuncTypeKeyV2))
	assert.False(t, isListPart(partFuncTypeHashV2))
	assert.False(t, isListPart(partFuncTypeKeyV3))
	assert.False(t, isListPart(partFuncTypeKeyImplV2))
	assert.False(t, isListPart(partFuncTypeRange))
	assert.False(t, isListPart(partFuncTypeRangeCol))
}

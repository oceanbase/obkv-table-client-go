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

type obPartFuncType int

const (
	partFuncTypeUnknown   obPartFuncType = -1
	partFuncTypeHash      obPartFuncType = 0
	partFuncTypeKey       obPartFuncType = 1
	partFuncTypeKeyImpl   obPartFuncType = 2
	partFuncTypeRange     obPartFuncType = 3
	partFuncTypeRangeCol  obPartFuncType = 4
	partFuncTypeList      obPartFuncType = 5
	partFuncTypeKeyV2     obPartFuncType = 6
	partFuncTypeListCol   obPartFuncType = 7
	partFuncTypeHashV2    obPartFuncType = 8
	partFuncTypeKeyV3     obPartFuncType = 9
	partFuncTypeKeyImplV2 obPartFuncType = 10
)

func isRangePart(partType obPartFuncType) bool {
	return partType == partFuncTypeRange || partType == partFuncTypeRangeCol
}

func isKeyPart(partType obPartFuncType) bool {
	return partType == partFuncTypeKey ||
		partType == partFuncTypeKeyImpl ||
		partType == partFuncTypeKeyV2 ||
		partType == partFuncTypeKeyV3 ||
		partType == partFuncTypeKeyImplV2
}

func isHashPart(partType obPartFuncType) bool {
	return partType == partFuncTypeHash || partType == partFuncTypeHashV2
}

func isListPart(partType obPartFuncType) bool {
	return partType == partFuncTypeList || partType == partFuncTypeListCol
}

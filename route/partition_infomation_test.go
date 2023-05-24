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

func TestObPartitionInfo(t *testing.T) {
	info := newObPartitionInfo(PartLevelTwo)
	assert.EqualValues(t, PartLevelTwo, info.Level())
	assert.Equal(t, "obPartitionInfo{level:2, firstPartDesc:nil, subPartDesc:nil, partTabletIdMap:{}}", info.String())

	m := make(map[uint64]uint64, 1)
	m[0] = 500032
	first := &obHashPartDesc{}
	sub := &obRangePartDesc{}
	info.firstPartDesc = first
	info.subPartDesc = sub
	info.partTabletIdMap = m
	assert.Equal(t, first, info.FirstPartDesc())
	assert.Equal(t, sub, info.SubPartDesc())
	tabletId, err := info.GetTabletId(0)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 500032, tabletId)
}

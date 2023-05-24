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

func TestObTableEntry(t *testing.T) {
	entry := ObTableEntry{}
	assert.Equal(t, "ObTableEntry{tableId:0, partNum:0, replicaNum:0, refreshTimeMills:0, tableEntryKey:nil, partitionInfo:nil, tableLocation:nil, partitionEntry:nil}", entry.String())
	info := newObPartitionInfo(PartLevelTwo)
	entry = ObTableEntry{
		tableId:           500021,
		partNum:           10,
		replicaNum:        16,
		refreshTimeMills:  0,
		tableEntryKey:     NewObTableEntryKey("cluster", "tenant", "database", "table"),
		partitionInfo:     info,
		tableLocation:     nil,
		partLocationEntry: nil,
	}
	entry.SetPartLocationEntry(newObPartLocationEntry(10))
	assert.EqualValues(t, 0, entry.RefreshTimeMills())
	assert.Equal(t, (*ObTableLocation)(nil), entry.TableLocation())
	assert.EqualValues(t, 500021, entry.TableId())
	assert.Equal(t, info, entry.PartitionInfo())
	assert.Equal(t, true, entry.IsPartitionTable())
	assert.Equal(t, "ObTableEntry{tableId:500021, partNum:10, replicaNum:16, refreshTimeMills:0, tableEntryKey:ObTableEntryKey{clusterName:cluster, tenantNane:database, databaseName:database, tableName:table}, partitionInfo:obPartitionInfo{level:2, firstPartDesc:nil, subPartDesc:nil, partTabletIdMap:{}}, tableLocation:nil, partitionEntry:ObPartLocationEntry{partLocations:{}}}", entry.String())
}

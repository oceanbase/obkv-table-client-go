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
	"strconv"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableEntry struct {
	tableId           uint64
	partNum           int
	replicaNum        int
	refreshTimeMills  int64
	tableEntryKey     ObTableEntryKey
	partitionInfo     *obPartitionInfo
	tableLocation     *ObTableLocation
	partLocationEntry *ObPartLocationEntry
}

func (e *ObTableEntry) SetPartLocationEntry(partLocationEntry *ObPartLocationEntry) {
	e.partLocationEntry = partLocationEntry
}

func (e *ObTableEntry) RefreshTimeMills() int64 {
	return e.refreshTimeMills
}

func (e *ObTableEntry) TableLocation() *ObTableLocation {
	return e.tableLocation
}

func (e *ObTableEntry) TableId() uint64 {
	return e.tableId
}

func (e *ObTableEntry) PartitionInfo() *obPartitionInfo {
	return e.partitionInfo
}

func (e *ObTableEntry) IsPartitionTable() bool {
	return e.partNum > 1
}

func (e *ObTableEntry) String() string {
	// partitionInfo to string
	var partitionInfoStr string
	if e.partitionInfo != nil {
		partitionInfoStr = e.partitionInfo.String()
	} else {
		partitionInfoStr = "nil"
	}

	// tableLocation to string
	var tableLocationStr string
	if e.tableLocation != nil {
		tableLocationStr = e.tableLocation.String()
	} else {
		tableLocationStr = "nil"
	}

	// partLocationEntry to string
	var partLocationEntryStr string
	if e.partLocationEntry != nil {
		partLocationEntryStr = e.partLocationEntry.String()
	} else {
		partLocationEntryStr = "nil"
	}
	return "ObTableEntry{" +
		"tableId:" + strconv.Itoa(int(e.tableId)) + ", " +
		"partNum:" + strconv.Itoa(int(e.partNum)) + ", " +
		"replicaNum:" + strconv.Itoa(int(e.replicaNum)) + ", " +
		"refreshTimeMills:" + strconv.Itoa(int(e.refreshTimeMills)) + ", " +
		"tableEntryKey:" + e.tableEntryKey.String() + ", " +
		"partitionInfo:" + partitionInfoStr + ", " +
		"tableLocation:" + tableLocationStr + ", " +
		"partitionEntry:" + partLocationEntryStr +
		"}"
}

func (e *ObTableEntry) extractSubpartIdx(id int64) int64 {
	// equal id & (^(0xffffffffffffffff << ObPartIdShift)) & (^(0xffffffffffffffff << ObPartIdBitNum))
	return id & ObSubPartIdMask
}

func (e *ObTableEntry) getPartitionLocation(partId int64, consistency ObConsistency) (*obReplicaLocation, error) {
	if util.ObVersion() >= 4 && e.IsPartitionTable() {
		tabletId, ok := e.partitionInfo.partTabletIdMap[partId]
		if !ok {
			return nil, errors.Errorf("tablet id not found, partId:%d, partInfo:%s", partId, e.partitionInfo.String())
		}
		partLoc, ok := e.partLocationEntry.partLocations[tabletId]
		if !ok {
			return nil, errors.Errorf("part location not found, tabletId:%d, partLocationEntry:%s",
				tabletId, e.partLocationEntry.String())
		}
		return partLoc.getReplica(consistency), nil
	} else {
		partLoc, ok := e.partLocationEntry.partLocations[partId]
		if !ok {
			return nil, errors.Errorf("part location not found, partId:%d, partLocationEntry:%s",
				partId, e.partLocationEntry.String())
		}
		return partLoc.getReplica(consistency), nil
	}
}

func (e *ObTableEntry) GetPartitionReplicaLocation(partId int64, consistency ObConsistency) (*obReplicaLocation, error) {
	logicId := partId
	if e.partitionInfo != nil && e.partitionInfo.level == PartLevelTwo {
		logicId = e.extractSubpartIdx(partId)
	}
	return e.getPartitionLocation(logicId, consistency)
}

func (e *ObTableEntry) SetRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	if e.partitionInfo != nil {
		e.partitionInfo.setRowKeyElement(rowKeyElement)
	}
}

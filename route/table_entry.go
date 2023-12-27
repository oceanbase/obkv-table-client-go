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
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/util"
)

const refreshInterval = 5 // 5 second

// ObTableEntry represents all the routing information of a table.
type ObTableEntry struct {
	tableId           uint64
	partNum           int
	replicaNum        int
	refreshTime       time.Time            // last refresh time
	tableEntryKey     *ObTableEntryKey     // clusterName/tenantName/databaseName/tableName
	partitionInfo     *obPartitionInfo     // partition key meta info
	tableLocation     *ObTableLocation     // location of table, all replica information of table
	partLocationEntry *ObPartLocationEntry // all partition location of table
}

func (e *ObTableEntry) SetPartLocationEntry(partLocationEntry *ObPartLocationEntry) {
	e.partLocationEntry = partLocationEntry
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

func (e *ObTableEntry) NeedRefresh() bool {
	return time.Now().Sub(e.refreshTime).Seconds() > refreshInterval
}

func (e *ObTableEntry) String() string {
	// tableEntryKey to string
	var tableEntryKeyStr string
	if e.tableEntryKey != nil {
		tableEntryKeyStr = e.tableEntryKey.String()
	} else {
		tableEntryKeyStr = "nil"
	}

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

	return "ObTableEntry{" +
		"tableId:" + strconv.Itoa(int(e.tableId)) + ", " +
		"partNum:" + strconv.Itoa(e.partNum) + ", " +
		"replicaNum:" + strconv.Itoa(e.replicaNum) + ", " +
		"refreshTime:" + e.refreshTime.String() + ", " +
		"tableEntryKey:" + tableEntryKeyStr + ", " +
		"partitionInfo:" + partitionInfoStr + ", " +
		"tableLocation:" + tableLocationStr +
		"}"
}

func (e *ObTableEntry) extractSubpartIdx(id uint64) uint64 {
	// equal id & (^(0xffffffffffffffff << ObPartIdShift)) & (^(0xffffffffffffffff << ObPartIdBitNum))
	return id & ObSubPartIdMask
}

// get part_id with PARTITION_LEVEL_TWO_MASK
func (e *ObTableEntry) extractPartId(id uint64) uint64 {
	return id >> ObPartIdShift
}

// get first part id
func (e *ObTableEntry) extractPartIdx(id uint64) uint64 {
	return e.extractPartId(id) & (0xfffffff)
}

// get part idx in second partition
func (e *ObTableEntry) getPartIdx(id uint64) uint64 {
	return e.extractPartIdx(id)*(uint64(e.partitionInfo.subPartDesc.PartNum())) + e.extractSubpartIdx(id)
}

// GetPartitionLocation get partition location by partId and consistency.
func (e *ObTableEntry) GetPartitionLocation(partId uint64, consistency ObConsistency) (*obReplicaLocation, error) {
	if util.ObVersion() >= 4 && e.IsPartitionTable() {
		logicId := partId
		if e.partitionInfo != nil && e.partitionInfo.level == PartLevelTwo {
			logicId = e.getPartIdx(partId)
		}
		// In ob version 4.0 and above, get tabletId firstly.
		// Because in version 4.0 and above we set up a relationship between tabletId and Location.
		tabletId, ok := e.partitionInfo.partTabletIdMap[logicId]
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
		// Below version 4.0 we set up a relationship between partId and Location.
		partLoc, ok := e.partLocationEntry.partLocations[partId]
		if !ok {
			return nil, errors.Errorf("part location not found, partId:%d, partLocationEntry:%s",
				partId, e.partLocationEntry.String())
		}
		return partLoc.getReplica(consistency), nil
	}
}

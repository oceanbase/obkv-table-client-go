package route

import (
	"errors"
	"strconv"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableEntry struct {
	tableId           uint64
	partNum           int
	replicaNum        int
	refreshTimeMills  int64
	tableEntryKey     ObTableEntryKey
	partitionInfo     *ObPartitionInfo
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

func (e *ObTableEntry) PartitionInfo() *ObPartitionInfo {
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

func (e *ObTableEntry) getPartitionLocation(partId int64, route *ObServerRoute) (*ObReplicaLocation, error) {
	if util.ObVersion() >= 4 && e.IsPartitionTable() {
		tabletId, ok := e.partitionInfo.partTabletIdMap[partId]
		if !ok {
			log.Warn("tablet id not found",
				log.Int64("part id", partId),
				log.String("part info", e.partitionInfo.String()))
			return nil, errors.New("tablet id not found")
		}
		partLoc, ok := e.partLocationEntry.partLocations[tabletId]
		if !ok {
			log.Warn("part location not found",
				log.Int64("tabletId", tabletId),
				log.String("part entry", e.partLocationEntry.String()))
			return nil, errors.New("part location not found")
		}
		return partLoc.getReplica(route), nil
	} else {
		partLoc, ok := e.partLocationEntry.partLocations[partId]
		if !ok {
			log.Warn("part location not found",
				log.Int64("partId", partId),
				log.String("part entry", e.partLocationEntry.String()))
			return nil, errors.New("part location not found")
		}
		return partLoc.getReplica(route), nil
	}
}

func (e *ObTableEntry) GetPartitionReplicaLocation(partId int64, route *ObServerRoute) (*ObReplicaLocation, error) {
	logicId := partId
	if e.partitionInfo != nil && e.partitionInfo.level.index == PartLevelTwoIndex {
		logicId = e.extractSubpartIdx(partId)
	}
	return e.getPartitionLocation(logicId, route)
}

func (e *ObTableEntry) SetRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	if e.partitionInfo != nil {
		e.partitionInfo.setRowKeyElement(rowKeyElement)
	}
}

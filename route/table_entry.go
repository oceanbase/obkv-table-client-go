package route

import (
	"errors"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
	"strconv"
)

type ObReplicaLocation struct {
	addr        ObServerAddr
	info        ObServerInfo
	role        ObServerRole
	replicaType ObReplicaType
}

func (l *ObReplicaLocation) Info() ObServerInfo {
	return l.info
}

func (l *ObReplicaLocation) Addr() *ObServerAddr {
	return &l.addr
}

func (l *ObReplicaLocation) isValid() bool {
	return !l.role.isInvalid() && l.info.IsActive()
}

func (l *ObReplicaLocation) isLeader() bool {
	return l.role.isLeader()
}

func (l *ObReplicaLocation) String() string {
	return "ObReplicaLocation{" +
		"addr:" + l.addr.String() + ", " +
		"info:" + l.info.String() + ", " +
		"role:" + l.role.String() + ", " +
		"replicaType:" + l.replicaType.String() +
		"}"
}

type ObTableLocation struct {
	replicaLocations []ObReplicaLocation
}

func (l *ObTableLocation) ReplicaLocations() []ObReplicaLocation {
	return l.replicaLocations
}

func (l *ObTableLocation) String() string {
	var replicaLocationsStr string
	replicaLocationsStr = replicaLocationsStr + "["
	for i := 0; i < len(l.replicaLocations); i++ {
		if i > 0 {
			replicaLocationsStr += ", "
		}
		replicaLocationsStr += l.replicaLocations[i].String()
	}
	replicaLocationsStr += "]"
	return "ObTableLocation{" +
		"replicaLocations:" + replicaLocationsStr +
		"}"
}

type ObPartitionInfo struct {
	level           ObPartitionLevel
	firstPartDesc   ObPartDesc
	subPartDesc     ObPartDesc
	partColumns     []*protocol.ObColumn
	partTabletIdMap map[int64]int64
	partNameIdMap   map[string]int64
}

func (p *ObPartitionInfo) SubPartDesc() ObPartDesc {
	return p.subPartDesc
}

func (p *ObPartitionInfo) FirstPartDesc() ObPartDesc {
	return p.firstPartDesc
}

func (p *ObPartitionInfo) GetTabletId(partId int64) (int64, error) {
	if p.partTabletIdMap == nil {
		log.Warn("partTabletIdMap is nil")
		return 0, errors.New("partTabletIdMap is nil")
	}
	return p.partTabletIdMap[partId], nil
}

func (p *ObPartitionInfo) Level() ObPartitionLevel {
	return p.level
}

func (p *ObPartitionInfo) setRowKeyElement(rowKeyElement *table.ObRowkeyElement) {
	if p.firstPartDesc != nil {
		p.firstPartDesc.setRowKeyElement(rowKeyElement)
	}
	if p.subPartDesc != nil {
		p.subPartDesc.setRowKeyElement(rowKeyElement)
	}
}

func (p *ObPartitionInfo) String() string {
	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(p.partColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += p.partColumns[i].String()
	}
	partColumnsStr += "]"

	// partTabletIdMap to string
	var partTabletIdMapStr string
	var i = 0
	partTabletIdMapStr = partTabletIdMapStr + "{"
	for k, v := range p.partTabletIdMap {
		if i > 0 {
			partTabletIdMapStr += ", "
		}
		i++
		partTabletIdMapStr += "m[" + strconv.Itoa(int(k)) + "]=" + strconv.Itoa(int(v))
	}
	partTabletIdMapStr += "}"

	// partNameIdMap to string
	var partNameIdMapStr string
	i = 0
	partNameIdMapStr = partNameIdMapStr + "{"
	for k, v := range p.partNameIdMap {
		if i > 0 {
			partNameIdMapStr += ", "
		}
		i++
		partNameIdMapStr += "m[" + k + "]=" + strconv.Itoa(int(v))
	}
	partNameIdMapStr += "}"

	// firstPartDesc to string
	var firstPartDescStr string
	if p.firstPartDesc != nil {
		firstPartDescStr = p.firstPartDesc.String()
	} else {
		firstPartDescStr = "nil"
	}

	// subPartDesc to string
	var subPartDescStr string
	if p.subPartDesc != nil {
		subPartDescStr = p.firstPartDesc.String()
	} else {
		subPartDescStr = "nil"
	}

	return "ObPartitionInfo{" +
		"level:" + p.level.String() + ", " +
		"firstPartDesc:" + firstPartDescStr + ", " +
		"subPartDesc:" + subPartDescStr + ", " +
		"partColumns:" + partColumnsStr + ", " +
		"partTabletIdMap:" + partTabletIdMapStr + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

type ObPartitionLocation struct {
	leader   ObReplicaLocation
	replicas []ObReplicaLocation
}

func (l *ObPartitionLocation) addReplicaLocation(replica *ObReplicaLocation) {
	if replica.isLeader() {
		l.leader = *replica
	}
	l.replicas = append(l.replicas, *replica)
}

func (l *ObPartitionLocation) getReplica(route *ObServerRoute) *ObReplicaLocation {
	if route.readConsistency == ObReadConsistencyStrong {
		return &l.leader
	}
	// todo:weak read by LDC
	return &l.leader
}

func (l *ObPartitionLocation) String() string {
	var replicasStr string
	replicasStr = replicasStr + "["
	for i := 0; i < len(l.replicas); i++ {
		if i > 0 {
			replicasStr += ", "
		}
		replicasStr += l.replicas[i].String()
	}
	replicasStr += "]"
	return "ObPartitionLocation{" +
		"leader:" + l.leader.String() + ", " +
		"replicas:" + replicasStr +
		"}"
}

type ObPartLocationEntry struct {
	partLocations map[int64]*ObPartitionLocation
}

func newObPartLocationEntry(partNum int) *ObPartLocationEntry {
	entry := new(ObPartLocationEntry)
	entry.partLocations = make(map[int64]*ObPartitionLocation, partNum)
	return entry
}

func (e *ObPartLocationEntry) String() string {
	var partitionLocationStr string
	var i = 0
	partitionLocationStr = partitionLocationStr + "{"
	for k, v := range e.partLocations {
		if i > 0 {
			partitionLocationStr += ", "
		}
		i++
		partitionLocationStr += "m[" + strconv.Itoa(int(k)) + "]="
		if v != nil {
			partitionLocationStr += v.String()
		} else {
			partitionLocationStr += "nil"
		}
	}
	partitionLocationStr += "}"
	return "ObPartLocationEntry{" +
		"partLocations:" + partitionLocationStr +
		"}"
}

type ObTableEntryKey struct {
	clusterName  string
	tenantName   string
	databaseName string
	tableName    string
}

func NewObTableEntryKey(
	clusterName string,
	tenantName string,
	databaseName string,
	tableName string) *ObTableEntryKey {
	return &ObTableEntryKey{clusterName, tenantName, databaseName, tableName}
}

func (k *ObTableEntryKey) String() string {
	return "ObTableEntryKey{" +
		"clusterName:" + k.clusterName + ", " +
		"tenantNane:" + k.databaseName + ", " +
		"databaseName:" + k.databaseName + ", " +
		"tableName:" + k.tableName +
		"}"
}

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

func (e *ObTableEntry) SetRowKeyElement(rowKeyElement *table.ObRowkeyElement) {
	if e.partitionInfo != nil {
		e.partitionInfo.setRowKeyElement(rowKeyElement)
	}
}

const (
	ObReadConsistencyStrong = 0
	ObReadConsistencyWeak   = 1
)

type ObReadConsistency int

type ObServerRoute struct {
	readConsistency ObReadConsistency
}

func (r *ObServerRoute) String() string {
	return "ObServerRoute{" +
		"readConsistency:" + strconv.Itoa(int(r.readConsistency)) +
		"}"
}

func NewObServerRoute(readOnly bool) *ObServerRoute {
	if readOnly {
		// todo: adapt java client
		return &ObServerRoute{ObReadConsistencyWeak}
	}
	return &ObServerRoute{ObReadConsistencyStrong}
}

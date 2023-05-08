package route

import (
	"errors"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/util"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	OceanbaseDatabase = "OCEANBASE"
	AllDummyTable     = "__all_dummy"
)

const (
	ObVersionSql     = "SELECT /*+READ_CONSISTENCY(WEAK)*/ OB_VERSION() AS CLUSTER_VERSION;"
	DummyLocationSql = "SELECT /*+READ_CONSISTENCY(WEAK)*/ A.partition_id as partition_id, A.svr_ip as svr_ip, " +
		"A.sql_port as sql_port, A.table_id as table_id, A.role as role, A.replica_num as replica_num, A.part_num as part_num, " +
		"B.svr_port as svr_port, B.status as status, B.stop_time as stop_time, A.spare1 as replica_type " +
		"FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip " +
		"and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ?;"
	DummyLocationSqlV4 = "SELECT /*+READ_CONSISTENCY(WEAK)*/ A.tablet_id as tablet_id, A.svr_ip as svr_ip, " +
		"A.sql_port as sql_port, A.table_id as table_id, A.role as role, A.replica_num as replica_num, A.part_num as part_num, " +
		"B.svr_port as svr_port, B.status as status, B.stop_time as stop_time, A.spare1 as replica_type " +
		"FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip " +
		"and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ?;"

	LocationSql = "SELECT /*+READ_CONSISTENCY(WEAK)*/ A.partition_id as partition_id, A.svr_ip as svr_ip, " +
		"A.sql_port as sql_port, A.table_id as table_id, A.role as role, A.replica_num as replica_num, A.part_num as part_num, " +
		"B.svr_port as svr_port, B.status as status, B.stop_time as stop_time, A.spare1 as replica_type " +
		"FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip " +
		"and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ? and partition_id = 0;"
	LocationSqlV4 = "SELECT /*+READ_CONSISTENCY(WEAK)*/ A.tablet_id as tablet_id, A.svr_ip as svr_ip, " +
		"A.sql_port as sql_port, A.table_id as table_id, A.role as role, A.replica_num as replica_num, A.part_num as part_num, " +
		"B.svr_port as svr_port, B.status as status, B.stop_time as stop_time, A.spare1 as replica_type " +
		"FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip " +
		"and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ? and tablet_id = 0;"

	PartitionLocationSql = "SELECT /*+READ_CONSISTENCY(WEAK)*/ A.partition_id as partition_id, A.svr_ip as svr_ip, " +
		"A.sql_port as sql_port, A.table_id as table_id, A.role as role, A.replica_num as replica_num, A.part_num as part_num, " +
		"B.svr_port as svr_port, B.status as status, B.stop_time as stop_time, A.spare1 as replica_type " +
		"FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip " +
		"and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ? and partition_id in"
	PartitionLocationSqlV4 = "SELECT /*+READ_CONSISTENCY(WEAK)*/ A.tablet_id as tablet_id, A.svr_ip as svr_ip, " +
		"A.sql_port as sql_port, A.table_id as table_id, A.role as role, A.replica_num as replica_num, A.part_num as part_num, " +
		"B.svr_port as svr_port, B.status as status, B.stop_time as stop_time, A.spare1 as replica_type " +
		"FROM oceanbase.__all_virtual_proxy_schema A inner join oceanbase.__all_server B on A.svr_ip = B.svr_ip " +
		"and A.sql_port = B.inner_port WHERE tenant_name = ? and database_name = ? and table_name = ? and tablet_id in"

	PartitionInfoSql = "SELECT /*+READ_CONSISTENCY(WEAK)*/ part_level, part_num, part_type, part_space, part_expr, " +
		"part_range_type, sub_part_num, sub_part_type, sub_part_space, sub_part_range_type, sub_part_expr, part_key_name, " +
		"part_key_type, part_key_idx, part_key_extra, spare1 FROM oceanbase.__all_virtual_proxy_partition_info " +
		"WHERE table_id = ? group by part_key_name order by part_key_name LIMIT ?;"
	PartitionInfoSqlV4 = "SELECT /*+READ_CONSISTENCY(WEAK)*/ part_level, part_num, part_type, part_space, part_expr, " +
		"part_range_type, sub_part_num, sub_part_type, sub_part_space, sub_part_range_type, sub_part_expr, part_key_name, " +
		"part_key_type, part_key_idx, part_key_extra, part_key_collation_type FROM oceanbase.__all_virtual_proxy_partition_info " +
		"WHERE tenant_name = ? and table_id = ? group by part_key_name order by part_key_name LIMIT ?;"

	FirstPartitionSql = "SELECT /*+READ_CONSISTENCY(WEAK)*/ part_id, part_name, high_bound_val " +
		"FROM oceanbase.__all_virtual_proxy_partition WHERE table_id = ? LIMIT ?;"
	FirstPartitionSqlV4 = "SELECT /*+READ_CONSISTENCY(WEAK)*/ part_id, part_name, tablet_id, high_bound_val, sub_part_num " +
		"FROM oceanbase.__all_virtual_proxy_partition WHERE tenant_name = ? and table_id = ? LIMIT ?;"

	SubPartitionSql = "SELECT /*+READ_CONSISTENCY(WEAK)*/ sub_part_id, part_name, high_bound_val " +
		"FROM oceanbase.__all_virtual_proxy_sub_partition WHERE table_id = ? LIMIT ?;"
	SubPartitionSqlV4 = "SELECT /*+READ_CONSISTENCY(WEAK)*/ sub_part_id, part_name, tablet_id, high_bound_val " +
		"FROM oceanbase.__all_virtual_proxy_sub_partition WHERE tenant_name = ? and table_id = ? LIMIT ?;"
)

var (
	proxySqlGuard             sync.Mutex
	proxyDummyLocationSql     string
	proxyLocationSql          string
	proxyPartitionLocationSql string
	proxyPartitionInfoSql     string
	proxyFirstPartitionSql    string
	proxySubPartitionSql      string
)

func InitSql(obVersion float32) {
	proxySqlGuard.Lock()
	if obVersion >= 4 {
		proxyDummyLocationSql = DummyLocationSqlV4
		proxyLocationSql = LocationSqlV4
		proxyPartitionLocationSql = PartitionLocationSqlV4
		proxyPartitionInfoSql = PartitionInfoSqlV4
		proxyFirstPartitionSql = FirstPartitionSqlV4
		proxySubPartitionSql = SubPartitionSqlV4
	} else {
		proxyDummyLocationSql = DummyLocationSql
		proxyLocationSql = LocationSql
		proxyPartitionLocationSql = PartitionLocationSql
		proxyPartitionInfoSql = PartitionInfoSql
		proxyFirstPartitionSql = FirstPartitionSql
		proxySubPartitionSql = SubPartitionSql
	}
	proxySqlGuard.Unlock()
}

// GetObVersionFromRemote get OceanBase cluster version by sql
// called when client init
func GetObVersionFromRemote(addr *ObServerAddr, sysUA *ObUserAuth) (float32, error) {
	// 1. Get db handle.
	db, err := NewDB(
		sysUA.userName,
		sysUA.password,
		addr.ip,
		strconv.Itoa(addr.sqlPort),
		OceanbaseDatabase,
	)
	if err != nil {
		log.Warn("failed to new db",
			log.String("sysUA", sysUA.String()),
			log.String("addr", addr.String()))
		return 0.0, err
	}
	defer func() {
		_ = db.Close()
	}()

	// 2. Prepare get observer version sql statement.
	stmt, err := db.Prepare(ObVersionSql)
	if err != nil {
		log.Warn("fail to prepare get observer version sql", log.String("sql", ObVersionSql))
		return 0.0, err
	}

	// 3. Get result from query row.
	var obVersionStr string
	err = stmt.QueryRow().Scan(&obVersionStr)
	if err != nil {
		log.Warn("fail to get observer version from query result", log.String("sql", ObVersionSql))
		return 0.0, err
	}

	// 4. parse ob version string
	// +-----------------+
	// | CLUSTER_VERSION |
	// +-----------------+
	// | 4.1.0.0         |
	str := strings.ReplaceAll(obVersionStr, ".", "") // 4100
	ver, err := strconv.Atoi(str)
	if err != nil {
		log.Warn("fail to convert string to int", log.String("str", str))
		return 0.0, err
	}
	res := float32(ver) / 1000.0 // ObVersion = 4.1
	return res, nil
}

func GetTableEntryFromRemote(
	addr *ObServerAddr,
	sysUA *ObUserAuth,
	key *ObTableEntryKey) (*ObTableEntry, error) {
	// 1. Get db handle
	db, err := NewDB(
		sysUA.userName,
		sysUA.password,
		addr.ip,
		strconv.Itoa(addr.sqlPort),
		OceanbaseDatabase,
	)
	if err != nil {
		log.Warn("failed to new db",
			log.String("sysUA", sysUA.String()),
			log.String("addr", addr.String()))
		return nil, err
	}
	defer func() {
		_ = db.Close()
	}()

	// 2. Do query with specific tenant name，database name and table name.
	// proxyDummyLocationSql for getting all tenant server ip address
	// proxyLocationSql for getting all table replicas ip address
	var sql string
	if key.tableName == AllDummyTable {
		sql = proxyDummyLocationSql
	} else {
		sql = proxyLocationSql
	}
	rows, err := db.Query(sql, key.tenantName, key.databaseName, key.tableName)
	if err != nil {
		log.Warn("failed to do query", log.String("sql", sql))
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	// 3. Create table entry by parsing query result set.
	entry, err := getTableEntryFromResultSet(rows)
	if err != nil {
		log.Warn("failed to get table entry from result set", log.String("key", key.String()))
		return nil, err
	}
	entry.tableEntryKey = *key

	// 4. Fetch partition info
	if entry.IsPartitionTable() {
		info, err := getPartitionInfoFromRemote(db, key.tenantName, entry.tableId)
		if err != nil {
			log.Warn("failed to fetch partition info",
				log.String("tenant", key.tenantName),
				log.Uint64("tableId", entry.tableId))
			return nil, err
		}
		entry.partitionInfo = info

		// 4.1. Fetch first partition info
		if info.level.index >= 1 {
			err = fetchFirstPart(db, info.firstPartDesc.partFuncType(), entry)
			if err != nil {
				log.Warn("failed to fetch first partition info",
					log.String("entry", entry.String()))
				return nil, err
			}
		}

		// 4.2. Fetch sub partition info
		if info.level.index == 2 {
			err = fetchSubPart(db, info.subPartDesc.partFuncType(), entry)
			if err != nil {
				log.Warn("failed to fetch sub partition info",
					log.String("entry", entry.String()))
				return nil, err
			}
		}

		entry.partitionInfo = info
	}

	// 5. Get partition location entry
	partLocationEntry, err := GetPartLocationEntryFromRemote(db, entry)
	if err != nil {
		log.Warn("failed to get table entry location", log.String("entry", entry.String()))
		return nil, err
	}
	entry.partLocationEntry = partLocationEntry
	entry.refreshTimeMills = time.Time{}.Unix()
	return entry, nil
}

func getTableEntryFromResultSet(rows *Rows) (*ObTableEntry, error) {
	tableLocation := new(ObTableLocation)
	entry := new(ObTableEntry)
	var (
		partitionId int
		svrIp       string
		sqlPort     int
		tableId     uint64
		role        int
		replicaNum  int
		partNum     int
		svrPort     int
		status      string
		stopTime    int64
		replicaType int
	)

	// 1. get replica location info from query rows
	for rows.Next() {
		err := rows.Scan(
			&partitionId,
			&svrIp,
			&sqlPort,
			&tableId,
			&role,
			&replicaNum,
			&partNum,
			&svrPort,
			&status,
			&stopTime,
			&replicaType,
		)
		if err != nil {
			log.Warn("failed to scan row")
			return nil, err
		}
		svrRole := newObServerRole(role)
		svrReplicaType := newObReplicaType(replicaType)
		svrAddr := ObServerAddr{ip: svrIp, sqlPort: sqlPort, svrPort: svrPort}
		svrInfo := ObServerInfo{stopTime: stopTime, status: status}
		replica := &ObReplicaLocation{addr: svrAddr, info: svrInfo, role: svrRole, replicaType: svrReplicaType}
		if !replica.isValid() {
			log.Warn("replica is invalid", log.String("replica", replica.String()))
			return nil, errors.New("replica is invalid")
		}
		tableLocation.replicaLocations = append(tableLocation.replicaLocations, *replica)
	}

	// 2. fill table entry
	entry.tableId = tableId
	entry.partNum = partNum
	entry.replicaNum = replicaNum
	entry.tableLocation = tableLocation

	return entry, nil
}

func GetPartLocationEntryFromRemote(db *DB, entry *ObTableEntry) (*ObPartLocationEntry, error) {
	// 1. Create inStatement "(0,1,2...partNum);".
	partIds := make([]int, 0, entry.partNum)
	if util.ObVersion() >= 4 && entry.IsPartitionTable() {
		for _, v := range entry.partitionInfo.partTabletIdMap {
			partIds = append(partIds, int(v))
		}
		sort.Ints(partIds) // partIds doesn't have to be ascending, so do sort
	} else {
		for i := 0; i < entry.partNum; i++ {
			partIds = append(partIds, i)
		}
	}
	inStatement := CreateInStatement(partIds)

	// 2. Do query with specific tenant name，database name and table name.
	sql := proxyPartitionLocationSql + inStatement
	key := entry.tableEntryKey
	rows, err := db.Query(sql, key.tenantName, key.databaseName, key.tableName)
	if err != nil {
		log.Warn("failed to do query", log.String("sql", sql))
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	// 3. create ObPartitionEntry by parsing query result set
	partLocationEntry, err := getPartLocationEntryFromResultSet(rows)
	if err != nil {
		log.Warn("failed to get partition location from result set")
		return nil, err
	}

	return partLocationEntry, nil
}

func getPartLocationEntryFromResultSet(rows *Rows) (*ObPartLocationEntry, error) {
	var partitionEntry *ObPartLocationEntry = nil
	isFirstRow := true
	var (
		partitionId int
		svrIp       string
		sqlPort     int
		tableId     uint64
		role        int
		replicaNum  int
		partNum     int
		svrPort     int
		status      string
		stopTime    int64
		replicaType int
	)
	for rows.Next() {
		err := rows.Scan(
			&partitionId,
			&svrIp,
			&sqlPort,
			&tableId,
			&role,
			&replicaNum,
			&partNum,
			&svrPort,
			&status,
			&stopTime,
			&replicaType,
		)
		if err != nil {
			log.Warn("failed to scan row")
			return nil, err
		}

		// create ObPartLocationEntry
		if isFirstRow {
			isFirstRow = false
			partitionEntry = newObPartLocationEntry(partNum)
		}

		// 1. create ObReplicaLocation
		svrRole := newObServerRole(role)
		svrReplicaType := newObReplicaType(replicaType)
		svrAddr := ObServerAddr{ip: svrIp, sqlPort: sqlPort, svrPort: svrPort}
		svrInfo := ObServerInfo{stopTime: stopTime, status: status}
		replica := &ObReplicaLocation{addr: svrAddr, info: svrInfo, role: svrRole, replicaType: svrReplicaType}

		// 2. find or create location, then add replica location
		location, ok := partitionEntry.partLocations[int64(partitionId)]
		if !ok {
			location = new(ObPartitionLocation)
			partitionEntry.partLocations[int64(partitionId)] = location
		}
		location.addReplicaLocation(replica)
	}
	return partitionEntry, nil
}

func getPartitionInfoFromRemote(db *DB, tenantName string, tableId uint64) (*ObPartitionInfo, error) {
	// 1. Do query with specific tenant name(ob version > 4), table id and limit.
	var rows *Rows
	var err error
	if util.ObVersion() >= 4 {
		rows, err = db.Query(proxyPartitionInfoSql, tenantName, tableId, math.MaxInt64)
	} else {
		rows, err = db.Query(proxyPartitionInfoSql, tableId, math.MaxInt64)
	}
	if err != nil {
		log.Warn("failed to do query",
			log.String("sql", proxyPartitionInfoSql),
			log.Uint64("tableId", tableId))
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	// 2. create ObPartitionInfo by parsing query result set
	info, err := getPartitionInfoFromResultSet(rows)
	if err != nil {
		log.Warn("failed to get partition info from result set")
		return nil, err
	}

	return info, nil
}

func getPartitionInfoFromResultSet(rows *Rows) (*ObPartitionInfo, error) {
	info := new(ObPartitionInfo)
	var (
		partLevel        int
		partNum          int
		partType         int
		partSpace        int
		partExpr         string
		partRangeType    string
		subPartNum       int
		subPartType      int
		subPartSpace     int
		subPartRangeType string
		subPartExpr      string
		partKeyName      string
		partKeyType      int
		partKeyIdx       int
		partKeyExtra     string
		spare1           int // Also "part_key_collation_type" when ob version >= 4. They mean collation type.
	)
	var isFirstRow = true
	for rows.Next() {
		err := rows.Scan(
			&partLevel,
			&partNum,
			&partType,
			&partSpace,
			&partExpr,
			&partRangeType,
			&subPartNum,
			&subPartType,
			&subPartSpace,
			&subPartRangeType,
			&subPartExpr,
			&partKeyName,
			&partKeyType,
			&partKeyIdx,
			&partKeyExtra,
			&spare1,
		)
		if err != nil {
			log.Warn("failed to scan row")
			return nil, err
		}
		if isFirstRow {
			isFirstRow = false
			info.level = newObPartitionLevel(partLevel)
			// build first part
			if info.level.index >= PartLevelOneIndex {
				firstPartDesc, err := buildPartDesc(
					partNum,
					partType,
					partSpace,
					partExpr,
					partRangeType,
				)
				if err != nil {
					log.Warn("failed to build first part desc")
					return nil, err
				}
				info.firstPartDesc = firstPartDesc
			}

			// build sub part
			if info.level.index == PartLevelTwoIndex {
				subPartDesc, err := buildPartDesc(
					subPartNum,
					subPartType,
					subPartSpace,
					subPartExpr,
					subPartRangeType,
				)
				if err != nil {
					log.Warn("failed to build sub part desc")
					return nil, err
				}
				info.subPartDesc = subPartDesc
			}
		}

		partKeyExtra = strings.ReplaceAll(partKeyExtra, "`", "") // '`' is not supported by druid
		partKeyExtra = strings.ReplaceAll(partKeyExtra, " ", "") // ' ' should be removed
		var column *protocol.ObColumn
		if partKeyExtra != "" {
			// todo: support generate column
			return nil, errors.New("not impl generate column")
		} else {
			objType, err := protocol.NewObObjType(partKeyType)
			if err != nil {
				log.Warn("failed to generate object type", log.Int("partKeyType", partKeyType))
				return nil, err
			}
			column = protocol.NewObSimpleColumn(
				partKeyName,
				partKeyIdx,
				objType,
				protocol.NewObCollationType(spare1),
			)
		}
		info.partColumns = append(info.partColumns, column)
	}
	info.partTabletIdMap = make(map[int64]int64, partNum)
	info.partNameIdMap = make(map[string]int64, partNum)

	var orderedPartedColumns1 []*protocol.ObColumn
	if info.level.index >= PartLevelOneIndex {
		if info.firstPartDesc.partFuncType().isListPart() ||
			info.firstPartDesc.partFuncType().isRangePart() {
			orderedPartedColumns1 = getOrderedPartColumns(info.partColumns, info.firstPartDesc)
		}
		// set the property of first part
		err := setPartDescProperty(info.firstPartDesc, info.partColumns, orderedPartedColumns1)
		if err != nil {
			log.Warn("failed to ser first part property", log.String("part info", info.String()))
			return nil, err
		}
	}

	var orderedPartedColumns2 []*protocol.ObColumn
	if info.level.index == PartLevelTwoIndex {
		if info.firstPartDesc.partFuncType().isListPart() ||
			info.firstPartDesc.partFuncType().isRangePart() {
			orderedPartedColumns2 = getOrderedPartColumns(info.partColumns, info.subPartDesc)
		}
		// set the property of sub part
		err := setPartDescProperty(info.subPartDesc, info.partColumns, orderedPartedColumns2)
		if err != nil {
			log.Warn("failed to ser sub part property", log.String("part info", info.String()))
			return nil, err
		}
	}

	return info, nil
}

func getOrderedPartColumns(
	partitionKeyColumns []*protocol.ObColumn,
	partDesc ObPartDesc) []*protocol.ObColumn {
	columns := make([]*protocol.ObColumn, 0, len(partitionKeyColumns))
	for _, partColumnName := range partDesc.orderedPartColumnNames() {
		for _, keyColumn := range partitionKeyColumns {
			if strings.EqualFold(keyColumn.ColumnName(), partColumnName) {
				columns = append(columns, keyColumn)
			}
		}
	}
	return columns
}

func setPartDescProperty(
	partDesc ObPartDesc,
	partColumns []*protocol.ObColumn,
	orderedCompareColumns []*protocol.ObColumn) error {
	partDesc.setPartColumns(partColumns)
	if partDesc.partFuncType().isKeyPart() {
		if len(partColumns) == 0 {
			log.Warn("part column is empty", log.String("part desc", partDesc.String()))
			return errors.New("part column is empty")
		}
	} else if partDesc.partFuncType().isListPart() {
		// todo: list part is not support now
		log.Warn("list part is not support now", log.String("part desc", partDesc.String()))
		return errors.New("list part is not support now")
	} else if partDesc.partFuncType().isRangePart() {
		if rangeDesc, ok := partDesc.(*ObRangePartDesc); ok {
			rangeDesc.orderedCompareColumns = orderedCompareColumns
		} else {
			log.Warn("failed to convert to ObRangePartDesc", log.String("part desc", partDesc.String()))
			return errors.New("failed to convert to ObRangePartDesc")
		}
	}
	return nil
}

func buildPartDesc(partNum int,
	partType int,
	partSpace int,
	partExpr string,
	partRangeType string) (ObPartDesc, error) {
	partFuncType := newObPartFuncType(partType)
	partExpr = strings.ReplaceAll(partExpr, "`", "") // '`' is not supported by druid
	if partFuncType.isRangePart() {
		rangeDesc := newObRangePartDesc()
		rangeDesc.partNum = partNum
		rangeDesc.PartFuncType = partFuncType
		rangeDesc.PartExpr = partExpr
		rangeDesc.partSpace = partSpace
		rangeDesc.setOrderedPartColumnNames(partExpr)
		for _, typeStr := range strings.Split(partRangeType, ",") { // todo: @林径 确认空格
			typeValue, err := strconv.Atoi(typeStr)
			if err != nil {
				log.Warn("failed to convert type string to type value", log.String("typeStr", typeStr))
				return nil, err
			}
			objType, err := protocol.NewObObjType(typeValue)
			if err != nil {
				log.Warn("failed to new object type", log.Int("typeValue", typeValue))
				return nil, err
			}
			rangeDesc.orderedCompareColumnTypes = append(rangeDesc.orderedCompareColumnTypes, objType)
		}
		return rangeDesc, nil
	} else if partFuncType.isHashPart() {
		hashDesc := newObHashPartDesc()
		hashDesc.partNum = partNum
		hashDesc.PartFuncType = partFuncType
		hashDesc.PartExpr = partExpr
		hashDesc.setOrderedPartColumnNames(partExpr)
		hashDesc.partSpace = partSpace
		if util.ObVersion() < 4 {
			hashDesc.partNameIdMap = buildDefaultPartNameIdMap(partNum)
		}
		return hashDesc, nil
	} else if partFuncType.isKeyPart() {
		keyDesc := newObKeyPartDesc()
		keyDesc.partNum = partNum
		keyDesc.PartFuncType = partFuncType
		keyDesc.PartExpr = partExpr
		keyDesc.setOrderedPartColumnNames(partExpr)
		keyDesc.partSpace = partSpace
		if util.ObVersion() < 4 {
			keyDesc.partNameIdMap = buildDefaultPartNameIdMap(partNum)
		}
		return keyDesc, nil
	} else {
		log.Warn("invalid part type", log.String("part type", partFuncType.String()))
		return nil, errors.New("invalid part type")
	}
}

func buildDefaultPartNameIdMap(partNum int) map[string]int64 {
	partNameIdMap := make(map[string]int64)
	for i := 0; i < partNum; i++ {
		partNameIdMap["p"+strconv.Itoa(i)] = int64(i)
	}
	return partNameIdMap
}

func fetchFirstPart(db *DB, partFuncType ObPartFuncType, entry *ObTableEntry) error {
	key := entry.tableEntryKey
	var rows *Rows
	var err error
	if util.ObVersion() >= 4 {
		rows, err = db.Query(proxyFirstPartitionSql, key.tenantName, entry.tableId, math.MaxInt64)
	} else {
		rows, err = db.Query(proxyFirstPartitionSql, entry.tableId, math.MaxInt64)
	}
	if err != nil {
		log.Warn("failed to db query",
			log.Float32("ob version", util.ObVersion()),
			log.String("tenant name", key.tenantName),
			log.Uint64("tableId", entry.tableId))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		partId       int64
		partName     string
		highBoundVal interface{}
		tabletId     int64
		subPartNum   int
	)
	var idx int64 = 0
	for rows.Next() {
		if util.ObVersion() >= 4 {
			err = rows.Scan(&partId, &partName, &tabletId, &highBoundVal, &subPartNum)
		} else {
			err = rows.Scan(&partId, &partName, &highBoundVal)
		}
		if err != nil {
			log.Warn("failed to scan row", log.Float32("ob version", util.ObVersion()))
			return err
		}

		if partFuncType.isRangePart() {
			// todo: handle range bounds
			// highBoundVal may be is nil
		} else if partFuncType.isListPart() {
			// todo: not support list part now
			log.Warn("not support list part now", log.String("partFuncType", partFuncType.String()))
			err = errors.New("not support list part now")
			return err
		} else if util.ObVersion() >= 4 && (partFuncType.isKeyPart() || partFuncType.isHashPart()) {
			entry.partitionInfo.partTabletIdMap[idx] = tabletId
			idx++
		}
	}
	return nil
}

func fetchSubPart(db *DB, partFuncType ObPartFuncType, entry *ObTableEntry) error {
	key := entry.tableEntryKey
	var rows *Rows
	var err error
	if util.ObVersion() >= 4 {
		rows, err = db.Query(proxySubPartitionSql, key.tenantName, entry.tableId, math.MaxInt64)
	} else {
		rows, err = db.Query(proxySubPartitionSql, entry.tableId, math.MaxInt64)
	}
	if err != nil {
		log.Warn("failed to db query",
			log.Float32("ob version", util.ObVersion()),
			log.String("tenant name", key.tenantName),
			log.Uint64("tableId", entry.tableId))
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		subPartId    int64
		partName     string
		highBoundVal interface{} // maybe is sql type NULL
		tabletId     int64
	)
	var idx int64 = 0
	for rows.Next() {
		if util.ObVersion() >= 4 {
			err = rows.Scan(&subPartId, &partName, &tabletId, &highBoundVal)
		} else {
			err = rows.Scan(&subPartId, &partName, &highBoundVal)
		}
		if err != nil {
			log.Warn("failed to scan row", log.Float32("ob version", util.ObVersion()))
			err = errors.New("failed to scan row")
			return err
		}

		if partFuncType.isRangePart() {
			// todo: handle range bounds
		} else if partFuncType.isListPart() {
			// todo: not support list part now
			log.Warn("not support list part now", log.String("partFuncType", partFuncType.String()))
			err = errors.New("not support list part now")
			return err
		} else if util.ObVersion() >= 4 && (partFuncType.isKeyPart() || partFuncType.isHashPart()) {
			entry.partitionInfo.partTabletIdMap[idx] = tabletId
			idx++
		}
	}
	return nil
}

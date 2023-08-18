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
	"context"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/util"
)

const (
	OceanBaseDatabase = "OCEANBASE"   // OceanBaseDatabase is used to obtain route.
	AllDummyTable     = "__all_dummy" // AllDummyTable is used to obtain all tenant server distributions.
)

const (
	obVersionSql     = "SELECT /*+READ_CONSISTENCY(WEAK)*/ OB_VERSION() AS CLUSTER_VERSION;"
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

// InitSql will be called after getting oceanbase cluster version.
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

// GetObVersionFromRemoteBySysUA get OceanBase cluster version by sysUA
// called when client init
func GetObVersionFromRemoteBySysUA(addr *ObServerAddr, sysUA *ObUserAuth) (float32, error) {
	db, err := NewDB(
		sysUA.userName,
		sysUA.password,
		addr.ip,
		strconv.Itoa(addr.sqlPort),
		OceanBaseDatabase,
	)
	if err != nil {
		return 0.0, errors.WithMessagef(err, "new db, sysUA:%s, addr:%s", sysUA.String(), addr.String())
	}
	defer func() {
		_ = db.Close()
	}()
	return GetObVersionFromRemote(db)
}

// GetObVersionFromRemoteByIpPort get OceanBase cluster version by sql
// called when client init
func GetObVersionFromRemoteByIpPort(ip string, port int, userName string, password string) (float32, error) {
	db, err := NewDB(
		userName,
		password,
		ip,
		strconv.Itoa(port),
		OceanBaseDatabase,
	)
	if err != nil {
		return 0.0, errors.WithMessagef(err, "new db, ip:%s, port:%d, userName:%s, password:%s", ip, port, userName, password)
	}
	defer func() {
		_ = db.Close()
	}()
	return GetObVersionFromRemote(db)
}

// GetObVersionFromOdp get OceanBase cluster version by odp
func GetObVersionFromOdp(ip string, port int, userName string, password string) (float32, error) {
	return 0, nil
}

// GetObVersionFromRemote get OceanBase cluster version by sql
func GetObVersionFromRemote(db *DB) (float32, error) {
	// 1. Prepare get observer version sql statement.
	stmt, err := db.Prepare(obVersionSql)
	if err != nil {
		return 0.0, errors.WithMessagef(err, "prepare get observer version sql, sql:%s", obVersionSql)
	}

	// 2. Get result from query row.
	var obVersionStr string
	err = stmt.QueryRow().Scan(&obVersionStr)
	if err != nil {
		return 0.0, errors.WithMessagef(err, "get observer version from query result, sql:%s", obVersionSql)
	}

	// 3. parse ob version string
	// +-----------------+
	// | CLUSTER_VERSION |
	// +-----------------+
	// | 4.1.0.0         |
	str := strings.ReplaceAll(obVersionStr, ".", "") // 4100
	ver, err := strconv.Atoi(str)
	if err != nil {
		return 0.0, errors.WithMessagef(err, "convert string to int, str:%s", str)
	}
	res := float32(ver) / 1000.0 // ObVersion = 4.1
	return res, nil
}

// GetTableEntryFromRemote obtain the route of a table in the ob cluster by querying the routing system table.
// ObTableEntry indicates the routing information of a table.
func GetTableEntryFromRemote(
	ctx context.Context,
	addr *ObServerAddr,
	sysUA *ObUserAuth,
	key *ObTableEntryKey) (*ObTableEntry, error) {
	// 1. Get db handle
	db, err := NewDB(
		sysUA.userName,
		sysUA.password,
		addr.ip,
		strconv.Itoa(addr.sqlPort),
		OceanBaseDatabase,
	)
	if err != nil {
		return nil, errors.WithMessagef(err, "new db, sysUA:%s, addr:%s", sysUA.String(), addr.String())
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
	rows, err := db.QueryContext(ctx, sql, key.tenantName, key.databaseName, key.tableName)
	if err != nil {
		return nil, errors.WithMessagef(err, "query partition location, sql:%s, tenantName:%s, "+
			"databaseName:%s, tableName:%s", sql, key.tenantName, key.databaseName, key.tableName)
	}
	defer func() {
		_ = rows.Close()
	}()

	// 3. Create table entry by parsing query result set.
	entry, err := getTableEntryFromResultSet(rows, key.tableName)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table entry from result set, key:%s", key.String())
	}
	entry.tableEntryKey = key

	// 4. Fetch partition info
	if entry.IsPartitionTable() {
		info, err := getPartitionInfoFromRemote(ctx, db, key.tenantName, entry.tableId)
		if err != nil {
			return nil, errors.WithMessagef(err, "get partition info, key:%s", key.String())
		}
		entry.partitionInfo = info

		// 4.1. Fetch first partition info
		if info.level >= 1 {
			err = fetchFirstPart(ctx, db, info.firstPartDesc.PartFuncType(), entry)
			if err != nil {
				return nil, errors.WithMessagef(err, "fetch first partition info, table entry:%s", entry.String())
			}
		}

		// 4.2. Fetch sub partition info
		if info.level == 2 {
			err = fetchSubPart(ctx, db, info.subPartDesc.PartFuncType(), entry)
			if err != nil {
				return nil, errors.WithMessagef(err, "fetch sub partition info, table entry:%s", entry.String())
			}
		}

		entry.partitionInfo = info
	}

	// 5. Get partition location entry
	partLocationEntry, err := GetPartLocationEntryFromRemote(ctx, db, entry)
	if err != nil {
		return nil, errors.WithMessagef(err, "get table entry location, table entry:%s", entry.String())
	}
	entry.partLocationEntry = partLocationEntry
	entry.refreshTimeMills = time.Time{}.Unix()
	return entry, nil
}

// getTableEntryFromResultSet get the queried routing information from sql rows, including ip, port, sqlPort and so on.
func getTableEntryFromResultSet(rows *Rows, tableName string) (*ObTableEntry, error) {
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
	isEmpty := true
	for rows.Next() {
		if isEmpty {
			isEmpty = false
		}
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
			return nil, errors.WithMessagef(err, "scan row")
		}
		replica := newReplicaLocation(
			NewObServerAddr(svrIp, sqlPort, svrPort),
			newServerStatus(stopTime, status),
			obServerRole(role),
			protocol.ObReplicaType(replicaType),
		)
		if !replica.isValid() {
			return nil, errors.Errorf("replica is invalid, replaca:%s", replica.String())
		}
		tableLocation.replicaLocations = append(tableLocation.replicaLocations, replica)
	}
	if isEmpty {
		return nil, errors.Errorf("Table not exist, tableName:%s", tableName)
	}

	// 2. fill table entry
	entry.tableId = tableId
	entry.partNum = partNum
	entry.replicaNum = replicaNum
	entry.tableLocation = tableLocation

	return entry, nil
}

// GetPartLocationEntryFromRemote get the location information for a table.
// Location information represents the distribution of all replica of the table.
func GetPartLocationEntryFromRemote(ctx context.Context, db *DB, entry *ObTableEntry) (*ObPartLocationEntry, error) {
	// 1. Create inStatement "(0,1,2...partNum);".
	partIds := make([]uint64, 0, entry.partNum)
	if util.ObVersion() >= 4 && entry.IsPartitionTable() {
		for _, v := range entry.partitionInfo.partTabletIdMap {
			partIds = append(partIds, v)
		}
		// partIds doesn't have to be ascending, so do sort
		sort.Slice(partIds, func(i, j int) bool {
			return partIds[i] < partIds[j]
		})
	} else {
		// use absolute partition id like（1152921509170249730）to get partition location
		if entry.partitionInfo != nil && entry.partitionInfo.level == PartLevelTwo {
			for i := 0; i < entry.partitionInfo.firstPartDesc.PartNum(); i++ {
				for j := 0; j < entry.partitionInfo.subPartDesc.PartNum(); j++ {
					partIds = append(partIds, generateSubPartId(uint64(i), uint64(j)))
				}
			}
		} else {
			for i := 0; i < entry.partNum; i++ {
				partIds = append(partIds, uint64(i))
			}
		}
	}
	inStatement := createInStatement(partIds)

	// 2. Do query with specific tenant name，database name and table name.
	sql := proxyPartitionLocationSql + inStatement
	key := entry.tableEntryKey
	rows, err := db.QueryContext(ctx, sql, key.tenantName, key.databaseName, key.tableName)
	if err != nil {
		return nil, errors.WithMessagef(err, "sql query, sql:%s", sql)
	}
	defer func() {
		_ = rows.Close()
	}()

	// 3. create ObPartitionEntry by parsing query result set
	partLocationEntry, err := getPartLocationEntryFromResultSet(rows, key.tableName)
	if err != nil {
		return nil, errors.WithMessagef(err, "get partition location from result set, sql:%s", sql)
	}

	return partLocationEntry, nil
}

// getPartLocationEntryFromResultSet get replica location information from sql rows,
// including partitionId, ip, port, sqlPort and so on.
func getPartLocationEntryFromResultSet(rows *Rows, tableName string) (*ObPartLocationEntry, error) {
	var partitionEntry *ObPartLocationEntry = nil
	isFirstRow := true
	var (
		partitionId uint64
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

	isEmpty := true
	for rows.Next() {
		if isEmpty {
			isEmpty = false
		}
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
			return nil, errors.WithMessagef(err, "scan row")
		}

		// create ObPartLocationEntry
		if isFirstRow {
			isFirstRow = false
			partitionEntry = newObPartLocationEntry(partNum)
		}

		// 1. create obReplicaLocation
		replica := newReplicaLocation(
			NewObServerAddr(svrIp, sqlPort, svrPort),
			newServerStatus(stopTime, status),
			obServerRole(role),
			protocol.ObReplicaType(replicaType),
		)

		// 2. find or create location, then add replica location
		location, ok := partitionEntry.partLocations[partitionId]
		if !ok {
			location = new(obPartitionLocation)
			partitionEntry.partLocations[partitionId] = location
		}
		location.addReplicaLocation(replica)
	}
	if isEmpty {
		return nil, errors.Errorf("Table not exist, tableName:%s", tableName)
	}

	return partitionEntry, nil
}

// getPartitionInfoFromRemote get the meta information for the partition key.
func getPartitionInfoFromRemote(ctx context.Context, db *DB, tenantName string, tableId uint64) (*obPartitionInfo, error) {
	// 1. Do query with specific tenant name(ob version > 4), table id and limit.
	var rows *Rows
	var err error
	if util.ObVersion() >= 4 {
		rows, err = db.QueryContext(ctx, proxyPartitionInfoSql, tenantName, tableId, math.MaxInt64)
	} else {
		rows, err = db.QueryContext(ctx, proxyPartitionInfoSql, tableId, math.MaxInt64)
	}
	if err != nil {
		return nil, errors.WithMessagef(err, "query partition info, "+
			"tenantName: %s, tableId:%d, obVersion:%f", tenantName, tableId, util.ObVersion())
	}
	defer func() {
		_ = rows.Close()
	}()

	// 2. create obPartitionInfo by parsing query result set
	info, err := getPartitionInfoFromResultSet(rows)
	if err != nil {
		return nil, errors.WithMessagef(err, "parse partition info from row, tableId: %d", tableId)
	}

	return info, nil
}

// getPartitionInfoFromResultSet get the meta information for the partition key from sql rows.
// primary key(`c1`,`c2`)) partition by range columns(`c1`, `c2`, `c3`) (
//
//	PARTITION p0 VALUES LESS THAN (10, 'aaa', 'bbb'),
//	PARTITION p1 VALUES LESS THAN (20, 'ccc', 'ddd'),
//	PARTITION p2 VALUES LESS THAN (MAXVALUE, MAXVALUE, MAXVALUE)
//
// );
// +------------+----------+-----------+------------+------------------+-----------------+--------------+---------------+----------------+---------------------+---------------+---------------+---------------+--------------+------------------+-------------------------+
// | part_level | part_num | part_type | part_space | part_expr        | part_range_type | sub_part_num | sub_part_type | sub_part_space | sub_part_range_type | sub_part_expr | part_key_name | part_key_type | part_key_idx | part_key_extra   | part_key_collation_type |
// +------------+----------+-----------+------------+------------------+-----------------+--------------+---------------+----------------+---------------------+---------------+---------------+---------------+--------------+------------------+-------------------------+
// |          1 |        3 |         4 |          0 | `c1`, `c2`, `c3` | 5,22,22         |            0 |             0 |              0 |                     |               | c1            |             5 |            0 |                  |                      63 |
// |          1 |        3 |         4 |          0 | `c1`, `c2`, `c3` | 5,22,22         |            0 |             0 |              0 |                     |               | c2            |            22 |            1 |                  |                      45 |
// |          1 |        3 |         4 |          0 | `c1`, `c2`, `c3` | 5,22,22         |            0 |             0 |              0 |                     |               | c3            |            22 |            2 | substr(`c2`,1,4) |                      45 |
// +------------+----------+-----------+------------+------------------+-----------------+--------------+---------------+----------------+---------------------+---------------+---------------+---------------+--------------+------------------+-------------------------+
func getPartitionInfoFromResultSet(rows *Rows) (*obPartitionInfo, error) {
	var info *obPartitionInfo = nil
	var partColumns []obColumn
	var firstPartColumnNames []string
	var subPartColumnNames []string
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
			return nil, errors.WithMessagef(err, "scan row")
		}
		if isFirstRow {
			isFirstRow = false
			partColumns = make([]obColumn, 0, 1)
			info = newObPartitionInfo(obPartLevel(partLevel))
			// build first part
			if info.level >= PartLevelOne {
				firstPartDesc, err := buildPartDesc(
					partNum,
					obPartFuncType(partType),
					partSpace,
				)
				if err != nil {
					return nil, errors.WithMessagef(err, "build first part desc, partNum:%d, partType:%d", partNum, partType)
				}
				info.firstPartDesc = firstPartDesc

				// eg:"c1, c2", need to remove ' '
				str := strings.ReplaceAll(partExpr, " ", "")
				str = strings.ReplaceAll(str, "`", "")
				firstPartColumnNames = strings.Split(str, ",") // firstPartColumnNames is ordered
			}

			// build sub part
			if info.level == PartLevelTwo {
				subPartDesc, err := buildPartDesc(
					subPartNum,
					obPartFuncType(subPartType),
					subPartSpace,
				)
				if err != nil {
					return nil, errors.WithMessagef(err, "build sub part desc, partNum:%d, partType:%d", partNum, subPartType)
				}
				info.subPartDesc = subPartDesc

				// eg:"`c1`, `c2`", need to remove ' ' and '`'
				str := strings.ReplaceAll(subPartExpr, " ", "")
				str = strings.ReplaceAll(str, "`", "")
				subPartColumnNames = strings.Split(str, ",") // subPartColumnNames is ordered
			}
		}

		partKeyExtra = strings.ReplaceAll(partKeyExtra, "`", "") // '`' is not supported by druid
		partKeyExtra = strings.ReplaceAll(partKeyExtra, " ", "") // ' ' should be removed
		var column obColumn
		if partKeyExtra != "" {
			return nil, errors.New("not support generate column now")
		} else {
			objType, err := protocol.NewObjType(protocol.ObObjTypeValue(partKeyType))
			if err != nil {
				return nil, errors.WithMessagef(err, "generate object type, partKeyType:%d", partKeyType)
			}
			column = newObSimpleColumn(
				partKeyName,
				objType,
				protocol.ObCollationType(spare1),
			)
		}
		partColumns = append(partColumns, column)
	}

	if info == nil {
		return nil, errors.New("empty set")
	}

	// order part columns
	if info.firstPartDesc != nil {
		firstColumns := make([]obColumn, 0, 1)
		for _, name := range firstPartColumnNames {
			for _, column := range partColumns {
				if strings.EqualFold(column.ColumnName(), name) {
					firstColumns = append(firstColumns, column)
				}
			}
		}
		info.firstPartDesc.SetPartColumns(firstColumns)
		if info.subPartDesc != nil {
			subColumns := make([]obColumn, 0, 1)
			for _, name := range subPartColumnNames {
				for _, column := range partColumns {
					if strings.EqualFold(column.ColumnName(), name) {
						subColumns = append(subColumns, column)
					}
				}
			}
			info.subPartDesc.SetPartColumns(subColumns)
		}
	}

	info.partTabletIdMap = make(map[uint64]uint64, partNum)

	return info, nil
}

// buildPartDesc generate partition key description information.
func buildPartDesc(partNum int,
	partFuncType obPartFuncType,
	partSpace int) (obPartDesc, error) {
	if isRangePart(partFuncType) {
		return newObRangePartDesc(partSpace, partNum, partFuncType), nil
	} else if isHashPart(partFuncType) {
		return newObHashPartDesc(partSpace, partNum, partFuncType), nil
	} else if isKeyPart(partFuncType) {
		return newObKeyPartDesc(partSpace, partNum, partFuncType), nil
	} else {
		return nil, errors.Errorf("invalid part type, partFuncType:%d", partFuncType)
	}
}

// fetchFirstPart get level 1 partition information.
func fetchFirstPart(ctx context.Context, db *DB, partFuncType obPartFuncType, entry *ObTableEntry) error {
	key := entry.tableEntryKey
	var rows *Rows
	var err error
	if util.ObVersion() >= 4 {
		rows, err = db.QueryContext(ctx, proxyFirstPartitionSql, key.tenantName, entry.tableId, math.MaxInt64)
	} else {
		rows, err = db.QueryContext(ctx, proxyFirstPartitionSql, entry.tableId, math.MaxInt64)
	}
	if err != nil {
		return errors.WithMessagef(err, "query first partition, "+
			"tenantName:%s, tableId:%d, obVersion:%f", key.tenantName, entry.tableId, util.ObVersion())
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		partId       uint64
		partName     string
		highBoundVal interface{}
		tabletId     uint64
		subPartNum   int
	)
	var idx uint64 = 0
	for rows.Next() {
		if util.ObVersion() >= 4 {
			err = rows.Scan(&partId, &partName, &tabletId, &highBoundVal, &subPartNum)
		} else {
			err = rows.Scan(&partId, &partName, &highBoundVal)
		}
		if err != nil {
			return errors.WithMessagef(err, "scan row, ob version:%f", util.ObVersion())
		}

		if isRangePart(partFuncType) {
			// +---------+-----------+-----------+----------------------------+--------------+
			// | part_id | part_name | tablet_id | high_bound_val             | sub_part_num |
			// +---------+-----------+-----------+----------------------------+--------------+
			// |  500006 | p0        |    200001 | 10,'aaa','bbb'             |            0 |
			// |  500007 | p1        |    200002 | 20,'aaa','bbb'             |            0 |
			// |  500008 | p2        |    200003 | MAXVALUE,MAXVALUE,MAXVALUE |            0 |
			// +---------+-----------+-----------+----------------------------+--------------+
			return errors.New("not support range partition now")
		} else if isListPart(partFuncType) {
			return errors.New("not support list partition now")
		} else if util.ObVersion() >= 4 && (isKeyPart(partFuncType) || isHashPart(partFuncType)) {
			// in version 4, subPartNum exists only in level 1 partitions, so you need to set subPartNum in advance.
			if entry.partitionInfo.subPartDesc != nil {
				entry.partitionInfo.subPartDesc.SetPartNum(subPartNum)
			}
			entry.partitionInfo.partTabletIdMap[idx] = tabletId
			idx++
		}
	}
	return nil
}

// fetchSubPart get level 2 partition information.
func fetchSubPart(ctx context.Context, db *DB, partFuncType obPartFuncType, entry *ObTableEntry) error {
	key := entry.tableEntryKey
	var rows *Rows
	var err error
	if util.ObVersion() >= 4 {
		rows, err = db.QueryContext(ctx, proxySubPartitionSql, key.tenantName, entry.tableId, math.MaxInt64)
	} else {
		rows, err = db.QueryContext(ctx, proxySubPartitionSql, entry.tableId, math.MaxInt64)
	}
	if err != nil {
		return errors.WithMessagef(err, "query sub partition, "+
			"tenantName:%s, tableId:%d, obVerdion:%f", key.tenantName, entry.tableId, util.ObVersion())
	}
	defer func() {
		_ = rows.Close()
	}()

	var (
		subPartId    uint64
		partName     string
		highBoundVal interface{} // maybe is sql type NULL
		tabletId     uint64
	)
	var idx uint64 = 0
	for rows.Next() {
		if util.ObVersion() >= 4 {
			err = rows.Scan(&subPartId, &partName, &tabletId, &highBoundVal)
		} else {
			err = rows.Scan(&subPartId, &partName, &highBoundVal)
		}
		if err != nil {
			err = errors.Errorf("failed to scan row, obVersion:%f", util.ObVersion())
			return err
		}

		if isRangePart(partFuncType) {
			return errors.New("not support range partition now")
		} else if isListPart(partFuncType) {
			return errors.New("not support list partition now")
		} else if util.ObVersion() >= 4 && (isKeyPart(partFuncType) || isHashPart(partFuncType)) {
			entry.partitionInfo.partTabletIdMap[idx] = tabletId
			idx++
		}
	}
	return nil
}

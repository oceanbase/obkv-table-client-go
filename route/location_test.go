package route

import (
	sql2 "database/sql"
	"math"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

// write your true config and create table by sql first
var (
	testClusterName = "test"
	testTenantName  = "mysql"
	testDatabase    = "testDatabaseName"
	testTableName   = "testTableName"
	testUserName    = "root"
	testPassword    = ""
	testIp          = "127.0.0.1"
	testSqlPort     = 41101
	testServerPort  = 41100
	testServerAddr  = ObServerAddr{testIp, testSqlPort, testServerPort}
	testUserAuth    = ObUserAuth{testUserName, testPassword}
)

// 3.x version
/*
CREATE TABLE t1(c1 INT(5),
                c2 int(10))
PARTITION BY RANGE COLUMNS(c1)
SUBPARTITION BY HASH(c2) SUBPARTITIONS 2
(PARTITION r1 VALUES LESS THAN(10),
 PARTITION r2 VALUES LESS THAN(20));
*/
func TestGetTableEntryFromRemoteV3(t *testing.T) {
	// 1. mock function 'GetObVersionFromRemote'
	patchVer := gomonkey.ApplyFunc(GetObVersionFromRemote, func(addr *ObServerAddr, sysUA *ObUserAuth) (float32, error) {
		return 3.24, nil
	})
	defer patchVer.Reset()
	// 2. check version
	ver, err := GetObVersionFromRemote(&testServerAddr, &testUserAuth)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, ver, float32(3.24))
	// 3. set version
	util.SetObVersion(ver)
	InitSql(ver)
	// 4. mock
	// 4.1 new mock DB
	mockDB, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// 4.2 mock function 'NewDB'
	patchDb := gomonkey.ApplyFunc(NewDB, func(userName string, password string, ip string, port string, database string) (*DB, error) {
		return mockDB, nil
	})
	defer patchDb.Reset()

	// 4.3 mock proxyLocationSql result
	sql := proxyLocationSql
	queryFields := []string{"partition_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows := sqlmock.NewRows(queryFields).
		AddRow(0, "127.0.0.1", 42705, 1099511677777, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.3", 42701, 1099511677777, 1, 3, 4, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, testDatabase, testTableName).WillReturnRows(mockRows)

	// 4.4 mock proxyPartitionInfoSql result
	sql = proxyPartitionInfoSql
	queryFields = []string{"part_level", "part_num", "part_type", "part_space",
		"part_expr", "part_range_type", "sub_part_num", "sub_part_type",
		"sub_part_space", "sub_part_range_type", "sub_part_expr", "part_key_name",
		"part_key_type", "part_key_idx", "part_key_extra", "spare1",
	}

	mockRows = sqlmock.NewRows(queryFields).
		AddRow(2, 2, 4, 0, "c1", 5, 2, 8, 0, "", "c2", "c1", 4, 0, "", 63).
		AddRow(2, 2, 4, 0, "c1", 5, 2, 8, 0, "", "c2", "c2", 4, 1, "", 63)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1099511677777, math.MaxInt64).WillReturnRows(mockRows)

	// 4.5 mock proxyFirstPartitionSql result
	sql = proxyFirstPartitionSql
	queryFields = []string{"part_id", "part_name", "high_bound_val"}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(0, "r1", 10).
		AddRow(1, "r2", 20)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1099511677777, math.MaxInt64).WillReturnRows(mockRows)

	// 4.6 mock proxySubPartitionSql result
	sql = proxySubPartitionSql
	queryFields = []string{"sub_part_id", "part_name", "high_bound_val"}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(0, "p0", sql2.NullString{}).
		AddRow(1, "p1", sql2.NullString{})
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1099511677777, math.MaxInt64).WillReturnRows(mockRows)

	// 4.7 mock proxyPartitionLocationSql result
	sql = proxyPartitionLocationSql
	sql += CreateInStatement([]int{0, 1, 2, 3})
	queryFields = []string{"partition_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(0, "127.0.0.1", 42705, 1099511677777, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.3", 42701, 1099511677777, 1, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(1, "127.0.0.1", 42705, 1099511677777, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(1, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(1, "127.0.0.3", 42701, 1099511677777, 1, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(2, "127.0.0.1", 42705, 1099511677777, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(2, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(2, "127.0.0.3", 42701, 1099511677777, 1, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(3, "127.0.0.1", 42705, 1099511677777, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(3, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(3, "127.0.0.3", 42701, 1099511677777, 1, 3, 4, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, testDatabase, testTableName).WillReturnRows(mockRows)

	// 5. test
	key := ObTableEntryKey{
		testClusterName,
		testTenantName,
		testDatabase,
		testTableName,
	}
	entry, err := GetTableEntryFromRemote(&testServerAddr, &testUserAuth, &key)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, entry.tableId, uint64(1099511677777))
	assert.Equal(t, entry.partNum, 4)
	assert.Equal(t, entry.replicaNum, 3)
	assert.Equal(t, entry.tableEntryKey, key)
	assert.Equal(t, entry.partitionInfo.level, ObPartitionLevel{"partLevelTwo", PartLevelTwoIndex})
	assert.Equal(t, entry.partitionInfo.firstPartDesc.String(), "ObRangePartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:RANGE_COLUMNS, index:4}, "+
		"partExpr:c1, orderedPartColumnNames:c1, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[ObColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"ObColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, partSpace:0, partNum:2, "+
		"orderedCompareColumns:[ObColumn{columnName:c1, index:0, "+
		"objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c1], isGenColumn:false, columnExpress:nil}], "+
		"orderedCompareColumnTypes:[ObObjType{type:ObInt64Type}]}")
	assert.Contains(t, entry.partitionInfo.subPartDesc.String(), "ObHashPartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH_V2, index:8}, "+
		"partExpr:c2, orderedPartColumnNames:c2, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[ObColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"ObColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, completeWorks:[], partSpace:0, partNum:2, "+
		"partNameIdMap:{m[p")
	assert.Equal(t, entry.tableLocation.String(), "ObTableLocation{"+
		"replicaLocations:["+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, "+
		"info:ObServerInfo{stopTime:0, status:ACTIVE}, "+
		"role:ObServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, "+
		"info:ObServerInfo{stopTime:0, status:ACTIVE}, "+
		"role:ObServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, "+
		"info:ObServerInfo{stopTime:0, status:ACTIVE}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, len(entry.partLocationEntry.partLocations), 4)
	assert.Equal(t, entry.partLocationEntry.partLocations[0].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[1].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[2].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[3].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	println(entry.String())
}

// 4.x version
/*
CREATE TABLE t1(c1 INT(5),
                c2 int(10))
PARTITION BY RANGE COLUMNS(c1)
SUBPARTITION BY HASH(c2) SUBPARTITIONS 2
(PARTITION r1 VALUES LESS THAN(10),
 PARTITION r2 VALUES LESS THAN(20));
*/
func TestGetTableEntryFromRemoteV4(t *testing.T) {
	// 1. mock function 'GetObVersionFromRemote'
	patchVer := gomonkey.ApplyFunc(GetObVersionFromRemote, func(addr *ObServerAddr, sysUA *ObUserAuth) (float32, error) {
		return 4.1, nil
	})
	defer patchVer.Reset()
	// 2. check version
	ver, err := GetObVersionFromRemote(&testServerAddr, &testUserAuth)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, ver, float32(4.1))
	// 3. set version
	util.SetObVersion(ver)
	InitSql(ver)
	// 4. mock
	// 4.1 new mock DB
	mockDB, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// 4.2 mock function 'NewDB'
	patchDb := gomonkey.ApplyFunc(NewDB, func(userName string, password string, ip string, port string, database string) (*DB, error) {
		return mockDB, nil
	})
	defer patchDb.Reset()

	// 4.3 mock proxyLocationSql result
	sql := proxyLocationSql
	queryFields := []string{"tablet_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows := sqlmock.NewRows(queryFields).
		AddRow(0, "127.0.0.1", 42705, 500012, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.3", 42701, 500012, 1, 3, 4, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, testDatabase, testTableName).WillReturnRows(mockRows)

	// 4.4 mock proxyPartitionInfoSql result
	sql = proxyPartitionInfoSql
	queryFields = []string{"part_level", "part_num", "part_type", "part_space",
		"part_expr", "part_range_type", "sub_part_num", "sub_part_type",
		"sub_part_space", "sub_part_range_type", "sub_part_expr", "part_key_name",
		"part_key_type", "part_key_idx", "part_key_extra", "part_key_collation_type",
	}

	mockRows = sqlmock.NewRows(queryFields).
		AddRow(2, 2, 4, 0, "c1", 5, 2, 8, 0, "", "c2", "c1", 4, 0, "", 63).
		AddRow(2, 2, 4, 0, "c1", 5, 2, 8, 0, "", "c2", "c2", 4, 1, "", 63)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, 500012, math.MaxInt64).WillReturnRows(mockRows)

	// 4.5 mock proxyFirstPartitionSql result
	sql = proxyFirstPartitionSql
	queryFields = []string{"part_id", "part_name", "tablet_id", "high_bound_val", "sub_part_num"}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(500013, "r1", 0, 10, 2).
		AddRow(500014, "r2", 0, 20, 2)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, 500012, math.MaxInt64).WillReturnRows(mockRows)

	// 4.6 mock proxySubPartitionSql result
	sql = proxySubPartitionSql
	queryFields = []string{"sub_part_id", "part_name", "tablet_id", "high_bound_val"}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(500015, "r1sp0", 200005, sql2.NullString{}).
		AddRow(500016, "r1sp1", 200006, sql2.NullString{}).
		AddRow(500017, "r2sp0", 200007, sql2.NullString{}).
		AddRow(500018, "r2sp1", 200008, sql2.NullString{})
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, 500012, math.MaxInt64).WillReturnRows(mockRows)

	// 4.7 mock proxyPartitionLocationSql result
	sql = proxyPartitionLocationSql
	sql += CreateInStatement([]int{200005, 200006, 200007, 200008})
	queryFields = []string{"tablet_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(200005, "127.0.0.1", 42705, 500012, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(200005, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(200005, "127.0.0.3", 42701, 500012, 1, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(200006, "127.0.0.1", 42705, 500012, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(200006, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(200006, "127.0.0.3", 42701, 500012, 1, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(200007, "127.0.0.1", 42705, 500012, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(200007, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(200007, "127.0.0.3", 42701, 500012, 1, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(200008, "127.0.0.1", 42705, 500012, 2, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(200008, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(200008, "127.0.0.3", 42701, 500012, 1, 3, 4, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(testTenantName, testDatabase, testTableName).WillReturnRows(mockRows)

	// 5. test
	key := ObTableEntryKey{
		testClusterName,
		testTenantName,
		testDatabase,
		testTableName,
	}
	entry, err := GetTableEntryFromRemote(&testServerAddr, &testUserAuth, &key)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, entry.tableId, uint64(500012))
	assert.Equal(t, entry.partNum, 4)
	assert.Equal(t, entry.replicaNum, 3)
	assert.Equal(t, entry.tableEntryKey, key)
	assert.Equal(t, entry.partitionInfo.level, ObPartitionLevel{"partLevelTwo", PartLevelTwoIndex})
	assert.Equal(t, entry.partitionInfo.firstPartDesc.String(), "ObRangePartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:RANGE_COLUMNS, index:4}, "+
		"partExpr:c1, orderedPartColumnNames:c1, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[ObColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"ObColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, partSpace:0, partNum:2, "+
		"orderedCompareColumns:[ObColumn{columnName:c1, index:0, "+
		"objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c1], isGenColumn:false, columnExpress:nil}], "+
		"orderedCompareColumnTypes:[ObObjType{type:ObInt64Type}]}")
	assert.Equal(t, entry.partitionInfo.subPartDesc.String(), "ObHashPartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH_V2, index:8}, "+
		"partExpr:c2, orderedPartColumnNames:c2, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[ObColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"ObColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, completeWorks:[], partSpace:0, partNum:2, "+
		"partNameIdMap:{}}")
	assert.Equal(t, entry.tableLocation.String(), "ObTableLocation{"+
		"replicaLocations:["+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, "+
		"info:ObServerInfo{stopTime:0, status:ACTIVE}, "+
		"role:ObServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, "+
		"info:ObServerInfo{stopTime:0, status:ACTIVE}, "+
		"role:ObServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, "+
		"info:ObServerInfo{stopTime:0, status:ACTIVE}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, len(entry.partLocationEntry.partLocations), 4)
	assert.Equal(t, entry.partLocationEntry.partLocations[200005].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[200006].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[200007].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[200008].String(), "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:ObServerInfo{stopTime:0, status:ACTIVE}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	println(entry.String())
}

func TestObUserAuth_ToString(t *testing.T) {
	auth := ObUserAuth{}
	assert.Equal(t, "ObUserAuth{userName:, password:}", auth.String())
	auth = ObUserAuth{"testUserName", "testPassword"}
	assert.Equal(t, "ObUserAuth{userName:testUserName, password:testPassword}", auth.String())
}

func TestObColumnIndexesPair_ToString(t *testing.T) {
	pair := ObColumnIndexesPair{}
	assert.Equal(t, "ObColumnIndexesPair{column:nil, indexes:[]}",
		pair.String())
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair = ObColumnIndexesPair{column, []int{1, 2, 3}}
	assert.Equal(t, "ObColumnIndexesPair{"+
		"column:ObColumn{"+
		"columnName:testColumnName, "+
		"index:0, "+
		"objType:ObObjType{type:ObTinyIntType}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[testColumnName], "+
		"isGenColumn:false, "+
		"columnExpress:nil}, "+
		"indexes:[1, 2, 3]}",
		pair.String(),
	)
}

func TestObPartDescCommon_ToString(t *testing.T) {
	comm := ObPartDescCommon{}
	assert.Equal(t, "ObPartDescCommon{"+
		"partFuncType:ObPartFuncType{name:, index:0}, "+
		"partExpr:, "+
		"orderedPartColumnNames:, "+
		"orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[], "+
		"rowKeyElement:nil}",
		comm.CommString(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := &ObColumnIndexesPair{column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeHashIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []*ObColumnIndexesPair{pair}
	partColumns := []*ObColumn{column}
	nameIdxMap := make(map[string]int, 3)
	nameIdxMap["c1"] = 0
	rowKeyElement := table.NewObRowKeyElement(nameIdxMap)
	comm = ObPartDescCommon{PartFuncType: partFuncType,
		PartExpr:                            partExpr,
		OrderedPartColumnNames:              orderedPartColumnNames,
		OrderedPartRefColumnRowKeyRelations: orderedPartRefColumnRowKeyRelations,
		PartColumns:                         partColumns,
		RowKeyElement:                       rowKeyElement,
	}
	assert.Equal(t, "ObPartDescCommon{"+
		"partFuncType:ObPartFuncType{name:HASH, index:0}, "+
		"partExpr:c1, c2, "+
		"orderedPartColumnNames:c1,c2, "+
		"orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], "+
		"partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:ObRowKeyElement{nameIdxMap:{m[c1]=0}}}",
		comm.CommString(),
	)
}

func TestObRangePartDesc_ToString(t *testing.T) {
	desc := ObRangePartDesc{}
	assert.Equal(t, "ObRangePartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:, index:0}, partExpr:, orderedPartColumnNames:, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:nil}, partSpace:0, partNum:0, "+
		"orderedCompareColumns:[], "+
		"orderedCompareColumnTypes:[]}",
		desc.String(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := &ObColumnIndexesPair{column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeRangeIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []*ObColumnIndexesPair{pair}
	partColumns := []*ObColumn{column}
	nameIdxMap := make(map[string]int, 3)
	nameIdxMap["c1"] = 0
	rowKeyElement := table.NewObRowKeyElement(nameIdxMap)
	desc = ObRangePartDesc{
		partNum:                   10,
		partSpace:                 0,
		orderedCompareColumns:     []*ObColumn{column, column},
		orderedCompareColumnTypes: []protocol.ObObjType{objType, objType},
	}
	desc.PartFuncType = partFuncType
	desc.PartExpr = partExpr
	desc.OrderedPartColumnNames = orderedPartColumnNames
	desc.OrderedPartRefColumnRowKeyRelations = orderedPartRefColumnRowKeyRelations
	desc.PartColumns = partColumns
	desc.RowKeyElement = rowKeyElement
	assert.Equal(t, "ObRangePartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:RANGE, index:3}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:ObRowKeyElement{nameIdxMap:{m[c1]=0}}}, partSpace:0, partNum:10, "+
		"orderedCompareColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], "+
		"orderedCompareColumnTypes:[ObObjType{type:ObTinyIntType}, ObObjType{type:ObTinyIntType}]}",
		desc.String(),
	)
}

func TestObHashPartDesc_GetPartId(t *testing.T) {
	desc := &ObHashPartDesc{}
	partId, err := desc.GetPartId(nil)
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)
	partId, err = desc.GetPartId([]interface{}{1, 2, 3})
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)
}

func TestObHashPartDesc_ToString(t *testing.T) {
	desc := &ObHashPartDesc{}
	assert.Equal(t, "ObHashPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:, index:0}, partExpr:, orderedPartColumnNames:, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:nil}, "+
		"completeWorks:[], "+
		"partSpace:0, "+
		"partNum:0, "+
		"partNameIdMap:{}}",
		desc.String(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := &ObColumnIndexesPair{column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeHashIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []*ObColumnIndexesPair{pair}
	partColumns := []*ObColumn{column}
	nameIdxMap := make(map[string]int, 3)
	nameIdxMap["c1"] = 0
	rowKeyElement := table.NewObRowKeyElement(nameIdxMap)
	partNameIdMap := make(map[string]int64)
	partNameIdMap["p0"] = 0
	desc = &ObHashPartDesc{
		completeWorks: []int64{1, 2, 3},
		partSpace:     0,
		partNum:       10,
		partNameIdMap: partNameIdMap,
	}
	desc.PartFuncType = partFuncType
	desc.PartExpr = partExpr
	desc.OrderedPartColumnNames = orderedPartColumnNames
	desc.OrderedPartRefColumnRowKeyRelations = orderedPartRefColumnRowKeyRelations
	desc.PartColumns = partColumns
	desc.RowKeyElement = rowKeyElement
	assert.Equal(t, "ObHashPartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH, index:0}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:ObRowKeyElement{nameIdxMap:{m[c1]=0}}}, "+
		"completeWorks:[1, 2, 3], "+
		"partSpace:0, "+
		"partNum:10, "+
		"partNameIdMap:{m[p0]=0}}",
		desc.String(),
	)
}

func TestObHashPartDesc_innerHash(t *testing.T) {
	hashDesc := ObHashPartDesc{partSpace: 0, partNum: 10}
	hashVal := hashDesc.innerHash(0)
	assert.Equal(t, int64(0), hashVal)
	hashDesc = ObHashPartDesc{partSpace: 0, partNum: 10}
	hashVal = hashDesc.innerHash(1)
	assert.Equal(t, int64(1), hashVal)
	hashVal = hashDesc.innerHash(11)
	assert.Equal(t, int64(1), hashVal)
	hashVal = hashDesc.innerHash(-1)
	assert.Equal(t, int64(1), hashVal)
}

func TestObKeyPartDesc_ToString(t *testing.T) {
	desc := ObKeyPartDesc{}
	assert.Equal(t, "ObKeyPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:, index:0}, partExpr:, orderedPartColumnNames:, orderedPartRefColumnRowKeyRelations:[], partColumns:[], rowKeyElement:nil}, "+
		"partSpace:0, "+
		"partNum:0, "+
		"partNameIdMap:{}}",
		desc.String(),
	)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := &ObColumnIndexesPair{column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeKeyV2Index)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []*ObColumnIndexesPair{pair}
	partColumns := []*ObColumn{column}
	nameIdxMap := make(map[string]int, 3)
	nameIdxMap["c1"] = 0
	rowKeyElement := table.NewObRowKeyElement(nameIdxMap)
	partNameIdMap := make(map[string]int64)
	partNameIdMap["p0"] = 0
	desc = ObKeyPartDesc{
		partSpace:     0,
		partNum:       10,
		partNameIdMap: partNameIdMap,
	}
	desc.PartFuncType = partFuncType
	desc.PartExpr = partExpr
	desc.OrderedPartColumnNames = orderedPartColumnNames
	desc.OrderedPartRefColumnRowKeyRelations = orderedPartRefColumnRowKeyRelations
	desc.PartColumns = partColumns
	desc.RowKeyElement = rowKeyElement
	assert.Equal(t, "ObKeyPartDesc{"+
		"comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:KEY_V2, index:6}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:ObRowKeyElement{nameIdxMap:{m[c1]=0}}}, "+
		"partSpace:0, "+
		"partNum:10, "+
		"partNameIdMap:{m[p0]=0}}",
		desc.String(),
	)
}

func TestObPartitionLevel_ToString(t *testing.T) {
	level := ObPartitionLevel{}
	assert.Equal(t, "ObPartitionLevel{name:, index:0}", level.String())
	level = newObPartitionLevel(PartLevelZeroIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelZero, index:0}", level.String())
	level = newObPartitionLevel(PartLevelOneIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelOne, index:1}", level.String())
	level = newObPartitionLevel(PartLevelTwoIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelTwo, index:2}", level.String())
	level = newObPartitionLevel(PartLevelUnknownIndex)
	assert.Equal(t, "ObPartitionLevel{name:partLevelUnknown, index:-1}", level.String())

}

func TestObPartFuncType_ToString(t *testing.T) {
	part := ObPartFuncType{}
	assert.Equal(t, "ObPartFuncType{name:, index:0}", part.String())
	part = newObPartFuncType(partFuncTypeHashIndex)
	assert.Equal(t, "ObPartFuncType{name:HASH, index:0}", part.String())
	part = newObPartFuncType(partFuncTypeKeyIndex)
	assert.Equal(t, "ObPartFuncType{name:KEY, index:1}", part.String())
	part = newObPartFuncType(partFuncTypeKeyImplIndex)
	assert.Equal(t, "ObPartFuncType{name:KEY_IMPLICIT, index:2}", part.String())
	part = newObPartFuncType(partFuncTypeRangeIndex)
	assert.Equal(t, "ObPartFuncType{name:RANGE, index:3}", part.String())
	part = newObPartFuncType(partFuncTypeRangeColIndex)
	assert.Equal(t, "ObPartFuncType{name:RANGE_COLUMNS, index:4}", part.String())
	part = newObPartFuncType(partFuncTypeListIndex)
	assert.Equal(t, "ObPartFuncType{name:LIST, index:5}", part.String())
	part = newObPartFuncType(partFuncTypeKeyV2Index)
	assert.Equal(t, "ObPartFuncType{name:KEY_V2, index:6}", part.String())
	part = newObPartFuncType(partFuncTypeListColIndex)
	assert.Equal(t, "ObPartFuncType{name:LIST_COLUMNS, index:7}", part.String())
	part = newObPartFuncType(partFuncTypeHashV2Index)
	assert.Equal(t, "ObPartFuncType{name:HASH_V2, index:8}", part.String())
	part = newObPartFuncType(partFuncTypeKeyV3Index)
	assert.Equal(t, "ObPartFuncType{name:KEY_V3, index:9}", part.String())
	part = newObPartFuncType(partFuncTypeUnknownIndex)
	assert.Equal(t, "ObPartFuncType{name:UNKNOWN, index:-1}", part.String())
}

func TestObServerAddr_ToString(t *testing.T) {
	addr := ObServerAddr{}
	assert.Equal(t, "ObServerAddr{ip:, sqlPort:0, svrPort:0}", addr.String())
	addr = ObServerAddr{"127.0.0.1", 8080, 1227}
	assert.Equal(t, "ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}", addr.String())
}

func TestObServerInfo_ToString(t *testing.T) {
	info := ObServerInfo{}
	assert.Equal(t, "ObServerInfo{stopTime:0, status:}", info.String())
	info = ObServerInfo{0, "Active"}
	assert.Equal(t, info.IsActive(), true)
	assert.Equal(t, "ObServerInfo{stopTime:0, status:Active}", info.String())
}

func TestObServerRole_ToString(t *testing.T) {
	role := ObServerRole{}
	assert.Equal(t, "ObServerRole{name:, index:0}", role.String())
	role = newObServerRole(ServerRoleLeaderIndex)
	assert.Equal(t, "ObServerRole{name:LEADER, index:1}", role.String())
	role = newObServerRole(ServerRoleFollowerIndex)
	assert.Equal(t, "ObServerRole{name:FOLLOWER, index:2}", role.String())
	role = newObServerRole(ServerRoleInvalidIndex)
	assert.Equal(t, "ObServerRole{name:INVALID_ROLE, index:-1}", role.String())
}

func TestObReplicaType_ToString(t *testing.T) {
	replica := ObReplicaType{}
	assert.Equal(t, "ObReplicaType{name:, index:0}", replica.String())
	replica = newObReplicaType(ReplicaTypeFullIndex)
	assert.Equal(t, "ObReplicaType{name:FULL, index:0}", replica.String())
	replica = newObReplicaType(ReplicaTypeLogOnlyIndex)
	assert.Equal(t, "ObReplicaType{name:LOGONLY, index:5}", replica.String())
	replica = newObReplicaType(ReplicaTypeReadOnlyIndex)
	assert.Equal(t, "ObReplicaType{name:READONLY, index:16}", replica.String())
	replica = newObReplicaType(ReplicaTypeInvalidIndex)
	assert.Equal(t, "ObReplicaType{name:INVALID, index:-1}", replica.String())
}

func TestObReplicaLocation_ToString(t *testing.T) {
	replica := ObReplicaLocation{}
	assert.Equal(t, "ObReplicaLocation{"+
		"addr:ObServerAddr{ip:, sqlPort:0, svrPort:0}, "+
		"info:ObServerInfo{stopTime:0, status:}, "+
		"role:ObServerRole{name:, index:0}, "+
		"replicaType:ObReplicaType{name:, index:0}}",
		replica.String(),
	)
	replica = ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	assert.Equal(t, "ObReplicaLocation{"+
		"addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, "+
		"info:ObServerInfo{stopTime:0, status:Active}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}",
		replica.String(),
	)
}

func TestObTableLocation_ToString(t *testing.T) {
	loc := ObTableLocation{}
	assert.Equal(t, "ObTableLocation{replicaLocations:[]}", loc.String())
	replica := &ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	loc = ObTableLocation{[]*ObReplicaLocation{replica, replica}}
	assert.Equal(t, "ObTableLocation{"+
		"replicaLocations:["+
		"ObReplicaLocation{"+
		"addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, "+
		"info:ObServerInfo{stopTime:0, status:Active}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"ObReplicaLocation{"+
		"addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, "+
		"info:ObServerInfo{stopTime:0, status:Active}, "+
		"role:ObServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}]}", loc.String(),
	)
}

func TestObPartitionInfo_ToString(t *testing.T) {
	info := ObPartitionInfo{}
	assert.Equal(t, "ObPartitionInfo{"+
		"level:ObPartitionLevel{name:, index:0}, "+
		"firstPartDesc:nil, "+
		"subPartDesc:nil, "+
		"partColumns:[], "+
		"partTabletIdMap:{}, "+
		"partNameIdMap:{}}",
		info.String(),
	)
	level := newObPartitionLevel(PartLevelZeroIndex)
	objType, _ := protocol.NewObObjType(1)
	collType := protocol.NewObCollationType(63)
	column := NewObSimpleColumn("testColumnName", 0, objType, collType)
	pair := &ObColumnIndexesPair{column, []int{0}}
	partFuncType := newObPartFuncType(partFuncTypeHashIndex)
	partExpr := "c1, c2"
	orderedPartColumnNames := []string{"c1", "c2"}
	orderedPartRefColumnRowKeyRelations := []*ObColumnIndexesPair{pair}
	partColumns := []*ObColumn{column}
	nameIdxMap := make(map[string]int, 3)
	nameIdxMap["c1"] = 0
	rowKeyElement := table.NewObRowKeyElement(nameIdxMap)
	partNameIdMap := make(map[string]int64)
	partNameIdMap["p0"] = 0
	desc := &ObHashPartDesc{
		completeWorks: []int64{1, 2, 3},
		partSpace:     0,
		partNum:       10,
		partNameIdMap: partNameIdMap,
	}
	desc.PartFuncType = partFuncType
	desc.PartExpr = partExpr
	desc.OrderedPartColumnNames = orderedPartColumnNames
	desc.OrderedPartRefColumnRowKeyRelations = orderedPartRefColumnRowKeyRelations
	desc.PartColumns = partColumns
	desc.RowKeyElement = rowKeyElement
	partTabletIdMap := make(map[int64]int64)
	partTabletIdMap[0] = 500021
	info = ObPartitionInfo{
		level:           level,
		firstPartDesc:   desc,
		subPartDesc:     desc,
		partColumns:     partColumns,
		partTabletIdMap: partTabletIdMap,
		partNameIdMap:   partNameIdMap,
	}
	assert.Equal(t, "ObPartitionInfo{"+
		"level:ObPartitionLevel{name:partLevelZero, index:0}, "+
		"firstPartDesc:ObHashPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH, index:0}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:ObRowKeyElement{nameIdxMap:{m[c1]=0}}}, completeWorks:[1, 2, 3], partSpace:0, partNum:10, partNameIdMap:{m[p0]=0}}, "+
		"subPartDesc:ObHashPartDesc{comm:ObPartDescCommon{partFuncType:ObPartFuncType{name:HASH, index:0}, partExpr:c1, c2, orderedPartColumnNames:c1,c2, orderedPartRefColumnRowKeyRelations:[ObColumnIndexesPair{column:ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}, indexes:[0]}], partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], rowKeyElement:ObRowKeyElement{nameIdxMap:{m[c1]=0}}}, completeWorks:[1, 2, 3], partSpace:0, partNum:10, partNameIdMap:{m[p0]=0}}, "+
		"partColumns:[ObColumn{columnName:testColumnName, index:0, objType:ObObjType{type:ObTinyIntType}, collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[testColumnName], isGenColumn:false, columnExpress:nil}], "+
		"partTabletIdMap:{m[0]=500021}, "+
		"partNameIdMap:{m[p0]=0}}",
		info.String(),
	)
}

func TestObPartitionLocation_ToString(t *testing.T) {
	loc := ObPartitionLocation{}
	assert.Equal(t, "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:, sqlPort:0, svrPort:0}, info:ObServerInfo{stopTime:0, status:}, role:ObServerRole{name:, index:0}, replicaType:ObReplicaType{name:, index:0}}, "+
		"replicas:[]}",
		loc.String(),
	)
	leader := ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	follower := ObReplicaLocation{
		ObServerAddr{"127.0.0.1", 8080, 1227},
		ObServerInfo{0, "Active"},
		newObServerRole(ServerRoleLeaderIndex),
		newObReplicaType(ReplicaTypeFullIndex),
	}
	loc = ObPartitionLocation{
		leader,
		[]ObReplicaLocation{follower, follower},
	}
	assert.Equal(t, "ObPartitionLocation{"+
		"leader:ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, info:ObServerInfo{stopTime:0, status:Active}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, info:ObServerInfo{stopTime:0, status:Active}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, ObReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:8080, svrPort:1227}, info:ObServerInfo{stopTime:0, status:Active}, role:ObServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}",
		loc.String(),
	)
}

func TestObPartLocationEntry_ToString(t *testing.T) {
	entry := ObPartLocationEntry{}
	assert.Equal(t, "ObPartLocationEntry{partLocations:{}}", entry.String())
	loc := ObPartitionLocation{}
	m := make(map[int64]*ObPartitionLocation, 2)
	m[0] = &loc
	entry = ObPartLocationEntry{m}
	assert.Equal(t, "ObPartLocationEntry{"+
		"partLocations:{m[0]=ObPartitionLocation{leader:ObReplicaLocation{addr:ObServerAddr{ip:, sqlPort:0, svrPort:0}, info:ObServerInfo{stopTime:0, status:}, role:ObServerRole{name:, index:0}, replicaType:ObReplicaType{name:, index:0}}, replicas:[]}}}",
		entry.String(),
	)
	m[1] = nil // test null pointer
	assert.Equal(t, 274, len(entry.String()))
}

func TestObTableEntryKey_ToString(t *testing.T) {
	key := ObTableEntryKey{}
	assert.Equal(t, "ObTableEntryKey{clusterName:, tenantNane:, databaseName:, tableName:}", key.String())
	key = ObTableEntryKey{
		"testClusterName",
		"testTenantNane",
		"testDatabaseName",
		"testTableName",
	}
	assert.Equal(t, "ObTableEntryKey{"+
		"clusterName:testClusterName, "+
		"tenantNane:testDatabaseName, "+
		"databaseName:testDatabaseName, "+
		"tableName:testTableName}",
		key.String(),
	)
}

func TestObTableEntry_ToString(t *testing.T) {
	entry := ObTableEntry{}
	assert.Equal(t, "ObTableEntry{"+
		"tableId:0, "+
		"partNum:0, "+
		"replicaNum:0, "+
		"refreshTimeMills:0, "+
		"tableEntryKey:ObTableEntryKey{clusterName:, tenantNane:, databaseName:, tableName:}, "+
		"partitionInfo:nil, "+
		"tableLocation:nil, "+
		"partitionEntry:nil}",
		entry.String(),
	)
}

func TestObServerRoster_ToString(t *testing.T) {
	r := ObServerRoster{}
	assert.Equal(t, r.String(), "ObServerRoster{maxPriority:0, roster:[]}")
	addr := NewObServerAddr("127.0.0.1", 8000, 8080)
	roster := make([]*ObServerAddr, 4)
	roster = append(roster, addr)
	roster = append(roster, addr)
	roster = append(roster, addr)
	r = ObServerRoster{roster: roster}
	r.maxPriority.Store(1)
	assert.Equal(t, r.String(), "ObServerRoster{maxPriority:1, "+
		"roster:[nil, nil, nil, nil, "+
		"ObServerAddr{ip:127.0.0.1, sqlPort:8000, svrPort:8080}, "+
		"ObServerAddr{ip:127.0.0.1, sqlPort:8000, svrPort:8080}, "+
		"ObServerAddr{ip:127.0.0.1, sqlPort:8000, svrPort:8080}]}")
}

func TestObServerRoute_ToString(t *testing.T) {
	r := ObServerRoute{}
	assert.Equal(t, r.String(), "ObServerRoute{readConsistency:0}")
	r = ObServerRoute{ObReadConsistencyStrong}
	assert.Equal(t, r.String(), "ObServerRoute{readConsistency:0}")
	r = ObServerRoute{ObReadConsistencyWeak}
	assert.Equal(t, r.String(), "ObServerRoute{readConsistency:1}")
}

func TestHash64a(t *testing.T) {
	result := MurmurHash64A([]byte{1}, len([]byte{1}), int64(0))
	assert.Equal(t, int64(-5720937396023583481), result)
	result = MurmurHash64A([]byte{1}, len([]byte{1}), int64(1))
	assert.Equal(t, int64(6351753276682545529), result)
	result = MurmurHash64A([]byte{1, 2, 3}, len([]byte{1, 2, 3}), int64(123456789))
	assert.Equal(t, int64(-4356950700900923028), result)
}

func TestHashSortUtf8Mb4(t *testing.T) {
	data := []byte{1}
	hashCode := 0
	res := hashSortUtf8Mb4(data, int64(hashCode), 10, true)
	assert.Equal(t, int64(-7030129012826305577), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 10, false)
	assert.Equal(t, int64(2570), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, true)
	assert.Equal(t, int64(-7030129012826305577), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, false)
	assert.Equal(t, int64(7062546676564130965), res)
	data = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res = hashSortUtf8Mb4(data, int64(hashCode), 10, true)
	assert.Equal(t, int64(-1916273894318764036), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 10, false)
	assert.Equal(t, int64(1013208367030014238), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, true)
	assert.Equal(t, int64(-1916273894318764036), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, false)
	assert.Equal(t, int64(7452881443849355883), res)
}

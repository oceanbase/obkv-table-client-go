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
	sql2 "database/sql"
	"math"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

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
	sql += createInStatement([]int{0, 1, 2, 3})
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
	entry, err := GetTableEntryFromRemote(context.TODO(), &testServerAddr, &testUserAuth, &key)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, entry.tableId, uint64(1099511677777))
	assert.Equal(t, entry.partNum, 4)
	assert.Equal(t, entry.replicaNum, 3)
	assert.Equal(t, entry.tableEntryKey, key)
	assert.EqualValues(t, entry.partitionInfo.level, PartLevelTwo)
	assert.Equal(t, entry.partitionInfo.firstPartDesc.String(), "obRangePartDesc{"+
		"comm:obPartDescCommon{partFuncType:obPartFuncType{name:RANGE_COLUMNS, index:4}, "+
		"partExpr:c1, orderedPartColumnNames:c1, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[obColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"obColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, partSpace:0, partNum:2, "+
		"orderedCompareColumns:[obColumn{columnName:c1, index:0, "+
		"objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c1], isGenColumn:false, columnExpress:nil}], "+
		"orderedCompareColumnTypes:[ObObjType{type:ObInt64Type}]}")
	assert.Contains(t, entry.partitionInfo.subPartDesc.String(), "obHashPartDesc{"+
		"comm:obPartDescCommon{partFuncType:obPartFuncType{name:HASH_V2, index:8}, "+
		"partExpr:c2, orderedPartColumnNames:c2, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[obColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"obColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, completeWorks:[], partSpace:0, partNum:2, "+
		"partNameIdMap:{m[p")
	assert.Equal(t, entry.tableLocation.String(), "ObTableLocation{"+
		"replicaLocations:["+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, "+
		"info:obServerStatus{stopTime:0, status:ACTIVE}, "+
		"role:obServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, "+
		"info:obServerStatus{stopTime:0, status:ACTIVE}, "+
		"role:obServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, "+
		"info:obServerStatus{stopTime:0, status:ACTIVE}, "+
		"role:obServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, len(entry.partLocationEntry.partLocations), 4)
	assert.Equal(t, entry.partLocationEntry.partLocations[0].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[1].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[2].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[3].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
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
	sql += createInStatement([]int{200005, 200006, 200007, 200008})
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
	entry, err := GetTableEntryFromRemote(context.TODO(), &testServerAddr, &testUserAuth, &key)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, entry.tableId, uint64(500012))
	assert.Equal(t, entry.partNum, 4)
	assert.Equal(t, entry.replicaNum, 3)
	assert.Equal(t, entry.tableEntryKey, key)
	assert.EqualValues(t, entry.partitionInfo.level, PartLevelTwo)
	assert.Equal(t, entry.partitionInfo.firstPartDesc.String(), "obRangePartDesc{"+
		"comm:obPartDescCommon{partFuncType:obPartFuncType{name:RANGE_COLUMNS, index:4}, "+
		"partExpr:c1, orderedPartColumnNames:c1, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[obColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"obColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, partSpace:0, partNum:2, "+
		"orderedCompareColumns:[obColumn{columnName:c1, index:0, "+
		"objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c1], isGenColumn:false, columnExpress:nil}], "+
		"orderedCompareColumnTypes:[ObObjType{type:ObInt64Type}]}")
	assert.Equal(t, entry.partitionInfo.subPartDesc.String(), "obHashPartDesc{"+
		"comm:obPartDescCommon{partFuncType:obPartFuncType{name:HASH_V2, index:8}, "+
		"partExpr:c2, orderedPartColumnNames:c2, orderedPartRefColumnRowKeyRelations:[], "+
		"partColumns:[obColumn{columnName:c1, index:0, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, refColumnNames:[c1], "+
		"isGenColumn:false, columnExpress:nil}, "+
		"obColumn{columnName:c2, index:1, objType:ObObjType{type:ObInt32Type}, "+
		"collationType:ObCollationType{collationType:CsTypeBinary}, "+
		"refColumnNames:[c2], isGenColumn:false, columnExpress:nil}], "+
		"rowKeyElement:nil}, completeWorks:[], partSpace:0, partNum:2, "+
		"partNameIdMap:{}}")
	assert.Equal(t, entry.tableLocation.String(), "ObTableLocation{"+
		"replicaLocations:["+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, "+
		"info:obServerStatus{stopTime:0, status:ACTIVE}, "+
		"role:obServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, "+
		"info:obServerStatus{stopTime:0, status:ACTIVE}, "+
		"role:obServerRole{name:FOLLOWER, index:2}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, "+
		"info:obServerStatus{stopTime:0, status:ACTIVE}, "+
		"role:obServerRole{name:LEADER, index:1}, "+
		"replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, len(entry.partLocationEntry.partLocations), 4)
	assert.Equal(t, entry.partLocationEntry.partLocations[200005].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[200006].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[200007].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	assert.Equal(t, entry.partLocationEntry.partLocations[200008].String(), "obPartitionLocation{"+
		"leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:42705, svrPort:42704}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.2, sqlPort:42707, svrPort:42706}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:FOLLOWER, index:2}, replicaType:ObReplicaType{name:FULL, index:0}}, "+
		"obReplicaLocation{addr:ObServerAddr{ip:127.0.0.3, sqlPort:42701, svrPort:42700}, info:obServerStatus{stopTime:0, status:ACTIVE}, role:obServerRole{name:LEADER, index:1}, replicaType:ObReplicaType{name:FULL, index:0}}]}")
	println(entry.String())
}

func TestObHashPartDesc_GetPartId(t *testing.T) {
	desc := &obHashPartDesc{}
	partId, err := desc.GetPartId(nil)
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)
	partId, err = desc.GetPartId([]interface{}{1, 2, 3})
	assert.NotEqual(t, nil, err)
	assert.EqualValues(t, ObInvalidPartId, partId)
}

func TestObHashPartDesc_innerHash(t *testing.T) {
	hashDesc := obHashPartDesc{partSpace: 0, partNum: 10}
	hashVal := hashDesc.innerHash(0)
	assert.Equal(t, int64(0), hashVal)
	hashDesc = obHashPartDesc{partSpace: 0, partNum: 10}
	hashVal = hashDesc.innerHash(1)
	assert.Equal(t, int64(1), hashVal)
	hashVal = hashDesc.innerHash(11)
	assert.Equal(t, int64(1), hashVal)
	hashVal = hashDesc.innerHash(-1)
	assert.Equal(t, int64(1), hashVal)
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

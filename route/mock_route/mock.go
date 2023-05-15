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

// Package mock_route only for test
package mock_route

import (
	"context"
	sql2 "database/sql"
	"math"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey/v2"

	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/oceanbase/obkv-table-client-go/util"
)

var (
	MockTestClusterName   = "test"
	MockTestTenantName    = "mysql"
	MockTestDatabase      = "MockTestDatabaseName"
	MockTestTableName     = "MockTestTableName"
	MockTestUserName      = "root"
	MockTestPassword      = ""
	MockTestIp            = "127.0.0.1"
	MockTestSqlPort       = 42705
	MockTestServerPort    = 42704
	MockTestServerAddr    = route.NewObServerAddr(MockTestIp, MockTestSqlPort, MockTestServerPort)
	MockTestUserAuth      = route.NewObUserAuth(MockTestUserName, MockTestPassword)
	MOckTestRowKeyElement = []string{"c1"}
)

// GetMockHashTableEntryV3 is for get 3.x version hash table entry.
// CREATE TABLE test(c1 INT, c2 int) PARTITION BY hash(c1) partitions 2;
func GetMockHashTableEntryV3() *route.ObTableEntry {
	proxyLocationSql := route.LocationSql
	proxyPartitionLocationSql := route.PartitionLocationSql
	proxyPartitionInfoSql := route.PartitionInfoSql
	proxyFirstPartitionSql := route.FirstPartitionSql
	// 1. mock function 'GetObVersionFromRemote'
	patchVer := gomonkey.ApplyFunc(route.GetObVersionFromRemote, func(addr *route.ObServerAddr, sysUA *route.ObUserAuth) (float32, error) {
		return 3.24, nil
	})
	defer patchVer.Reset()
	// 2. check version
	ver, err := route.GetObVersionFromRemote(MockTestServerAddr, MockTestUserAuth)
	if err != nil {
		panic(err)
	}
	// 3. set version
	util.SetObVersion(ver)
	route.InitSql(ver)
	// 4. mock
	// 4.1 new mock DB
	mockDB, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// 4.2 mock function 'NewDB'
	patchDb := gomonkey.ApplyFunc(route.NewDB, func(userName string, password string, ip string, port string, database string) (*route.DB, error) {
		return mockDB, nil
	})
	defer patchDb.Reset()

	// 4.3 mock proxyLocationSql result
	sql := proxyLocationSql
	queryFields := []string{"partition_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows := sqlmock.NewRows(queryFields).
		AddRow(0, "127.0.0.1", 42705, 1099511677777, 1, 3, 2, 42704, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.2", 42707, 1099511677777, 2, 3, 2, 42706, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.3", 42701, 1099511677777, 2, 3, 2, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(MockTestTenantName, MockTestDatabase, MockTestTableName).WillReturnRows(mockRows)

	// 4.4 mock proxyPartitionInfoSql result
	sql = proxyPartitionInfoSql
	queryFields = []string{"part_level", "part_num", "part_type", "part_space",
		"part_expr", "part_range_type", "sub_part_num", "sub_part_type",
		"sub_part_space", "sub_part_range_type", "sub_part_expr", "part_key_name",
		"part_key_type", "part_key_idx", "part_key_extra", "spare1",
	}

	mockRows = sqlmock.NewRows(queryFields).
		AddRow(1, 2, 8, 0, "c1", "", 1, 0, 0, "", "", "c1", 4, 0, "", 63)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1099511677777, math.MaxInt64).WillReturnRows(mockRows)

	// 4.5 mock proxyFirstPartitionSql result
	sql = proxyFirstPartitionSql
	queryFields = []string{"part_id", "part_name", "high_bound_val"}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(0, "p0", sql2.NullString{}).
		AddRow(1, "p1", sql2.NullString{})
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1099511677777, math.MaxInt64).WillReturnRows(mockRows)

	// 4.7 mock proxyPartitionLocationSql result
	sql = proxyPartitionLocationSql
	sql += "(0, 1);"
	queryFields = []string{"partition_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(0, "127.0.0.1", 42705, 1099511677777, 1, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.3", 42701, 1099511677777, 2, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(1, "127.0.0.1", 42705, 1099511677777, 1, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(1, "127.0.0.2", 42707, 1099511677777, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(1, "127.0.0.3", 42701, 1099511677777, 2, 3, 4, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(MockTestTenantName, MockTestDatabase, MockTestTableName).WillReturnRows(mockRows)

	// 5. test
	key := route.NewObTableEntryKey(
		MockTestClusterName,
		MockTestTenantName,
		MockTestDatabase,
		MockTestTableName,
	)
	entry, err := route.GetTableEntryFromRemote(context.TODO(), MockTestServerAddr, MockTestUserAuth, key)
	if err != nil {
		panic(err)
	}
	return entry
}

// GetMockHashTableEntryV4 is for get 4.x version hash table entry.
// CREATE TABLE test(c1 INT, c2 int) PARTITION BY hash(c1) partitions 2;
func GetMockHashTableEntryV4() *route.ObTableEntry {
	proxyLocationSql := route.LocationSqlV4
	proxyPartitionLocationSql := route.PartitionLocationSqlV4
	proxyPartitionInfoSql := route.PartitionInfoSqlV4
	proxyFirstPartitionSql := route.FirstPartitionSqlV4
	// 1. mock function 'GetObVersionFromRemote'
	patchVer := gomonkey.ApplyFunc(route.GetObVersionFromRemote, func(addr *route.ObServerAddr, sysUA *route.ObUserAuth) (float32, error) {
		return 4.1, nil
	})
	defer patchVer.Reset()
	// 2. check version
	ver, err := route.GetObVersionFromRemote(MockTestServerAddr, MockTestUserAuth)
	if err != nil {
		panic(err)
	}
	// 3. set version
	util.SetObVersion(ver)
	route.InitSql(ver)
	// 4. mock
	// 4.1 new mock DB
	mockDB, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	// 4.2 mock function 'NewDB'
	patchDb := gomonkey.ApplyFunc(route.NewDB, func(userName string, password string, ip string, port string, database string) (*route.DB, error) {
		return mockDB, nil
	})
	defer patchDb.Reset()

	// 4.3 mock proxyLocationSql result
	sql := proxyLocationSql
	queryFields := []string{"tablet_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows := sqlmock.NewRows(queryFields).
		AddRow(0, "127.0.0.1", 42705, 500012, 1, 3, 2, 42704, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.2", 42707, 500012, 2, 3, 2, 42706, "ACTIVE", 0, 0).
		AddRow(0, "127.0.0.3", 42701, 500012, 2, 3, 2, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(MockTestTenantName, MockTestDatabase, MockTestTableName).WillReturnRows(mockRows)

	// 4.4 mock proxyPartitionInfoSql result
	sql = proxyPartitionInfoSql
	queryFields = []string{"part_level", "part_num", "part_type", "part_space",
		"part_expr", "part_range_type", "sub_part_num", "sub_part_type",
		"sub_part_space", "sub_part_range_type", "sub_part_expr", "part_key_name",
		"part_key_type", "part_key_idx", "part_key_extra", "part_key_collation_type",
	}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(1, 2, 0, 0, "c1", "", 0, 0, 0, "", "", "c1", 4, 0, "", 63)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(MockTestTenantName, 500012, math.MaxInt64).WillReturnRows(mockRows)

	// 4.5 mock proxyFirstPartitionSql result
	sql = proxyFirstPartitionSql
	queryFields = []string{"part_id", "part_name", "tablet_id", "high_bound_val", "sub_part_num"}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(500013, "p1", 200009, sql2.NullString{}, 0).
		AddRow(500014, "p2", 200010, sql2.NullString{}, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(MockTestTenantName, 500012, math.MaxInt64).WillReturnRows(mockRows)

	// 4.6 mock proxyPartitionLocationSql result
	sql = proxyPartitionLocationSql
	sql += "(0, 1);"
	queryFields = []string{"tablet_id", "svr_ip", "sql_port", "table_id",
		"role", "replica_num", "part_num", "svr_port", "status", "stop_time", "replica_type",
	}
	mockRows = sqlmock.NewRows(queryFields).
		AddRow(200009, "127.0.0.1", 42705, 500012, 1, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(200009, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(200009, "127.0.0.3", 42701, 500012, 2, 3, 4, 42700, "ACTIVE", 0, 0).
		AddRow(200010, "127.0.0.1", 42705, 500012, 1, 3, 4, 42704, "ACTIVE", 0, 0).
		AddRow(200010, "127.0.0.2", 42707, 500012, 2, 3, 4, 42706, "ACTIVE", 0, 0).
		AddRow(200010, "127.0.0.3", 42701, 500012, 2, 3, 4, 42700, "ACTIVE", 0, 0)
	sqlMock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(MockTestTenantName, MockTestDatabase, MockTestTableName).WillReturnRows(mockRows)

	// 5. test
	key := route.NewObTableEntryKey(
		MockTestClusterName,
		MockTestTenantName,
		MockTestDatabase,
		MockTestTableName,
	)
	entry, err := route.GetTableEntryFromRemote(context.TODO(), MockTestServerAddr, MockTestUserAuth, key)
	if err != nil {
		panic(err)
	}
	return entry
}

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

package login

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

const (
	sqlUser     = "root@mysql"
	sqlPassWord = ""
	sqlIp       = "..."
	sqlPort     = "..."
)

const (
	database1 = "test_database1"
	database2 = "test_database2"
)

var sqlDB1 *sql.DB
var sqlDB2 *sql.DB

func createDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", sqlUser, sqlPassWord, sqlIp, sqlPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createDbWithDatabase(database string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sqlUser, sqlPassWord, sqlIp, sqlPort, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createDatabase(db *sql.DB, database string) {
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database)
	_, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func setup() {
	// 1. create sql.DB
	sqlDB := createDB()

	// 2. create database1, database2
	createDatabase(sqlDB, database1)
	createDatabase(sqlDB, database2)

	// 3. create table in each database
	sqlDB1 = createDbWithDatabase(database1)
	sqlDB2 = createDbWithDatabase(database2)
	sqlDB1.Exec(testLoginCreateStatement)
	sqlDB2.Exec(testLoginCreateStatement)

}

func teardown() {
	sqlDB1.Exec(fmt.Sprintf("DROP TABLE %s;", testLoginTableName))
	sqlDB2.Exec(fmt.Sprintf("DROP TABLE %s;", testLoginTableName))
}

func TestMain(m *testing.M) {
	if passLoginTest {
		fmt.Println("Please run login tests manually!!!")
		fmt.Println("Change passLoginTest to false in test/login/login_test.go to run login tests.")
		return
	}
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

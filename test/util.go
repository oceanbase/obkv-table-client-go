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

package test

import (
	"database/sql"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
)

const (
	configUrl    = "..."
	fullUserName = "..."
	passWord     = ""
	sysUserName  = "root"
	sysPassWord  = ""
)

// use for create table by sql driver
const (
	sqlUser     = "root"
	sqlPassWord = ""
	sqlIp       = "..."
	sqlPort     = "..."
	sqlDatabase = "test"
)

var globalDb *sql.DB = nil

func newClient() client.Client {
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, config.NewDefaultClientConfig())
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func newDB() *sql.DB {
	if globalDb == nil {
		dsn := sqlUser + ":" + sqlPassWord + "@tcp(" + sqlIp + ":" + sqlPort + ")/" + sqlDatabase
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		globalDb = db
	}
	return globalDb
}

func createTable(createTableStatement string) {
	db := newDB()
	_, err := db.Exec(createTableStatement)
	if err != nil {
		panic(err.Error())
	}
}

func dropTable(tableName string) {
	db := newDB()
	_, _ = db.Exec("DROP TABLE " + tableName + ";")
}

func deleteTable(tableName string) {
	db := newDB()
	_, _ = db.Exec("DELETE FROM " + tableName + ";")
}

func selectTable(selectStatement string) *sql.Rows {
	db := newDB()
	rows, err := db.Query(selectStatement)
	if err != nil {
		panic(err.Error())
	}
	return rows
}

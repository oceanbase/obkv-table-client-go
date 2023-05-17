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
	"fmt"

	_ "github.com/go-sql-driver/mysql"

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

func CreateClient() client.Client {
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, config.NewDefaultClientConfig())
	if err != nil {
		panic(err.Error())
	}
	return cli
}

var globalDB *sql.DB

func CreateDB() {
	if globalDB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sqlUser, passWord, sqlIp, sqlPort, sqlDatabase)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		globalDB = db
	}
}

func CloseDB() {
	globalDB.Close()
}

func CreateTable(createTableStatement string) {
	_, err := globalDB.Exec(createTableStatement)
	if err != nil {
		panic(err.Error())
	}
}

func DropTable(tableName string) {
	_, err := globalDB.Exec(fmt.Sprintf("drop table %s;", tableName))
	if err != nil {
		panic(err.Error())
	}
}

func DeleteTable(tableName string) {
	_, err := globalDB.Exec(fmt.Sprintf("delete from %s;", tableName))
	if err != nil {
		panic(err.Error())
	}
}

func SelectTable(selectStatement string) *sql.Rows {
	rows, err := globalDB.Query(selectStatement)
	if err != nil {
		panic(err.Error())
	}
	return rows
}

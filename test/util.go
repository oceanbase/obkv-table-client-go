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
	"time"
)

const (
	configUrl    = "..."
	fullUserName = "..."
	passWord     = ""
	sysUserName  = "root"
	sysPassWord  = ""

	// odp
	odpIP           = ""
	odpFullUserName = "..."
	odpPassWord     = ""
	odpRpcPort      = 0
	database        = "..."

	// toml config
	tomlConfigPath = ""
)

const (
	sqlUser     = "root"
	sqlPassWord = ""
	sqlIp       = "..."
	sqlPort     = "..."
	sqlDatabase = "test"
)

func CreateClient() client.Client {
	var cli client.Client
	var err error
	if tomlConfigPath != "" {
		cli, err = client.NewClientWithTomlConfig(tomlConfigPath)
	} else if odpIP != "" {
		cli, err = client.NewOdpClient(odpFullUserName, odpPassWord, odpIP, odpRpcPort, database, config.NewDefaultClientConfig())
	} else {
		cli, err = client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, config.NewDefaultClientConfig())
	}
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func CreateMoveClient() client.Client {
	cfg := config.NewDefaultClientConfig()
	cfg.EnableRerouting = true
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func CreateConnectionBalanceClient(maxConnectionAge time.Duration, enableSLBLoadBalance bool, connectionPoolSize int) client.Client {
	cfg := config.NewDefaultClientConfig()
	cfg.MaxConnectionAge = maxConnectionAge
	cfg.ConnPoolMaxConnSize = connectionPoolSize
	cfg.EnableSLBLoadBalance = enableSLBLoadBalance

	cli, err := client.NewOdpClient(odpFullUserName, odpPassWord, odpIP, odpRpcPort, database, cfg)
	if err != nil {
		panic(err.Error())
	}
	println("connection Balance Client Created")
	return cli
}

var GlobalDB *sql.DB

func CreateDB() {
	if GlobalDB == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", sqlUser, sqlPassWord, sqlIp, sqlPort, sqlDatabase)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		GlobalDB = db
	}
}

func CloseDB() {
	GlobalDB.Close()
}

func CreateTable(createTableStatement string) {
	_, err := GlobalDB.Exec(createTableStatement)
	if err != nil {
		panic(err.Error())
	}
}

func DropTable(tableName string) {
	_, err := GlobalDB.Exec(fmt.Sprintf("drop table %s;", tableName))
	if err != nil {
		panic(err.Error())
	}
}

func DeleteTable(tableName string) {
	_, err := GlobalDB.Exec(fmt.Sprintf("delete from %s;", tableName))
	if err != nil {
		panic(err.Error())
	}
}

func DeleteTables(tableNames []string) {
	for _, name := range tableNames {
		DeleteTable(name)
	}
}

func InsertTable(insertStatement string) {
	_, err := GlobalDB.Exec(insertStatement)
	if err != nil {
		panic(err.Error())
	}
}

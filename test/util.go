package test

import (
	"database/sql"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
)

const (
	configUrl    = "http://ocp-cfg.alibaba.net:8080/services?User_ID=alibaba&UID=test&Action=ObRootServiceInfo&ObCluster=ob10.chenweixin.cwx.11.158.97.240&database=test"
	fullUserName = "root@sys#ob10.chenweixin.cwx.11.158.97.240"
	passWord     = ""
	sysUserName  = "root"
	sysPassWord  = ""
)

// use for create table by sql driver
const (
	sqlUser     = "root"
	sqlPassWord = ""
	sqlIp       = "11.158.97.240"
	sqlPort     = "41101"
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
	_, err := db.Exec("DROP TABLE " + tableName + ";")
	if err != nil {
		panic(err.Error())
	}
}

func deleteTable(tableName string) {
	db := newDB()
	_, err := db.Exec("DELETE FROM " + tableName + ";")
	if err != nil {
		panic(err.Error())
	}
}

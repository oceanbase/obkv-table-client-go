package test

import (
	"database/sql"
	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
)

const (
	configUrl    = ""
	fullUserName = ""
	passWord     = ""
	sysUserName  = "root"
	sysPassWord  = ""
)

// use for create table by sql driver
const (
	sqlUser     = "root"
	sqlPassWord = ""
	sqlIp       = ""
	sqlPort     = ""
	sqlDatabase = "test"
)

func newClient() client.Client {
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, config.NewDefaultClientConfig())
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func createTable(createTableStatement string) {
	dsn := sqlUser + ":" + sqlPassWord + "@tcp(" + sqlIp + ":" + sqlPort + ")/" + sqlDatabase
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(createTableStatement)
	if err != nil {
		panic(err.Error())
	}
}

func dropTable(tableName string) {
	dsn := sqlUser + ":" + sqlPassWord + "@tcp(" + sqlIp + ":" + sqlPort + ")/" + sqlDatabase
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE " + tableName + ";")
	if err != nil {
		panic(err.Error())
	}
}

func deleteTable(tableName string) {
	dsn := sqlUser + ":" + sqlPassWord + "@tcp(" + sqlIp + ":" + sqlPort + ")/" + sqlDatabase
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM " + tableName + ";")
	if err != nil {
		panic(err.Error())
	}
}

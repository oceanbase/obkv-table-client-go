package main

import (
	"fmt"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

func main() {
	const (
		configUrl    = "xxx"
		fullUserName = "root@sys#obcluster"
		passWord     = ""
		sysUserName  = "root"
		sysPassWord  = ""
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	if err != nil {
		panic(err)
	}

	err = cli.AddRowkey(tableName, []string{"c1"})
	if err != nil {
		panic(err)
	}
	rowkey := []*table.Column{table.NewColumn("c1", int64(1))}
	selectColumns := []string{"c1", "c2"}
	m, err := cli.Get(
		tableName,
		rowkey,
		selectColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(m["c1"])
	fmt.Println(m["c2"])
}

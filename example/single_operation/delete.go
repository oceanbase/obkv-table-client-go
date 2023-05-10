package main

import (
	"context"
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

	err = cli.AddRowKey(tableName, []string{"c1"})
	if err != nil {
		panic(err)
	}
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	affectRows, err := cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)
}

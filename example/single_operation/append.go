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

package main

import (
	"context"
	"fmt"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// CREATE TABLE test(c1 bigint, c2 varchar(20), PRIMARY KEY(c1)) PARTITION BY hash(c1) partitions 2;
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

	// insert
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	insertColumns := []*table.Column{table.NewColumn("c2", "hello")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		insertColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// append c2(2) -> c2("hello") + (" oceanbase") = c2("hello oceanbase")
	appendColumns := []*table.Column{table.NewColumn("c2", " oceanbase")}
	res, err := cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		appendColumns,
		client.WithReturnRowKey(true),         // return rowKey if you need
		client.WithReturnAffectedEntity(true), // return result entity if you need
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.AffectedRows())
	fmt.Println(res.RowKey())
	fmt.Println(res.Value("c2"))
}

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
	"time"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// CREATE TABLE test(c1 bigint, c2 varchar(20), PRIMARY KEY(c1)) PARTITION BY hash(c1) partitions 2;
func main() {
	const (
		configUrl    = "xxx"
		fullUserName = "user@tenant#cluster"
		passWord     = ""
		sysUserName  = "sysUser"
		sysPassWord  = ""
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cli, err := client.NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	if err != nil {
		panic(err)
	}

	// insert
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	insertColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	affectRows, err := cli.Insert(
		ctx,
		tableName,
		rowKey,
		insertColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// get
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		selectColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Value("c1"))
	fmt.Println(result.Value("c2"))
}

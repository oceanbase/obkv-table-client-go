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
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
	"time"
)

// CREATE TABLE test(c1 bigint, c2 bigint, PRIMARY KEY(c1));
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

	// insert, prepare data
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	insertColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	affectRows, err := cli.InsertOrUpdate(
		ctx,
		tableName,
		rowKey,
		insertColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// set key range
	startRowKey := []*table.Column{table.NewColumn("c1", "1")}
	endRowKey := []*table.Column{table.NewColumn("c1", "10")}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	rowKey = []*table.Column{table.NewColumn("c1", "4")}
	mutationColumns := []*table.Column{table.NewColumn("c2", int64(3))}

	// satisfy the filter, insert (4, 3)
	affectRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c2", int64(2))), // where c2 = 2
		option.WithScanRange(keyRanges),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)
}

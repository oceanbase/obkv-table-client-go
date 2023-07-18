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
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"time"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// CREATE TABLE test(c1 bigint, c2 bigint, PRIMARY KEY(c1)) PARTITION BY hash(c1) partitions 2;
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

	// update c2(2) -> c2(3)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns := []*table.Column{table.NewColumn("c2", int64(3))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.LessThan, "c2", int64(3))), // where c2 < 3
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// update c2(3) -> c2(4)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns = []*table.Column{table.NewColumn("c2", int64(4))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.GreaterThan, "c2", int64(2))), // where c2 > 2
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// update c2(4) -> c2(5)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns = []*table.Column{table.NewColumn("c2", int64(5))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.LessOrEqualThan, "c2", int64(4))), // where c2 <= 4
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// update c2(5) -> c2(6)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns = []*table.Column{table.NewColumn("c2", int64(6))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.GreaterOrEqualThan, "c2", int64(5))), // where c2 >= 5
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// update c2(6) -> c2(7)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns = []*table.Column{table.NewColumn("c2", int64(7))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.NotEqual, "c2", int64(7))), // where c2 != 7
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// update c2(7) -> c2(9)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns = []*table.Column{table.NewColumn("c2", int64(9))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c2", int64(7))), // where c2 = 7
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// update c2(9) -> c2(10)
	var andFilterList []filter.ObTableFilter
	andFilterList = append(andFilterList, filter.CompareVal(filter.LessThan, "c2", int64(10)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.GreaterThan, "c2", int64(8)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.LessOrEqualThan, "c2", int64(9)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.GreaterOrEqualThan, "c2", int64(9)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.NotEqual, "c2", int64(0)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.Equal, "c2", int64(9)))
	andList := filter.AndList(andFilterList...)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) // 10s
	updateColumns = []*table.Column{table.NewColumn("c2", int64(10))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(andList), // where c2 < 10 and c2 > 8 and c2 <= 9 and c2 >= 9 and c2 != 0 and c2 = 9
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)
}

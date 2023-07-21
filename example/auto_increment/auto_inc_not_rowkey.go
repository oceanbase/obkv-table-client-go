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

package auto_increment

import (
	"context"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
	"time"
)

// create table if not exists autoIncNotRowkeyTable(`c1` bigint(20) not null, c2 bigint(20) auto_increment, c3 varchar(20) default 'hello', c4 bigint(20) default 0, primary key (`c1`)) partition by key(c1) partitions 100;
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

	// insert. set 0, use auto increment value
	// insert (1, 1, hello, 0)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	affectRows, err := cli.Insert(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// get and check
	startRowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)

	if err != nil {
		panic(err)
	}

	res, err := resSet.Next()

	println(res.Value("c1").(int64)) // get 1
	println(res.Value("c2").(int64)) // get 1

	// insert. set 50, use specified value
	// insert (2, 50, hello, null)
	rowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	mutationColumns := []*table.Column{table.NewColumn("c2", int64(50))}
	affectRows, err = cli.Insert(
		ctx,
		tableName,
		rowKey,
		mutationColumns,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)

	// get and check
	startRowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)

	if err != nil {
		panic(err)
	}

	res, err = resSet.Next()

	println(res.Value("c1").(int64)) // get 2
	println(res.Value("c2").(int64)) // get 50
}

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

package aggregate

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// create table if not exists aggregateTable(`c1` bigint(20) not null, c2 bigint(20) not null, c3 varchar(20) default 'hello', c4 double default null, primary key (`c1`, `c2`))
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

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	// create agg executor
	aggExecutor := cli.NewAggExecutor(tableName, keyRanges).Sum("c3").Min("c2").Max("c1").Count()

	// get agg result
	res, err := aggExecutor.Execute(context.TODO())

	if err != nil {
		panic(err)
	}

	// get agg result value
	println(res.Value("sum(c3)").(int64))
	println(res.Value("min(c2)").(int64))
	println(res.Value("max(c1)").(int64))
	println(res.Value("count(*)").(int64))
}

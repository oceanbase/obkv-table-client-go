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
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
	"time"
)

// create table if not exists autoIncRowkeyTable(`c1` bigint(20) not null auto_increment, c2 bigint(20) not null, c3 varchar(20) default 'hello', c4 bigint(20) default 0, primary key (`c1`, `c2`)) partition by hash(c2) partitions 15;
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
	rowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}
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

	// insert. set 50, use specified value
	// insert (50, 1, hello, 0)
	rowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}
	affectRows, err = cli.Insert(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(affectRows)
}

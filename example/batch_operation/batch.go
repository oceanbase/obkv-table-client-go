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

package batch_operation

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

	rowKey1 := []*table.Column{table.NewColumn("c1", int64(1))}
	rowKey2 := []*table.Column{table.NewColumn("c1", int64(2))}
	selectColumns1 := []string{"c1"}
	selectColumns2 := []string{"c2"}
	mutateColumns1 := []*table.Column{table.NewColumn("c2", int64(1))}
	mutateColumns2 := []*table.Column{table.NewColumn("c2", int64(2))}

	batchExecutor := cli.NewBatchExecutor(tableName)
	err = batchExecutor.AddInsertOp(rowKey1, mutateColumns1)
	if err != nil {
		panic(err)
	}
	err = batchExecutor.AddInsertOp(rowKey2, mutateColumns2)
	if err != nil {
		panic(err)
	}
	err = batchExecutor.AddGetOp(rowKey1, selectColumns1)
	if err != nil {
		panic(err)
	}
	err = batchExecutor.AddGetOp(rowKey2, selectColumns2)
	if err != nil {
		panic(err)
	}
	batchRes, err := batchExecutor.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("size:%d, success:%d, fail:%d", batchRes.Size(), batchRes.CorrectCount(), batchRes.WrongCount())
	allResults := batchRes.GetResults()
	println(allResults[0].AffectedRows())
	println(allResults[1].AffectedRows())
}

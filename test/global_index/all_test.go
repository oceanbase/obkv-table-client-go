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
package global_index

import (
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
	"os"
	"testing"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()
	test.CreateDB()
	test.ExecStatement("set global ob_enable_index_direct_select = 1;")
	test.CreateTable(testGlabalIndexCreateHashTable)
	test.CreateTable(testGlabalIndexCreateKeyTable)
	test.CreateTable(testGlobalIndexNoPartCreateStat)
	test.CreateTable(testGlobalAllNoPartCreateStat)
	test.CreateTable(testGlobalPrimaryNoPartCreateStat)
}

func teardown() {
	cli.Close()
	test.DropTable(testGlabalIndexHashTableName)
	test.DropTable(testGlabalIndexKeyTableName)
	test.DropTable(testGlobalIndexNoPart)
	test.DropTable(testGlobalAllNoPart)
	test.DropTable(testGlobalPrimaryNoPart)
	test.CloseDB()
}

func createTableAndGlobalIndex(createTable string, createIndex string) {
	test.ExecStatement(createTable)
	test.ExecStatement(createIndex)
}

func getGlobalIndexTableName(tableName string) []string {
	var res []string
	if test.GlobalDB == nil {
		panic("db is nil")
	}
	sql := fmt.Sprintf("select table_name from oceanbase.__all_virtual_table where data_table_id = (select table_id from oceanbase.__all_virtual_table where table_name = '%s');", tableName)
	rows, err := test.GlobalDB.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	var indexTableName string
	for rows.Next() {
		err = rows.Scan(&indexTableName)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, indexTableName)
	}
	return res
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

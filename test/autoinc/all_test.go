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

package autoinc

import (
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()

	test.CreateDB()
	test.CreateTable(autoIncRowkeyZeroFillTableCreateStatement)
	test.CreateTable(autoIncRowkeyNotFillTableCreateStatement)
	test.CreateTable(autoIncNormalNotFillTableCreateStatement)
	test.CreateTable(autoIncNormalFillTableCreateStatement)
	test.CreateTable(autoIncRowkeyNilFillTableCreateStatement)
	test.CreateTable(autoIncNormalNilFillTableCreateStatement)
}

func teardown() {
	cli.Close()

	test.DropTable(autoIncRowkeyZeroFillTableTableName)
	test.DropTable(autoIncRowkeyNotFillTableTableName)
	test.DropTable(autoIncNormalNotFillTableTableName)
	test.DropTable(autoIncNormalFillTableTableName)
	test.DropTable(autoIncRowkeyNilFillTableTableName)
	test.DropTable(autoIncNormalNilFillTableTableName)
	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

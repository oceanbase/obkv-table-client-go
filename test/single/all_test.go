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

package single

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

	test.CreateTable(testTinyintCreateStatement)
	test.CreateTable(testUTinyintCreateStatement)
	test.CreateTable(testSmallintCreateStatement)
	test.CreateTable(testUSmallintCreateStatement)
	test.CreateTable(testInt32CreateStatement)
	test.CreateTable(testUInt32CreateStatement)
	test.CreateTable(testInt64CreateStatement)
	test.CreateTable(testUInt64CreateStatement)
	test.CreateTable(testFloatCreateStatement)
	test.CreateTable(testDoubleCreateStatement)
	test.CreateTable(testVarcharCreateStatement)
	test.CreateTable(testVarbinaryCreateStatement)
	test.CreateTable(testDatetimeCreateStatement)
	test.CreateTable(testTimestampCreateStatement)
	test.CreateTable(testDateCreateStatement)
	test.CreateTable(testYearCreateStatement)
}

func teardown() {
	cli.Close()

	test.DropTable(testTinyintTableName)
	test.DropTable(testUTinyintTableName)
	test.DropTable(testSmallintTableName)
	test.DropTable(testUSmallintTableName)
	test.DropTable(testInt32TableName)
	test.DropTable(testUInt32TableName)
	test.DropTable(testInt64TableName)
	test.DropTable(testUInt64TableName)
	test.DropTable(testFloatTableName)
	test.DropTable(testDoubleTableName)
	test.DropTable(testVarcharTableName)
	test.DropTable(testVarbinaryTableName)
	test.DropTable(testDatetimeTableName)
	test.DropTable(testTimestampTableName)
	test.DropTable(testDateTableName)
	test.DropTable(testYearTableName)

	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

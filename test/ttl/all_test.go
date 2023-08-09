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

package ttl

import (
	"fmt"
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

var cli client.Client

func setup() {
	if passTTLTest {
		fmt.Println("Please run TTL tests manually!!!")
		fmt.Println("Change passTTLTest to false in test/ttl/ttl_test.go to run ttl tests.")
		return
	}

	cli = test.CreateClient()

	test.CreateDB()

	test.CreateTable(testTTLCreateStatement)
}

func teardown() {
	if passTTLTest {
		fmt.Println("Please run TTL tests manually!!!")
		fmt.Println("Change passTTLTest to false in test/ttl/ttl_test.go to run ttl tests.")
		return
	}

	cli.Close()

	test.DropTable(testTTLTableName)

	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

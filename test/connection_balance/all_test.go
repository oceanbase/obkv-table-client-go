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

package connection_balance

import (
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
	"os"
	"testing"
	"time"
)

var cli client.Client

const (
	testConnectionBalanceTableName = "test_connection_balance"
	// NOTE: cannot create table directly in obkv cluster
	testConnectionBalanceCreateStatement = "create table if not exists `test_connection_balance`(`c1` varchar(1024) primary key,`c2` int);"
	concurrencyNum                       = 300
	// NOTE: make sure test timeout is greatter than testDuration, e.g, go test -timeout 10m
	testDuration         = time.Duration(8) * time.Minute
	maxConnectionAge     = time.Duration(10) * time.Second
	connectionPoolSize   = 150
	enableSLBLoadBalance = true
)

func setup() {
	cli = test.CreateConnectionBalanceClient(maxConnectionAge, enableSLBLoadBalance, connectionPoolSize)
	fmt.Println("connection balance setup")
}

func teardown() {
	cli.Close()
	fmt.Println("connection balance teardown")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

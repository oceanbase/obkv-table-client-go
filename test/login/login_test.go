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

package login

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

const (
	testLoginTableName       = "test_login"
	testLoginCreateStatement = "create table if not exists test_login(c1 int primary key, c2 int) ;"
)

const (
	passLoginTest = false
)

const (
	configUrl    = "url without database"
	fullUserName = "userName@tenantName#clusterName"
	passWord     = ""
	sysUserName  = "root"
	sysPassWord  = ""
)

func TestLogin_login(t *testing.T) {
	cfg := config.NewDefaultClientConfig()
	url1 := configUrl + "&" + "database=" + database1
	cli1, err := client.NewClient(url1, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	assert.Equal(t, nil, err)

	rowKey := []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err := cli1.Insert(
		context.TODO(),
		testLoginTableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	url2 := configUrl + "&" + "database=" + database2
	cli2, err := client.NewClient(url2, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	assert.Equal(t, nil, err)
	affectRows, err = cli2.Insert(
		context.TODO(),
		testLoginTableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	res, err := cli1.Get(
		context.TODO(),
		testLoginTableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.Value("c1"))

	res, err = cli2.Get(
		context.TODO(),
		testLoginTableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.Value("c1"))
}

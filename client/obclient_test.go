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

package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/config"
)

func TestNewObClient(t *testing.T) {
	const (
		testConfigUrl    = "http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx"
		testFullUserName = "user@mysql#obkv_cluster"
		testPassWord     = ""
		testSysUserName  = "sys"
		testSysPassWord  = ""
	)
	cfg := config.NewDefaultClientConfig()
	cli, err := newObClient(testConfigUrl, testFullUserName, testPassWord, testSysUserName, testSysPassWord, cfg)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, cli)
}

func TestObClient_String(t *testing.T) {
	cli := &ObClient{}
	assert.Equal(t, "ObClient{configUrl:, fullUserName:, userName:, tenantName:, clusterName:, sysUA:ObUserAuth{userName:, password:}, config:nil}", cli.String())
	const (
		testConfigUrl    = "http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx"
		testFullUserName = "user@mysql#obkv_cluster"
		testPassWord     = ""
		testSysUserName  = "sys"
		testSysPassWord  = ""
	)
	cfg := config.NewDefaultClientConfig()
	cli, err := newObClient(testConfigUrl, testFullUserName, testPassWord, testSysUserName, testSysPassWord, cfg)
	assert.Equal(t, nil, err)
	assert.Equal(t, "ObClient{configUrl:http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx, fullUserName:user@mysql#obkv_cluster, userName:user, tenantName:mysql, clusterName:obkv_cluster, sysUA:ObUserAuth{userName:sys, password:}, config:ClientConfig{ConnPoolMaxConnSize:1, OperationTimeOut:10000000000, LogLevel:7}}", cli.String())
}

func TestObClient_parseFullUserName(t *testing.T) {
	testFullUserName := "user@mysql#obkv_cluster"
	cli := &ObClient{}
	err := cli.parseFullUserName(testFullUserName)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "user", cli.userName)
	assert.EqualValues(t, "mysql", cli.tenantName)
	assert.EqualValues(t, "obkv_cluster", cli.clusterName)
	assert.EqualValues(t, testFullUserName, cli.fullUserName)
	testFullUserName = "@user@mysql#obkv_cluster"
	err = cli.parseFullUserName(testFullUserName)
	assert.NotEqual(t, nil, err)
	testFullUserName = "user@@mysql#obkv_cluster"
	err = cli.parseFullUserName(testFullUserName)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "user", cli.userName)
	assert.EqualValues(t, "@mysql", cli.tenantName)
	assert.EqualValues(t, "obkv_cluster", cli.clusterName)
	assert.EqualValues(t, testFullUserName, cli.fullUserName)
	testFullUserName = "user@@mysql#####obkv_cluster"
	err = cli.parseFullUserName(testFullUserName)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "user", cli.userName)
	assert.EqualValues(t, "@mysql", cli.tenantName)
	assert.EqualValues(t, "####obkv_cluster", cli.clusterName)
	assert.EqualValues(t, testFullUserName, cli.fullUserName)
	testFullUserName = "12434234jaFFhfkj@##$$ddfkFFShf#obkv_cluster"
	err = cli.parseFullUserName(testFullUserName)
	assert.NotEqual(t, nil, err)
}

func TestObClient_parseConfigUrl(t *testing.T) {
	testConfigUrl := "http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx"
	cli := &ObClient{}
	err := cli.parseConfigUrl(testConfigUrl)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, testConfigUrl, cli.configUrl)
	assert.EqualValues(t, "xxx", cli.database)
	testConfigUrl = "database=xxx&database=xxx"
	err = cli.parseConfigUrl(testConfigUrl)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, testConfigUrl, cli.configUrl)
	assert.EqualValues(t, "xxx&database=xxx", cli.database)
	testConfigUrl = "DATABASE=xxx"
	err = cli.parseConfigUrl(testConfigUrl)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, testConfigUrl, cli.configUrl)
	assert.EqualValues(t, "xxx", cli.database)
}

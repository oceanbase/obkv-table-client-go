package client

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/oceanbase/obkv-table-client-go/route/mock_route"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TestConfigUrl    = "http://127.0.0.1:8080/services?User_ID=xxx&UID=xxx&Action=ObRootServiceInfo&ObCluster=xxx&database=xxx"
	TestFullUserName = "user@mysql#obkv_cluster"
	TestPassWord     = ""
	TestSysUserName  = "sys"
	TestSysPassWord  = ""
)

func getMockObClient() (*ObClient, error) {
	// CREATE TABLE test(c1 INT, c2 int) PARTITION BY hash(c1) partitions 2;
	obCli, err := newObClient(
		TestConfigUrl,
		TestFullUserName,
		TestPassWord,
		TestSysUserName,
		TestSysPassWord,
		config.NewDefaultClientConfig(),
	)
	if err != nil {
		return nil, err
	}
	tb, err := table.NewObTable(
		mock_route.MockTestServerAddr.Ip(),
		mock_route.MockTestServerAddr.SvrPort(),
		mock_route.MockTestTenantName,
		obCli.userName,
		obCli.password,
		obCli.database,
		obCli.config.ConnPoolMaxConnSize,
	)
	if err != nil {
		return nil, err
	}
	obCli.tableRoster.Store(*mock_route.MockTestServerAddr, tb)
	obCli.serverRoster.Reset([]*route.ObServerAddr{mock_route.MockTestServerAddr})
	return obCli, nil
}

func TestObClient_Insert(t *testing.T) {
	// 1. create client
	cli, err := getMockObClient()
	assert.Equal(t, nil, err)
	// 2. mock route
	entry := mock_route.GetMockHashTableEntryV4()
	patch := gomonkey.ApplyFunc(
		route.GetTableEntryFromRemote, func(
			addr *route.ObServerAddr,
			sysUA *route.ObUserAuth,
			key *route.ObTableEntryKey) (*route.ObTableEntry, error) {
			return entry, nil
		},
	)
	defer patch.Reset()
	// 3. add rowkey element
	err = cli.AddRowkey(mock_route.MockTestTableName, []string{"c1"})
	// 4. insert
	rowkey := []table.Column{{"c1", 1}}
	mutateColumns := []table.Column{{"c2", 1}}
	affectRows, err := cli.Insert(
		mock_route.MockTestTableName,
		rowkey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, affectRows)
}

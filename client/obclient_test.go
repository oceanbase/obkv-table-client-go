package client

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/route"
	"github.com/oceanbase/obkv-table-client-go/route/mock_route"
	"github.com/oceanbase/obkv-table-client-go/table"
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
	cfg := config.NewDefaultClientConfig()
	obCli, err := newObClient(
		TestConfigUrl,
		TestFullUserName,
		TestPassWord,
		TestSysUserName,
		TestSysPassWord,
		cfg,
	)
	if err != nil {
		return nil, err
	}
	tb := NewObTable(
		mock_route.MockTestServerAddr.Ip(),
		mock_route.MockTestServerAddr.SvrPort(),
		mock_route.MockTestTenantName,
		obCli.userName,
		obCli.password,
		obCli.database,
	)
	err = tb.init(cfg.ConnPoolMaxConnSize, cfg.RpcConnectTimeOut)
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
	// 3. add rowKey element
	err = cli.AddRowKey(mock_route.MockTestTableName, []string{"c1"})
	assert.Equal(t, nil, err)
	// 4. insert
	rowKey := []*table.Column{table.NewColumn("c1", 1)}
	mutateColumns := []*table.Column{table.NewColumn("c2", 1)}
	affectRows, err := cli.Insert(
		context.TODO(),
		mock_route.MockTestTableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(0), affectRows)
}

func TestObClient_Get(t *testing.T) {
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
	// 3. add rowKey element
	err = cli.AddRowKey(mock_route.MockTestTableName, []string{"c1"})
	assert.Equal(t, nil, err)
	// 4. get
	rowKey := []*table.Column{table.NewColumn("c1", 1)}
	selectColumns := []string{"c1", "c2"}
	res, err := cli.Get(
		context.TODO(),
		mock_route.MockTestTableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(res))
}

func TestObClientInsertConcurrent(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
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
	// 3. add rowKey element
	err = cli.AddRowKey(mock_route.MockTestTableName, []string{"c1"})
	assert.Equal(t, nil, err)
	// 4. test
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rowKey := []*table.Column{table.NewColumn("c1", 1)}
			mutateColumns := []*table.Column{table.NewColumn("c2", 1)}
			_, err := cli.Insert(
				context.TODO(),
				mock_route.MockTestTableName,
				rowKey,
				mutateColumns,
			)
			assert.Equal(t, nil, err)
		}(i)
	}
}

func TestObTableParam_ToString(t *testing.T) {
	param := ObTableParam{}
	assert.Equal(t, param.String(), "ObTableParam{table:nil, tableId:0, partitionId:0}")
	param = ObTableParam{nil, 500023, 500012}
	assert.Equal(t, param.String(), "ObTableParam{table:nil, tableId:500023, partitionId:500012}")
}

func TestInsert(t *testing.T) {
	// CREATE TABLE test(c1 bigint(20), c2 bigint(20), primary key (c1)) PARTITION BY hash(c1) partitions 2;
	const (
		configUrl    = "..."
		fullUserName = "..."
		passWord     = ""
		sysUserName  = "root"
		sysPassWord  = "..."
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cfg.OperationTimeOut = time.Duration(1000000) * time.Millisecond
	cli, err := NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	assert.Equal(t, nil, err)

	err = cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

func TestGet(t *testing.T) {
	const (
		configUrl    = "..."
		fullUserName = "..."
		passWord     = ""
		sysUserName  = "root"
		sysPassWord  = ""
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cli, err := NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	assert.Equal(t, nil, err)

	err = cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	selectColumns := []string{"c1", "c2"}
	m, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c1"])
	assert.EqualValues(t, 1, m["c2"])
}

func TestUpdate(t *testing.T) {
	// CREATE TABLE test(c1 bigint(20), c2 bigint(20), primary key (c1)) PARTITION BY hash(c1) partitions 2;
	const (
		configUrl    = "..."
		fullUserName = "..."
		passWord     = ""
		sysUserName  = "root"
		sysPassWord  = "..."
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cfg.OperationTimeOut = time.Duration(1000000) * time.Millisecond
	cli, err := NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	assert.Equal(t, nil, err)

	err = cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	affectRows, err := cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.EqualValues(t, 1, affectRows)
	selectColumns := []string{"c1", "c2"}
	m, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c2"])
	assert.EqualValues(t, 2, m["c2"])
}

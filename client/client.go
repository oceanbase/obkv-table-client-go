package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// NewClient create a client.
// configUrl is used to obtain rslist by using the http service.
// fullUserName format id "userName@tenantName#clusterName".
// passWord is the password of the fullUserName user.
// sysUserName and sysPassWord is used to access the routing table. Therefore, user a must belong to the system tenant.
// cliConfig is the configuration of the client, you can configure based on your business.
func NewClient(
	configUrl string,
	fullUserName string,
	passWord string,
	sysUserName string,
	sysPassWord string,
	cliConfig *config.ClientConfig) (Client, error) {
	// 1. Check args
	if configUrl == "" {
		return nil, errors.New("config url is empty")
	}
	if fullUserName == "" || sysUserName == "" {
		return nil, errors.New("full user name is null")
	}
	if sysUserName == "" {
		return nil, errors.New("system user name is null")
	}
	if cliConfig == nil {
		return nil, errors.New("client config is nil")
	}

	// 2. New client and init
	cli, err := newObClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cliConfig)
	if err != nil {
		return nil, errors.WithMessagef(err, "new ob client, configUrl:%s, fullUserName:%s", configUrl, fullUserName)
	}
	err = cli.init()
	if err != nil {
		return nil, errors.WithMessagef(err, "init client, client:%s", cli.String())
	}

	return cli, nil
}

type Client interface {
	// AddRowKey add the primary key of the current table to the client.
	AddRowKey(tableName string, rowKey []string) error
	// Insert a record by rowkey.
	Insert(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	// Update a record by rowkey.
	Update(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	// InsertOrUpdate insert or update a record by rowkey.
	// insert if the primary key does not exist, update if it does.
	InsertOrUpdate(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	// Delete a record by rowkey.
	Delete(ctx context.Context, tableName string, rowKey []*table.Column, opts ...ObkvOption) (int64, error)
	// Get a record by rowkey.
	Get(ctx context.Context, tableName string, rowKey []*table.Column, getColumns []string, opts ...ObkvOption) (map[string]interface{}, error)
	// NewBatchExecutor create a batch executor
	NewBatchExecutor(tableName string) BatchExecutor
}

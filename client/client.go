package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
)

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
	AddRowKey(tableName string, rowKey []string) error
	Insert(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	Update(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	InsertOrUpdate(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	Delete(ctx context.Context, tableName string, rowKey []*table.Column, opts ...ObkvOption) (int64, error)
	Get(ctx context.Context, tableName string, rowKey []*table.Column, getColumns []string, opts ...ObkvOption) (map[string]interface{}, error)
	NewBatchExecutor(tableName string) BatchExecutor
}

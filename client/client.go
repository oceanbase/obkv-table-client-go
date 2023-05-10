package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/log"
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
		log.Warn("config url is empty")
		return nil, errors.New("config url is empty")
	}
	if fullUserName == "" || sysUserName == "" {
		log.Warn("user name is empty",
			log.String("fullUserName", fullUserName),
			log.String("sysUserName", sysUserName))
		return nil, errors.New("user name is null")
	}
	if cliConfig == nil {
		log.Warn("client config is nil")
		return nil, errors.New("client config is nil")
	}

	// 2. New client and init
	cli, err := newObClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cliConfig)
	if err != nil {
		log.Warn("failed to new obClient",
			log.String("configUrl", configUrl),
			log.String("fullUserName", fullUserName))
		return nil, err
	}
	err = cli.init()
	if err != nil {
		log.Warn("failed to init client", log.String("client", cli.String()))
		return nil, err
	}

	return cli, nil
}

type Client interface {
	AddRowKey(tableName string, rowKey []string) error
	Insert(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	InsertOrUpdate(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...ObkvOption) (int64, error)
	Delete(ctx context.Context, tableName string, rowKey []*table.Column, opts ...ObkvOption) (int64, error)
	Get(ctx context.Context, tableName string, rowKey []*table.Column, getColumns []string, opts ...ObkvOption) (map[string]interface{}, error)
	NewBatchExecutor(tableName string) BatchExecutor
}

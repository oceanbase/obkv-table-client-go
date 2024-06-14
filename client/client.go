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
	"context"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/log"
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
	// init log
	err := initLogProcess(cliConfig)
	if err != nil {
		return nil, err
	}
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
		return nil, err
	}

	return cli, nil
}

// NewOdpClient create a odp client .
// fullUserName format id "userName@tenantName#clusterName".
// passWord is the password of the fullUserName user.
// odpIP is the ip of the odp server.
// odpRpcPort is the rpc port of the odp server.
// database is the database name which you want to access.
// cliConfig is the configuration of the client, you can configure based on your business.
func NewOdpClient(
	fullUserName string,
	passWord string,
	odpIP string,
	odpRpcPort int,
	database string,
	cliConfig *config.ClientConfig) (Client, error) {
	// init log
	err := initLogProcess(cliConfig)
	if err != nil {
		return nil, err
	}

	// 1. Check args
	if fullUserName == "" {
		return nil, errors.New("full user name is null")
	}
	if odpIP == "" {
		return nil, errors.New("Odp IP is null")
	}
	if cliConfig == nil {
		return nil, errors.New("client config is nil")
	}

	// 2. New odp client and init
	cli, err := newOdpClient(fullUserName, passWord, odpIP, odpRpcPort, database, cliConfig)
	if err != nil {
		return nil, errors.WithMessagef(err, "new odp client, odpIP:%s, fullUserName:%s", odpIP, fullUserName)
	}
	err = cli.initOdp()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

// NewClientWithTomlConfig create a client with toml config.
func NewClientWithTomlConfig(configFilePath string) (Client, error) {
	clientConfig, err := config.GetClientConfigurationFromTOML(configFilePath)
	if err != nil {
		return nil, errors.WithMessagef(err, "get client config from toml, configFilePath:%s", configFilePath)
	}
	// init log
	err = initLogProcess(clientConfig.GetClientConfig())
	if err != nil {
		return nil, err
	}
	switch clientConfig.Mode {
	case "direct":
		return NewClient(
			clientConfig.DirectClientConfig.ConfigUrl,
			clientConfig.DirectClientConfig.FullUserName,
			clientConfig.DirectClientConfig.Password,
			clientConfig.DirectClientConfig.SysUserName,
			clientConfig.DirectClientConfig.SysPassword,
			clientConfig.GetClientConfig(),
		)
	case "proxy":
		return NewOdpClient(
			clientConfig.OdpClientConfig.FullUserName,
			clientConfig.OdpClientConfig.Password,
			clientConfig.OdpClientConfig.OdpIp,
			clientConfig.OdpClientConfig.OdpRpcPort,
			clientConfig.OdpClientConfig.Database,
			clientConfig.GetClientConfig(),
		)
	default:
		return nil, errors.Errorf("invalid mode:%s", clientConfig.Mode)
	}
}

func initLogProcess(c *config.ClientConfig) error {
	var logConfig log.LogConfig
	if c != nil {
		logConfig.LogFileName = c.LogFileName
		logConfig.MaxAgeFileRem = c.MaxAgeFileRem
		logConfig.MaxBackupFileSize = c.MaxBackupFileSize
		logConfig.SingleFileMaxSize = c.SingleFileMaxSize
		logConfig.Compress = c.Compress
		logConfig.SlowQueryThreshold = c.SlowQueryThreshold
	} else {
		return errors.New("client config is null")
	}
	return log.InitLoggerWithConfig(logConfig)
}

type Client interface {
	// Insert a record by rowKey.
	Insert(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (int64, error)
	// Update a record by rowKey.
	Update(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (int64, error)
	// InsertOrUpdate insert or update a record by rowKey.
	// insert if the primary key does not exist, update if it does.
	InsertOrUpdate(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (int64, error)
	// InsertOrUpdateWithResult insert or update a record by rowKey.
	// insert if the primary key does not exist, update if it does.
	// IsInsertOrUpdateDoInsert() in SingleResult tells whether the insert operation has been performed.
	InsertOrUpdateWithResult(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (SingleResult, error)
	// Replace a record by rowKey.
	Replace(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (int64, error)
	// Increment a record by rowKey.
	Increment(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (SingleResult, error)
	// Append a record by rowKey.
	Append(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (SingleResult, error)
	// Delete a record by rowKey.
	Delete(ctx context.Context, tableName string, rowKey []*table.Column, opts ...option.ObOperationOption) (int64, error)
	// Get a record by rowKey.
	Get(ctx context.Context, tableName string, rowKey []*table.Column, getColumns []string, opts ...option.ObOperationOption) (SingleResult, error)
	// Query records by rangePairs.
	Query(ctx context.Context, tableName string, rangePairs []*table.RangePair, opts ...option.ObQueryOption) (QueryResultIterator, error)
	// NewBatchExecutor create a batch executor.
	NewBatchExecutor(tableName string, opts ...option.ObBatchOption) BatchExecutor
	// NewAggExecutor create an aggregate executor.
	NewAggExecutor(tableName string, rangePairs []*table.RangePair, opts ...option.ObQueryOption) AggExecutor
	// Close closes the client.
	// close will disconnect all connections and exit all goroutines.
	Close()
	Redis(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, opts ...option.ObOperationOption) (SingleResult, error)
}

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
	// Insert a record by rowKey.
	Insert(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, operationOpts ...OperationOption) (int64, error)
	// Update a record by rowKey.
	Update(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, operationOpts ...OperationOption) (int64, error)
	// InsertOrUpdate insert or update a record by rowKey.
	// insert if the primary key does not exist, update if it does.
	InsertOrUpdate(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, operationOpts ...OperationOption) (int64, error)
	// Replace a record by rowKey.
	Replace(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, operationOpts ...OperationOption) (int64, error)
	// Increment a record by rowKey.
	Increment(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, operationOpts ...OperationOption) (SingleResult, error)
	// Append a record by rowKey.
	Append(ctx context.Context, tableName string, rowKey []*table.Column, mutateColumns []*table.Column, operationOpts ...OperationOption) (SingleResult, error)
	// Delete a record by rowKey.
	Delete(ctx context.Context, tableName string, rowKey []*table.Column, operationOpts ...OperationOption) (int64, error)
	// Get a record by rowKey.
	Get(ctx context.Context, tableName string, rowKey []*table.Column, getColumns []string, operationOpts ...OperationOption) (SingleResult, error)
	// Query records by rangePairs.
	Query(ctx context.Context, tableName string, rangePairs []*table.RangePair, opts ...ObkvQueryOption) (QueryResultIterator, error)
	// NewBatchExecutor create a batch executor.
	NewBatchExecutor(tableName string) BatchExecutor
	// Close closes the client.
	// close will disconnect all connections and exit all goroutines.
	Close()
}

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

package obkvrpc

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
)

type RpcClientOption struct {
	ip                  string
	port                int
	connPoolMaxConnSize int
	connectTimeout      time.Duration
	loginTimeout        time.Duration

	tenantName   string
	databaseName string
	userName     string
	password     string
}

func NewRpcClientOption(ip string, port int, connPoolMaxConnSize int, connectTimeout time.Duration, loginTimeout time.Duration,
	tenantName string, databaseName string, userName string, password string) *RpcClientOption {
	return &RpcClientOption{
		ip:                  ip,
		port:                port,
		connPoolMaxConnSize: connPoolMaxConnSize,
		connectTimeout:      connectTimeout,
		loginTimeout:        loginTimeout,
		tenantName:          tenantName,
		databaseName:        databaseName,
		userName:            userName,
		password:            password,
	}
}

func (o *RpcClientOption) String() string {
	return "RpcClientOption{" +
		"ip:" + o.ip + ", " +
		"port:" + strconv.Itoa(o.port) + ", " +
		"connPoolMaxConnSize:" + strconv.Itoa(o.connPoolMaxConnSize) + ", " +
		"connectTimeout:" + o.connectTimeout.String() + ", " +
		"loginTimeout:" + o.loginTimeout.String() + ", " +
		"tenantName:" + o.tenantName + ", " +
		"databaseName:" + o.databaseName + ", " +
		"userName:" + o.userName + ", " +
		"password:" + o.password +
		"}"
}

type RpcClient struct {
	option *RpcClientOption

	connectionPool *ConnectionPool
}

func NewRpcClient(rpcClientOption *RpcClientOption) (*RpcClient, error) {
	client := &RpcClient{option: rpcClientOption}

	poolOption := NewPoolOption(client.option.ip, client.option.port, client.option.connPoolMaxConnSize, client.option.connectTimeout, client.option.loginTimeout,
		client.option.tenantName, client.option.databaseName, client.option.userName, client.option.password)
	connectionPool, err := NewConnectionPool(poolOption)
	if err != nil {
		return nil, errors.WithMessage(err, "create connection pool")
	}

	client.connectionPool = connectionPool
	return client, nil
}

func (c *RpcClient) Execute(ctx context.Context, request protocol.ObPayload, response protocol.ObPayload) error {
	connection, err := c.connectionPool.GetConnection(ctx)
	if err != nil {
		return errors.WithMessage(err, "connection pool get connection")
	}

	err = connection.Execute(ctx, request, response)
	if err != nil {
		return errors.WithMessage(err, "connection execute")
	}

	return nil
}

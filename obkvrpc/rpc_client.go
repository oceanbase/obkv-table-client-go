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

	maxConnectionAge     time.Duration
	enableSLBLoadBalance bool
}

func (o *RpcClientOption) ConnPoolMaxConnSize() int {
	return o.connPoolMaxConnSize
}

func (o *RpcClientOption) ConnectTimeout() time.Duration {
	return o.connectTimeout
}

func (o *RpcClientOption) LoginTimeout() time.Duration {
	return o.loginTimeout
}

func NewRpcClientOption(ip string, port int, connPoolMaxConnSize int, connectTimeout time.Duration, loginTimeout time.Duration,
	tenantName string, databaseName string, userName string, password string, maxConnectionAge time.Duration, enableSLBLoadBalance bool) *RpcClientOption {
	return &RpcClientOption{
		ip:                   ip,
		port:                 port,
		connPoolMaxConnSize:  connPoolMaxConnSize,
		connectTimeout:       connectTimeout,
		loginTimeout:         loginTimeout,
		tenantName:           tenantName,
		databaseName:         databaseName,
		userName:             userName,
		password:             password,
		maxConnectionAge:     maxConnectionAge,
		enableSLBLoadBalance: enableSLBLoadBalance,
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
		"password:" + o.password + ", " +
		"maxConnectionAge:" + o.maxConnectionAge.String() + ", " +
		"enableSLBLoadBalance:" + strconv.FormatBool(o.enableSLBLoadBalance) +
		"}"
}

type RpcClient struct {
	option *RpcClientOption

	connectionPool *ConnectionPool
}

func (c *RpcClient) IsDisconnected() bool {
	return c.connectionPool.IsDisconnected()
}

func (c *RpcClient) Option() *RpcClientOption {
	return c.option
}

func NewRpcClient(rpcClientOption *RpcClientOption) (*RpcClient, error) {
	client := &RpcClient{option: rpcClientOption}

	poolOption := NewPoolOption(client.option.ip, client.option.port, client.option.connPoolMaxConnSize, client.option.connectTimeout, client.option.loginTimeout,
		client.option.tenantName, client.option.databaseName, client.option.userName, client.option.password, client.option.maxConnectionAge, client.option.enableSLBLoadBalance)
	connectionPool, err := NewConnectionPool(poolOption)
	if err != nil {
		return nil, errors.WithMessage(err, "create connection pool")
	}

	client.connectionPool = connectionPool
	return client, nil
}

func (c *RpcClient) Execute(
	ctx context.Context,
	request protocol.ObPayload,
	response protocol.ObPayload) (*protocol.ObTableMoveResponse, error) {

	var (
		connection *Connection
		index      int
		err        error
	)

	connection, index = c.connectionPool.GetConnection()
	if connection == nil {
		// maybe this connection has been disconnected，we need to reconnect it.
		connection, err = c.connectionPool.RecreateConnection(ctx, index)
		if err != nil {
			return nil, errors.WithMessage(err, "recreate connection")
		}
	}

	return connection.Execute(ctx, request, response)
}

func (c *RpcClient) Close() {
	c.connectionPool.Close()
}

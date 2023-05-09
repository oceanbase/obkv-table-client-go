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

	tenantName   string
	databaseName string
	userName     string
	password     string
}

func NewRpcClientOption(ip string, port int, connPoolMaxConnSize int, connectTimeout time.Duration, tenantName string, databaseName string, userName string, password string) *RpcClientOption {
	return &RpcClientOption{
		ip:                  ip,
		port:                port,
		connPoolMaxConnSize: connPoolMaxConnSize,
		connectTimeout:      connectTimeout,
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

	poolOption := NewPoolOption(client.option.ip, client.option.port, client.option.connPoolMaxConnSize, client.option.connectTimeout, client.option.tenantName, client.option.databaseName, client.option.userName, client.option.password)
	connectionPool, err := NewConnectionPool(poolOption)
	if err != nil {
		return nil, errors.Wrap(err, "create rpc client")
	}

	client.connectionPool = connectionPool
	return client, nil
}

func (c *RpcClient) Execute(ctx context.Context, request protocol.Payload, response protocol.Payload) error {
	connection, err := c.connectionPool.GetConnection()
	if err != nil {
		return errors.Wrap(err, "rpc client execute failed")
	}

	err = connection.Execute(ctx, request, response)
	if err != nil {
		return errors.Wrap(err, "rpc client execute failed")
	}

	return nil
}

package obkvrpc

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
)

type RpcClientOption struct {
	ip             string
	port           int
	connectTimeout time.Duration

	tenantName   string
	databaseName string
	userName     string
	password     string
}

func (o *RpcClientOption) String() string {
	return "RpcClientOption{" +
		"ip:" + o.ip + ", " +
		"port:" + strconv.Itoa(o.port) + ", " +
		"connectTimeout:" + o.connectTimeout.String() + ", " +
		"tenantName:" + o.tenantName + ", " +
		"databaseName:" + o.databaseName + ", " +
		"userName:" + o.userName + ", " +
		"password:" + o.password +
		"}"
}

func NewRpcClientOption(
	ip string,
	port int,
	connectTimeout time.Duration,
	tenantName string,
	databaseName string,
	userName string,
	password string) *RpcClientOption {
	return &RpcClientOption{
		ip:             ip,
		port:           port,
		connectTimeout: connectTimeout,
		tenantName:     tenantName,
		databaseName:   databaseName,
		userName:       userName,
		password:       password,
	}
}

type RpcClient struct {
	connection *Connection

	rpcClientOption *RpcClientOption
}

func NewRpcClient(rpcClientOption *RpcClientOption) (*RpcClient, error) {
	client := &RpcClient{rpcClientOption: rpcClientOption}
	err := client.Connect()
	if err != nil {
		return nil, errors.Wrap(err, "create rpc client")
	}
	return client, nil
}

func (c *RpcClient) Connect() error {
	option := NewOption(c.rpcClientOption.ip, c.rpcClientOption.port, c.rpcClientOption.connectTimeout,
		c.rpcClientOption.tenantName, c.rpcClientOption.databaseName, c.rpcClientOption.userName, c.rpcClientOption.password)

	connection := NewConnection(option, uuid.New())

	err := connection.Connect()
	if err != nil {
		return errors.Wrap(err, "rpc client connect")
	}

	err = connection.Login()
	if err != nil {
		return errors.Wrap(err, "rpc client login")
	}

	c.connection = connection
	return nil
}

func (c *RpcClient) Execute(ctx context.Context, request protocol.Payload, response protocol.Payload) error {
	if c.connection.active.Load() == true {
		err := c.connection.Execute(ctx, request, response)
		if err != nil {
			return errors.Wrap(err, "rpc client execute")
		}
	} else {
		err := c.Connect()
		if err != nil {
			return errors.Wrap(err, "rpc client reconnect")
		}
		err = c.connection.Execute(ctx, request, response)
		if err != nil {
			return errors.Wrap(err, "rpc client execute")
		}
	}
	return nil
}

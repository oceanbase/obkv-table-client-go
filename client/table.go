package client

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/obkvrpc"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"strconv"
	"sync"
)

type ObTable struct {
	ip         string
	port       int
	tenantName string
	userName   string
	password   string
	database   string
	rpcClient  *obkvrpc.RpcClient

	isClosed bool
	mutex    sync.Mutex
}

func NewObTable(
	ip string,
	port int,
	tenantName string,
	userName string,
	password string,
	database string) *ObTable {
	return &ObTable{
		ip:         ip,
		port:       port,
		tenantName: tenantName,
		userName:   userName,
		password:   password,
		database:   database,
		isClosed:   false,
	}
}

func (t *ObTable) init(config *config.ClientConfig) error {
	opt := obkvrpc.NewRpcClientOption(
		t.ip,
		t.port,
		config.RpcConnectTimeOut,
		t.tenantName,
		t.database,
		t.userName,
		t.password,
	)
	cli, err := obkvrpc.NewRpcClient(opt)
	if err != nil {
		log.Warn("failed to new rpc client", log.String("rpc client opt", opt.String()))
		return err
	}
	t.rpcClient = cli
	return nil
}

func (t *ObTable) execute(request protocol.Payload, result protocol.Payload) error {
	return t.rpcClient.Execute(context.TODO(), request, result)
}

func (t *ObTable) close() {
	if !t.isClosed {
		t.mutex.Lock()
		if !t.isClosed { // double check after lock
			//if t.rpcClient != nil {
			// todo: t.rpcClient.Close()
			//}
			t.isClosed = true
		}
		t.mutex.Unlock()
	}
}

func (t *ObTable) String() string {
	return "ObTable{" +
		"ip:" + t.ip + ", " +
		"port:" + strconv.Itoa(t.port) + ", " +
		"tenantName:" + t.tenantName + ", " +
		"userName:" + t.userName + ", " +
		"password:" + t.password + ", " +
		"database:" + t.database + ", " +
		"isClosed:" + strconv.FormatBool(t.isClosed) +
		"}"
}

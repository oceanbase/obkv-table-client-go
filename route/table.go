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

package route

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/oceanbase/obkv-table-client-go/obkvrpc"
	"github.com/oceanbase/obkv-table-client-go/protocol"
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

	maxConnectionAge     time.Duration
	enableSLBLoadBalance bool
}

func (t *ObTable) Ip() string {
	return t.ip
}

func (t *ObTable) Port() int {
	return t.port
}

func NewObTable(
	ip string,
	port int,
	tenantName string,
	userName string,
	password string,
	database string) *ObTable {
	return &ObTable{
		ip:                   ip,
		port:                 port,
		tenantName:           tenantName,
		userName:             userName,
		password:             password,
		database:             database,
		isClosed:             false,
		maxConnectionAge:     time.Duration(0),
		enableSLBLoadBalance: false,
	}
}

func (t *ObTable) Init(connPoolSize int, connectTimeout time.Duration, loginTimeout time.Duration) error {
	opt := obkvrpc.NewRpcClientOption(
		t.ip,
		t.port,
		connPoolSize,
		connectTimeout,
		loginTimeout,
		t.tenantName,
		t.database,
		t.userName,
		t.password,
		t.maxConnectionAge,
		t.enableSLBLoadBalance,
	)
	cli, err := obkvrpc.NewRpcClient(opt)
	if err != nil {
		return err
	}
	t.rpcClient = cli
	return nil
}

func (t *ObTable) IsDisconnected() bool {
	return t.rpcClient.IsDisconnected()
}

func (t *ObTable) SetMaxConnectionAge(duration time.Duration) {
	t.maxConnectionAge = duration
}

func (t *ObTable) SetEnableSLBLoadBalance(b bool) {
	t.enableSLBLoadBalance = b
}

func (t *ObTable) Execute(
	ctx context.Context,
	request protocol.ObPayload,
	result protocol.ObPayload) (*protocol.ObTableMoveResponse, error) {
	return t.rpcClient.Execute(ctx, request, result)
}

func (t *ObTable) Close() {
	if !t.isClosed {
		t.mutex.Lock()
		if !t.isClosed { // double check after lock
			if t.rpcClient != nil {
				t.rpcClient.Close()
			}
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

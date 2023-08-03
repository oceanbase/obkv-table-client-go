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
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
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

func (t *ObTable) init(connPoolSize int, connectTimeout time.Duration, loginTimeout time.Duration) error {
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
	)
	cli, err := obkvrpc.NewRpcClient(opt)
	if err != nil {
		return errors.WithMessagef(err, "new rpc client, opt:%s", opt.String())
	}
	t.rpcClient = cli
	return nil
}

func (t *ObTable) retry(
	ctx context.Context,
	request protocol.ObPayload,
	result protocol.ObPayload,
	moveResult *protocol.ObTableMoveResponse) error {

	// 1. create new table by move result
	newTable := NewObTable(
		moveResult.ReplicaInfo().Server().IpToString(),
		int(moveResult.ReplicaInfo().Server().Port()),
		t.tenantName,
		t.userName,
		t.password,
		t.database,
	)

	err := newTable.init(
		t.rpcClient.Option().ConnPoolMaxConnSize(),
		t.rpcClient.Option().ConnectTimeout(),
		t.rpcClient.Option().LoginTimeout())
	if err != nil {
		return errors.WithMessagef(err, "new table init fail, new table:%s", newTable.String())
	}

	// 2. execute until timeout
	// We'll set the default timeout to 500 milliseconds if the user doesn't set a timeout
	err = newTable.execute(ctx, request, result)
	if err != nil {
		log.Info("retry execute failed by first times")
	}
	for i := 0; err != nil; i++ { // retry until context timeout
		err = newTable.execute(ctx, request, result)
		if err != nil {
			log.Info("retry execute failed",
				log.Int("times", i), log.String("ip", newTable.ip), log.String("err", err.Error()))
		}
	}

	return err
}

func (t *ObTable) execute(
	ctx context.Context,
	request protocol.ObPayload,
	result protocol.ObPayload) error {

	moveResult := protocol.NewObTableMoveResponse()
	err := t.rpcClient.Execute(ctx, request, result, moveResult)
	if err != nil {
		if moveResult.Valid() { // move response, retry
			err = t.retry(ctx, request, result, moveResult)
			if err != nil {
				return errors.WithMessagef(err, "retry execute request, move result%s", moveResult.String())
			}
		} else {
			return errors.WithMessagef(err, "obtable execute")
		}
	}
	return nil
}

func (t *ObTable) close() {
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

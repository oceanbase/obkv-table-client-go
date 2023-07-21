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
	"github.com/oceanbase/obkv-table-client-go/table"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"

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

func (t *ObTable) execute(ctx context.Context, request protocol.ObPayload, result protocol.ObPayload) error {
	return t.rpcClient.Execute(ctx, request, result)
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

// transferQueryRange sets the query range into tableQuery.
func transferQueryRange(rangePair []*table.RangePair) ([]*protocol.ObNewRange, error) {
	queryRanges := make([]*protocol.ObNewRange, 0, len(rangePair))
	for _, rangePair := range rangePair {
		if len(rangePair.Start()) != len(rangePair.End()) {
			return nil, errors.New("startRange and endRange key length is not equal")
		}
		startObjs := make([]*protocol.ObObject, 0, len(rangePair.Start()))
		endObjs := make([]*protocol.ObObject, 0, len(rangePair.End()))
		for i := 0; i < len(rangePair.Start()); i++ {
			// append start obj
			objMeta, err := protocol.DefaultObjMeta(rangePair.Start()[i].Value())
			if err != nil {
				return nil, errors.WithMessage(err, "create obj meta by Range key")
			}
			startObjs = append(startObjs, protocol.NewObObjectWithParams(objMeta, rangePair.Start()[i].Value()))

			// append end obj
			objMeta, err = protocol.DefaultObjMeta(rangePair.End()[i].Value())
			if err != nil {
				return nil, errors.WithMessage(err, "create obj meta by Range key")
			}
			endObjs = append(endObjs, protocol.NewObObjectWithParams(objMeta, rangePair.End()[i].Value()))
		}
		borderFlag := protocol.NewObBorderFlag()
		if rangePair.IncludeStart() {
			borderFlag.SetInclusiveStart()
		}
		if rangePair.IncludeEnd() {
			borderFlag.SetInclusiveEnd()
		}
		queryRanges = append(queryRanges, protocol.NewObNewRangeWithParams(startObjs, endObjs, borderFlag))
	}
	return queryRanges, nil
}

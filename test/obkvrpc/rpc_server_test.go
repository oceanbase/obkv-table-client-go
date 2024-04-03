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
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/oceanbase/obkv-table-client-go/obkvrpc"
	"github.com/stretchr/testify/assert"
)

var (
	MaxMsgNum       = 3
	MaxClientNum    = 10
	RoutinePoolSize = 10000
	ExpiredDuration = time.Second * 60
)

type MockCodec struct {
	Conn    net.Conn
	RunTime int
	Wg      *sync.WaitGroup
	IDMap   sync.Map
	Test    *testing.T
}

func (mcs *MockCodec) ReadRequest(req *obkvrpc.Request) error {
	if mcs.RunTime++; mcs.RunTime > MaxMsgNum {
		// fmt.Println("ReadRequest return err", mcs.RunTime)
		return io.EOF
	}
	req.Method = "testMethod"
	req.ID = uuid.New().String()

	arg1 := []byte{'a', 'r', 'g'}
	arg2 := []byte{'1', '2', '3'}
	req.Args = append(req.Args, arg1)
	req.Args = append(req.Args, arg2)
	mcs.IDMap.Store(req.ID, true)
	// fmt.Println("ReadRequest: ", req.ID, obkvrpc.GetGID())
	return nil
}

func (mcs *MockCodec) WriteResponse(resp *obkvrpc.Response) error {
	_, ok := mcs.IDMap.Load(resp.ID)
	assert.Equal(mcs.Test, true, ok)
	mcs.IDMap.Delete(resp.ID)
	// fmt.Println("WriteResponse: ", resp.ID, obkvrpc.GetGID())
	return nil
}

func (mcs *MockCodec) Call(req *obkvrpc.Request, resp *obkvrpc.Response) error {
	resp.ID = req.ID
	resp.RspContent = []byte("content")
	// fmt.Println("Call: ", req.ID, obkvrpc.GetGID())
	return nil
}

func (mcs *MockCodec) Close() {
	// fmt.Println("MockCodec closed")
	mcs.Conn.Close()
	mcs.Wg.Done()
	mcs.IDMap.Range(func(key any, value any) bool {
		mcs.Test.Error("IDMap is not empty when MockCodec closing", key)
		return false
	})
}

func MockTCPServer(t *testing.T, lis net.Listener, doneChan chan struct{}) {
	rpcSrv, err := obkvrpc.NewServer(RoutinePoolSize, ExpiredDuration, new(chan struct{}))
	assert.Equal(t, nil, err)
	wg := new(sync.WaitGroup)
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("listener closed")
			break
		}
		codec := MockCodec{Conn: conn, RunTime: 0, Wg: wg, Test: t}
		wg.Add(1)
		go rpcSrv.ServeCodec(&codec)
	}
	wg.Wait()
	close(doneChan)
	rpcSrv.Close()
}

func Dial(t *testing.T) {
	_, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, nil, err)
}

func ServerAndClients(t *testing.T, maxMsgNum int, maxClientNum int) {
	// fmt.Println("start to test tcp server")
	MaxMsgNum = maxMsgNum
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	assert.Equal(t, nil, err)
	doneChan := make(chan struct{})
	go MockTCPServer(t, lis, doneChan)
	for runTime := 0; runTime < maxClientNum; runTime++ {
		Dial(t)
	}
	lis.Close()
	_, ok := <-doneChan
	fmt.Println("finish test tcp server, isOpen = ", ok)
}

func TestRPCServer(t *testing.T) {
	// defer goleak.VerifyNone(t)
	// single message single client
	ServerAndClients(t, 1, 1)
	time.Sleep(time.Millisecond * 500)

	// single message multi client
	ServerAndClients(t, 1, 10000)
	time.Sleep(time.Millisecond * 500)

	// multi message single client
	ServerAndClients(t, 10000, 1)
	time.Sleep(time.Millisecond * 500)

	// multi message multi client
	ServerAndClients(t, 39000, 256)
	time.Sleep(time.Millisecond)
}

func TestCopy(t *testing.T) {
	b := []byte("123")
	data := binary.BigEndian.Uint16(b)
	fmt.Println(data)
	fmt.Println(string(b))
	fmt.Println(strconv.Atoi(string(b)))
}

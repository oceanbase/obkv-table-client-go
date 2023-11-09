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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
	route "github.com/oceanbase/obkv-table-client-go/test/route/util"
	"github.com/oceanbase/obkv-table-client-go/util"
)

const (
	testInt32RouteTableName       = "test_int32_route"
	testInt32RouteCreateStatement = "create table if not exists `test_int32_route`(`c1` int(12) not null,`c2` int(12) default null,primary key (`c1`)) partition by hash(c1) partitions 2;"
)

const (
	passRoutingTest = true
	sshHost         = ""
	sshUser         = ""
	sshPassword     = ""
	killPid         = -1 // fill it carefully
)

func singleGet() error {
	tableName := testInt32RouteTableName
	cli := test.CreateClient()
	defer cli.Close()

	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	_, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func TestRoute_SwitchLeaderNormalClient(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)
	cli := test.CreateClientWithoutRouting()
	defer cli.Close()

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
				_, err := cli.Get(
					ctx,
					tableName,
					rowKey,
					nil,
				)
				if err != nil {
					println("get fail")
				} else {
					println("get success")
				}
			}
		}()
	}

	// 2. switch leader
	go func() {
		if util.ObVersion() < 4 {
			err = route.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
			assert.Equal(t, nil, err)
		} else {
			err = route.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
			assert.Equal(t, nil, err)
		}
	}()

	time.Sleep(10 * time.Second)

	// 3. get record with new route
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 0, res.Value("c2"))

	assert.Equal(t, nil, singleGet())
}

func TestRoute_SwitchLeaderMoveClient(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
				_, err := cli.Get(
					ctx,
					tableName,
					rowKey,
					nil,
				)
				if err != nil {
					println("get fail")
				} else {
					println("get success")
				}
			}
		}()
	}

	// 2. switch leader
	go func() {
		if util.ObVersion() < 4 {
			err = route.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
			assert.Equal(t, nil, err)
		} else {
			err = route.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
			assert.Equal(t, nil, err)
		}
	}()

	time.Sleep(10 * time.Second)

	// 3. get record with new route
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 0, res.Value("c2"))

	assert.Equal(t, nil, singleGet())
}

func killServer() {
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", sshHost+":22", config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	command := fmt.Sprintf("kill -9 %d", killPid)
	_, err = session.CombinedOutput(command)
	if err != nil {
		panic(err)
	}
}

func TestRoute_OfflineNormalClient(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)
	cli := test.CreateClientWithoutRouting()
	defer cli.Close()

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
				_, err := cli.Get(
					ctx,
					tableName,
					rowKey,
					nil,
				)
				if err != nil {
					fmt.Printf("get fail, err:%s\n", err.Error())
				} else {
					println("get success")
				}
			}
		}()
	}

	// 2. kill observer
	killServer()

	time.Sleep(10 * time.Second)

	// 3. get record with new route
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 0, res.Value("c2"))

	time.Sleep(10 * time.Second)
	assert.Equal(t, nil, singleGet())
}

func TestRoute_OfflineMoveClient(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
				_, err := cli.Get(
					ctx,
					tableName,
					rowKey,
					nil,
				)
				if err != nil {
					fmt.Printf("get fail, err:%s\n", err.Error())
				} else {
					println("get success")
				}
			}
		}()
	}

	// 2. kill observer
	killServer()

	time.Sleep(10 * time.Second)

	// 3. get record with new route
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 0, res.Value("c2"))

	time.Sleep(10 * time.Second)
	assert.Equal(t, nil, singleGet())
}

func TestRoute_OfflineCheckRslist(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(5 * time.Second)

	// 2. kill observer
	killServer()

	time.Sleep(6 * time.Minute)

	// 3. get record
	assert.Equal(t, nil, singleGet())
}

func TestRoute_OnlineNormalClient(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)
	cli := test.CreateClientWithoutRouting()
	defer cli.Close()

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
				_, err := cli.Get(
					ctx,
					tableName,
					rowKey,
					nil,
				)
				if err != nil {
					fmt.Printf("get fail, err:%s\n", err.Error())
				} else {
					println("get success")
				}
			}
		}()
	}

	// 2. kill observer
	killServer()

	// 3. time for you to up server
	time.Sleep(120 * time.Second)

	// 4. get record with new route
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 0, res.Value("c2"))

	time.Sleep(10 * time.Second)
	assert.Equal(t, nil, singleGet())
}

func TestRoute_OnlineMoveClient(t *testing.T) {
	if passRoutingTest {
		fmt.Println("Please run Routing tests manually!!!")
		fmt.Println("Change passRoutingTest to false in test/route/route_test.go to run Routing tests.")
		assert.Equal(t, passRoutingTest, false)
		return
	}
	tableName := testInt32RouteTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(5 * time.Second)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
				_, err := cli.Get(
					ctx,
					tableName,
					rowKey,
					nil,
				)
				if err != nil {
					fmt.Printf("get fail, err:%s\n", err.Error())
				} else {
					println("get success")
				}
			}
		}()
	}

	// 2. kill observer
	killServer()

	// 3. time for you to up server
	time.Sleep(120 * time.Second)

	// 4. get record with new route
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 0, res.Value("c2"))

	time.Sleep(10 * time.Second)
	assert.Equal(t, nil, singleGet())
}

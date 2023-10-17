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

	"github.com/oceanbase/obkv-table-client-go/util"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
	reroute "github.com/oceanbase/obkv-table-client-go/test/route/util"
)

const (
	testInt32RerouteTableName       = "test_int32_reroute"
	testInt32RerouteCreateStatement = "create table if not exists `test_int32_reroute`(`c1` int(12) not null,`c2` int(12) default null,primary key (`c1`)) partition by hash(c1) partitions 2;"
)

const (
	passReroutingTest = true
	tenantName        = "sys"
	databaseName      = "test"
	partNum           = 2
)

func TestMoveReplica_singleOp(t *testing.T) {
	if passReroutingTest {
		fmt.Println("Please run Rerouting tests manually!!!")
		fmt.Println("Change passReroutingTest to false in test/route/reroute_test.go to run rerouting tests.")
		assert.Equal(t, passReroutingTest, false)
		return
	}
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)
	test.SetReroutingEnable(true)
	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. get
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := moveCli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, int32(0), res.Value("c1"))
	assert.EqualValues(t, int32(0), res.Value("c2"))
}

func TestMoveReplica_singleOp_insertUp(t *testing.T) {
	if passReroutingTest {
		fmt.Println("Please run Rerouting tests manually!!!")
		fmt.Println("Change passReroutingTest to false in test/route/reroute_test.go to run rerouting tests.")
		assert.Equal(t, passReroutingTest, false)
		return
	}
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. get
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	affectRows, err = moveCli.InsertOrUpdate(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

func TestMoveReplica_batch(t *testing.T) {
	if passReroutingTest {
		fmt.Println("Please run Rerouting tests manually!!!")
		fmt.Println("Change passReroutingTest to false in test/route/reroute_test.go to run rerouting tests.")
		assert.Equal(t, passReroutingTest, false)
		return
	}
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)
	test.SetReroutingEnable(true)
	// 1. insert
	ctx1, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		ctx1,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. multi get
	batchExecutor := moveCli.NewBatchExecutor(
		tableName,
		option.WithBatchSamePropertiesNames(true), // Strongly recommend you to set this option to true if all names of properties are the same in this batch.
	)
	err = batchExecutor.AddGetOp(rowKey, []string{"c1", "c2"})
	assert.Equal(t, nil, err)

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := batchExecutor.Execute(ctx)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.Size())
	assert.EqualValues(t, int32(0), res.GetResults()[0].Value("c1"))
	assert.EqualValues(t, int32(0), res.GetResults()[0].Value("c2"))
}

func TestMoveReplica_batch_insertUp(t *testing.T) {
	if passReroutingTest {
		fmt.Println("Please run Rerouting tests manually!!!")
		fmt.Println("Change passReroutingTest to false in test/route/reroute_test.go to run rerouting tests.")
		assert.Equal(t, passReroutingTest, false)
		return
	}
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)
	test.SetReroutingEnable(true)

	// 1. insert
	ctx1, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		ctx1,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. multi get
	batchExecutor := moveCli.NewBatchExecutor(
		tableName,
		option.WithBatchSamePropertiesNames(true), // Strongly recommend you to set this option to true if all names of properties are the same in this batch.
	)
	insertUpColumns := []*table.Column{table.NewColumn("c2", int32(5))}
	err = batchExecutor.AddInsertOrUpdateOp(rowKey, insertUpColumns)
	assert.Equal(t, nil, err)

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := batchExecutor.Execute(ctx)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.Size())
	assert.EqualValues(t, int64(1), res.GetResults()[0].AffectedRows())
}

func TestMoveReplica_query(t *testing.T) {
	if passReroutingTest {
		fmt.Println("Please run Rerouting tests manually!!!")
		fmt.Println("Change passReroutingTest to false in test/route/reroute_test.go to run rerouting tests.")
		assert.Equal(t, passReroutingTest, false)
		return
	}
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	ctx1, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		ctx1,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. query
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	startRowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := moveCli.Query(
		ctx,
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2"}),
	)
	assert.Equal(t, nil, err)

	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.EqualValues(t, int32(0), res.Value("c1"))
		assert.EqualValues(t, int32(0), res.Value("c2"))
		fmt.Printf("get value\n")
	}
	assert.Equal(t, nil, err)
}

func TestMoveReplica_queryAndMutate(t *testing.T) {
	if passReroutingTest {
		fmt.Println("Please run Rerouting tests manually!!!")
		fmt.Println("Change passReroutingTest to false in test/route/reroute_test.go to run rerouting tests.")
		assert.Equal(t, passReroutingTest, false)
		return
	}
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)
	test.SetReroutingEnable(true)

	// 1. insert
	ctx1, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		ctx1,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. query and mutate
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	updateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err = moveCli.Update(
		ctx,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c2", int32(0))), // where c2 = 0
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

func TestMoveReplica_serverReroutingOff(t *testing.T) {
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)
	test.SetReroutingEnable(false)
	// 1. insert
	ctx1, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := moveCli.Insert(
		ctx1,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. do operator
	// single get
	ctx2, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res2, err2 := moveCli.Get(
		ctx2,
		tableName,
		rowKey,
		nil,
	)
	assert.NotEqual(t, nil, err2)
	assert.Equal(t, nil, res2)
	// multi get
	batchExecutor := moveCli.NewBatchExecutor(
		tableName,
		option.WithBatchSamePropertiesNames(true), // Strongly recommend you to set this option to true if all names of properties are the same in this batch.
	)
	err = batchExecutor.AddGetOp(rowKey, nil)
	assert.Equal(t, nil, err)
	ctx3, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	_, err3 := batchExecutor.Execute(ctx3)
	assert.NotEqual(t, nil, err3)

	// query
	ctx4, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	startRowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err4 := moveCli.Query(
		ctx4,
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2"}),
	)
	assert.Equal(t, nil, err4)
	res4, _ := resSet.Next()
	assert.Equal(t, nil, res4)

	ctx5, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	updateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows5, err5 := moveCli.Update(
		ctx5,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c2", int32(0))), // where c2 = 0
	)
	assert.NotEqual(t, nil, err5)
	assert.EqualValues(t, -1, affectRows5)
}

func TestMoveReplica_clientReroutingOff(t *testing.T) {
	tableName := testInt32RerouteTableName
	defer test.DeleteTable(tableName)
	test.SetReroutingEnable(true)
	//client := moveCli // client move res on
	client := cli // client move res off
	// 1. insert
	ctx1, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := client.Insert(
		ctx1,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. switch leader
	if util.ObVersion() < 4 {
		err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
		assert.Equal(t, nil, err)
	} else {
		err = reroute.SwitchReplicaLeaderRandomly4x(tenantName, databaseName, tableName)
		assert.Equal(t, nil, err)
	}
	time.Sleep(5 * time.Second)

	// 3. do operator
	// single get
	ctx2, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res2, err2 := client.Get(
		ctx2,
		tableName,
		rowKey,
		nil,
	)
	assert.NotEqual(t, nil, err2)
	assert.Equal(t, nil, res2)
	// multi get
	batchExecutor := client.NewBatchExecutor(
		tableName,
		option.WithBatchSamePropertiesNames(true), // Strongly recommend you to set this option to true if all names of properties are the same in this batch.
	)
	err = batchExecutor.AddGetOp(rowKey, nil)
	assert.Equal(t, nil, err)
	ctx3, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	_, err3 := batchExecutor.Execute(ctx3)
	assert.NotEqual(t, nil, err3)

	// query
	ctx4, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	startRowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err4 := client.Query(
		ctx4,
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2"}),
	)
	assert.Equal(t, nil, err4)
	res4, _ := resSet.Next()
	assert.Equal(t, nil, res4)

	ctx5, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	updateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows5, err5 := client.Update(
		ctx5,
		tableName,
		rowKey,
		updateColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c2", int32(0))), // where c2 = 0
	)
	assert.NotEqual(t, nil, err5)
	assert.EqualValues(t, -1, affectRows5)
}

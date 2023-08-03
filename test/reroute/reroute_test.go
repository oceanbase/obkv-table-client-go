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

package reroute

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
	reroute "github.com/oceanbase/obkv-table-client-go/test/reroute/util"
)

const (
	testInt32RerouteTableName       = "test_int32_reroute"
	testInt32RerouteCreateStatement = "create table if not exists `test_int32_reroute`(`c1` int(12) not null,`c2` int(12) default null,primary key (`c1`)) partition by hash(c1) partitions 2;"
)

const (
	tenantName   = "sys"
	databaseName = "test"
	partNum      = 2
)

func TestMoveReplica_singleOp(t *testing.T) {
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
	err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
	assert.Equal(t, nil, err)
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

func TestMoveReplica_batch(t *testing.T) {
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
	err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
	assert.Equal(t, nil, err)
	time.Sleep(5 * time.Second)

	// 3. multi get
	batchExecutor := moveCli.NewBatchExecutor(
		tableName,
		option.WithBatchSamePropertiesNames(true), // Strongly recommend you to set this option to true if all names of properties are the same in this batch.
	)
	err = batchExecutor.AddGetOp(rowKey, nil)
	assert.Equal(t, nil, err)

	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := batchExecutor.Execute(ctx)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.Size())
	assert.EqualValues(t, int32(0), res.GetResults()[0].Value("c1"))
	assert.EqualValues(t, int32(0), res.GetResults()[0].Value("c2"))
}

func TestMoveReplica_query(t *testing.T) {
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
	err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
	assert.Equal(t, nil, err)
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
	err = reroute.SwitchReplicaLeaderRandomly(tenantName, databaseName, tableName, partNum)
	assert.Equal(t, nil, err)
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

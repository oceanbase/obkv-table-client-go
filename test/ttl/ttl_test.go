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

package ttl

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testTTLTableName       = "test_ttl"
	testTTLCreateStatement = "create table if not exists test_ttl(c1 int(12) primary key, c2 int(12), c3 timestamp default current_timestamp on update current_timestamp) TTL(c3 + INTERVAL 2 second);"
)

func TestTTL_insert(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. insert, conflict, not expired, ret=OB_ERR_PRIMARY_KEY_DUPLICATE
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)

	time.Sleep(2 * time.Second)

	// 3. insert, conflict, expired, delete old, insert new
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 4. get, not expired, get successfully
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, int32(0), res.Value("c1"))
	assert.EqualValues(t, int32(1), res.Value("c2"))
}

func TestTTL_delete(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. delete, not expired, delete successfully, affectRows = 1
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	affectRows, err = cli.Delete(
		ctx,
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 3. insert
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(2 * time.Second)

	// 4. delete, expired, delete failed, affectRows = 0
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	affectRows, err = cli.Delete(
		ctx,
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, affectRows)

	// 5. get, expired, get failed
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, nil, res.Value("c1"))
}

func TestTTL_update(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. update, not expired, update successfully, affectRows = 1
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(2 * time.Second)

	// 3. update, expired, update failed, affectRows = 0
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(2)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Update(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, affectRows)

	// 4. get, expired, get failed
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, nil, res.Value("c1"))
}

func TestTTL_replace(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. replace, not expired, replace successfully, affectRows = 2
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Replace(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	time.Sleep(2 * time.Second)

	// 3. replace, expired, replace successfully, affectRows = 2
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(2)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err = cli.Replace(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	// 4. get, not expired, get successfully
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
}

func TestTTL_insertUp(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	// 1. insertUp(insert)
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. insertUp(insert), not expired, update
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err = cli.InsertOrUpdate(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	time.Sleep(2 * time.Second)

	// 3. insertUp(insert), expired, delete expired, insert new, success
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(2))}
	affectRows, err = cli.InsertOrUpdate(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 4. get, not expired, get successfully
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err := cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 2, res.Value("c2"))
}

func TestTTL_increment(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. increment, not expired, success
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
	res, err := cli.Increment(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
		option.WithReturnAffectedEntity(true),
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	assert.EqualValues(t, 1, res.Value("c2"))

	time.Sleep(2 * time.Second)

	// 3. increment, expired, delete expired, insert new, success
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	rowKey = []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(2))}
	res, err = cli.Increment(
		ctx,
		tableName,
		rowKey,
		mutateColumns,
		option.WithReturnAffectedEntity(true),
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	assert.EqualValues(t, 2, res.Value("c2"))

	// 4. get, not expired, get successfully
	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second) // 10s
	res, err = cli.Get(
		ctx,
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, 2, res.Value("c2"))
}

func TestTTL_batch(t *testing.T) {
	tableName := testTTLTableName
	defer test.DeleteTable(tableName)

	batchExecutor := cli.NewBatchExecutor(tableName)
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0)), table.NewColumn("c3", table.TimeStamp(time.Now().Local().Truncate(time.Second)))}
	err := batchExecutor.AddInsertOp(rowKey, mutateColumns)
	assert.EqualValues(t, nil, err)

	res, err := batchExecutor.Execute(context.TODO())
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, 1, len(res.GetResults()))
	assert.EqualValues(t, 1, res.GetResults()[0].AffectedRows())

	time.Sleep(2 * time.Second)

	// expired, expected insert success
	batchExecutor = cli.NewBatchExecutor(tableName)
	err = batchExecutor.AddInsertOp(rowKey, mutateColumns)
	assert.EqualValues(t, nil, err)
	res, err = batchExecutor.Execute(context.TODO())
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, 1, len(res.GetResults()))
	assert.EqualValues(t, 1, res.GetResults()[0].AffectedRows())

	time.Sleep(2 * time.Second)

	// expired, expected insertUp success
	batchExecutor = cli.NewBatchExecutor(tableName)
	err = batchExecutor.AddInsertOrUpdateOp(rowKey, mutateColumns)
	assert.EqualValues(t, nil, err)
	res, err = batchExecutor.Execute(context.TODO())
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, 1, len(res.GetResults()))
	assert.EqualValues(t, 1, res.GetResults()[0].AffectedRows())

	time.Sleep(2 * time.Second)

	// expired, expected increment success
	batchExecutor = cli.NewBatchExecutor(tableName)
	err = batchExecutor.AddIncrementOp(rowKey, mutateColumns)
	assert.EqualValues(t, nil, err)
	res, err = batchExecutor.Execute(context.TODO())
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, 1, len(res.GetResults()))
	assert.EqualValues(t, 1, res.GetResults()[0].AffectedRows())
}

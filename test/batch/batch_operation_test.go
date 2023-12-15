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

package batch

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	batchOpTableTableName       = "batchOpTable"
	batchOpTableCreateStatement = "create table if not exists batchOpTable(`c1` bigint(20) not null, c2 bigint(20), c3 varchar(20) default 'hello', primary key (`c1`)) partition by hash(c1) partitions 15;"
)

var getColumns = []string{"c1", "c2"}

func prepareRecord(recordCount int) {
	for i := 0; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s(c1, c2) values(%d, %d);", batchOpTableTableName, i, i)
		test.InsertTable(insertStatement)
	}
}

func TestBatch_MultiInsert(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		mutateColumns := []*table.Column{table.NewColumn("c2", int64(i))}
		err := batchExecutor.AddInsertOp(rowKey, mutateColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)

	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.NotEqual(t, nil, res.GetResults()[i])
		if res.GetResults()[i] != nil {
			assert.EqualValues(t, 1, res.GetResults()[i].AffectedRows())
		}
	}
}

func TestBatch_MultiGet(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		err := batchExecutor.AddGetOp(rowKey, getColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)
	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.EqualValues(t, i, res.GetResults()[i].Value("c1"))
		assert.EqualValues(t, i, res.GetResults()[i].Value("c2"))
	}
}

func TestBatch_MultiGetEmpty(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := recordCount; i < recordCount+recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		err := batchExecutor.AddGetOp(rowKey, getColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)
	assert.EqualValues(t, true, res.IsEmptySet())
}

func TestBatch_MultiDelete(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		err := batchExecutor.AddDeleteOp(rowKey)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)
	assert.EqualValues(t, recordCount, res.Size())
}

func TestBatch_MultiUpdate(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		updateColumns := []*table.Column{table.NewColumn("c2", int64(i+i))}
		err := batchExecutor.AddUpdateOp(rowKey, updateColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)

	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.EqualValues(t, 1, res.GetResults()[i].AffectedRows())
	}
}

func TestBatch_MultiReplace(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		replaceColumns := []*table.Column{table.NewColumn("c2", int64(i+i))}
		err := batchExecutor.AddReplaceOp(rowKey, replaceColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)

	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.EqualValues(t, 2, res.GetResults()[i].AffectedRows())
	}
}

func TestBatch_MultiInsertOrUpdate(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		insertUpColumns := []*table.Column{table.NewColumn("c2", int64(i+i))}
		err := batchExecutor.AddInsertOrUpdateOp(rowKey, insertUpColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)

	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.EqualValues(t, 1, res.GetResults()[i].AffectedRows())
	}
}

func TestBatch_MultiIncrement(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		incrementColumns := []*table.Column{table.NewColumn("c2", int64(i+i))}
		err := batchExecutor.AddIncrementOp(rowKey, incrementColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)

	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.EqualValues(t, 1, res.GetResults()[i].AffectedRows())
	}
}

func TestBatch_MultiAppend(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareRecord(recordCount)

	batchExecutor := cli.NewBatchExecutor(tableName)

	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		appendColumns := []*table.Column{table.NewColumn("c3", "world")}
		err := batchExecutor.AddAppendOp(rowKey, appendColumns)
		assert.Equal(t, nil, err)
	}

	res, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)

	assert.EqualValues(t, recordCount, res.Size())
	for i := 0; i < res.Size(); i++ {
		assert.EqualValues(t, 1, res.GetResults()[i].AffectedRows())
	}
}

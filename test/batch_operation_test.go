package test

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	batchOpTable = "CREATE TABLE IF NOT EXISTS batchOpTable(`c1` bigint(20) NOT NULL, c2 bigint(20) NOT NULL, PRIMARY KEY (`c1`)) PARTITION BY HASH(c1) PARTITIONS 2;"
)

func TestBatch(t *testing.T) {
	tableName := "batchOpTable"
	cli := newClient()
	createTable(batchOpTable)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	rowKey1 := []*table.Column{table.NewColumn("c1", int64(1))}
	rowKey2 := []*table.Column{table.NewColumn("c1", int64(2))}
	selectColumns1 := []string{"c1"}
	selectColumns2 := []string{"c2"}
	mutateColumns1 := []*table.Column{table.NewColumn("c2", int64(1))}
	mutateColumns2 := []*table.Column{table.NewColumn("c2", int64(2))}

	batchExecutor := cli.NewBatchExecutor(tableName)
	err = batchExecutor.AddInsertOp(rowKey1, mutateColumns1)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddInsertOp(rowKey2, mutateColumns2)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddGetOp(rowKey1, selectColumns1)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddGetOp(rowKey2, selectColumns2)
	assert.Equal(t, nil, err)
	batchRes, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)
	allResults := batchRes.GetResults()
	assert.Equal(t, 4, len(allResults))
	assert.EqualValues(t, 1, allResults[0].AffectedRows())
	assert.EqualValues(t, 1, allResults[1].AffectedRows())
	assert.EqualValues(t, 1, allResults[2].Entity().GetSimpleProperties()["c1"])
	assert.EqualValues(t, 2, allResults[3].Entity().GetSimpleProperties()["c2"])
}

package client

import (
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatch(t *testing.T) {
	// CREATE TABLE test(c1 INT, c2 int) PARTITION BY hash(c1) partitions 2;
	const (
		configUrl    = "xxx"
		fullUserName = "xxx"
		passWord     = ""
		sysUserName  = "root"
		sysPassWord  = ""
		tableName    = "test"
	)

	cfg := config.NewDefaultClientConfig()
	cli, err := NewClient(configUrl, fullUserName, passWord, sysUserName, sysPassWord, cfg)
	assert.Equal(t, nil, err)

	err = cli.AddRowkey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)
	batchExecutor := cli.NewBatchExecutor(tableName)

	rowkey1 := []*table.Column{table.NewColumn("c1", int64(1))}
	rowkey2 := []*table.Column{table.NewColumn("c1", int64(2))}
	selectColumns1 := []string{"c1"}
	selectColumns2 := []string{"c2"}
	mutateColumns1 := []*table.Column{table.NewColumn("c2", int64(1))}
	mutateColumns2 := []*table.Column{table.NewColumn("c2", int64(2))}

	err = batchExecutor.AddInsertOp(rowkey1, mutateColumns1)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddInsertOp(rowkey2, mutateColumns2)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddGetOp(rowkey1, selectColumns1)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddGetOp(rowkey2, selectColumns2)
	assert.Equal(t, nil, err)
	batchRes, err := batchExecutor.Execute()
	assert.Equal(t, nil, err)
	allResults := batchRes.GetResults()
	assert.Equal(t, 4, len(allResults))
	assert.EqualValues(t, 1, allResults[0].AffectedRows())
	assert.EqualValues(t, 1, allResults[1].AffectedRows())
	assert.EqualValues(t, 1, allResults[2].Entity().GetSimpleProperties()["c1"])
	assert.EqualValues(t, 2, allResults[3].Entity().GetSimpleProperties()["c2"])
}

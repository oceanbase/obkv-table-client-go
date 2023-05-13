package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
)

// create table statement
const (
	hashTable = "CREATE TABLE IF NOT EXISTS hashTable(`c1` bigint(20) NOT NULL, c2 bigint(20) NOT NULL, PRIMARY KEY (`c1`)) PARTITION BY HASH(c1) PARTITIONS 2;"
)

func TestInsert(t *testing.T) {
	tableName := "hashTable"
	cli := newClient()
	createTable(hashTable)
	defer func() {
		deleteTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"}
	m, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c1"])
	assert.EqualValues(t, 1, m["c2"])
}

func TestUpdate(t *testing.T) {
	tableName := "hashTable"
	cli := newClient()
	createTable(hashTable)
	defer func() {
		deleteTable(tableName)
	}()

	err := cli.AddRowKey("hashTable", []string{"c1"})
	assert.Equal(t, nil, err)

	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	affectRows, err := cli.Insert(
		context.TODO(),
		"hashTable",
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"}
	m, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c1"])
	assert.EqualValues(t, 2, m["c2"])
}

func TestGet(t *testing.T) {
	tableName := "hashTable"
	cli := newClient()
	createTable(hashTable)
	defer func() {
		deleteTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int64(1))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"} // select c1, c2
	m, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c1"])
	assert.EqualValues(t, 1, m["c2"])

	selectColumns = []string{"c1"} // select c1
	m, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c1"])
	assert.EqualValues(t, nil, m["c2"])

	selectColumns = nil // default select all when selectColumns is nil
	m, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, m["c1"])
	assert.EqualValues(t, 1, m["c2"])

	deleteTable(tableName)
	m, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, nil, m["c1"])
	assert.EqualValues(t, nil, m["c2"])
}

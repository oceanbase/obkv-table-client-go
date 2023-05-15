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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
)

// create table statement
const (
	hashPartitionL1 = "CREATE TABLE IF NOT EXISTS hashPartitionL1(`c1` bigint(20) NOT NULL, c2 bigint(20) NOT NULL, " +
		"PRIMARY KEY (`c1`)) PARTITION BY HASH(c1) PARTITIONS 2;"
	keyPartitionIntL1 = "CREATE TABLE IF NOT EXISTS keyPartitionIntL1(`c1` bigint(20) NOT NULL, c2 bigint(20) NOT NULL, " +
		"PRIMARY KEY (`c1`)) PARTITION BY KEY(c1) PARTITIONS 15;"
	// todo test date type
	keyPartitionVarcharL1 = "CREATE TABLE IF NOT EXISTS keyPartitionVarcharL1(`c1` varchar(20) NOT NULL, c2 bigint(20) NOT NULL, " +
		"PRIMARY KEY (`c1`)) PARTITION BY KEY(c1) PARTITIONS 15;"
	hashPartitionL2 = "CREATE TABLE IF NOT EXISTS hashPartitionL2(`c1` bigint(20) NOT NULL, `c2` bigint(20) NOT NULL, `c3` bigint(20) NOT NULL, " +
		"PRIMARY KEY (`c1`, `c2`)) PARTITION BY HASH(`c1`) SUBPARTITION BY hash(`c2`) SUBPARTITIONS 4 PARTITIONS 16;"
	keyPartitionIntL2 = "CREATE TABLE IF NOT EXISTS keyPartitionIntL2(`c1` bigint(20) NOT NULL, c2 bigint(20) NOT NULL, `c3` bigint(20) NOT NULL, " +
		"PRIMARY KEY (`c1`, `c2`)) PARTITION BY KEY(`c1`) SUBPARTITION BY KEY(`c2`) SUBPARTITIONS 4 PARTITIONS 16;"
	keyPartitionVarcharL2 = "CREATE TABLE IF NOT EXISTS keyPartitionVarcharL2(`c1` varchar(20) NOT NULL, `c2` varchar(20) NOT NULL, c3 bigint(20) NOT NULL, " +
		"PRIMARY KEY (`c1`, `c2`)) PARTITION BY KEY(`c1`) SUBPARTITION BY KEY(`c2`) SUBPARTITIONS 4 PARTITIONS 16;"
)

func TestHashPartitionL1(t *testing.T) {
	tableName := "hashPartitionL1"
	cli := newClient()
	createTable(hashPartitionL1)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		mutateColumns := []*table.Column{table.NewColumn("c2", int64(i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// get
	selectColumns := []string{"c1", "c2"}
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, i, m["c1"])
		assert.EqualValues(t, i, m["c2"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2 from %s where c1 = %d;", tableName, i)
		rows := selectTable(selectStatement)
		var c1Val int
		var c2Val int
		for rows.Next() {
			err = rows.Scan(&c1Val, &c2Val)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, c1Val, m["c1"])
			assert.EqualValues(t, c2Val, m["c2"])
		}
	}
}

func TestKeyPartitionIntL1(t *testing.T) {
	tableName := "keyPartitionIntL1"
	cli := newClient()
	createTable(keyPartitionIntL1)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		mutateColumns := []*table.Column{table.NewColumn("c2", int64(i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// get
	selectColumns := []string{"c1", "c2"}
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, i, m["c1"])
		assert.EqualValues(t, i, m["c2"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2 from %s where c1 = %d;", tableName, i)
		rows := selectTable(selectStatement)
		var c1Val int
		var c2Val int
		for rows.Next() {
			err = rows.Scan(&c1Val, &c2Val)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, c1Val, m["c1"])
			assert.EqualValues(t, c2Val, m["c2"])
		}
	}
}

func TestKeyPartitionVarcharL1(t *testing.T) {
	tableName := "keyPartitionVarcharL1"
	cli := newClient()
	createTable(keyPartitionVarcharL1)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowkeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowkeyVal)}
		mutateColumns := []*table.Column{table.NewColumn("c2", int64(i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// get
	selectColumns := []string{"c1", "c2"}
	for i := 0; i < rowCount; i++ {
		rowkeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowkeyVal)}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, rowkeyVal, m["c1"])
		assert.EqualValues(t, i, m["c2"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2 from %s where c1 = '%s';", tableName, rowkeyVal)
		rows := selectTable(selectStatement)
		var c1Val string
		var c2Val int
		for rows.Next() {
			err = rows.Scan(&c1Val, &c2Val)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, c1Val, m["c1"])
			assert.EqualValues(t, c2Val, m["c2"])
		}
	}
}

func TestHashPartitionL2(t *testing.T) {
	tableName := "hashPartitionL2"
	cli := newClient()
	createTable(hashPartitionL2)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1", "c2"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i)), table.NewColumn("c2", int64(i))}
		mutateColumns := []*table.Column{table.NewColumn("c3", int64(i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// get
	selectColumns := []string{"c1", "c2", "c3"}
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i)), table.NewColumn("c2", int64(i))}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, i, m["c1"])
		assert.EqualValues(t, i, m["c2"])
		assert.EqualValues(t, i, m["c3"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2, c3 from %s where c1 = %d and c2 = %d;", tableName, i, i)
		rows := selectTable(selectStatement)
		var c1Val int
		var c2Val int
		var c3Val int
		for rows.Next() {
			err = rows.Scan(&c1Val, &c2Val, &c3Val)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, c1Val, m["c1"])
			assert.EqualValues(t, c2Val, m["c2"])
			assert.EqualValues(t, c3Val, m["c3"])
		}
	}
}

func TestKeyPartitionIntL2(t *testing.T) {
	tableName := "keyPartitionIntL2"
	cli := newClient()
	createTable(keyPartitionIntL2)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1", "c2"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i)), table.NewColumn("c2", int64(i))}
		mutateColumns := []*table.Column{table.NewColumn("c3", int64(i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// get
	selectColumns := []string{"c1", "c2", "c3"}
	for i := 0; i < rowCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i)), table.NewColumn("c2", int64(i))}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, i, m["c1"])
		assert.EqualValues(t, i, m["c2"])
		assert.EqualValues(t, i, m["c3"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2, c3 from %s where c1 = %d and c2 = %d;", tableName, i, i)
		rows := selectTable(selectStatement)
		var c1Val int
		var c2Val int
		var c3Val int
		for rows.Next() {
			err = rows.Scan(&c1Val, &c2Val, &c3Val)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, c1Val, m["c1"])
			assert.EqualValues(t, c2Val, m["c2"])
			assert.EqualValues(t, c3Val, m["c3"])
		}
	}
}

func TestKeyPartitionVarcharL2(t *testing.T) {
	tableName := "keyPartitionVarcharL2"
	cli := newClient()
	createTable(keyPartitionVarcharL2)
	defer func() {
		dropTable(tableName)
	}()

	err := cli.AddRowKey(tableName, []string{"c1", "c2"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowkeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowkeyVal), table.NewColumn("c2", rowkeyVal)}
		mutateColumns := []*table.Column{table.NewColumn("c3", int64(i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// get
	selectColumns := []string{"c1", "c2", "c3"}
	for i := 0; i < rowCount; i++ {
		rowkeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowkeyVal), table.NewColumn("c2", rowkeyVal)}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, rowkeyVal, m["c1"])
		assert.EqualValues(t, rowkeyVal, m["c2"])
		assert.EqualValues(t, i, m["c3"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2, c3 from %s where c1 = '%s' and c2 = '%s';", tableName, rowkeyVal, rowkeyVal)
		rows := selectTable(selectStatement)
		var c1Val string
		var c2Val string
		var c3Val int
		for rows.Next() {
			err = rows.Scan(&c1Val, &c2Val, &c3Val)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, c1Val, m["c1"])
			assert.EqualValues(t, c2Val, m["c2"])
			assert.EqualValues(t, c3Val, m["c3"])
		}
	}
}

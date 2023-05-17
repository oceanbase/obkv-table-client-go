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

package partition

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

func TestHashPartitionL1(t *testing.T) {
	tableName := hashPartitionL1TableName
	defer test.DeleteTable(tableName)

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
		rows := test.SelectTable(selectStatement)
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
	tableName := keyPartitionIntL1TableName
	defer test.DeleteTable(tableName)

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
		rows := test.SelectTable(selectStatement)
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
	tableName := keyPartitionVarcharL1TableName
	defer test.DeleteTable(tableName)

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowKeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowKeyVal)}
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
		rowKeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowKeyVal)}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, rowKeyVal, m["c1"])
		assert.EqualValues(t, i, m["c2"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2 from %s where c1 = '%s';", tableName, rowKeyVal)
		rows := test.SelectTable(selectStatement)
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
	tableName := hashPartitionL2TableName
	defer test.DeleteTable(tableName)

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
		rows := test.SelectTable(selectStatement)
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
	tableName := keyPartitionIntL2TableName
	defer test.DeleteTable(tableName)

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
		rows := test.SelectTable(selectStatement)
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
	tableName := keyPartitionVarcharL2TableName
	defer test.DeleteTable(tableName)

	err := cli.AddRowKey(tableName, []string{"c1", "c2"})
	assert.Equal(t, nil, err)

	// insert
	rowCount := 10
	for i := 0; i < rowCount; i++ {
		rowKeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowKeyVal), table.NewColumn("c2", rowKeyVal)}
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
		rowKeyVal := fmt.Sprintf("oceanbase%d", i)
		rowKey := []*table.Column{table.NewColumn("c1", rowKeyVal), table.NewColumn("c2", rowKeyVal)}
		m, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			selectColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, rowKeyVal, m["c1"])
		assert.EqualValues(t, rowKeyVal, m["c2"])
		assert.EqualValues(t, i, m["c3"])

		// select by sql
		selectStatement := fmt.Sprintf("select c1, c2, c3 from %s where c1 = '%s' and c2 = '%s';", tableName, rowKeyVal, rowKeyVal)
		rows := test.SelectTable(selectStatement)
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

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

package single

import (
	"context"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client/option"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testVarcharTableName       = "test_varchar"
	testVarcharCreateStatement = "create table if not exists `test_varchar`(`c1` varchar(20) not null,`c2` varchar(20) default null,primary key (`c1`));"
)

func TestInsertVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))
}

func TestUpdateVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns := []*table.Column{table.NewColumn("c2", "2")}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "2", result.Value("c2"))
}

func TestInsertOrUpdateVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))

	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns = []string{"c1", "c2"}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))
}

func TestDeleteVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, affectRows)
}

func TestGetVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	selectColumns := []string{"c1", "c2"} // select c1, c2
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))

	selectColumns = []string{"c1"} // select c1
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, nil, result.Value("c2"))

	selectColumns = nil // default select all when selectColumns is nil
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))

	test.DeleteTable(tableName)
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, nil, result.Value("c1"))
	assert.EqualValues(t, nil, result.Value("c2"))
}

func TestIncrementVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// Increment
	incrementColumns := []*table.Column{table.NewColumn("c2", "2")}
	res, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		incrementColumns,
		option.WithReturnRowKey(true),
		option.WithReturnAffectedEntity(true), // return affected entity
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())
	assert.EqualValues(t, "3", res.Value("c2"))
	assert.EqualValues(t, "1", res.RowKey()[0])

	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "3", result.Value("c2"))
}

func TestAppendVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// append
	appendColumns := []*table.Column{table.NewColumn("c2", "2")}
	res, err := cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		appendColumns,
		option.WithReturnRowKey(true),
		option.WithReturnAffectedEntity(true), // return affected entity
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())
	assert.EqualValues(t, "12", res.Value("c2"))
	assert.EqualValues(t, "1", res.RowKey()[0])

	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "12", result.Value("c2"))
}

func TestReplaceVarchar(t *testing.T) {
	tableName := testVarcharTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}
	mutateColumns := []*table.Column{table.NewColumn("c2", "1")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// replace
	replaceColumns := []*table.Column{table.NewColumn("c2", "2")}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		replaceColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	selectColumns := []string{"c1", "c2"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "1", result.Value("c1"))
	assert.EqualValues(t, "2", result.Value("c2"))
}

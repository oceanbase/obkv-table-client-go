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
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testDateTableName       = "test_date"
	testDateCreateStatement = "create table if not exists `test_date`(`c1` date not null,`c2` date default null,primary key (`c1`));"
)

func TestInsertDate(t *testing.T) {
	tableName := testDateTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.Date{Value: time.Date(1990, 5, 25, 1, 0, 0, 0, time.Local)})}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
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
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c2"))
}

func TestUpdateDate(t *testing.T) {
	tableName := testDateTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns := []*table.Column{table.NewColumn("c2", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 1, 1, time.Local)})}
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
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c2"))
}

func TestInsertOrUpdateDate(t *testing.T) {
	tableName := testDateTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
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
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c2"))

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
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c2"))

}

func TestDeleteDate(t *testing.T) {
	tableName := testDateTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
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

func TestGetDate(t *testing.T) {
	tableName := testDateTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.Date{Value: time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)})}
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
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c2"))

	selectColumns = []string{"c1"} // select c1
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, nil, result.Value("c2"))

	selectColumns = nil // default select all when selectColumns is nil
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c1"))
	assert.EqualValues(t, time.Date(1990, time.May, 25, 0, 0, 0, 0, time.Local), result.Value("c2"))

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

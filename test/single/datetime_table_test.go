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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testDatetimeTableName       = "test_datetime"
	testDatetimeCreateStatement = "create table if not exists `test_datetime`(`c1` datetime(6) not null,`c2` datetime(6) default null,primary key (`c1`));"
)

func TestInsertDatetime(t *testing.T) {
	tableName := testDatetimeTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.DateTime(time.Now().Local()))}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.DateTime(time.Now().Local()))}
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
	fmt.Println(result.Value("c1").(time.Time).String())
	fmt.Println(result.Value("c2").(time.Time).String())
}

func TestUpdateDatetime(t *testing.T) {
	tableName := testDatetimeTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.DateTime(time.Now()))}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.DateTime(time.Now()))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns := []*table.Column{table.NewColumn("c2", table.DateTime(time.Now()))}
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
	println(result.Value("c1").(time.Time).String())
	println(result.Value("c2").(time.Time).String())
}

func TestInsertOrUpdateDatetime(t *testing.T) {
	tableName := testDatetimeTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.DateTime(time.Now()))}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.DateTime(time.Now()))}
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
	println(result.Value("c1").(time.Time).String())
	println(result.Value("c2").(time.Time).String())

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
	println(result.Value("c1").(time.Time).String())
	println(result.Value("c2").(time.Time).String())
}

func TestDeleteDatetime(t *testing.T) {
	tableName := testDatetimeTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.DateTime(time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)))}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.DateTime(time.Date(1990, 5, 25, 0, 0, 0, 0, time.Local)))}
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

func TestGetDatetime(t *testing.T) {
	tableName := testDatetimeTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", table.DateTime(time.Now()))}
	mutateColumns := []*table.Column{table.NewColumn("c2", table.DateTime(time.Now()))}
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
	println(result.Value("c1").(time.Time).String())
	println(result.Value("c2").(time.Time).String())

	selectColumns = []string{"c1"} // select c1
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	println(result.Value("c1").(time.Time).String())

	selectColumns = nil // default select all when selectColumns is nil
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	println(result.Value("c1").(time.Time).String())
	println(result.Value("c2").(time.Time).String())

	test.DeleteTable(tableName)
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, result.IsEmptySet())
}

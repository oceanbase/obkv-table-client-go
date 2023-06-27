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

package filter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testFilterOpTableName       = "test_filter_operation"
	testFilterOpCreateStatement = "create table if not exists `test_filter_operation`(`c1` varchar(20) not null,`c2` bigint(20) default null,`c3` varchar(20) default null,primary key (`c1`));"
)

func TestAppend(t *testing.T) {
	tableName := testFilterOpTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}

	insertColumns := []*table.Column{table.NewColumn("c2", int64(1)), table.NewColumn("c3", "hello")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		insertColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	appendColumns := []*table.Column{table.NewColumn("c3", " oceanbase")}
	res, err := cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		appendColumns,
		client.WithFilter(filter.CompareVal(filter.Equal, "c3", "hello")), // where c2 = hello
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())
}

func TestIncrement(t *testing.T) {
	tableName := testFilterOpTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}

	insertColumns := []*table.Column{table.NewColumn("c2", int64(1)), table.NewColumn("c3", "hello")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		insertColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	appendColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	res, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		appendColumns,
		client.WithFilter(filter.CompareVal(filter.Equal, "c3", "hello")), // where c2 = hello
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())
}

func TestUpdate(t *testing.T) {
	tableName := testFilterOpTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}

	insertColumns := []*table.Column{table.NewColumn("c2", int64(1)), table.NewColumn("c3", "hello")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		insertColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	appendColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		appendColumns,
		client.WithFilter(filter.CompareVal(filter.Equal, "c3", "hello")), // where c2 = hello
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

func TestDelete(t *testing.T) {
	tableName := testFilterOpTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}

	insertColumns := []*table.Column{table.NewColumn("c2", int64(1)), table.NewColumn("c3", "hello")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		insertColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
		client.WithFilter(filter.CompareVal(filter.Equal, "c3", "hello")), // where c2 = hello
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

func TestFilter(t *testing.T) {
	tableName := testFilterOpTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("c1", "1")}

	insertColumns := []*table.Column{table.NewColumn("c2", int64(1)), table.NewColumn("c3", "hello")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		insertColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns := []*table.Column{table.NewColumn("c2", int64(2))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.LessThan, "c2", int64(2))),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(3))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.GreaterThan, "c2", int64(1))),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(4))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.LessOrEqualThan, "c2", int64(3))),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(5))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.GreaterOrEqualThan, "c2", int64(4))),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(6))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.NotEqual, "c2", int64(0))),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(7))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.Equal, "c2", int64(6))),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", nil)}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(8))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.IsNull, "c2", nil)),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	updateColumns = []*table.Column{table.NewColumn("c2", int64(9))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(filter.CompareVal(filter.IsNotNull, "c2", nil)),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	var andFilterList []filter.ObTableFilter
	andFilterList = append(andFilterList, filter.CompareVal(filter.LessThan, "c2", int64(10)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.GreaterThan, "c2", int64(8)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.LessOrEqualThan, "c2", int64(9)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.GreaterOrEqualThan, "c2", int64(9)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.NotEqual, "c2", int64(0)))
	andFilterList = append(andFilterList, filter.CompareVal(filter.Equal, "c2", int64(9)))

	andList := filter.AndList(andFilterList...)
	updateColumns = []*table.Column{table.NewColumn("c2", int64(10))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(andList),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	var orFilterList []filter.ObTableFilter
	orFilterList = append(orFilterList, filter.CompareVal(filter.LessThan, "c2", int64(9)))
	orFilterList = append(orFilterList, filter.CompareVal(filter.GreaterThan, "c2", int64(11)))
	orFilterList = append(orFilterList, filter.CompareVal(filter.LessOrEqualThan, "c2", int64(9)))
	orFilterList = append(orFilterList, filter.CompareVal(filter.GreaterOrEqualThan, "c2", int64(11)))
	orFilterList = append(orFilterList, filter.CompareVal(filter.NotEqual, "c2", int64(10)))
	orFilterList = append(orFilterList, filter.CompareVal(filter.Equal, "c2", int64(0)))

	orList := filter.OrList(orFilterList...)
	updateColumns = []*table.Column{table.NewColumn("c2", int64(11))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		updateColumns,
		client.WithFilter(orList),
	)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, affectRows)
}

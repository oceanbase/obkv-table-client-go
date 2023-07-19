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

package autoinc

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/test"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/table"
)

const (
	autoIncRowkeyTableTableName       = "autoIncRowkeyTable"
	autoIncRowkeyTableCreateStatement = "create table if not exists autoIncRowkeyTable(`c1` bigint(20) not null auto_increment, c2 bigint(20) not null, c3 varchar(20) default 'hello', c4 bigint(20) default 0, primary key (`c1`, `c2`)) partition by hash(c2) partitions 15;"
)

func TestAuto_IncRowkey(t *testing.T) {
	tableName := autoIncRowkeyTableTableName
	defer test.DeleteTable(tableName)

	// test insert.
	// test auto inc value.
	rowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}

	res, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), res)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", int64(1))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(1), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	// test assign value.
	rowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}

	res, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), res)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(50), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	// test update auto inc value.
	rowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}

	res, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), res)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(51)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(51), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	// test delete
	rowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(1))}
	res, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)

	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(50), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	// test update
	rowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}
	mutationColumns := []*table.Column{table.NewColumn("c3", "update")}
	res, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c1", int64(50))),
	)

	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(50)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(50), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "update", res.Value("c3"))
	}

	// test replace not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "replace")}
	res, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(52)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(52)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(52), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "replace", res.Value("c3"))
	}

	// test replace exist, replace
	rowKey = []*table.Column{table.NewColumn("c1", int64(52)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "replace exist")}
	res, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(52)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(52)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(52), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "replace exist", res.Value("c3"))
	}

	// test insertup not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "insertup")}
	res, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(53)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(53)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(53), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "insertup", res.Value("c3"))
	}

	// test insertup exist, update
	rowKey = []*table.Column{table.NewColumn("c1", int64(53)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "insertup exist")}
	res, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(53)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(53)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(53), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "insertup exist", res.Value("c3"))
	}

	// test increment not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c4", int64(10))}
	resultSet, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	startRowKey = []*table.Column{table.NewColumn("c1", int64(54)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(54)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3", "c4"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(54), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, int64(10), res.Value("c4"))
	}

	// test increment exist, increment
	rowKey = []*table.Column{table.NewColumn("c1", int64(54)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c4", int64(10))}
	resultSet, err = cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	startRowKey = []*table.Column{table.NewColumn("c1", int64(54)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(54)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3", "c4"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(54), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, int64(20), res.Value("c4"))
	}

	// test append not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "append")}
	resultSet, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	startRowKey = []*table.Column{table.NewColumn("c1", int64(55)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(55)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3", "c4"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(55), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "append", res.Value("c3"))
	}

	// test append exist, append
	rowKey = []*table.Column{table.NewColumn("c1", int64(55)), table.NewColumn("c2", int64(1))}
	mutationColumns = []*table.Column{table.NewColumn("c3", " exist")}
	resultSet, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	startRowKey = []*table.Column{table.NewColumn("c1", int64(55)), table.NewColumn("c2", int64(1))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(55)), table.NewColumn("c2", int64(1))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithSelectColumns([]string{"c1", "c2", "c3", "c4"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 1; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, int64(55), res.Value("c1"))
		assert.EqualValues(t, int64(1), res.Value("c2"))
		assert.EqualValues(t, "append exist", res.Value("c3"))
	}
}

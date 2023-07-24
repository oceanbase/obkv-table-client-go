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
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	autoIncNotRowkeyTableTableName       = "autoIncNotRowkeyTable"
	autoIncNotRowkeyTableCreateStatement = "create table if not exists autoIncNotRowkeyTable(`c1` bigint(20) not null, c2 bigint(20) auto_increment, c3 varchar(20) default 'hello', c4 bigint(20) default 0, primary key (`c1`)) partition by key(c1) partitions 100;"
)

func TestAuto_IncNotRowkey(t *testing.T) {
	tableName := autoIncNotRowkeyTableTableName
	defer test.DeleteTable(tableName)

	// test insert.
	// test auto inc value.
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}

	res, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), res)

	rowKey = []*table.Column{table.NewColumn("c1", int64(1))}
	selectColumns := []string{"c1", "c2", "c3", "c4"}
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(1), result.Value("c1"))
	assert.EqualValues(t, int64(1), result.Value("c2"))
	assert.EqualValues(t, "hello", result.Value("c3"))

	// test assign value.
	rowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	mutationColumns := []*table.Column{table.NewColumn("c2", int64(50))}

	res, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), res)

	rowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(2), result.Value("c1"))
	assert.EqualValues(t, int64(50), result.Value("c2"))
	assert.EqualValues(t, "hello", result.Value("c3"))

	// test update auto inc value.
	rowKey = []*table.Column{table.NewColumn("c1", int64(3))}

	res, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), res)

	rowKey = []*table.Column{table.NewColumn("c1", int64(3))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(3), result.Value("c1"))
	assert.EqualValues(t, int64(51), result.Value("c2"))
	assert.EqualValues(t, "hello", result.Value("c3"))

	// test delete.
	rowKey = []*table.Column{table.NewColumn("c1", int64(1))}
	res, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)

	assert.Equal(t, nil, err)

	rowKey = []*table.Column{table.NewColumn("c1", int64(1))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, nil, result.Value("c1"))

	// test update.
	rowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "update")}
	res, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
		option.WithFilter(filter.CompareVal(filter.Equal, "c1", int64(2))),
	)

	assert.Equal(t, nil, err)

	rowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(2), result.Value("c1"))
	assert.EqualValues(t, int64(50), result.Value("c2"))
	assert.EqualValues(t, "update", result.Value("c3"))

	// test replace not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(4))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "replace")}
	res, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	rowKey = []*table.Column{table.NewColumn("c1", int64(4))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(4), result.Value("c1"))
	assert.EqualValues(t, int64(52), result.Value("c2"))
	assert.EqualValues(t, "replace", result.Value("c3"))

	// test replace exist, replace
	rowKey = []*table.Column{table.NewColumn("c1", int64(4))}
	mutationColumns = []*table.Column{table.NewColumn("c2", int64(20)), table.NewColumn("c3", "replace exist")}
	res, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	rowKey = []*table.Column{table.NewColumn("c1", int64(4))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(4), result.Value("c1"))
	assert.EqualValues(t, int64(20), result.Value("c2"))
	assert.EqualValues(t, "replace exist", result.Value("c3"))

	// test insertup not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(5))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "insertup")}
	res, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	rowKey = []*table.Column{table.NewColumn("c1", int64(5))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(5), result.Value("c1"))
	assert.EqualValues(t, int64(53), result.Value("c2"))
	assert.EqualValues(t, "insertup", result.Value("c3"))

	// test insertup exist, update
	rowKey = []*table.Column{table.NewColumn("c1", int64(5))}
	mutationColumns = []*table.Column{table.NewColumn("c2", int64(20)), table.NewColumn("c3", "insertup exist")}
	res, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)

	rowKey = []*table.Column{table.NewColumn("c1", int64(5))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(5), result.Value("c1"))
	assert.EqualValues(t, int64(20), result.Value("c2"))
	assert.EqualValues(t, "insertup exist", result.Value("c3"))

	// test increment not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(6))}
	mutationColumns = []*table.Column{table.NewColumn("c4", int64(10))}
	resultSet, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	rowKey = []*table.Column{table.NewColumn("c1", int64(6))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(6), result.Value("c1"))
	assert.EqualValues(t, int64(54), result.Value("c2"))
	assert.EqualValues(t, int64(10), result.Value("c4"))

	// test increment exist, increment
	rowKey = []*table.Column{table.NewColumn("c1", int64(6))}
	mutationColumns = []*table.Column{table.NewColumn("c4", int64(10))}
	resultSet, err = cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	rowKey = []*table.Column{table.NewColumn("c1", int64(6))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(6), result.Value("c1"))
	assert.EqualValues(t, int64(54), result.Value("c2"))
	assert.EqualValues(t, int64(20), result.Value("c4"))

	// test append not exist, insert
	rowKey = []*table.Column{table.NewColumn("c1", int64(7))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "append")}
	resultSet, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	rowKey = []*table.Column{table.NewColumn("c1", int64(7))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(7), result.Value("c1"))
	assert.EqualValues(t, int64(56), result.Value("c2"))
	assert.EqualValues(t, "append", result.Value("c3"))

	// test append exist, append
	rowKey = []*table.Column{table.NewColumn("c1", int64(7))}
	mutationColumns = []*table.Column{table.NewColumn("c3", " exist")}
	resultSet, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), resultSet.AffectedRows())

	rowKey = []*table.Column{table.NewColumn("c1", int64(7))}
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		selectColumns,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, int64(7), result.Value("c1"))
	assert.EqualValues(t, int64(56), result.Value("c2"))
	assert.EqualValues(t, "append exist", result.Value("c3"))
}

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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	autoIncNormalNotFillTableTableName       = "autoIncNormalNotFillTable"
	autoIncNormalNotFillTableCreateStatement = "create table if not exists autoIncNormalNotFillTable(`c1` bigint(20) not null, c2 bigint(20) auto_increment, c3 varchar(20) default 'hello', c4 bigint(20) default 0, primary key (`c1`)) partition by key(c1) partitions 100;"
)

func TestAuto_NormalNotFill(t *testing.T) {
	tableName := autoIncNormalNotFillTableTableName
	defer test.DeleteTable(tableName)

	// test insert.
	// c1-1 c2-auto=1 c3-default("hello") c4-default(0)
	rowKey := []*table.Column{table.NewColumn("c1", int64(1))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, "hello", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// test update.
	// c1-1 c2-1 c3-"update" c4-default(0)
	rowKey = []*table.Column{table.NewColumn("c1", int64(1))}
	mutationColumns := []*table.Column{table.NewColumn("c3", "update")}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, "update", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// test replace not exist, insert
	// c1-2 c2-auto-2 c3-"replace" c4-default(0)
	rowKey = []*table.Column{table.NewColumn("c1", int64(2))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "replace")}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, result.Value("c1"))
	assert.EqualValues(t, 2, result.Value("c2"))
	assert.EqualValues(t, "replace", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// test replace exist, replace
	// c1-2 c2-auto-3 c3-"replace exist" c4-default(0)
	mutationColumns = []*table.Column{table.NewColumn("c3", "replace exist")}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, result.Value("c1"))
	assert.EqualValues(t, 3, result.Value("c2"))
	assert.EqualValues(t, "replace exist", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// test insertup not exist, insert
	// c1-3 c2-auto-4 c3-"insertup-insert" c4-default(0)
	rowKey = []*table.Column{table.NewColumn("c1", int64(3))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "insertup-insert")}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)

	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, result.Value("c1"))
	assert.EqualValues(t, 4, result.Value("c2"))
	assert.EqualValues(t, "insertup-insert", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// test insertup exist, update
	// c1-3 c2-auto-4 c3-"insertup-update" c4-default(0)
	mutationColumns = []*table.Column{table.NewColumn("c3", "insertup-update")}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, result.Value("c1"))
	assert.EqualValues(t, 4, result.Value("c2"))
	assert.EqualValues(t, "insertup-update", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// insert up cause c2 to be 5, global auto value is 5 now

	// test increment not exist, insert
	// c1-4 c2-auto-6 c3-default("hello") c4-1
	rowKey = []*table.Column{table.NewColumn("c1", int64(4))}
	mutationColumns = []*table.Column{table.NewColumn("c4", int64(1))}
	resultSet, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, resultSet.AffectedRows())

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 4, result.Value("c1"))
	assert.EqualValues(t, 6, result.Value("c2"))
	assert.EqualValues(t, "hello", result.Value("c3"))
	assert.EqualValues(t, 1, result.Value("c4"))

	// test increment exist, increment
	// c1-4 c2-auto-6 c3-default("hello") c4-1+1
	mutationColumns = []*table.Column{table.NewColumn("c4", int64(1))}
	resultSet, err = cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, resultSet.AffectedRows())

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 4, result.Value("c1"))
	assert.EqualValues(t, 6, result.Value("c2"))
	assert.EqualValues(t, "hello", result.Value("c3"))
	assert.EqualValues(t, 2, result.Value("c4"))

	// insert up(increment) cause c2 to be 5, global auto value is 7 now

	// test append not exist, insert
	// c1-5 c2-auto-8 c3-"append" c4-default(0)
	rowKey = []*table.Column{table.NewColumn("c1", int64(5))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "append")}
	resultSet, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, resultSet.AffectedRows())

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 5, result.Value("c1"))
	assert.EqualValues(t, 7, result.Value("c2"))
	assert.EqualValues(t, "append", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))

	// test append exist, append
	// c1-5 c2-auto-8 c3-"append-exist" c4-default(0)
	rowKey = []*table.Column{table.NewColumn("c1", int64(5))}
	mutationColumns = []*table.Column{table.NewColumn("c3", "-exist")}
	resultSet, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutationColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, resultSet.AffectedRows())

	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 5, result.Value("c1"))
	assert.EqualValues(t, 7, result.Value("c2"))
	assert.EqualValues(t, "append-exist", result.Value("c3"))
	assert.EqualValues(t, 0, result.Value("c4"))
}

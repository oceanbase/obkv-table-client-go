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

package primary_key_order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testPrimaryKeyTableName1        = "testPrimaryKey1"
	testPrimaryKeyTCreateStatement1 = "create table if not exists `testPrimaryKey1`(`c1` varchar(64), `c2` bigint, `c3` bigint, c4 bigint default null, primary key (`c1`, `c3`)) partition by key(c1) partitions 16;"

	testPrimaryKeyTableName2        = "testPrimaryKey2"
	testPrimaryKeyTCreateStatement2 = "create table if not exists `testPrimaryKey2`(`c1` varchar(64), `c2` bigint, `c3` bigint, c4 bigint default null, primary key (`c3`, `c1`)) partition by key(c1) partitions 16;"

	testPrimaryKeyTableName3        = "testPrimaryKey3"
	testPrimaryKeyTCreateStatement3 = "create table if not exists `testPrimaryKey3`(`c1` varchar(64), `c2` bigint, `c3` bigint, c4 bigint default null, primary key (`c2`, `c3`)) partition by key(c2) partitions 16;"
)

func TestPrimaryKeyOrder_test1(t *testing.T) {
	tableName := testPrimaryKeyTableName1
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{
		table.NewColumn("c1", "c1_value"),
		table.NewColumn("c3", int64(1)),
	}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", int64(2)),
	}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// get
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "c1_value", result.Value("c1"))
	assert.EqualValues(t, 2, result.Value("c2"))
	assert.EqualValues(t, 1, result.Value("c3"))

	// update
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int64(3)),
	}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "c1_value", result.Value("c1"))
	assert.EqualValues(t, 3, result.Value("c2"))
	assert.EqualValues(t, 1, result.Value("c3"))
}

func TestPrimaryKeyOrder_test2(t *testing.T) {
	tableName := testPrimaryKeyTableName2
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{
		table.NewColumn("c3", int64(1)),
		table.NewColumn("c1", "c1_value"),
	}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", int64(2)),
	}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// get
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "c1_value", result.Value("c1"))
	assert.EqualValues(t, 2, result.Value("c2"))
	assert.EqualValues(t, 1, result.Value("c3"))

	// update
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int64(3)),
	}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "c1_value", result.Value("c1"))
	assert.EqualValues(t, 3, result.Value("c2"))
	assert.EqualValues(t, 1, result.Value("c3"))
}

func TestPrimaryKeyOrder_test3(t *testing.T) {
	tableName := testPrimaryKeyTableName3
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{
		table.NewColumn("c2", int64(1)),
		table.NewColumn("c3", int64(2)),
	}
	mutateColumns := []*table.Column{
		table.NewColumn("c1", "c1_value"),
	}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// get
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "c1_value", result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, 2, result.Value("c3"))

	// update
	mutateColumns = []*table.Column{
		table.NewColumn("c1", "c1_value_update"),
	}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "c1_value_update", result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, 2, result.Value("c3"))
}

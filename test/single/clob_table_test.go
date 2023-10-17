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

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testClobTableName       = "test_clob"
	testClobCreateStatement = "create table if not exists test_clob(c1 int(12) primary key, c2 int(12) default null, c3 text not null);"
)

func TestClob_insert(t *testing.T) {
	tableName := testClobTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c3", "abc")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. get
	res, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, "abc", res.Value("c3"))
}

func TestClob_delete(t *testing.T) {
	tableName := testClobTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c3", "abc")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. delete
	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 3. get
	res, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, nil, res.Value("c3"))
}

func TestClob_update(t *testing.T) {
	tableName := testClobTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c3", "abc")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. update
	mutateColumns = []*table.Column{table.NewColumn("c3", "efg")}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 3. get
	res, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "efg", res.Value("c3"))
}

func TestClob_replace(t *testing.T) {
	tableName := testClobTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c3", "abc")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. replace
	mutateColumns = []*table.Column{table.NewColumn("c3", "efg")}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	// 3. get
	res, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, "efg", res.Value("c3"))
}

func TestClob_insertUp(t *testing.T) {
	tableName := testClobTableName
	defer test.DeleteTable(tableName)

	// 1. insertUp(insert)
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c3", "abc")}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. insertUp(update)
	mutateColumns = []*table.Column{table.NewColumn("c3", "efg")}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 3. get
	res, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, "efg", res.Value("c3"))
}

func TestClob_append(t *testing.T) {
	tableName := testClobTableName
	defer test.DeleteTable(tableName)

	// 1. insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c3", "abc")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// 2. append, not expired, success
	mutateColumns = []*table.Column{table.NewColumn("c3", "efg")}
	res, err := cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
		option.WithReturnAffectedEntity(true),
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	assert.EqualValues(t, "abcefg", res.Value("c3"))

	// 3. get
	res, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, res.Value("c1"))
	assert.EqualValues(t, "abcefg", res.Value("c3"))
}

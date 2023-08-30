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

package current_timestamp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testStoredGenColumnTableName                = "test_stored_gen_column"
	testStoredGenColumnTableNameCreateStatement = "create table if not exists `test_stored_gen_column`(`c1` int(12) not null, `c2` varchar(20), `c3` varchar(20), `g` varchar(30) generated always as (concat(`c2`,`c3`)) stored, primary key (`c1`)) partition by key(c1) partitions 16;"
)

func TestGenerateColumn_stored(t *testing.T) {
	tableName := testStoredGenColumnTableName
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", "1"),
		table.NewColumn("c3", "1"),
	}
	affectRows, err := cli.Insert(
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
	assert.EqualValues(t, 0, result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))
	assert.EqualValues(t, "1", result.Value("c3"))
	assert.EqualValues(t, "11", result.Value("g"))

	// update
	mutateColumns = []*table.Column{
		table.NewColumn("c3", "2"),
	}
	affectRows, err = cli.Update(
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
	assert.EqualValues(t, 0, result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))
	assert.EqualValues(t, "2", result.Value("c3"))
	assert.EqualValues(t, "12", result.Value("g"))

	// insertUp insert
	rowKey = []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "1"),
		table.NewColumn("c3", "1"),
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
	assert.EqualValues(t, 1, result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))
	assert.EqualValues(t, "1", result.Value("c3"))
	assert.EqualValues(t, "11", result.Value("g"))

	// insertUp update
	mutateColumns = []*table.Column{table.NewColumn("c3", "2")}
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
	assert.EqualValues(t, 1, result.Value("c1"))
	assert.EqualValues(t, "1", result.Value("c2"))
	assert.EqualValues(t, "2", result.Value("c3"))
	assert.EqualValues(t, "12", result.Value("g"))

	// delete
	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

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
	testVirtualGenColumnTableName                = "test_virtual_gen_column"
	testVirtualGenColumnTableNameCreateStatement = "create table if not exists `test_virtual_gen_column`(" +
		"`c1` int(12) not null, " +
		"`c2` varchar(20), " +
		"`c3` varchar(20), " +
		"`c4` timestamp default current_timestamp on update current_timestamp, " +
		"`g` varchar(30) generated always as (concat(`c2`,`c3`)), " +
		"primary key (`c1`)) partition by key(c1) partitions 16;"

	testVirtualGenColumnWithGlobalIndexTableName                = "test_virtual_gen_column_with_g_index"
	testVirtualGenColumnWithGlobalIndexTableNameCreateStatement = "create table if not exists `test_virtual_gen_column_with_g_index`(" +
		"`id` int(11) NOT NULL," +
		"`num1` int(11) DEFAULT NULL," +
		"`num2` int(11) DEFAULT NULL," +
		"`num3` int(11) NOT NULL," +
		"`num4` int(11) NOT NULL AUTO_INCREMENT," +
		"`num5` int(11) GENERATED ALWAYS AS ((`id` + 400)) VIRTUAL," +
		"`description` varchar(255) DEFAULT NULL," +
		"`gmt_modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"`name` varchar(255) DEFAULT NULL," +
		"`part` int(11) NOT NULL," +
		"PRIMARY KEY (`num3`, `id`, `part`)," +
		"UNIQUE KEY `index_g_u` (`part`, `name`, `num2`) BLOCK_SIZE 16384 GLOBAL, " +
		"KEY `index_g` (`part`, `name`, `num1`) BLOCK_SIZE 16384 GLOBAL);"
)

func TestGenerateColumn_virtual(t *testing.T) {
	tableName := testVirtualGenColumnTableName
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
}

func TestGenerateColumn_virtual_append(t *testing.T) {
	tableName := testVirtualGenColumnTableName
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", "1"),
		table.NewColumn("c3", "1"),
	}
	res, err := cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

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
}

func TestGenerateColumn_virtual_with_index(t *testing.T) {
	tableName := testVirtualGenColumnWithGlobalIndexTableName
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{
		table.NewColumn("num3", int32(0)),
		table.NewColumn("part", int32(0)),
		table.NewColumn("id", int32(0)),
	}
	mutateColumns := []*table.Column{
		table.NewColumn("description", "description"),
		table.NewColumn("name", "name"),
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
	assert.EqualValues(t, 0, result.Value("id"))
	assert.EqualValues(t, nil, result.Value("num1"))
	assert.EqualValues(t, nil, result.Value("num2"))
	assert.EqualValues(t, 0, result.Value("num3"))
	assert.EqualValues(t, 1, result.Value("num4"))
	assert.EqualValues(t, 400, result.Value("num5"))
	assert.EqualValues(t, "description", result.Value("description"))
	assert.EqualValues(t, "name", result.Value("name"))
	assert.EqualValues(t, 0, result.Value("part"))

	// update
	mutateColumns = []*table.Column{
		table.NewColumn("description", "description1"),
		table.NewColumn("name", "name1"),
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
	assert.EqualValues(t, 0, result.Value("id"))
	assert.EqualValues(t, nil, result.Value("num1"))
	assert.EqualValues(t, nil, result.Value("num2"))
	assert.EqualValues(t, 0, result.Value("num3"))
	assert.EqualValues(t, 1, result.Value("num4"))
	assert.EqualValues(t, 400, result.Value("num5"))
	assert.EqualValues(t, "description1", result.Value("description"))
	assert.EqualValues(t, "name1", result.Value("name"))
	assert.EqualValues(t, 0, result.Value("part"))

	// delete
	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// insertUp-insert
	mutateColumns = []*table.Column{
		table.NewColumn("description", "description"),
		table.NewColumn("name", "name"),
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
	assert.EqualValues(t, 0, result.Value("id"))
	assert.EqualValues(t, nil, result.Value("num1"))
	assert.EqualValues(t, nil, result.Value("num2"))
	assert.EqualValues(t, 0, result.Value("num3"))
	assert.EqualValues(t, 2, result.Value("num4"))
	assert.EqualValues(t, 400, result.Value("num5"))
	assert.EqualValues(t, "description", result.Value("description"))
	assert.EqualValues(t, "name", result.Value("name"))
	assert.EqualValues(t, 0, result.Value("part"))

	// insertUp-update
	mutateColumns = []*table.Column{
		table.NewColumn("description", "description1"),
		table.NewColumn("name", "name1"),
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
	assert.EqualValues(t, 0, result.Value("id"))
	assert.EqualValues(t, nil, result.Value("num1"))
	assert.EqualValues(t, nil, result.Value("num2"))
	assert.EqualValues(t, 0, result.Value("num3"))
	assert.EqualValues(t, 2, result.Value("num4"))
	assert.EqualValues(t, 400, result.Value("num5"))
	assert.EqualValues(t, "description1", result.Value("description"))
	assert.EqualValues(t, "name1", result.Value("name"))
	assert.EqualValues(t, 0, result.Value("part"))
}

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

package nullable

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	nullableTableName       = "nullableTable"
	nullableCreateStatement = "create table if not exists nullableTable(" +
		"c1 bigint(20) not null, " +
		"c2 varchar(20) not null, " +
		"c3 varchar(20) not null default 'hello', " +
		"c4 varchar(20) default 'hello', " +
		"primary key (c1)) partition by hash(c1) partitions 15;"
)

func TestNullable_Insert(t *testing.T) {
	tableName := nullableTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}
	// c2 not null but not fill, expect report error
	mutateColumns := []*table.Column{
		table.NewColumn("c3", "a"),
		table.NewColumn("c4", "b"),
	}
	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString := err.Error()
	errMsg := "errCode:-4227"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c2 not null but fill null, expect report error
	mutateColumns = []*table.Column{
		table.NewColumn("c2", nil),
	}
	_, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString = err.Error()
	errMsg = "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c3 not null but fill null, expect report error
	mutateColumns = []*table.Column{
		table.NewColumn("c3", nil),
	}
	_, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString = err.Error()
	errMsg = "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c3 not null but has default value, expect success
	rowKey = []*table.Column{
		table.NewColumn("c1", int64(2)),
	}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
	}
	affectedRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	rowKey = []*table.Column{
		table.NewColumn("c1", int64(3)),
	}
	// c4 has default value, expect success
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c3", "a"),
	}
	affectedRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)
}

func TestNullable_Update(t *testing.T) {
	tableName := nullableTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", "a"),
	}
	affectedRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	// c2 not null, can not update to null
	mutateColumns = []*table.Column{
		table.NewColumn("c2", nil),
	}
	affectedRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString := err.Error()
	errMsg := "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c3 not null, can not update to null
	mutateColumns = []*table.Column{
		table.NewColumn("c3", nil),
	}
	affectedRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString = err.Error()
	errMsg = "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c4 nullable, can update to null
	mutateColumns = []*table.Column{
		table.NewColumn("c4", nil),
	}
	affectedRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)
}

func TestNullable_InsertUp_Insert(t *testing.T) {
	tableName := nullableTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}
	// c2 not null but not fill, expect report error
	mutateColumns := []*table.Column{
		table.NewColumn("c3", "a"),
		table.NewColumn("c4", "b"),
	}
	_, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString := err.Error()
	errMsg := "errCode:-4227"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c3 not null but has default value, expect success
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
	}
	affectedRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	// c3 not null but fill null, expect error
	rowKey = []*table.Column{
		table.NewColumn("c1", int64(2)),
	}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c3", nil),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString = err.Error()
	errMsg = "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c4 has default value, expect success
	rowKey = []*table.Column{
		table.NewColumn("c1", int64(3)),
	}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c3", "a"),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	// c4 has nullable, expect success
	rowKey = []*table.Column{
		table.NewColumn("c1", int64(4)),
	}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c4", nil),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)
}

func TestNullable_InsertUp_Update(t *testing.T) {
	tableName := nullableTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c3", "a"),
		table.NewColumn("c4", "a"),
	}
	affectedRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	// c2 not null can not update to null
	mutateColumns = []*table.Column{
		table.NewColumn("c2", nil),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString := err.Error()
	errMsg := "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c3 not null can not update to null
	mutateColumns = []*table.Column{
		table.NewColumn("c3", nil),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString = err.Error()
	errMsg = "errCode:-4235"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c4 nullable can update to null， but c2 not has default value, expect report error
	mutateColumns = []*table.Column{
		table.NewColumn("c4", nil),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	errString = err.Error()
	errMsg = "errCode:-4227"
	assert.EqualValues(t, true, strings.Contains(errString, errMsg))

	// c4 nullable can update to null
	mutateColumns = []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c4", nil),
	}
	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)
}

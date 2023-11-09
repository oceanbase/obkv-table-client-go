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
	autoIncRowkeyNilFillTableTableName       = "autoIncRowkeyNilFillTable"
	autoIncRowkeyNilFillTableCreateStatement = "create table if not exists autoIncRowkeyNilFillTable(" +
		"c1 bigint(20) not null auto_increment, " +
		"c2 bigint(20) not null, " +
		"c3 varchar(20) default 'hello', " +
		"primary key (`c1`, `c2`)) partition by hash(c2) partitions 15;"
)

func TestAuto_IncRowkeyFillNil(t *testing.T) {
	tableName := autoIncRowkeyNilFillTableTableName
	defer test.DeleteTable(tableName)

	// test insert.
	rowKey := []*table.Column{table.NewColumn("c1", nil), table.NewColumn("c2", int64(1))}
	affectedRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), affectedRows)

	rowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(1))}
	selectColumns := []string{"c1", "c2", "c3"}
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
}

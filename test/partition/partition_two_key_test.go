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

package partition

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	keyBigIntL2TableName       = "keyBigIntL2"
	keyBigIntL2CreateStatement = "create table if not exists keyBigIntL2(`c1` bigint(20) not null, c2 bigint(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(`c1`) subpartition by key(`c2`) subpartitions 4 partitions 16;"

	keyVarcharL2TableName       = "keyVarcharL2"
	keyVarcharL2CreateStatement = "create table if not exists keyVarcharL2(`c1` varchar(20) not null, `c2` varchar(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(`c1`) subpartition by key(`c2`) subpartitions 4 partitions 16;"

	keyVarBinaryL2TableName       = "keyVarBinaryL2"
	keyVarBinaryL2CreateStatement = "create table if not exists keyVarBinaryL2(`c1` varchar(20) not null, `c2` varchar(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(`c1`) subpartition by key(`c2`) subpartitions 4 partitions 16;"
)

func TestKeyPartitionL2_BIGINT(t *testing.T) {
	tableName := keyBigIntL2TableName
	defer test.DeleteTable(tableName)
	recordCount := 10

	// insert by sql
	for i := -recordCount; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s values(%d, %d);", tableName, i, i)
		test.InsertTable(insertStatement)
	}

	// get by obkv
	for i := -recordCount; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i)), table.NewColumn("c2", int64(i))}
		res, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			nil,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, i, res.Value("c1"))
		assert.EqualValues(t, i, res.Value("c2"))
	}
}

func TestKeyPartitionL2_VARCHAR(t *testing.T) {
	tableName := keyVarcharL2TableName
	defer test.DeleteTable(tableName)
	recordCount := 10

	// insert by sql
	for i := -recordCount; i < recordCount; i++ {
		v := "key" + strconv.Itoa(i)
		insertStatement := fmt.Sprintf("insert into %s values('%s', '%s');", tableName, v, v)
		test.InsertTable(insertStatement)
	}

	// get by obkv
	for i := -recordCount; i < recordCount; i++ {
		v := "key" + strconv.Itoa(i)
		rowKey := []*table.Column{table.NewColumn("c1", v), table.NewColumn("c2", v)}
		res, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			nil,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, v, res.Value("c1"))
		assert.EqualValues(t, v, res.Value("c2"))
	}
}

func TestKeyVarcharL2_VARBINARY(t *testing.T) {
	tableName := keyVarBinaryL2TableName
	defer test.DeleteTable(tableName)
	recordCount := 10

	// insert by sql
	for i := -recordCount; i < recordCount; i++ {
		v := "key" + strconv.Itoa(i)
		insertStatement := fmt.Sprintf("insert into %s values('%s', '%s');", tableName, v, v)
		test.InsertTable(insertStatement)
	}

	// get by obkv
	for i := -recordCount; i < recordCount; i++ {
		v := "key" + strconv.Itoa(i)
		rowKey := []*table.Column{table.NewColumn("c1", v), table.NewColumn("c2", v)}
		res, err := cli.Get(
			context.TODO(),
			tableName,
			rowKey,
			nil,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, v, res.Value("c1"))
		assert.EqualValues(t, v, res.Value("c2"))
	}
}

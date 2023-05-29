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
	keyBigIntL1TableName       = "keyBigIntL1"
	keyBigIntL1CreateStatement = "create table if not exists keyBigIntL1(`c1` bigint(20) not null, c2 bigint(20) not null, " +
		"primary key (`c1`)) partition by key(c1) partitions 15;"

	keyVarcharL1TableName       = "keyVarcharL1"
	keyVarcharL1CreateStatement = "create table if not exists keyVarcharL1(`c1` varchar(20) not null, c2 varchar(20) not null, " +
		"primary key (`c1`)) partition by key(c1) partitions 15;"

	keyVarBinaryL1TableName       = "keyVarBinaryL1"
	keyVarBinaryL1CreateStatement = "create table if not exists keyVarBinaryL1(`c1` varbinary(20) not null, c2 varbinary(20) not null, " +
		"primary key (`c1`)) partition by key(c1) partitions 15;"
)

func TestKeyPartitionL1_BIGINT(t *testing.T) {
	tableName := keyBigIntL1TableName
	defer test.DeleteTable(tableName)
	recordCount := 10

	// insert by sql
	for i := -recordCount; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s values(%d, %d);", tableName, i, i)
		test.InsertTable(insertStatement)
	}

	// get by obkv
	for i := -recordCount; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
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

func TestKeyPartitionL1_VARCHAR(t *testing.T) {
	tableName := keyVarcharL1TableName
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
		rowKey := []*table.Column{table.NewColumn("c1", v)}
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

func TestKeyVarcharL1_VARBINARY(t *testing.T) {
	tableName := keyVarBinaryL1TableName
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
		rowKey := []*table.Column{table.NewColumn("c1", v)}
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

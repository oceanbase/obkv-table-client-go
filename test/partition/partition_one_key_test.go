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
	keyBigintL1TableName       = "keyBigintL1"
	keyBigintL1CreateStatement = "create table if not exists keyBigintL1(`c1` bigint(20) not null, c2 bigint(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(c1) partitions 15;"
	keyMultiBigintL1TableName       = "keyMultiBigintL1"
	keyMultiBigintL1CreateStatement = "create table if not exists keyMultiBigintL1(`c1` bigint(20) not null, c2 bigint(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(c1, c2) partitions 15;"

	keyVarcharL1TableName       = "keyVarcharL1"
	keyVarcharL1CreateStatement = "create table if not exists keyVarcharL1(`c1` varchar(20) not null, c2 varchar(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(c1) partitions 15;"
	keyMultiVarcharL1TableName       = "keyMultiVarcharL1"
	keyMultiVarcharL1CreateStatement = "create table if not exists keyMultiVarcharL1(`c1` varchar(20) not null, c2 varchar(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(c1, c2) partitions 15;"

	keyVarBinaryL1TableName       = "keyVarBinaryL1"
	keyVarBinaryL1CreateStatement = "create table if not exists keyVarBinaryL1(`c1` varbinary(20) not null, c2 varbinary(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(c1) partitions 15;"
	keyMultiVarBinaryL1TableName       = "keyMultiVarBinaryL1"
	keyMultiVarBinaryL1CreateStatement = "create table if not exists keyMultiVarBinaryL1(`c1` varbinary(20) not null, c2 varbinary(20) not null, " +
		"primary key (`c1`, `c2`)) partition by key(c1, c2) partitions 15;"
)

var l1KeyIntegerTableNames = []string{
	keyBigintL1TableName,
	keyMultiBigintL1TableName,
}

var l1KeyStringTableNames = []string{
	keyVarcharL1TableName,
	keyMultiVarcharL1TableName,
	keyVarBinaryL1TableName,
	keyMultiVarBinaryL1TableName,
}

func TestKeyPartitionL1_INTEGER(t *testing.T) {
	defer func() {
		test.DeleteTables(l1KeyIntegerTableNames)
	}()
	for _, tableName := range l1KeyIntegerTableNames {
		// insert by sql
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
			insertStatement := fmt.Sprintf("insert into %s values(%d, %d);", tableName, i, i)
			test.InsertTable(insertStatement)
		}

		// get by obkv
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
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
}

func TestKeyPartitionL1_STRING(t *testing.T) {
	defer func() {
		test.DeleteTables(l1KeyStringTableNames)
	}()

	for _, tableName := range l1KeyStringTableNames {
		// insert by sql
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
			v := "key" + strconv.Itoa(i)
			insertStatement := fmt.Sprintf("insert into %s values('%s', '%s');", tableName, v, v)
			test.InsertTable(insertStatement)
		}

		// get by obkv
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
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
}

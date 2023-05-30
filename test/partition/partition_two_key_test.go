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
	keyBigintL2TableName       = "keyBigintL2"
	keyBigintL2CreateStatement = "create table if not exists keyBigintL2(`c1` bigint(20) not null, c2 bigint(20) not null, " +
		"c3 bigint(20) not null, primary key (`c1`, `c2`, `c3`)) partition by key(`c1`, `c2`) subpartition by key(`c3`) subpartitions 4 partitions 16;"
	keyMultiBigintL2TableName       = "keyMultiBigintL2"
	keyMultiBigintL2CreateStatement = "create table if not exists keyMultiBigintL2(`c1` bigint(20) not null, c2 bigint(20) not null, " +
		"c3 bigint(20) not null, primary key (`c1`, `c2`, `c3`)) partition by key(`c1`, `c2`) subpartition by key(`c3`) subpartitions 4 partitions 16;"

	keyVarcharL2TableName       = "keyVarcharL2"
	keyVarcharL2CreateStatement = "create table if not exists keyVarcharL2(`c1` varchar(20) not null, `c2` varchar(20) not null, " +
		"c3 varchar(20) not null, primary key (`c1`, `c2`, `c3`)) partition by key(`c1`, `c2`) subpartition by key(`c3`) subpartitions 4 partitions 16;"
	keyMultiVarcharL2TableName       = "keyMultiVarcharL2"
	keyMultiVarcharL2CreateStatement = "create table if not exists keyMultiVarcharL2(`c1` varchar(20) not null, `c2` varchar(20) not null, " +
		"c3 varchar(20) not null, primary key (`c1`, `c2`, `c3`)) partition by key(`c1`, `c2`) subpartition by key(`c3`) subpartitions 4 partitions 16;"

	keyVarBinaryL2TableName       = "keyVarBinaryL2"
	keyVarBinaryL2CreateStatement = "create table if not exists keyVarBinaryL2(`c1` varbinary(20) not null, `c2` varbinary(20) not null, " +
		"c3 varbinary(20) not null, primary key (`c1`, `c2`, `c3`)) partition by key(`c1`, `c2`) subpartition by key(`c3`) subpartitions 4 partitions 16;"
	keyMultiVarBinaryL2TableName       = "keyMultiVarBinaryL2"
	keyMultiVarBinaryL2CreateStatement = "create table if not exists keyMultiVarBinaryL2(`c1` varbinary(20) not null, `c2` varbinary(20) not null, " +
		"c3 varbinary(20) not null, primary key (`c1`, `c2`, `c3`)) partition by key(`c1`, `c2`) subpartition by key(`c3`) subpartitions 4 partitions 16;"
)

var l2KeyIntegerTableNames = []string{
	keyBigintL2TableName,
	keyMultiBigintL2TableName,
}

var l2KeyStringTableNames = []string{
	keyVarcharL2TableName,
	keyMultiVarcharL2TableName,
	keyVarBinaryL2TableName,
	keyMultiVarBinaryL2TableName,
}

func TestKeyPartitionL2_INTEGER(t *testing.T) {
	defer func() {
		test.DeleteTables(l2KeyIntegerTableNames)
	}()
	for _, tableName := range l2KeyIntegerTableNames {
		// insert by sql
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
			insertStatement := fmt.Sprintf("insert into %s values(%d, %d, %d);", tableName, i, i, i)
			test.InsertTable(insertStatement)
		}

		// get by obkv
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
			rowKey := []*table.Column{
				table.NewColumn("c1", int64(i)),
				table.NewColumn("c2", int64(i)),
				table.NewColumn("c3", int64(i)),
			}
			res, err := cli.Get(
				context.TODO(),
				tableName,
				rowKey,
				nil,
			)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, i, res.Value("c1"))
			assert.EqualValues(t, i, res.Value("c2"))
			assert.EqualValues(t, i, res.Value("c3"))
		}
	}
}

func TestKeyPartitionL2_STRING(t *testing.T) {
	defer func() {
		test.DeleteTables(l2KeyStringTableNames)
	}()

	for _, tableName := range l2KeyStringTableNames {
		// insert by sql
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
			v := "key" + strconv.Itoa(i)
			insertStatement := fmt.Sprintf("insert into %s values('%s', '%s', '%s');", tableName, v, v, v)
			test.InsertTable(insertStatement)
		}

		// get by obkv
		for i := -partitionTestRecordCount; i < partitionTestRecordCount; i++ {
			v := "key" + strconv.Itoa(i)
			rowKey := []*table.Column{
				table.NewColumn("c1", v),
				table.NewColumn("c2", v),
				table.NewColumn("c3", v),
			}
			res, err := cli.Get(
				context.TODO(),
				tableName,
				rowKey,
				nil,
			)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, v, res.Value("c1"))
			assert.EqualValues(t, v, res.Value("c2"))
			assert.EqualValues(t, v, res.Value("c3"))
		}
	}
}

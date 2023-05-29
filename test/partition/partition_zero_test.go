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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	hashBigintL0TableName       = "hashBigintL0"
	hashBigintL0CreateStatement = "create table if not exists hashBigintL0(`c1` bigint(20) not null, c2 bigint(20) not null, primary key (`c1`));"
)

func TestHashPartitionL0(t *testing.T) {
	tableName := hashBigintL0TableName
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

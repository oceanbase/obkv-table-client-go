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

package current_time

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testTimestampCurrentTimeTableName       = "test_timestamp_current_time"
	testTimestampCurrentTimeCreateStatement = "create table if not exists `test_timestamp_current_time`(`c1` int(12) not null,`c2` int(12) default 0, `c3` timestamp default current_timestamp, `c4` timestamp default current_timestamp on update current_timestamp, primary key (`c1`)) partition by key(c1) partitions 16;"
	// testTimestampCurrentTimeCreateStatement = "create table if not exists `test_timestamp_current_time`(`c1` int(12) not null,`c2` int(12) default 0, `c3` timestamp default current_timestamp, primary key (`c1`)) partition by key(c1) partitions 16;"
)

func TestCurrentTimeTimstamp(t *testing.T) {
	tableName := testTimestampCurrentTimeTableName
	defer test.DeleteTable(tableName)

	// insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// update
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// insertUp insert
	rowKey = []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// insertUp update
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
}

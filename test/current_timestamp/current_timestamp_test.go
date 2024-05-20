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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testTimestampCurrentTimeTableName       = "test_timestamp_current_time"
	testTimestampCurrentTimeCreateStatement = "create table if not exists `test_timestamp_current_time`(" +
		"`c1` int(12) not null," +
		"`c2` int(12) default 0, " +
		"`c3` timestamp default current_timestamp, " +
		"`c4` timestamp default current_timestamp on update current_timestamp, " +
		"`c5` timestamp default null on update current_timestamp, " +
		"`c6` varchar(20) default 'hello', primary key (`c1`)) " +
		"partition by key(c1) partitions 16;"
)

func TestCurrentTimestamp_common(t *testing.T) {
	tableName := testTimestampCurrentTimeTableName
	defer test.DeleteTable(tableName)

	// insert
	fmt.Println("insert")
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

	// get
	result, err := cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 0, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, result.Value("c3"), result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// update
	time.Sleep(1 * time.Second)
	fmt.Println("update")
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
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
	assert.EqualValues(t, 1, result.Value("c2"))
	fmt.Println(result.Value("c3"))
	assert.EqualValues(t, result.Value("c4"), result.Value("c5"))

	// insertUp insert
	time.Sleep(1 * time.Second)
	fmt.Println("insertUp-insert")
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

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, result.Value("c3"), result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// insertUp update
	time.Sleep(1 * time.Second)
	fmt.Println("insertUp-update")
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
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
	assert.EqualValues(t, 1, result.Value("c2"))
	fmt.Println(result.Value("c3"))
	assert.EqualValues(t, result.Value("c4"), result.Value("c5"))

	// replace insert
	time.Sleep(1 * time.Second)
	fmt.Println("replace-insert")
	rowKey = []*table.Column{table.NewColumn("c1", int32(2))}
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(0))}
	affectRows, err = cli.Replace(
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
	assert.EqualValues(t, 2, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, result.Value("c3"), result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// replace replace
	time.Sleep(1 * time.Second)
	fmt.Println("replace-replace")
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	fmt.Println(result.Value("c3"))
	fmt.Println(result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// increment insert
	time.Sleep(1 * time.Second)
	fmt.Println("increment-insert")
	rowKey = []*table.Column{table.NewColumn("c1", int32(3))}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
	}
	res, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, result.Value("c3"), result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// increment inc
	time.Sleep(1 * time.Second)
	fmt.Println("increment-inc")
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
	}
	res, err = cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, result.Value("c1"))
	assert.NotEqual(t, result.Value("c3"), result.Value("c4"))
	assert.EqualValues(t, result.Value("c4"), result.Value("c5"))

	// append insert
	time.Sleep(1 * time.Second)
	fmt.Println("append-insert")
	rowKey = []*table.Column{table.NewColumn("c1", int32(4))}
	mutateColumns = []*table.Column{
		table.NewColumn("c6", "abc"),
	}
	res, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 4, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, "abc", result.Value("c6"))
	assert.EqualValues(t, result.Value("c3"), result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// append append
	time.Sleep(1 * time.Second)
	fmt.Println("append-append")
	mutateColumns = []*table.Column{
		table.NewColumn("c6", "efg"),
	}
	res, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 4, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, "abcefg", result.Value("c6"))
	assert.NotEqual(t, result.Value("c3"), result.Value("c4"))
	assert.EqualValues(t, result.Value("c4"), result.Value("c5"))
}

func TestCurrentTimestamp_fillValue(t *testing.T) {
	tableName := testTimestampCurrentTimeTableName
	defer test.DeleteTable(tableName)

	// insert
	fmt.Println("insert")
	var now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	rowKey := []*table.Column{table.NewColumn("c1", int32(0))}
	mutateColumns := []*table.Column{
		table.NewColumn("c2", int32(0)),
		table.NewColumn("c3", table.TimeStamp(now)),
		table.NewColumn("c4", table.TimeStamp(now)),
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
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c3"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c4"))
	assert.Equal(t, nil, result.Value("c5"))

	// update
	time.Sleep(1 * time.Second)
	fmt.Println("update")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
		table.NewColumn("c3", table.TimeStamp(now)),
		table.NewColumn("c4", table.TimeStamp(now)),
		table.NewColumn("c5", table.TimeStamp(now)),
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
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c3"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c4"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c5"))

	// insertUp insert
	time.Sleep(1 * time.Second)
	fmt.Println("insertUp-insert")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	rowKey = []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(0)),
		table.NewColumn("c3", table.TimeStamp(now)),
		table.NewColumn("c4", table.TimeStamp(now)),
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
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c3"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c4"))
	assert.EqualValues(t, nil, result.Value("c5"))

	// insertUp update
	time.Sleep(1 * time.Second)
	fmt.Println("insertUp-update")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
		table.NewColumn("c3", table.TimeStamp(now)),
		table.NewColumn("c4", table.TimeStamp(now)),
		table.NewColumn("c5", table.TimeStamp(now)),
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
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c3"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c4"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c5"))

	// replace insert
	time.Sleep(1 * time.Second)
	fmt.Println("replace-insert")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	rowKey = []*table.Column{table.NewColumn("c1", int32(2))}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(0)),
		table.NewColumn("c3", table.TimeStamp(now)),
		table.NewColumn("c4", table.TimeStamp(now)),
	}
	affectRows, err = cli.Replace(
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
	assert.EqualValues(t, 2, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c3"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c4"))
	assert.EqualValues(t, nil, result.Value("c5"))

	// replace replace
	time.Sleep(1 * time.Second)
	fmt.Println("replace-replace")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
		table.NewColumn("c3", table.TimeStamp(now)),
		table.NewColumn("c4", table.TimeStamp(now)),
		table.NewColumn("c5", table.TimeStamp(now)),
	}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c3"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c4"))
	assert.EqualValues(t, table.TimeStamp(now), result.Value("c5"))

	// increment insert
	time.Sleep(1 * time.Second)
	fmt.Println("increment-insert")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	rowKey = []*table.Column{table.NewColumn("c1", int32(3))}
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
	}
	res, err := cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, result.Value("c1"))
	assert.EqualValues(t, 1, result.Value("c2"))

	// increment inc
	time.Sleep(1 * time.Second)
	fmt.Println("increment-inc")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	mutateColumns = []*table.Column{
		table.NewColumn("c2", int32(1)),
	}
	res, err = cli.Increment(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, result.Value("c1"))
	assert.EqualValues(t, 2, result.Value("c2"))

	// append insert
	time.Sleep(1 * time.Second)
	fmt.Println("append-insert")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	rowKey = []*table.Column{table.NewColumn("c1", int32(4))}
	mutateColumns = []*table.Column{
		table.NewColumn("c6", "abc"),
	}
	res, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 4, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, "abc", result.Value("c6"))
	assert.EqualValues(t, nil, result.Value("c5"))

	// append append
	time.Sleep(1 * time.Second)
	fmt.Println("append-append")
	now = time.Now().Local().Truncate(time.Second)
	fmt.Println(now)
	mutateColumns = []*table.Column{
		table.NewColumn("c6", "efg"),
	}
	res, err = cli.Append(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, res.AffectedRows())

	// get
	result, err = cli.Get(
		context.TODO(),
		tableName,
		rowKey,
		nil,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 4, result.Value("c1"))
	assert.EqualValues(t, 0, result.Value("c2"))
	assert.EqualValues(t, "abcefg", result.Value("c6"))
}

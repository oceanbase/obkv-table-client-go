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

package aggregate

import (
	"context"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	aggregateTableName            = "aggregateTable"
	aggregateTableCreateStatement = "create table if not exists aggregateTable(`c1` bigint(20) not null, c2 bigint(20) not null, c3 varchar(20) default 'hello', c4 double default null, c5 int default null, primary key (`c1`, `c2`)) partition by hash(c1) partitions 15;"
)

func prepareAggRecord(recordCount int) {
	for i := 0; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s(c1, c2, c4) values(1, %d, %f);", aggregateTableName, i, (float64)(i))
		test.InsertTable(insertStatement)
	}
}

func addAggRecord(recordCount int) {
	for i := 0; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s(c1, c2, c4, c5) values(1, %d, %f, %d);", aggregateTableName, i+10, (float64)(i), i)
		test.InsertTable(insertStatement)
	}
}

func TestAggregate(t *testing.T) {
	tableName := aggregateTableName
	defer test.DeleteTable(tableName)

	// aggregate empty table.
	startRowKey := []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(9))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	aggExecutor := cli.NewAggExecutor(tableName, keyRanges).Min("c2")

	res, err := aggExecutor.Execute(context.TODO())

	assert.Equal(t, nil, err)
	assert.Equal(t, nil, res.Value("min(c2)"))

	recordCount := 10
	prepareAggRecord(recordCount)

	// normal aggregate
	startRowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(0))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(9))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	aggExecutor = cli.NewAggExecutor(tableName, keyRanges).
		Min("c2").
		Max("c2").
		Count().
		Sum("c2").
		Avg("c2").
		Min("c4").
		Max("c4").
		Sum("c4").
		Avg("c4").
		Min("c5").
		Max("c5").
		Sum("c5").
		Avg("c5")

	res, err = aggExecutor.Execute(context.TODO())

	assert.Equal(t, nil, err)
	assert.Equal(t, int64(0), res.Value("min(c2)").(int64))
	assert.Equal(t, int64(9), res.Value("max(c2)").(int64))
	assert.Equal(t, int64(45), res.Value("sum(c2)").(int64))
	assert.Equal(t, 4.5, res.Value("avg(c2)").(float64))
	assert.Equal(t, int64(10), res.Value("count(*)").(int64))

	assert.Equal(t, 0.0, res.Value("min(c4)").(float64))
	assert.Equal(t, 9.0, res.Value("max(c4)").(float64))
	assert.Equal(t, 45.0, res.Value("sum(c4)").(float64))
	assert.Equal(t, 4.5, res.Value("avg(c4)").(float64))

	// test null value
	assert.Equal(t, nil, res.Value("min(c5)"))
	assert.Equal(t, nil, res.Value("max(c5)"))
	assert.Equal(t, nil, res.Value("sum(c5)"))
	assert.Equal(t, nil, res.Value("avg(c5)"))

	// test defense for multiple parts aggregation
	startRowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(0))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(2)), table.NewColumn("c2", int64(9))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	aggExecutor = cli.NewAggExecutor(tableName, keyRanges).Min("c2")

	res, err = aggExecutor.Execute(context.TODO())

	assert.Equal(t, nil, res)
	assert.NotNil(t, err)

	// test column has null value
	addAggRecord(recordCount)
	startRowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(0))}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(1)), table.NewColumn("c2", int64(19))}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}

	aggExecutor = cli.NewAggExecutor(tableName, keyRanges).
		Min("c5").
		Max("c5").
		Sum("c5").
		Avg("c5")

	res, err = aggExecutor.Execute(context.TODO())

	assert.Equal(t, nil, err)
	assert.Equal(t, int32(0), res.Value("min(c5)").(int32))
	assert.Equal(t, int32(9), res.Value("max(c5)").(int32))
	assert.Equal(t, int64(45), res.Value("sum(c5)").(int64))
	assert.Equal(t, 4.5, res.Value("avg(c5)").(float64))

	// invalid aggregation
	aggExecutor = cli.NewAggExecutor(tableName, keyRanges).Sum("c3")

	res, err = aggExecutor.Execute(context.TODO())

	assert.Equal(t, nil, res)
	assert.NotNil(t, err)

	// aggregate not exist column
	aggExecutor = cli.NewAggExecutor(tableName, keyRanges).Min("c9")

	res, err = aggExecutor.Execute(context.TODO())

	assert.Equal(t, nil, res)
	assert.NotNil(t, err)
}

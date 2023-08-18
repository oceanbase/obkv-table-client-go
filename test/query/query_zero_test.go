/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
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

package query

import (
	"context"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/client/filter"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	queryZeroTableName            = "queryZeroTable"
	queryZeroTableCreateStatement = "create table if not exists queryZeroTable(`c1` bigint(20) not null, c2 bigint(20) not null, c3 varchar(20) default 'hello', index i1(`c1`, `c3`) local, primary key (`c1`, `c2`));"
)

func prepareZeroRecord(recordCount int) {
	for i := 0; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s(c1, c2) values(%d, %d);", queryZeroTableName, i, i)
		test.InsertTable(insertStatement)
	}
}

func prepareZeroSinglePartitionRecord(pk int, recordCount int) {
	for i := 0; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s(c1, c2, c3) values(%d, %d, %s);", queryZeroTableName, pk, i, fmt.Sprintf("'hello_%d'", i))
		test.InsertTable(insertStatement)
	}
}

func TestQueryZeroSimple(t *testing.T) {
	tableName := queryZeroTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	prepareZeroRecord(recordCount)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < recordCount; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	startRowKey = []*table.Column{table.NewColumn("c1", int64(5)), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(10)), table.NewColumn("c2", table.Max)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < 5; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	// test range pair inclusive
	startRowKey = []*table.Column{table.NewColumn("c1", int64(9)), table.NewColumn("c2", table.Max)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(11)), table.NewColumn("c2", table.Min)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey, false)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.Equal(t, nil, err)
		assert.Equal(t, nil, res)
	}
	assert.Equal(t, nil, err)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(8)), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(9)), table.NewColumn("c2", table.Min)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey, true, false)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	res, err = resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.EqualValues(t, 8, res.Value("c1"))
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}
	assert.Equal(t, nil, err)
}

func TestQueryZeroSinglePartition(t *testing.T) {
	tableName := queryZeroTableName
	defer test.DeleteTable(tableName)

	pk := 0
	recordCount := 10
	prepareZeroSinglePartitionRecord(pk, recordCount)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	i := 0
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.EqualValues(t, "hello", res.Value("c3").(string)[:5])
		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, recordCount, i)

	// wrong range
	// Max - Min
	startRowKey = []*table.Column{table.NewColumn("c1", table.Max), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", table.Min), table.NewColumn("c2", table.Max)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	res, err = resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.Equal(t, nil, err)
		assert.Equal(t, nil, res)
	}
	// assert.Equal(t, nil, err) ODP will return err now

	// test NextBatch()
	startRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Max)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
	)
	assert.Equal(t, nil, err)
	batchRes, err := resSet.NextBatch()
	for ; batchRes != nil && err == nil; batchRes, err = resSet.NextBatch() {
		assert.Equal(t, nil, err)
		for i := 0; i < len(batchRes); i++ {
			assert.Equal(t, int64(pk), batchRes[i].Value("c1"))
			assert.Equal(t, int64(i), batchRes[i].Value("c2"))
			assert.Equal(t, "hello", batchRes[i].Value("c3").(string)[:5])
		}
	}
	assert.Equal(t, nil, err)

	// test batchSize
	batchSize := 1
	startRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Max)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
	)
	assert.Equal(t, nil, err)
	i = 0
	res, err = resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.Equal(t, nil, err)
		assert.EqualValues(t, "hello", res.Value("c3").(string)[:5])
		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, recordCount, i)

	// test filter
	lt6 := filter.CompareVal(filter.LessThan, "c2", int64(6))
	gt4 := filter.CompareVal(filter.GreaterThan, "c2", int64(4))
	filterList := filter.AndList(lt6, gt4)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c2", table.Max)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryFilter(filterList),
	)
	assert.Equal(t, nil, err)
	i = 0
	res, err = resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.Equal(t, nil, err)
		assert.EqualValues(t, "hello", res.Value("c3").(string)[:5])
		assert.Equal(t, int64(5), res.Value("c2"))
		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, i)

	// test local index
	startRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c3", "hello_3")}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(pk)), table.NewColumn("c3", "hello_5")}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryIndexName("i1"),
	)
	assert.Equal(t, nil, err)
	i = 0
	res, err = resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.EqualValues(t, int64(pk), res.Value("c1"))
		assert.EqualValues(t, "hello", res.Value("c3").(string)[:5])
		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 3, i)

	startRowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c3", "not exist")}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c3", "not exist")}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryIndexName("i1"),
	)
	assert.Equal(t, nil, err)
	res, err = resSet.Next()
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, res)
}

func TestQueryZeroBatchSize(t *testing.T) {
	tableName := queryZeroTableName
	defer test.DeleteTable(tableName)

	recordCount := 50
	batchSize := 1
	prepareZeroRecord(recordCount)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < recordCount; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	// test NextBatch
	startRowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
	)
	assert.Equal(t, nil, err)
	res, err := resSet.NextBatch()
	for ; res != nil && err == nil; res, err = resSet.NextBatch() {
		for i := 0; i < len(res); i++ {
			assert.EqualValues(t, res[i].Value("c1"), res[i].Value("c2"))
			assert.EqualValues(t, "hello", res[i].Value("c3"))
			assert.EqualValues(t, res[i].Values(), []interface{}{res[i].Value("c1"), res[i].Value("c1"), "hello"})
		}
	}
	assert.Equal(t, nil, err)
}

func TestQueryZeroIndex(t *testing.T) {
	tableName := queryZeroTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	batchSize := 1
	prepareZeroRecord(recordCount)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c3", "hello")}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c3", "hello")}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryIndexName("i1"),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < recordCount; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}

	startRowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c3", "not exist")}
	endRowKey = []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c3", "not exist")}
	keyRanges = []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryIndexName("i1"),
	)
	assert.Equal(t, nil, err)
	res, err := resSet.Next()
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, res)
}

func TestQueryZeroMixture(t *testing.T) {
	tableName := queryZeroTableName
	defer test.DeleteTable(tableName)

	recordCount := 10
	batchSize := 1
	limit := 5
	scanOrder := table.Reverse
	offset := 3
	prepareZeroRecord(recordCount)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c3", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c3", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryLimit(limit),
		option.WithQueryScanOrder(scanOrder),
		option.WithQueryOffset(offset),
	)
	assert.Equal(t, nil, err)
	for i := 0; i < limit; i++ {
		res, err := resSet.Next()
		assert.Equal(t, nil, err)
		assert.Equal(t, int64(recordCount-offset-i-1), res.Value("c1"))
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
	}
}

func TestQueryZeroFilter(t *testing.T) {
	tableName := queryZeroTableName
	defer test.DeleteTable(tableName)

	recordCount := 50
	batchSize := 1
	prepareZeroRecord(recordCount)

	// test Next()
	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	lt30 := filter.CompareVal(filter.LessThan, "c2", int64(30))
	gt10 := filter.CompareVal(filter.GreaterThan, "c2", int64(10))
	filterList := filter.AndList(lt30, gt10)
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryBatchSize(batchSize),
		option.WithQueryFilter(filterList),
	)
	assert.Equal(t, nil, err)
	i := 0
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.EqualValues(t, res.Value("c1"), res.Value("c2"))
		assert.EqualValues(t, "hello", res.Value("c3"))
		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 19, i)
}

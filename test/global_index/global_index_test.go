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
package global_index

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testGlabalIndexHashTableName      = "test_global_index_hash"
	testGlabalIndexKeyTableName       = "test_global_index_key"
	testGlabalIndexCreateHashTable    = "CREATE TABLE IF NOT EXISTS `test_global_index_hash` (  `c1` int(11) NOT NULL,  `c2` int(11) DEFAULT NULL,  `c3` int(11) DEFAULT NULL,  PRIMARY KEY (`c1`),  KEY `idx` (`c2`) GLOBAL partition by hash(`c2`)(partition `p0`,partition `p1`,partition `p2`,partition `p3`,partition `p4`)) partition by hash(`c1`)(partition `p0`,partition `p1`,partition `p2`,partition `p3`,partition `p4`)"
	testGlabalIndexCreateKeyTable     = " CREATE TABLE IF NOT EXISTS `test_global_index_key` (`c1` int(11) NOT NULL,`c2` int(11) DEFAULT NULL,  `c3` int(11) DEFAULT NULL,  PRIMARY KEY (`c1`),  KEY `idx` (`c2`)  GLOBAL partition by key(`c2`)(partition `p0`,partition `p1`,partition `p2`,partition `p3`,partition `p4`))  partition by key(`c1`)(partition `p0`,partition `p1`,partition `p2`,partition `p3`,partition `p4`)"
	testGlobalIndexNoPart             = "test_global_index_no_part"
	testGlobalAllNoPart               = "test_global_all_no_part"
	testGlobalPrimaryNoPart           = "test_global_primary_no_part"
	testGlobalIndexNoPartCreateStat   = "CREATE TABLE IF NOT EXISTS `test_global_index_no_part` (`c1` int(11) NOT NULL, `c2` int(11) DEFAULT NULL, `c3` int(11) DEFAULT NULL,  PRIMARY KEY (`c1`),  KEY `idx` (`c2`) GLOBAL) partition by hash(`c1`) (partition p0, partition p1, partition p2, partition p3, partition p4, partition p6);"
	testGlobalAllNoPartCreateStat     = "CREATE TABLE IF NOT EXISTS `test_global_all_no_part` (`c1` int(11) NOT NULL, `c2` int(11) DEFAULT NULL, `c3` int(11) DEFAULT NULL, PRIMARY KEY (`c1`), KEY `idx` (`c2`) GLOBAL);"
	testGlobalPrimaryNoPartCreateStat = "CREATE TABLE IF NOT EXISTS `test_global_primary_no_part` (`c1` int(11) NOT NULL, `c2` int(11) DEFAULT NULL, `c3` int(11) DEFAULT NULL, PRIMARY KEY (`c1`), KEY `idx` (`c2`) GLOBAL partition by hash(`c2`) (partition p0, partition p1, partition p2, partition p3, partition p4));"
	testGlobalTwoLevelPart            = "test_global_two_level_part"
	testGlobalTwoLevelPartCreateStat  = "CREATE TABLE IF NOT EXISTS `test_global_two_level_part` (`c1` int NOT NULL, `c2` int NOT NULL, `c3` int NOT NULL, `c4` int NOT NULL, `c5` varchar(20) default NULL,PRIMARY KEY (`c1`, `c2`), KEY `idx` (`c3`,`c4`) GLOBAL partition by hash(`c3`) partitions 8) partition by hash(`c1`) subpartition by hash(`c2`) subpartitions 4 partitions 16;"
	testGlobalUniqueIndex             = "testGlobalUniqueIndex"
	testGlobalUniqueIndexCreateStat   = "CREATE TABLE IF NOT EXISTS `testGlobalUniqueIndex` (`C1` int NOT NULL,`C2` int NOT NULL,`C3` int NOT NULL,`C4` int NOT NULL,`C5` varchar(20) default NULL,PRIMARY KEY (`C1`),UNIQUE KEY `idx` (`C3`) GLOBAL partition by hash(`C3`) partitions 8) partition by key(`C1`) partitions 16;"
	testGlobalIndexWithTTL            = "test_global_index_with_ttl"
	testGlobalIndexWithTTLCreateStat  = "CREATE TABLE IF NOT EXISTS `test_global_index_with_ttl` (`c1` varchar(20) NOT NULL,`c2` bigint NOT NULL,`c3` bigint DEFAULT NULL,`c4` bigint DEFAULT NULL,`expired_ts` timestamp(6),PRIMARY KEY (`c1`, `c2`),KEY `idx`(`c1`, `c4`) local,KEY `idx2`(`c3`) global partition by hash(`c3`) partitions 4) TTL(expired_ts + INTERVAL 0 SECOND) partition by key(`c1`) partitions 4;"
)

func checkIndexTableRow(t *testing.T, indexTableName string, rowKey []*table.Column, mutateColumns []*table.Column, affectedCount int) {
	// check index table row
	checkIndexSql := fmt.Sprintf("select c1, c2 from %s where c1=%d", indexTableName, rowKey[0].Value())
	rows, err := test.GlobalDB.Query(checkIndexSql)
	assert.Equal(t, nil, err)
	var c1, c2 int32
	recordCount := 0
	for rows.Next() {
		err = rows.Scan(&c1, &c2)
		assert.Equal(t, nil, err)
		assert.Equal(t, rowKey[0].Value(), c1)
		assert.Equal(t, mutateColumns[0].Value(), c2)
		recordCount++
	}
	assert.Equal(t, affectedCount, recordCount)
}

func DoInsDel(t *testing.T, tableName string, indexTableName string) {
	fmt.Println("insert")
	rowKey := []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(1)), table.NewColumn("c3", int32(3))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

	fmt.Println("delete")
	affectRows, err = cli.Delete(
		context.TODO(),
		tableName,
		rowKey,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 0)

}

func TestKeyGlobalIndexInsDel(t *testing.T) {
	tableName := testGlabalIndexKeyTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	DoInsDel(t, tableName, indexTableName[0])

}

func TestHashGlobalIndexInsDel(t *testing.T) {
	tableName := testGlabalIndexHashTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	DoInsDel(t, tableName, indexTableName[0])
}

func doInsertup(t *testing.T, tableName string, indexTableName string) {
	fmt.Println("do insert_or_update...")
	rowKey := []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err := cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

	fmt.Println("do insert_or_update again...")
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(3))}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

}

func TestKeyGlobalIndexInsertup(t *testing.T) {
	tableName := testGlabalIndexKeyTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	doInsertup(t, tableName, indexTableName[0])
}

func TestHashGlobalIndexInsertup(t *testing.T) {
	tableName := testGlabalIndexHashTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	doInsertup(t, tableName, indexTableName[0])
}

func doRepalce(t *testing.T, tableName string, indexTableName string) {
	fmt.Println("do replace...")
	rowKey := []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err := cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

	fmt.Println("do replace again...")
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(2))}
	affectRows, err = cli.Replace(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 2, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

}

func TestHashGlobalIndexReplace(t *testing.T) {
	tableName := testGlabalIndexHashTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	doRepalce(t, tableName, indexTableName[0])
}

func TestKeyGlobalIndexReplace(t *testing.T) {
	tableName := testGlabalIndexKeyTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	doRepalce(t, tableName, indexTableName[0])
}

func doUpdate(t *testing.T, tableName string, indexTableName string) {
	fmt.Println("do insert...")
	rowKey := []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(1))}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

	fmt.Println("do update...")
	mutateColumns = []*table.Column{table.NewColumn("c2", int32(3))}
	affectRows, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)
	checkIndexTableRow(t, indexTableName, rowKey, mutateColumns, 1)

}

func TestHashGlobalIndexUpdate(t *testing.T) {
	tableName := testGlabalIndexHashTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	doUpdate(t, tableName, indexTableName[0])
}

func TestKeyGlobalIndexUpdate(t *testing.T) {
	tableName := testGlabalIndexKeyTableName
	indexTableName := getGlobalIndexTableName(tableName)
	defer test.DeleteTable(tableName)

	doUpdate(t, tableName, indexTableName[0])
}

func doGlobalIndexQuery(t *testing.T, tableName string) {
	// prepare data
	recordCount := 20
	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int32(i))}
		mutateColumns := []*table.Column{table.NewColumn("c2", int32(3*i)),
			table.NewColumn("c3", int32(5*i))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	// test global index
	startRowKey := []*table.Column{table.NewColumn("c2", int32(0))}
	endRowKey := []*table.Column{table.NewColumn("c2", int32(3*recordCount))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3"}),
		option.WithQueryIndexName("idx"),
	)
	assert.Equal(t, nil, err)
	i := 0
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.Equal(t, nil, err)
		pk := res.Value("c1").(int32)
		assert.EqualValues(t, 3*pk, res.Value("c2"))
		assert.EqualValues(t, 5*pk, res.Value("c3"))
		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, recordCount, i)

	//delete data
	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int32(i))}
		affectRows, err := cli.Delete(
			context.TODO(),
			tableName,
			rowKey)
		assert.EqualValues(t, 1, affectRows)
		assert.Equal(t, nil, err)
	}
}

func TestHashGlobalIndexQuery(t *testing.T) {
	tableName := testGlabalIndexHashTableName
	defer test.DeleteTable(tableName)
	// insert -> query -> delete
	doGlobalIndexQuery(t, tableName)
}

func TestKeyGlobalIndexQuery(t *testing.T) {
	tableName := testGlabalIndexKeyTableName
	defer test.DeleteTable(tableName)

	doGlobalIndexQuery(t, tableName)
}

func TestPrimaryAndGlobalIndexNoPartition(t *testing.T) {
	defer test.DeleteTable(testGlobalIndexNoPart)
	doGlobalIndexQuery(t, testGlobalIndexNoPart)
}

func TestPrimaryNoPartitionAndGlobalIndexPartition(t *testing.T) {
	defer test.DeleteTable(testGlobalAllNoPart)
	doGlobalIndexQuery(t, testGlobalAllNoPart)
}

func TestPrimaryPartitionAndGlobalIndexNoPartition(t *testing.T) {
	defer test.DeleteTable(testGlobalPrimaryNoPart)
	doGlobalIndexQuery(t, testGlobalPrimaryNoPart)
}

func TestTwoLevelPartitionPrimaryTable(t *testing.T) {
	defer test.DeleteTable(testGlobalTwoLevelPart)
	doDmlAndQuery(t, testGlobalTwoLevelPart)
}

func verifyResult(t *testing.T, tableName string,
	rowKey []*table.Column, mutateColumns []*table.Column) {
	res, err := cli.Get(context.TODO(),
		tableName,
		rowKey,
		[]string{"c1", "c2", "c3", "c4", "c5"})
	assert.Equal(t, nil, err)
	assert.Equal(t, rowKey[0].Value(), res.Value("c1"))
	assert.Equal(t, rowKey[1].Value(), res.Value("c2"))
	assert.Equal(t, mutateColumns[0].Value(), res.Value("c3"))
	assert.Equal(t, mutateColumns[1].Value(), res.Value("c4"))
	assert.Equal(t, mutateColumns[2].Value(), res.Value("c5"))

}

func doDmlAndQuery(t *testing.T, tableName string) {
	recordCount := 20
	// prepare data
	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int32(i)), table.NewColumn("c2", int32(3*i))}
		mutateColumns := []*table.Column{table.NewColumn("c3", int32(5*i)),
			table.NewColumn("c4", int32(7*i)), table.NewColumn("c5", "hello~")}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.EqualValues(t, 1, affectRows)
		assert.Equal(t, nil, err)
		verifyResult(t, tableName, rowKey, mutateColumns)
	}

	// update
	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int32(i)), table.NewColumn("c2", int32(3*i))}
		mutateColumns := []*table.Column{table.NewColumn("c3", int32(4*i)),
			table.NewColumn("c4", int32(6*i)), table.NewColumn("c5", "hello~hello~")}
		affectRows, err := cli.Update(context.TODO(),
			tableName,
			rowKey,
			mutateColumns)
		assert.EqualValues(t, 1, affectRows)
		assert.Equal(t, nil, err)
		verifyResult(t, tableName, rowKey, mutateColumns)
	}
	// repalce
	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int32(i)), table.NewColumn("c2", int32(3*i))}
		mutateColumns := []*table.Column{table.NewColumn("c3", int32(5*i)),
			table.NewColumn("c4", int32(7*i)), table.NewColumn("c5", "hello~")}
		affectRows, err := cli.Replace(context.TODO(),
			tableName,
			rowKey,
			mutateColumns)
		assert.EqualValues(t, 2, affectRows)
		assert.Equal(t, nil, err)
		verifyResult(t, tableName, rowKey, mutateColumns)
	}
	// insert_or_update
	for i := 0; i < recordCount; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int32(i)), table.NewColumn("c2", int32(3*i))}
		mutateColumns := []*table.Column{table.NewColumn("c3", int32(5*i)),
			table.NewColumn("c4", int32(6*i)), table.NewColumn("c5", "hello~hi~")}
		affectRows, err := cli.InsertOrUpdate(context.TODO(),
			tableName,
			rowKey,
			mutateColumns)
		assert.EqualValues(t, 1, affectRows)
		assert.Equal(t, nil, err)
		verifyResult(t, tableName, rowKey, mutateColumns)
	}
}

func TestGlobalUniqueIndexWithConflict(t *testing.T) {
	tableName := testGlobalUniqueIndex
	defer test.DeleteTable(tableName)

	// first insert
	rowKey := []*table.Column{table.NewColumn("c1", int32(1))}
	mutateColumns := []*table.Column{table.NewColumn("c2", int32(1)), table.NewColumn("c3", int32(1)),
		table.NewColumn("c4", int32(1)), table.NewColumn("c5", "hello~")}
	affectRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns)
	assert.EqualValues(t, 1, affectRows)
	assert.Equal(t, nil, err)

	// will cause unique conflict
	rowKey2 := []*table.Column{table.NewColumn("c1", int32(2))}
	mutateColumns2 := []*table.Column{table.NewColumn("c2", int32(2)), table.NewColumn("c3", int32(1)),
		table.NewColumn("c4", int32(8)), table.NewColumn("c5", "hello~")}
	affectRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey2,
		mutateColumns2)
	assert.EqualValues(t, -1, affectRows)
	assert.NotEqual(t, nil, err)

	// query
	startRowKey := []*table.Column{table.NewColumn("c3", int32(0))}
	endRowKey := []*table.Column{table.NewColumn("c3", int32(10))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3", "c4", "c5"}),
		option.WithQueryIndexName("idx"),
	)
	assert.Equal(t, nil, err)
	i := 0
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, res.Value("c1"))
		assert.EqualValues(t, 1, res.Value("c2"))
		assert.EqualValues(t, 1, res.Value("c3"))
		assert.EqualValues(t, 1, res.Value("c4"))
		assert.EqualValues(t, "hello~", res.Value("c5"))

		i++
	}
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, i)

	// insert_or_update with global unique index
	rowKey3 := []*table.Column{table.NewColumn("c1", int32(3))}
	mutateColumns3 := []*table.Column{table.NewColumn("c2", int32(4)), table.NewColumn("c3", int32(1)),
		table.NewColumn("c4", int32(16)), table.NewColumn("c5", "hi~")}
	affectRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey3,
		mutateColumns3)
	assert.EqualValues(t, 1, affectRows)
	assert.Equal(t, nil, err)
}

func TestGlobalIndexWithTTL(t *testing.T) {
	tableName := testGlobalIndexWithTTL
	defer test.DeleteTable(tableName)
	prefixKey := "test"
	keyIds := []int64{1, 2}
	// prepare data
	for i := 0; i < len(keyIds); i++ {
		rowKey := []*table.Column{table.NewColumn("c1", prefixKey), table.NewColumn("c2", keyIds[i])}
		mutateColumns := []*table.Column{table.NewColumn("c3", keyIds[i]+100),
			table.NewColumn("c3", keyIds[i]+100),
			table.NewColumn("c4", keyIds[i]+200),
			table.NewColumn("expired_ts", nil)}
		affectRows, err := cli.Insert(context.TODO(), tableName, rowKey, mutateColumns)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}
	// query
	startRowKey := []*table.Column{table.NewColumn("c3", int64(101))}
	endRowKey := []*table.Column{table.NewColumn("c3", int64(102))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3", "c4", "expired_ts"}),
		option.WithQueryIndexName("idx2"),
	)
	assert.Equal(t, nil, err)
	res, err := resSet.Next()
	assert.Equal(t, nil, err)
	count := 0
	for ; res != nil && err == nil; res, err = resSet.Next() {
		count++
	}
	assert.EqualValues(t, 2, count)

	// update
	curTime := time.Now().Local()
	rowKey := []*table.Column{table.NewColumn("c1", prefixKey), table.NewColumn("c2", keyIds[1])}
	mutateColumns := []*table.Column{table.NewColumn("expired_ts", table.TimeStamp(curTime))}
	affectRows, err := cli.Update(context.TODO(), tableName, rowKey, mutateColumns)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectRows)

	// requery
	resSet, err = cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2", "c3", "c4", "expired_ts"}),
		option.WithQueryIndexName("idx2"),
	)
	assert.Equal(t, nil, err)
	count = 0
	res, err = resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		count++
		assert.EqualValues(t, "test", res.Value("c1"))
		assert.EqualValues(t, 1, res.Value("c2"))
		assert.EqualValues(t, 101, res.Value("c3"))
		assert.EqualValues(t, 201, res.Value("c4"))
		assert.EqualValues(t, nil, res.Value("expired_ts"))
	}
	assert.EqualValues(t, 1, count)
}

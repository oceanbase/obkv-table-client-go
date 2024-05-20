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

package error_message

import (
	"context"
	"fmt"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	errMsgTableName       = "error_message_table"
	errMsgCreateStatement = "create table if not exists error_message_table(" +
		"c1 bigint(20) not null, " +
		"c2 varchar(5) not null, " +
		"c3 datetime default current_timestamp," +
		"c4 varchar(5) generated always as (substr(c2, 1))," +
		"c5 double default 0," +
		"primary key (c1));"
)

func TestErrMsg_TypeNotMatch(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", int64(1)), // c2 varchar(5) not null
	}

	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-10511, " +
		"errCodeName:ObKvColumnTypeNotMatch, " +
		"errMsg:Column type for 'c2' not match, schema column type is 'VARCHAR', input column type is 'BIGINT'"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_ColumnNotExist(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("cx", int64(1)), // cx not exist
	}

	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-5217, errCodeName:ObErrBadFieldError, errMsg:Unknown column 'cx' in 'error_message_table'"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_DataTooLong(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", "123456"), // c2 varchar(5) not null
	}

	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-5167, errCodeName:ObErrDataTooLong, errMsg:Data too long for column 'c2' at row 0"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_DataOverflow(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", "a"),
		table.NewColumn("c3", table.DateTime(time.Now().Local())), // c3 datetime default current_timestamp
	}

	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-4157, errCodeName:ObOperateOverflow, errMsg:DATETIME value is out of range in 'c3'"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_RowkeyCountNotMatch(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
		table.NewColumn("c2", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", "a"),
	}

	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-10510, errCodeName:ObKvRowkeyCountNotMatch, errMsg:Rowkey column count not match, schema rowkey count is '1', input rowkey count is '2'"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_MutateRowkeyNotSupport(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", "a"),
	}

	affectedRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	mutateColumns = []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	affectedRows, err = cli.InsertOrUpdate(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-4007, errCodeName:ObNotSupported, errMsg:mutate rowkey column not supported"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_PrimaryKeyDuplicate(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", "a"),
	}

	affectedRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	affectedRows, err = cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-5024, errCodeName:ObErrPrimaryKeyDuplicate, errMsg:Duplicate entry '1' for key 'PRIMARY'"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_ScanRangeMissing(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0)), table.NewColumn("c2", table.Min)}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(100)), table.NewColumn("c2", table.Max)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	resSet, err := cli.Query(
		context.TODO(),
		tableName,
		keyRanges,
		option.WithQuerySelectColumns([]string{"c1", "c2"}),
	)
	res, err := resSet.Next()
	for ; res != nil && err == nil; res, err = resSet.Next() {
		values := res.Values()
		println(values[0].(int64), values[1].(string))
		println(res.Value("c1").(int64), res.Value("c2").(string))
	}
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-10513, errCodeName:ObKvScanRangeMissing, errMsg:Scan range missing, input scan range cell count is '2', which should equal to rowkey count '1'"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_NotHasDefaultValue(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c3", table.DateTime(time.Now().Local().Truncate(time.Second))), // c2 varchar(5) not null, but not set
	}

	_, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-4227, errCodeName:ObErrNoDefaultForField, errMsg:Field 'c2' doesn't have a default value"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_UpdateVirtualColumnNotSupport(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{
		table.NewColumn("c1", int64(1)),
	}

	mutateColumns := []*table.Column{
		table.NewColumn("c2", "abc"),
	}

	affectedRows, err := cli.Insert(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.Equal(t, nil, err)
	assert.EqualValues(t, 1, affectedRows)

	mutateColumns = []*table.Column{
		table.NewColumn("c4", "abc"), // c4 varchar(5) generated always as (substr(c2, 1)),
	}

	_, err = cli.Update(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
	)
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-4007, errCodeName:ObNotSupported, errMsg:The specified value for generated column not supported"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_AggOutOfRange(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	for i := 0; i < 5; i++ {
		rowKey := []*table.Column{
			table.NewColumn("c1", int64(i)),
		}

		mutateColumns := []*table.Column{
			table.NewColumn("c2", "abc"),
			table.NewColumn("c5", math.MaxFloat64),
		}

		affectedRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectedRows)
	}

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(5))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	aggExecutor := cli.NewAggExecutor(tableName, keyRanges).Sum("c5")

	_, err := aggExecutor.Execute(context.TODO())
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-4224, errCodeName:ObDataOutOfRange, errMsg:Out of range value for column 'sum(c5)' at row 0"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

func TestErrMsg_NotSupportTypeForAgg(t *testing.T) {
	tableName := errMsgTableName
	defer test.DeleteTable(tableName)

	for i := 0; i < 5; i++ {
		rowKey := []*table.Column{
			table.NewColumn("c1", int64(i)),
		}

		mutateColumns := []*table.Column{
			table.NewColumn("c2", "abc"),
			table.NewColumn("c5", math.MaxFloat64),
		}

		affectedRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectedRows)
	}

	startRowKey := []*table.Column{table.NewColumn("c1", int64(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(5))}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey)}
	aggExecutor := cli.NewAggExecutor(tableName, keyRanges).Sum("c2")

	_, err := aggExecutor.Execute(context.TODO())
	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
	expectContainErrMsg := "errCode:-4007, errCodeName:ObNotSupported, errMsg:VARCHAR type for sum aggregation not supported"
	assert.EqualValues(t, true, strings.Contains(err.Error(), expectContainErrMsg))
}

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
	"strings"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client/option"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	queryLobTableName            = "queryLobTable"
	queryLobTableCreateStatement = "create table if not exists queryLobTable(" +
		"c1 bigint(20) not null, " +
		"c2 bigint(20) not null, " +
		"c3 longtext not null, " +
		"primary key (c1));"
)

const (
	passQueryLobTest = true
	lobLength        = 2 * 1024 * 1024 // 2M
)

func prepareLobRecord(recordCount int) {
	lobVal := strings.Repeat("a", lobLength) // 2M
	for i := 0; i < recordCount; i++ {
		insertStatement := fmt.Sprintf("insert into %s(c1, c2, c3) values(%d, %d, '%s');", queryLobTableName, i, i, lobVal)
		test.InsertTable(insertStatement)
	}
}

func TestQueryLob_test1(t *testing.T) {
	if passQueryLobTest {
		fmt.Println("Please run query lob tests manually!!!")
		fmt.Println("Change passQueryLobTest to false in test/query/query_lob_test.go to run query lob tests.")
		assert.Equal(t, passQueryLobTest, false)
		return
	}

	tableName := queryLobTableName
	defer test.DeleteTable(tableName)

	recordCount := 32

	// insert
	prepareLobRecord(recordCount)

	// query
	startRowKey := []*table.Column{table.NewColumn("c1", int64(0))}
	endRowKey := []*table.Column{table.NewColumn("c1", int64(recordCount))}
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
		assert.EqualValues(t, i, res.Value("c1"))
		str, ok := res.Value("c3").(string)
		if ok {
			assert.EqualValues(t, lobLength, len(str))
		}
	}
}

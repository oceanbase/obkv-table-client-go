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

package compress_rpc_result

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	tableName       = "test_compress"
	createTableStat = "create table if not exists test_compress (`c1` bigint(20) not null, c2 varchar(256) not null, primary key (`c1`));"
)

func makeBigString(str string, length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func TestQueryResultWithCompress(t *testing.T) {
	defer test.DeleteTable(tableName)
	// prepare data
	for i := 1; i < 11; i++ {
		rowKey := []*table.Column{table.NewColumn("c1", int64(i))}
		mutateColumns := []*table.Column{table.NewColumn("c2", makeBigString("A", 256))}
		affectRows, err := cli.Insert(
			context.TODO(),
			tableName,
			rowKey,
			mutateColumns,
		)
		assert.Equal(t, nil, err)
		assert.EqualValues(t, 1, affectRows)
	}

	compressTypeList := []string{"lz4_1.0", "snappy_1.0", "zstd_1.0", "zlib_1.0", "zstd_1.3.8", "lz4_1.9.1"}
	for _, compressType := range compressTypeList {
		err := setCompressType(compressType)
		assert.Equal(t, nil, err)
		DoQuery(t)
	}
}

func DoQuery(t *testing.T) {
	startKey := []*table.Column{table.NewColumn("c1", int64(1))}
	endKey := []*table.Column{table.NewColumn("c1", table.Max)}

	keyRanges := []*table.RangePair{table.NewRangePair(startKey, endKey)}
	resSet, resErr := cli.Query(context.TODO(), tableName, keyRanges)
	assert.Equal(t, nil, resErr)
	count := 0
	res, err := resSet.Next()
	assert.Equal(t, err, nil)
	for ; res != nil && err == nil; res, err = resSet.Next() {
		count++
		assert.Equal(t, nil, err)
		assert.EqualValues(t, count, res.Value("c1"))
	}
	assert.EqualValues(t, 10, count)
}

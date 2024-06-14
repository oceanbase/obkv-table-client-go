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

package single

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	testRedisTableName = "modis_list_table"
)

func TestRedis(t *testing.T) {
	tableName := testRedisTableName
	defer test.DeleteTable(tableName)

	rowKey := []*table.Column{table.NewColumn("db", int64(1)), table.NewColumn("rkey", []byte("key1")), table.NewColumn("index", int64(-1))}
	mutateColumns := []*table.Column{table.NewColumn("REDIS_CODE_STR", []byte("*2\r\n$5\r\nLpUSH\r\n$6\r\nfoobar\r\n"))}
	result, err := cli.Redis(
		context.TODO(),
		tableName,
		rowKey,
		mutateColumns,
		option.WithReturnAffectedEntity(true),
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, "test hello ob_redis list operator!", result.Value("REDIS_CODE_STR"))
	assert.EqualValues(t, 10086, result.AffectedRows())
}

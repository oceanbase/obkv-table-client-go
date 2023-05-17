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

package batch

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/test"
)

func TestBatch(t *testing.T) {
	tableName := batchOpTableTableName
	defer test.DeleteTable(tableName)

	err := cli.AddRowKey(tableName, []string{"c1"})
	assert.Equal(t, nil, err)

	rowKey1 := []*table.Column{table.NewColumn("c1", int64(1))}
	rowKey2 := []*table.Column{table.NewColumn("c1", int64(2))}
	selectColumns1 := []string{"c1"}
	selectColumns2 := []string{"c2"}
	mutateColumns1 := []*table.Column{table.NewColumn("c2", int64(1))}
	mutateColumns2 := []*table.Column{table.NewColumn("c2", int64(2))}

	batchExecutor := cli.NewBatchExecutor(tableName)
	err = batchExecutor.AddInsertOp(rowKey1, mutateColumns1)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddInsertOp(rowKey2, mutateColumns2)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddGetOp(rowKey1, selectColumns1)
	assert.Equal(t, nil, err)
	err = batchExecutor.AddGetOp(rowKey2, selectColumns2)
	assert.Equal(t, nil, err)
	batchRes, err := batchExecutor.Execute(context.TODO())
	assert.Equal(t, nil, err)
	allResults := batchRes.GetResults()

	assert.Equal(t, 4, len(allResults))
	assert.EqualValues(t, 1, allResults[0].AffectedRows())
	assert.EqualValues(t, 1, allResults[1].AffectedRows())
	assert.EqualValues(t, 1, allResults[2].Entity().GetSimpleProperties()["c1"])
	assert.EqualValues(t, 2, allResults[3].Entity().GetSimpleProperties()["c2"])
}

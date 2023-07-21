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

package client

import (
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestObTransferRange(t *testing.T) {

	start := rand.Uint64()
	end := start + rand.Uint64()

	inclusiveStart := true

	inclusiveEnd := false

	startRowKey := []*table.Column{table.NewColumn("c1", start), table.NewColumn("c2", start)}
	endRowKey := []*table.Column{table.NewColumn("c1", end), table.NewColumn("c2", end)}
	keyRanges := []*table.RangePair{table.NewRangePair(startRowKey, endRowKey, inclusiveStart, inclusiveEnd)}

	transferRanges, err := TransferQueryRange(keyRanges)

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(transferRanges))

	transferRange := transferRanges[0]

	flag := transferRange.BorderFlag()

	assert.Equal(t, true, flag.IsInclusiveStart())
	assert.Equal(t, false, flag.IsInclusiveEnd())

	for _, startKey := range transferRange.StartKey() {
		assert.Equal(t, start, startKey.Value().(uint64))
	}

	for _, endKey := range transferRange.EndKey() {
		assert.Equal(t, end, endKey.Value().(uint64))
	}
}

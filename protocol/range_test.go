/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS WITHOUT WARRANTIES OF ANY KIND
 * EITHER EXPRESS OR IMPLIED INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObNewRangeEncodeDecode(t *testing.T) {
	util.SetObVersion(4)
	obNewRange := NewObNewRange()
	obNewRange.SetTableId(rand.Uint64())
	obNewRange.SetBorderFlag(ObBorderFlag(rand.Intn(255)))

	randomLen := rand.Intn(100)
	startKey := make([]*ObObject, 0, randomLen)
	endKey := make([]*ObObject, 0, randomLen)
	columns := make([]*table.Column, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		columns = append(columns, table.NewColumn(util.String(10), int64(rand.Intn(10000))))
	}
	for _, column := range columns {
		objMeta, _ := DefaultObjMeta(column.Value())
		startKey = append(startKey, NewObObjectWithParams(objMeta, column.Value()))
		endKey = append(endKey, NewObObjectWithParams(objMeta, column.Value()))
	}

	obNewRange.SetStartKey(startKey)
	obNewRange.SetEndKey(endKey)
	obNewRange.SetFlag(int64(rand.Uint64()))

	encodedLength := obNewRange.EncodedLength()
	buf := make([]byte, encodedLength)
	buffer := bytes.NewBuffer(buf)
	obNewRange.Encode(buffer)

	newObNewRange := NewObNewRange()
	newBuffer := bytes.NewBuffer(buf)
	newObNewRange.Decode(newBuffer)

	assert.EqualValues(t, obNewRange.TableId(), newObNewRange.TableId())
	assert.EqualValues(t, obNewRange.BorderFlag(), newObNewRange.BorderFlag())
	assert.EqualValues(t, obNewRange.StartKey(), newObNewRange.StartKey())
	assert.EqualValues(t, obNewRange.EndKey(), newObNewRange.EndKey())
	assert.EqualValues(t, obNewRange.Flag(), newObNewRange.Flag())
	assert.EqualValues(t, obNewRange, newObNewRange)
}

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

package protocol

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObTableAggregationSingleEncodeDecode(t *testing.T) {
	obTableAggregationSingle := NewObTableAggregationSingle()
	obTableAggregationSingle.SetVersion(1)
	obTableAggregationSingle.SetContentLength(0)
	obTableAggregationSingle.SetAggType(ObTableAggregationType(rand.Intn(255)))
	obTableAggregationSingle.SetAggColumn(util.String(10))

	payloadLen := obTableAggregationSingle.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableAggregationSingle.Encode(buffer)

	newObTableAggregationSingle := NewObTableAggregationSingle()
	newBuffer := bytes.NewBuffer(buf)
	newObTableAggregationSingle.Decode(newBuffer)

	assert.EqualValues(t, obTableAggregationSingle.Version(), newObTableAggregationSingle.Version())
	assert.EqualValues(t, obTableAggregationSingle.ContentLength(), newObTableAggregationSingle.ContentLength())
	assert.EqualValues(t, obTableAggregationSingle.AggType(), newObTableAggregationSingle.AggType())
	assert.EqualValues(t, obTableAggregationSingle.AggColumn(), newObTableAggregationSingle.AggColumn())
	assert.EqualValues(t, obTableAggregationSingle, newObTableAggregationSingle)
}

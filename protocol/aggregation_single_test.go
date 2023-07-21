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

func TestObTableAggregationEncodeDecode(t *testing.T) {
	ObTableAggregation := NewObTableAggregation()
	ObTableAggregation.SetVersion(1)
	ObTableAggregation.SetContentLength(0)
	ObTableAggregation.SetAggType(ObTableAggregationType(rand.Intn(255)))
	ObTableAggregation.SetAggColumn(util.String(10))

	payloadLen := ObTableAggregation.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	ObTableAggregation.Encode(buffer)

	newObTableAggregation := NewObTableAggregation()
	newBuffer := bytes.NewBuffer(buf)
	newObTableAggregation.Decode(newBuffer)

	assert.EqualValues(t, ObTableAggregation.Version(), newObTableAggregation.Version())
	assert.EqualValues(t, ObTableAggregation.ContentLength(), newObTableAggregation.ContentLength())
	assert.EqualValues(t, ObTableAggregation.AggType(), newObTableAggregation.AggType())
	assert.EqualValues(t, ObTableAggregation.AggColumn(), newObTableAggregation.AggColumn())
	assert.EqualValues(t, ObTableAggregation, newObTableAggregation)
}

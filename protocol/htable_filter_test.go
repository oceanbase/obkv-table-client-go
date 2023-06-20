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

	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObHTableFilterEncodeDecode(t *testing.T) {
	obHTableFilter := NewObHTableFilter()
	obHTableFilter.SetVersion(1)
	obHTableFilter.SetContentLength(0)
	obHTableFilter.SetIsValid(util.ByteToBool(byte(rand.Intn(2))))
	selectColumnQualifierLen := rand.Intn(10)
	selectColumnQualifier := make([][]byte, 0, rand.Intn(selectColumnQualifierLen))
	for i := 0; i < selectColumnQualifierLen; i++ {
		selectColumnQualifier = append(selectColumnQualifier, []byte(util.String(10)))
	}
	obHTableFilter.SetSelectColumnQualifier(selectColumnQualifier)
	obHTableFilter.SetMinStamp(int64(rand.Uint64()))
	obHTableFilter.SetMaxStamp(int64(rand.Uint64()))
	obHTableFilter.SetMaxVersions(int32(rand.Uint32()))
	obHTableFilter.SetLimitPerRowPerCf(int32(rand.Uint32()))
	obHTableFilter.SetOffsetPerRowPerCf(int32(rand.Uint32()))
	obHTableFilter.SetFilterString(util.String(10))

	payloadLen := obHTableFilter.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obHTableFilter.Encode(buffer)

	newObHTableFilter := NewObHTableFilter()
	newBuffer := bytes.NewBuffer(buf)
	newObHTableFilter.Decode(newBuffer)

	assert.EqualValues(t, obHTableFilter.IsValid(), newObHTableFilter.IsValid())
	assert.EqualValues(t, obHTableFilter.SelectColumnQualifier(), newObHTableFilter.SelectColumnQualifier())
	assert.EqualValues(t, obHTableFilter.MinStamp(), newObHTableFilter.MinStamp())
	assert.EqualValues(t, obHTableFilter.MaxStamp(), newObHTableFilter.MaxStamp())
	assert.EqualValues(t, obHTableFilter.MaxVersions(), newObHTableFilter.MaxVersions())
	assert.EqualValues(t, obHTableFilter.LimitPerRowPerCf(), newObHTableFilter.LimitPerRowPerCf())
	assert.EqualValues(t, obHTableFilter.OffsetPerRowPerCf(), newObHTableFilter.OffsetPerRowPerCf())
	assert.EqualValues(t, obHTableFilter.FilterString(), newObHTableFilter.FilterString())
	assert.EqualValues(t, obHTableFilter, newObHTableFilter)
}

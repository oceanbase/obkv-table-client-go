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
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObRpcHeaderEncodeDecode(t *testing.T) {
	util.SetObVersion(4)
	obRpcHeader := NewObRpcHeader()
	obRpcHeader.SetPCode(rand.Uint32())
	obRpcHeader.SetPriority(uint8(rand.Uint32()))
	obRpcHeader.SetFlag(uint16(rand.Uint32()))
	obRpcHeader.SetChecksum(int64(rand.Uint64()))
	obRpcHeader.SetTenantId(rand.Uint64())
	obRpcHeader.SetPrevTenantId(rand.Uint64())
	obRpcHeader.SetSessionId(rand.Uint64())
	obRpcHeader.SetTraceId0(rand.Uint64())
	obRpcHeader.SetTraceId1(rand.Uint64())
	obRpcHeader.SetTimeout(time.Duration(rand.Uint64()))
	obRpcHeader.SetTimestamp(int64(rand.Uint64()))

	obRpcCostTime := NewObRpcCostTime()
	obRpcCostTime.SetLen(int32(rand.Uint32()))
	obRpcCostTime.SetArrivalPushDiff(int32(rand.Uint32()))
	obRpcCostTime.SetPushPopDiff(int32(rand.Uint32()))
	obRpcCostTime.SetPopProcessStartDiff(int32(rand.Uint32()))
	obRpcCostTime.SetProcessStartEndDiff(int32(rand.Uint32()))
	obRpcCostTime.SetProcessEndResponseDiff(int32(rand.Uint32()))
	obRpcCostTime.SetPacketId(int64(rand.Uint64()))
	obRpcCostTime.SetRequestArriveTime(int64(rand.Uint64()))

	obRpcHeader.SetObRpcCostTime(obRpcCostTime)

	obRpcHeader.SetDstClusterId(int64(rand.Uint64()))
	obRpcHeader.SetCompressType(ObCompressType(rand.Uint32()))
	obRpcHeader.SetOriginalLen(int32(rand.Uint32()))
	obRpcHeader.SetSrcClusterId(int64(rand.Uint64()))
	obRpcHeader.SetUnisVersion(int64(rand.Uint64()))
	obRpcHeader.SetRequestLevel(int32(rand.Uint32()))
	obRpcHeader.SetSeqNo(int64(rand.Uint64()))
	obRpcHeader.SetGroupId(int32(rand.Uint32()))
	obRpcHeader.SetTraceId2(int64(rand.Uint64()))
	obRpcHeader.SetTraceId3(int64(rand.Uint64()))
	obRpcHeader.SetClusterNameHash(int64(rand.Uint64()))

	obRpcHeader.SetHLen(RpcHeaderEncodeSizeV4)
	buf := make([]byte, RpcHeaderEncodeSizeV4)
	buffer := bytes.NewBuffer(buf)
	obRpcHeader.Encode(buffer)

	newObRpcHeader := NewObRpcHeader()
	newBuffer := bytes.NewBuffer(buf)
	newObRpcHeader.Decode(newBuffer)

	assert.EqualValues(t, obRpcHeader.PCode(), newObRpcHeader.PCode())
	assert.EqualValues(t, obRpcHeader.HLen(), newObRpcHeader.HLen())
	assert.EqualValues(t, obRpcHeader.Priority(), newObRpcHeader.Priority())
	assert.EqualValues(t, obRpcHeader.Flag(), newObRpcHeader.Flag())
	assert.EqualValues(t, obRpcHeader.Checksum(), newObRpcHeader.Checksum())
	assert.EqualValues(t, obRpcHeader.TenantId(), newObRpcHeader.TenantId())
	assert.EqualValues(t, obRpcHeader.PrevTenantId(), newObRpcHeader.PrevTenantId())
	assert.EqualValues(t, obRpcHeader.SessionId(), newObRpcHeader.SessionId())
	assert.EqualValues(t, obRpcHeader.TraceId0(), newObRpcHeader.TraceId0())
	assert.EqualValues(t, obRpcHeader.TraceId1(), newObRpcHeader.TraceId1())
	assert.EqualValues(t, obRpcHeader.Timeout().Microseconds(), newObRpcHeader.Timeout()/time.Microsecond)
	assert.EqualValues(t, obRpcHeader.Timestamp(), newObRpcHeader.Timestamp())
	assert.EqualValues(t, obRpcHeader.ObRpcCostTime(), newObRpcHeader.ObRpcCostTime())
	assert.EqualValues(t, obRpcHeader.DstClusterId(), newObRpcHeader.DstClusterId())
	assert.EqualValues(t, obRpcHeader.CompressType(), newObRpcHeader.CompressType())
	assert.EqualValues(t, obRpcHeader.OriginalLen(), newObRpcHeader.OriginalLen())
	assert.EqualValues(t, obRpcHeader.SrcClusterId(), newObRpcHeader.SrcClusterId())
	assert.EqualValues(t, obRpcHeader.UnisVersion(), newObRpcHeader.UnisVersion())
	assert.EqualValues(t, obRpcHeader.RequestLevel(), newObRpcHeader.RequestLevel())
	assert.EqualValues(t, obRpcHeader.SeqNo(), newObRpcHeader.SeqNo())
	assert.EqualValues(t, obRpcHeader.GroupId(), newObRpcHeader.GroupId())
	assert.EqualValues(t, obRpcHeader.TraceId2(), newObRpcHeader.TraceId2())
	assert.EqualValues(t, obRpcHeader.TraceId3(), newObRpcHeader.TraceId3())
	assert.EqualValues(t, obRpcHeader.ClusterNameHash(), newObRpcHeader.ClusterNameHash())

	util.SetObVersion(3)
	obRpcHeader = NewObRpcHeader()
	obRpcHeader.SetPCode(rand.Uint32())
	obRpcHeader.SetPriority(uint8(rand.Uint32()))
	obRpcHeader.SetFlag(uint16(rand.Uint32()))
	obRpcHeader.SetChecksum(int64(rand.Uint64()))
	obRpcHeader.SetTenantId(rand.Uint64())
	obRpcHeader.SetPrevTenantId(rand.Uint64())
	obRpcHeader.SetSessionId(rand.Uint64())
	obRpcHeader.SetTraceId0(rand.Uint64())
	obRpcHeader.SetTraceId1(rand.Uint64())
	obRpcHeader.SetTimeout(time.Duration(rand.Uint64()))
	obRpcHeader.SetTimestamp(int64(rand.Uint64()))

	obRpcCostTime = NewObRpcCostTime()
	obRpcCostTime.SetLen(int32(rand.Uint32()))
	obRpcCostTime.SetArrivalPushDiff(int32(rand.Uint32()))
	obRpcCostTime.SetPushPopDiff(int32(rand.Uint32()))
	obRpcCostTime.SetPopProcessStartDiff(int32(rand.Uint32()))
	obRpcCostTime.SetProcessStartEndDiff(int32(rand.Uint32()))
	obRpcCostTime.SetProcessEndResponseDiff(int32(rand.Uint32()))
	obRpcCostTime.SetPacketId(int64(rand.Uint64()))
	obRpcCostTime.SetRequestArriveTime(int64(rand.Uint64()))

	obRpcHeader.SetObRpcCostTime(obRpcCostTime)

	obRpcHeader.SetDstClusterId(int64(rand.Uint64()))
	obRpcHeader.SetCompressType(ObCompressType(rand.Uint32()))
	obRpcHeader.SetOriginalLen(int32(rand.Uint32()))

	obRpcHeader.SetHLen(RpcHeaderEncodeSizeV3)
	buf = make([]byte, RpcHeaderEncodeSizeV3)
	buffer = bytes.NewBuffer(buf)
	obRpcHeader.Encode(buffer)

	newObRpcHeader = NewObRpcHeader()
	newBuffer = bytes.NewBuffer(buf)
	newObRpcHeader.Decode(newBuffer)

	assert.EqualValues(t, obRpcHeader.PCode(), newObRpcHeader.PCode())
	assert.EqualValues(t, obRpcHeader.HLen(), newObRpcHeader.HLen())
	assert.EqualValues(t, obRpcHeader.Priority(), newObRpcHeader.Priority())
	assert.EqualValues(t, obRpcHeader.Flag(), newObRpcHeader.Flag())
	assert.EqualValues(t, obRpcHeader.Checksum(), newObRpcHeader.Checksum())
	assert.EqualValues(t, obRpcHeader.TenantId(), newObRpcHeader.TenantId())
	assert.EqualValues(t, obRpcHeader.PrevTenantId(), newObRpcHeader.PrevTenantId())
	assert.EqualValues(t, obRpcHeader.SessionId(), newObRpcHeader.SessionId())
	assert.EqualValues(t, obRpcHeader.TraceId0(), newObRpcHeader.TraceId0())
	assert.EqualValues(t, obRpcHeader.TraceId1(), newObRpcHeader.TraceId1())
	assert.EqualValues(t, obRpcHeader.Timeout().Microseconds(), newObRpcHeader.Timeout()/time.Microsecond)
	assert.EqualValues(t, obRpcHeader.Timestamp(), newObRpcHeader.Timestamp())
	assert.EqualValues(t, obRpcHeader.ObRpcCostTime(), newObRpcHeader.ObRpcCostTime())
	assert.EqualValues(t, obRpcHeader.DstClusterId(), newObRpcHeader.DstClusterId())
	assert.EqualValues(t, obRpcHeader.CompressType(), newObRpcHeader.CompressType())
	assert.EqualValues(t, obRpcHeader.OriginalLen(), newObRpcHeader.OriginalLen())
}

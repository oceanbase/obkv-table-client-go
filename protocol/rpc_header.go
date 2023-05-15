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
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

const (
	defaultFlag uint16 = 7

	defaultOperationTimeout = 10 * 1000 * time.Millisecond

	headerEncodeSize = 72

	encodeSizeWithCostTime = headerEncodeSize +
		obCostTimeEncodeSize

	encodeSizeWithCostTimeAndDstClusterId = headerEncodeSize +
		obCostTimeEncodeSize +
		8 // dstClusterId

	encodeSize = headerEncodeSize +
		obCostTimeEncodeSize +
		8 + // dstClusterId
		4 + // compressType
		4 // originalLen

	encodeSizeV4 = headerEncodeSize +
		obCostTimeEncodeSize +
		8 + // dstClusterId
		4 + // compressType
		4 + // originalLen
		8 + // srcClusterId
		8 + // unis version
		4 + // request level
		8 + // seq no
		4 + // group id
		8 + // trace id2
		8 + // trace id3
		8 // clusterNameHash
)

type ObCompressType int32

const (
	ObCompressTypeInvalid ObCompressType = iota
	ObCompressTypeNone
	ObCompressTypeLZ4
	ObCompressTypeSnappy
	ObCompressTypeZlib
	ObCompressTypeZstd
)

type ObRpcHeader struct {
	pCode         uint32
	hLen          uint8
	priority      uint8
	flag          uint16
	checksum      int64
	tenantId      uint64
	prevTenantId  uint64
	sessionId     uint64
	traceId0      uint64 // uniqueId
	traceId1      uint64 // sequence
	timeout       time.Duration
	timestamp     int64
	obRpcCostTime *ObRpcCostTime
	dstClusterId  int64
	compressType  ObCompressType
	// original length before compression.
	originalLen int32
	// v4
	srcClusterId    int64
	unisVersion     int64
	requestLevel    int32
	seqNo           int64
	groupId         int32
	traceId2        int64
	traceId3        int64
	clusterNameHash int64
}

func NewObRpcHeader() *ObRpcHeader {
	return &ObRpcHeader{
		pCode:           0,
		hLen:            0,
		priority:        5,
		flag:            defaultFlag,
		checksum:        0,
		tenantId:        1,
		prevTenantId:    1,
		sessionId:       1,
		traceId0:        0,
		traceId1:        0,
		timeout:         defaultOperationTimeout,
		timestamp:       time.Now().Unix(),
		obRpcCostTime:   NewObRpcCostTime(),
		dstClusterId:    -1,
		compressType:    ObCompressTypeInvalid,
		originalLen:     0,
		srcClusterId:    -1,
		unisVersion:     0,
		requestLevel:    0,
		seqNo:           0,
		groupId:         0,
		traceId2:        0,
		traceId3:        0,
		clusterNameHash: 0,
	}
}

func (h *ObRpcHeader) PCode() uint32 {
	return h.pCode
}

func (h *ObRpcHeader) SetPCode(pCode uint32) {
	h.pCode = pCode
}

func (h *ObRpcHeader) HLen() uint8 {
	return h.hLen
}

func (h *ObRpcHeader) SetHLen(hLen uint8) {
	h.hLen = hLen
}

func (h *ObRpcHeader) Priority() uint8 {
	return h.priority
}

func (h *ObRpcHeader) SetPriority(priority uint8) {
	h.priority = priority
}

func (h *ObRpcHeader) Flag() uint16 {
	return h.flag
}

func (h *ObRpcHeader) SetFlag(flag uint16) {
	h.flag = flag
}

func (h *ObRpcHeader) Checksum() int64 {
	return h.checksum
}

func (h *ObRpcHeader) SetChecksum(checksum int64) {
	h.checksum = checksum
}

func (h *ObRpcHeader) TenantId() uint64 {
	return h.tenantId
}

func (h *ObRpcHeader) SetTenantId(tenantId uint64) {
	h.tenantId = tenantId
}

func (h *ObRpcHeader) PrevTenantId() uint64 {
	return h.prevTenantId
}

func (h *ObRpcHeader) SetPrevTenantId(prevTenantId uint64) {
	h.prevTenantId = prevTenantId
}

func (h *ObRpcHeader) SessionId() uint64 {
	return h.sessionId
}

func (h *ObRpcHeader) SetSessionId(sessionId uint64) {
	h.sessionId = sessionId
}

func (h *ObRpcHeader) TraceId0() uint64 {
	return h.traceId0
}

func (h *ObRpcHeader) SetTraceId0(traceId0 uint64) {
	h.traceId0 = traceId0
}

func (h *ObRpcHeader) TraceId1() uint64 {
	return h.traceId1
}

func (h *ObRpcHeader) SetTraceId1(traceId1 uint64) {
	h.traceId1 = traceId1
}

func (h *ObRpcHeader) Timeout() time.Duration {
	return h.timeout
}

func (h *ObRpcHeader) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *ObRpcHeader) Timestamp() int64 {
	return h.timestamp
}

func (h *ObRpcHeader) SetTimestamp(timestamp int64) {
	h.timestamp = timestamp
}

func (h *ObRpcHeader) ObRpcCostTime() *ObRpcCostTime {
	return h.obRpcCostTime
}

func (h *ObRpcHeader) SetObRpcCostTime(obRpcCostTime *ObRpcCostTime) {
	h.obRpcCostTime = obRpcCostTime
}

func (h *ObRpcHeader) DstClusterId() int64 {
	return h.dstClusterId
}

func (h *ObRpcHeader) SetDstClusterId(dstClusterId int64) {
	h.dstClusterId = dstClusterId
}

func (h *ObRpcHeader) CompressType() ObCompressType {
	return h.compressType
}

func (h *ObRpcHeader) SetCompressType(compressType ObCompressType) {
	h.compressType = compressType
}

func (h *ObRpcHeader) OriginalLen() int32 {
	return h.originalLen
}

func (h *ObRpcHeader) SetOriginalLen(originalLen int32) {
	h.originalLen = originalLen
}

func (h *ObRpcHeader) SrcClusterId() int64 {
	return h.srcClusterId
}

func (h *ObRpcHeader) SetSrcClusterId(srcClusterId int64) {
	h.srcClusterId = srcClusterId
}

func (h *ObRpcHeader) UnisVersion() int64 {
	return h.unisVersion
}

func (h *ObRpcHeader) SetUnisVersion(unisVersion int64) {
	h.unisVersion = unisVersion
}

func (h *ObRpcHeader) RequestLevel() int32 {
	return h.requestLevel
}

func (h *ObRpcHeader) SetRequestLevel(requestLevel int32) {
	h.requestLevel = requestLevel
}

func (h *ObRpcHeader) SeqNo() int64 {
	return h.seqNo
}

func (h *ObRpcHeader) SetSeqNo(seqNo int64) {
	h.seqNo = seqNo
}

func (h *ObRpcHeader) GroupId() int32 {
	return h.groupId
}

func (h *ObRpcHeader) SetGroupId(groupId int32) {
	h.groupId = groupId
}

func (h *ObRpcHeader) TraceId2() int64 {
	return h.traceId2
}

func (h *ObRpcHeader) SetTraceId2(traceId2 int64) {
	h.traceId2 = traceId2
}

func (h *ObRpcHeader) TraceId3() int64 {
	return h.traceId3
}

func (h *ObRpcHeader) SetTraceId3(traceId3 int64) {
	h.traceId3 = traceId3
}

func (h *ObRpcHeader) ClusterNameHash() int64 {
	return h.clusterNameHash
}

func (h *ObRpcHeader) SetClusterNameHash(clusterNameHash int64) {
	h.clusterNameHash = clusterNameHash
}

func (h *ObRpcHeader) Encode() []byte {
	var rpcHeaderBuf []byte
	// Maybe it would be better to use the version number to judge
	if util.ObVersion() >= 4 {
		rpcHeaderBuf = make([]byte, encodeSizeV4)
		h.hLen = encodeSizeV4
	} else { // v3
		rpcHeaderBuf = make([]byte, encodeSize)
		h.hLen = encodeSize
	}

	rpcHeaderBuffer := bytes.NewBuffer(rpcHeaderBuf)

	util.PutUint32(rpcHeaderBuffer, h.pCode)
	util.PutUint8(rpcHeaderBuffer, h.hLen)
	util.PutUint8(rpcHeaderBuffer, h.priority)

	util.PutUint16(rpcHeaderBuffer, h.flag)
	util.PutUint64(rpcHeaderBuffer, uint64(h.checksum))
	util.PutUint64(rpcHeaderBuffer, h.tenantId)
	util.PutUint64(rpcHeaderBuffer, h.prevTenantId)
	util.PutUint64(rpcHeaderBuffer, h.sessionId)
	util.PutUint64(rpcHeaderBuffer, h.traceId0)
	util.PutUint64(rpcHeaderBuffer, h.traceId1)
	util.PutUint64(rpcHeaderBuffer, uint64(h.timeout))
	util.PutUint64(rpcHeaderBuffer, uint64(h.timestamp))

	h.obRpcCostTime.Encode(rpcHeaderBuffer)

	util.PutUint64(rpcHeaderBuffer, uint64(h.dstClusterId))
	util.PutUint32(rpcHeaderBuffer, uint32(h.compressType))
	util.PutUint32(rpcHeaderBuffer, uint32(h.originalLen))

	if util.ObVersion() >= 4 {
		util.PutUint64(rpcHeaderBuffer, uint64(h.srcClusterId))
		util.PutUint64(rpcHeaderBuffer, uint64(h.unisVersion))
		util.PutUint32(rpcHeaderBuffer, uint32(h.requestLevel))
		util.PutUint64(rpcHeaderBuffer, uint64(h.seqNo))
		util.PutUint32(rpcHeaderBuffer, uint32(h.groupId))
		util.PutUint64(rpcHeaderBuffer, uint64(h.traceId2))
		util.PutUint64(rpcHeaderBuffer, uint64(h.traceId3))
		util.PutUint64(rpcHeaderBuffer, uint64(h.clusterNameHash))
	}

	return rpcHeaderBuf
}

func (h *ObRpcHeader) Decode(buffer *bytes.Buffer) {
	h.pCode = util.Uint32(buffer)
	h.hLen = util.Uint8(buffer)
	h.priority = util.Uint8(buffer)
	h.flag = util.Uint16(buffer)
	h.checksum = int64(util.Uint64(buffer))
	h.tenantId = util.Uint64(buffer)
	h.prevTenantId = util.Uint64(buffer)
	h.sessionId = util.Uint64(buffer)
	h.traceId0 = util.Uint64(buffer)
	h.traceId1 = util.Uint64(buffer)
	h.timeout = time.Duration(util.Uint64(buffer))
	h.timestamp = int64(util.Uint64(buffer))

	// Maybe it would be better to use the version number to judge
	if h.hLen >= encodeSizeV4 {
		h.obRpcCostTime.Decode(buffer)

		h.dstClusterId = int64(util.Uint64(buffer))
		h.compressType = ObCompressType(util.Uint32(buffer))
		h.originalLen = int32(util.Uint32(buffer))

		h.srcClusterId = int64(util.Uint64(buffer))
		h.unisVersion = int64(util.Uint64(buffer))
		h.requestLevel = int32(util.Uint32(buffer))
		h.seqNo = int64(util.Uint64(buffer))
		h.groupId = int32(util.Uint32(buffer))
		h.traceId2 = int64(util.Uint64(buffer))
		h.traceId3 = int64(util.Uint64(buffer))
		h.clusterNameHash = int64(util.Uint64(buffer))

		util.SkipBytes(buffer, int(h.hLen-encodeSizeV4))
	} else if h.hLen >= encodeSize {
		h.obRpcCostTime.Decode(buffer)

		h.dstClusterId = int64(util.Uint64(buffer))
		h.compressType = ObCompressType(util.Uint32(buffer))
		h.originalLen = int32(util.Uint32(buffer))

		util.SkipBytes(buffer, int(h.hLen-encodeSize))
	} else if h.hLen >= encodeSizeWithCostTimeAndDstClusterId {
		h.obRpcCostTime.Decode(buffer)

		h.dstClusterId = int64(util.Uint64(buffer))

		util.SkipBytes(buffer, int(h.hLen-encodeSizeWithCostTimeAndDstClusterId))
	} else if h.hLen >= encodeSizeWithCostTime {
		h.obRpcCostTime.Decode(buffer)

		util.SkipBytes(buffer, int(h.hLen-encodeSizeWithCostTime))
	} else {
		util.SkipBytes(buffer, int(h.hLen-headerEncodeSize))
	}
}

type ObRpcCostTime struct {
	len                    int32
	arrivalPushDiff        int32
	pushPopDiff            int32
	popProcessStartDiff    int32
	processStartEndDiff    int32
	processEndResponseDiff int32
	packetId               int64
	requestArriveTime      int64
}

const obCostTimeEncodeSize = 40

func NewObRpcCostTime() *ObRpcCostTime {
	return &ObRpcCostTime{
		len:                    obCostTimeEncodeSize,
		arrivalPushDiff:        0,
		pushPopDiff:            0,
		popProcessStartDiff:    0,
		processStartEndDiff:    0,
		processEndResponseDiff: 0,
		packetId:               0,
		requestArriveTime:      0,
	}
}

func (t *ObRpcCostTime) Len() int32 {
	return t.len
}

func (t *ObRpcCostTime) SetLen(len int32) {
	t.len = len
}

func (t *ObRpcCostTime) ArrivalPushDiff() int32 {
	return t.arrivalPushDiff
}

func (t *ObRpcCostTime) SetArrivalPushDiff(arrivalPushDiff int32) {
	t.arrivalPushDiff = arrivalPushDiff
}

func (t *ObRpcCostTime) PushPopDiff() int32 {
	return t.pushPopDiff
}

func (t *ObRpcCostTime) SetPushPopDiff(pushPopDiff int32) {
	t.pushPopDiff = pushPopDiff
}

func (t *ObRpcCostTime) PopProcessStartDiff() int32 {
	return t.popProcessStartDiff
}

func (t *ObRpcCostTime) SetPopProcessStartDiff(popProcessStartDiff int32) {
	t.popProcessStartDiff = popProcessStartDiff
}

func (t *ObRpcCostTime) ProcessStartEndDiff() int32 {
	return t.processStartEndDiff
}

func (t *ObRpcCostTime) SetProcessStartEndDiff(processStartEndDiff int32) {
	t.processStartEndDiff = processStartEndDiff
}

func (t *ObRpcCostTime) ProcessEndResponseDiff() int32 {
	return t.processEndResponseDiff
}

func (t *ObRpcCostTime) SetProcessEndResponseDiff(processEndResponseDiff int32) {
	t.processEndResponseDiff = processEndResponseDiff
}

func (t *ObRpcCostTime) PacketId() int64 {
	return t.packetId
}

func (t *ObRpcCostTime) SetPacketId(packetId int64) {
	t.packetId = packetId
}

func (t *ObRpcCostTime) RequestArriveTime() int64 {
	return t.requestArriveTime
}

func (t *ObRpcCostTime) SetRequestArriveTime(requestArriveTime int64) {
	t.requestArriveTime = requestArriveTime
}

func (t *ObRpcCostTime) Encode(buffer *bytes.Buffer) {
	util.PutUint32(buffer, uint32(t.len))
	util.PutUint32(buffer, uint32(t.arrivalPushDiff))
	util.PutUint32(buffer, uint32(t.pushPopDiff))
	util.PutUint32(buffer, uint32(t.popProcessStartDiff))
	util.PutUint32(buffer, uint32(t.processStartEndDiff))
	util.PutUint32(buffer, uint32(t.processEndResponseDiff))
	util.PutUint64(buffer, uint64(t.packetId))
	util.PutUint64(buffer, uint64(t.requestArriveTime))
}

func (t *ObRpcCostTime) Decode(buffer *bytes.Buffer) {
	t.len = int32(util.Uint32(buffer))
	t.arrivalPushDiff = int32(util.Uint32(buffer))
	t.pushPopDiff = int32(util.Uint32(buffer))
	t.popProcessStartDiff = int32(util.Uint32(buffer))
	t.processStartEndDiff = int32(util.Uint32(buffer))
	t.processEndResponseDiff = int32(util.Uint32(buffer))
	t.packetId = int64(util.Uint64(buffer))
	t.requestArriveTime = int64(util.Uint64(buffer))
}

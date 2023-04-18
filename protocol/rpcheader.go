package protocol

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

// TODO Init global version
var globalVersion = 4

const (
	defaultFlag             uint16 = 7
	defaultOperationTimeout        = 10 * 1000 * time.Millisecond

	headerEncodeSize = 72

	encodeSizeWithCostTime = headerEncodeSize +
		costTimeEncodeSize

	encodeSizeWithCostTimeAndDstClusterId = headerEncodeSize +
		costTimeEncodeSize +
		8 // dstClusterId

	encodeSize = headerEncodeSize +
		costTimeEncodeSize +
		8 + // dstClusterId
		4 + // compressType
		4 // originalLen

	encodeSizeV4 = headerEncodeSize +
		costTimeEncodeSize +
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

type CompressType int32

const (
	CompressTypeInvalid CompressType = iota
	CompressTypeNone
	CompressTypeLZ4
	CompressTypeSnappy
	CompressTypeZlib
	CompressTypeZstd
)

type RpcHeader struct {
	pCode        uint32
	hLen         uint8
	priority     uint8
	flag         uint16
	checksum     int64
	tenantId     uint64
	prevTenantId uint64
	sessionId    uint64
	traceId0     uint64
	traceId1     uint64
	timeout      time.Duration
	timestamp    int64
	rpcCostTime  *RpcCostTime
	dstClusterId int64
	compressType CompressType
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

func NewRpcHeader() *RpcHeader {
	return &RpcHeader{
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
		rpcCostTime:     NewRpcCostTime(),
		dstClusterId:    -1,
		compressType:    CompressTypeInvalid,
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

func (h *RpcHeader) PCode() uint32 {
	return h.pCode
}

func (h *RpcHeader) SetPCode(pCode uint32) {
	h.pCode = pCode
}

func (h *RpcHeader) HLen() uint8 {
	return h.hLen
}

func (h *RpcHeader) SetHLen(hLen uint8) {
	h.hLen = hLen
}

func (h *RpcHeader) Priority() uint8 {
	return h.priority
}

func (h *RpcHeader) SetPriority(priority uint8) {
	h.priority = priority
}

func (h *RpcHeader) Flag() uint16 {
	return h.flag
}

func (h *RpcHeader) SetFlag(flag uint16) {
	h.flag = flag
}

func (h *RpcHeader) Checksum() int64 {
	return h.checksum
}

func (h *RpcHeader) SetChecksum(checksum int64) {
	h.checksum = checksum
}

func (h *RpcHeader) TenantId() uint64 {
	return h.tenantId
}

func (h *RpcHeader) SetTenantId(tenantId uint64) {
	h.tenantId = tenantId
}

func (h *RpcHeader) PrevTenantId() uint64 {
	return h.prevTenantId
}

func (h *RpcHeader) SetPrevTenantId(prevTenantId uint64) {
	h.prevTenantId = prevTenantId
}

func (h *RpcHeader) SessionId() uint64 {
	return h.sessionId
}

func (h *RpcHeader) SetSessionId(sessionId uint64) {
	h.sessionId = sessionId
}

func (h *RpcHeader) TraceId0() uint64 {
	return h.traceId0
}

func (h *RpcHeader) SetTraceId0(traceId0 uint64) {
	h.traceId0 = traceId0
}

func (h *RpcHeader) TraceId1() uint64 {
	return h.traceId1
}

func (h *RpcHeader) SetTraceId1(traceId1 uint64) {
	h.traceId1 = traceId1
}

func (h *RpcHeader) Timeout() time.Duration {
	return h.timeout
}

func (h *RpcHeader) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *RpcHeader) Timestamp() int64 {
	return h.timestamp
}

func (h *RpcHeader) SetTimestamp(timestamp int64) {
	h.timestamp = timestamp
}

func (h *RpcHeader) RpcCostTime() *RpcCostTime {
	return h.rpcCostTime
}

func (h *RpcHeader) SetRpcCostTime(rpcCostTime *RpcCostTime) {
	h.rpcCostTime = rpcCostTime
}

func (h *RpcHeader) DstClusterId() int64 {
	return h.dstClusterId
}

func (h *RpcHeader) SetDstClusterId(dstClusterId int64) {
	h.dstClusterId = dstClusterId
}

func (h *RpcHeader) CompressType() CompressType {
	return h.compressType
}

func (h *RpcHeader) SetCompressType(compressType CompressType) {
	h.compressType = compressType
}

func (h *RpcHeader) OriginalLen() int32 {
	return h.originalLen
}

func (h *RpcHeader) SetOriginalLen(originalLen int32) {
	h.originalLen = originalLen
}

func (h *RpcHeader) SrcClusterId() int64 {
	return h.srcClusterId
}

func (h *RpcHeader) SetSrcClusterId(srcClusterId int64) {
	h.srcClusterId = srcClusterId
}

func (h *RpcHeader) UnisVersion() int64 {
	return h.unisVersion
}

func (h *RpcHeader) SetUnisVersion(unisVersion int64) {
	h.unisVersion = unisVersion
}

func (h *RpcHeader) RequestLevel() int32 {
	return h.requestLevel
}

func (h *RpcHeader) SetRequestLevel(requestLevel int32) {
	h.requestLevel = requestLevel
}

func (h *RpcHeader) SeqNo() int64 {
	return h.seqNo
}

func (h *RpcHeader) SetSeqNo(seqNo int64) {
	h.seqNo = seqNo
}

func (h *RpcHeader) GroupId() int32 {
	return h.groupId
}

func (h *RpcHeader) SetGroupId(groupId int32) {
	h.groupId = groupId
}

func (h *RpcHeader) TraceId2() int64 {
	return h.traceId2
}

func (h *RpcHeader) SetTraceId2(traceId2 int64) {
	h.traceId2 = traceId2
}

func (h *RpcHeader) TraceId3() int64 {
	return h.traceId3
}

func (h *RpcHeader) SetTraceId3(traceId3 int64) {
	h.traceId3 = traceId3
}

func (h *RpcHeader) ClusterNameHash() int64 {
	return h.clusterNameHash
}

func (h *RpcHeader) SetClusterNameHash(clusterNameHash int64) {
	h.clusterNameHash = clusterNameHash
}

func (h *RpcHeader) Encode() []byte {
	var rpcHeaderBuf []byte
	// TODO Maybe it would be better to use the version number to judge
	if globalVersion >= 4 { // v4
		rpcHeaderBuf = make([]byte, encodeSizeV4)
	} else { // v3
		rpcHeaderBuf = make([]byte, encodeSize)
	}

	binary.BigEndian.PutUint32(rpcHeaderBuf[:4], h.pCode)
	// TODO hLen = encodeSizeV4
	util.PutUint8(rpcHeaderBuf[4:5], encodeSizeV4)
	util.PutUint8(rpcHeaderBuf[5:6], h.priority)
	binary.BigEndian.PutUint16(rpcHeaderBuf[6:8], h.flag)
	binary.BigEndian.PutUint64(rpcHeaderBuf[8:16], uint64(h.checksum))
	binary.BigEndian.PutUint64(rpcHeaderBuf[16:24], h.tenantId)
	binary.BigEndian.PutUint64(rpcHeaderBuf[24:32], h.prevTenantId)
	binary.BigEndian.PutUint64(rpcHeaderBuf[32:40], h.sessionId)
	binary.BigEndian.PutUint64(rpcHeaderBuf[40:48], h.traceId0)
	binary.BigEndian.PutUint64(rpcHeaderBuf[48:56], h.traceId1)
	binary.BigEndian.PutUint64(rpcHeaderBuf[56:64], uint64(h.timeout))
	binary.BigEndian.PutUint64(rpcHeaderBuf[64:headerEncodeSize], uint64(h.timestamp))

	h.rpcCostTime.Encode(rpcHeaderBuf[headerEncodeSize:encodeSizeWithCostTime])

	binary.BigEndian.PutUint64(rpcHeaderBuf[encodeSizeWithCostTime:120], uint64(h.dstClusterId))
	binary.BigEndian.PutUint32(rpcHeaderBuf[120:124], uint32(h.compressType))
	binary.BigEndian.PutUint32(rpcHeaderBuf[124:encodeSize], uint32(h.originalLen))

	if globalVersion >= 4 {
		binary.BigEndian.PutUint64(rpcHeaderBuf[encodeSize:136], uint64(h.srcClusterId))
		binary.BigEndian.PutUint64(rpcHeaderBuf[136:144], uint64(h.unisVersion))
		binary.BigEndian.PutUint32(rpcHeaderBuf[144:148], uint32(h.requestLevel))
		binary.BigEndian.PutUint64(rpcHeaderBuf[148:156], uint64(h.seqNo))
		binary.BigEndian.PutUint32(rpcHeaderBuf[156:160], uint32(h.groupId))
		binary.BigEndian.PutUint64(rpcHeaderBuf[160:168], uint64(h.traceId2))
		binary.BigEndian.PutUint64(rpcHeaderBuf[168:176], uint64(h.traceId3))
		binary.BigEndian.PutUint64(rpcHeaderBuf[176:encodeSizeV4], uint64(h.clusterNameHash))
	}

	return rpcHeaderBuf
}

func (h *RpcHeader) Decode(buffer *bytes.Buffer) {
	h.pCode = binary.BigEndian.Uint32(buffer.Next(4))
	h.hLen = util.Uint8(buffer.Next(1))
	h.priority = util.Uint8(buffer.Next(1))
	h.flag = binary.BigEndian.Uint16(buffer.Next(2))
	h.checksum = int64(binary.BigEndian.Uint64(buffer.Next(8)))
	h.tenantId = binary.BigEndian.Uint64(buffer.Next(8))
	h.prevTenantId = binary.BigEndian.Uint64(buffer.Next(8))
	h.sessionId = binary.BigEndian.Uint64(buffer.Next(8))
	h.traceId0 = binary.BigEndian.Uint64(buffer.Next(8))
	h.traceId1 = binary.BigEndian.Uint64(buffer.Next(8))
	h.timeout = time.Duration(binary.BigEndian.Uint64(buffer.Next(8)))
	h.timestamp = int64(binary.BigEndian.Uint64(buffer.Next(8)))

	// TODO Maybe it would be better to use the version number to judge
	if h.hLen >= encodeSizeV4 {
		h.rpcCostTime.Decode(buffer)

		h.dstClusterId = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.compressType = CompressType(binary.BigEndian.Uint32(buffer.Next(4)))
		h.originalLen = int32(binary.BigEndian.Uint32(buffer.Next(4)))

		h.srcClusterId = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.unisVersion = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.requestLevel = int32(binary.BigEndian.Uint32(buffer.Next(4)))
		h.seqNo = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.groupId = int32(binary.BigEndian.Uint32(buffer.Next(4)))
		h.traceId2 = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.traceId3 = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.clusterNameHash = int64(binary.BigEndian.Uint64(buffer.Next(8)))

		util.SkipBytes(buffer, int(h.hLen-encodeSizeV4))
	} else if h.hLen >= encodeSize {
		h.rpcCostTime.Decode(buffer)

		h.dstClusterId = int64(binary.BigEndian.Uint64(buffer.Next(8)))
		h.compressType = CompressType(binary.BigEndian.Uint32(buffer.Next(4)))
		h.originalLen = int32(binary.BigEndian.Uint32(buffer.Next(4)))

		util.SkipBytes(buffer, int(h.hLen-encodeSize))
	} else if h.hLen >= encodeSizeWithCostTimeAndDstClusterId {
		h.rpcCostTime.Decode(buffer)

		h.dstClusterId = int64(binary.BigEndian.Uint64(buffer.Next(8)))

		util.SkipBytes(buffer, int(h.hLen-encodeSizeWithCostTimeAndDstClusterId))
	} else if h.hLen >= encodeSizeWithCostTime {
		h.rpcCostTime.Decode(buffer)

		util.SkipBytes(buffer, int(h.hLen-encodeSizeWithCostTime))
	} else {
		util.SkipBytes(buffer, int(h.hLen-headerEncodeSize))
	}
}

type RpcCostTime struct {
	len                    int32
	arrivalPushDiff        int32
	pushPopDiff            int32
	popProcessStartDiff    int32
	processStartEndDiff    int32
	processEndResponseDiff int32
	packetId               int64
	requestArriveTime      int64
}

const costTimeEncodeSize = 40

func NewRpcCostTime() *RpcCostTime {
	return &RpcCostTime{
		len:                    costTimeEncodeSize,
		arrivalPushDiff:        0,
		pushPopDiff:            0,
		popProcessStartDiff:    0,
		processStartEndDiff:    0,
		processEndResponseDiff: 0,
		packetId:               0,
		requestArriveTime:      0,
	}
}

func (t *RpcCostTime) Len() int32 {
	return t.len
}

func (t *RpcCostTime) SetLen(len int32) {
	t.len = len
}

func (t *RpcCostTime) ArrivalPushDiff() int32 {
	return t.arrivalPushDiff
}

func (t *RpcCostTime) SetArrivalPushDiff(arrivalPushDiff int32) {
	t.arrivalPushDiff = arrivalPushDiff
}

func (t *RpcCostTime) PushPopDiff() int32 {
	return t.pushPopDiff
}

func (t *RpcCostTime) SetPushPopDiff(pushPopDiff int32) {
	t.pushPopDiff = pushPopDiff
}

func (t *RpcCostTime) PopProcessStartDiff() int32 {
	return t.popProcessStartDiff
}

func (t *RpcCostTime) SetPopProcessStartDiff(popProcessStartDiff int32) {
	t.popProcessStartDiff = popProcessStartDiff
}

func (t *RpcCostTime) ProcessStartEndDiff() int32 {
	return t.processStartEndDiff
}

func (t *RpcCostTime) SetProcessStartEndDiff(processStartEndDiff int32) {
	t.processStartEndDiff = processStartEndDiff
}

func (t *RpcCostTime) ProcessEndResponseDiff() int32 {
	return t.processEndResponseDiff
}

func (t *RpcCostTime) SetProcessEndResponseDiff(processEndResponseDiff int32) {
	t.processEndResponseDiff = processEndResponseDiff
}

func (t *RpcCostTime) PacketId() int64 {
	return t.packetId
}

func (t *RpcCostTime) SetPacketId(packetId int64) {
	t.packetId = packetId
}

func (t *RpcCostTime) RequestArriveTime() int64 {
	return t.requestArriveTime
}

func (t *RpcCostTime) SetRequestArriveTime(requestArriveTime int64) {
	t.requestArriveTime = requestArriveTime
}

func (t *RpcCostTime) Encode(buf []byte) {
	binary.BigEndian.PutUint32(buf[:4], uint32(t.len))
	binary.BigEndian.PutUint32(buf[4:8], uint32(t.arrivalPushDiff))
	binary.BigEndian.PutUint32(buf[8:12], uint32(t.pushPopDiff))
	binary.BigEndian.PutUint32(buf[12:16], uint32(t.popProcessStartDiff))
	binary.BigEndian.PutUint32(buf[16:20], uint32(t.processStartEndDiff))
	binary.BigEndian.PutUint32(buf[20:24], uint32(t.processEndResponseDiff))
	binary.BigEndian.PutUint64(buf[24:32], uint64(t.packetId))
	binary.BigEndian.PutUint64(buf[32:40], uint64(t.requestArriveTime))
}

func (t *RpcCostTime) Decode(buffer *bytes.Buffer) {
	t.len = int32(binary.BigEndian.Uint32(buffer.Next(4)))
	t.arrivalPushDiff = int32(binary.BigEndian.Uint32(buffer.Next(4)))
	t.pushPopDiff = int32(binary.BigEndian.Uint32(buffer.Next(4)))
	t.popProcessStartDiff = int32(binary.BigEndian.Uint32(buffer.Next(4)))
	t.processStartEndDiff = int32(binary.BigEndian.Uint32(buffer.Next(4)))
	t.processEndResponseDiff = int32(binary.BigEndian.Uint32(buffer.Next(4)))
	t.packetId = int64(binary.BigEndian.Uint64(buffer.Next(8)))
	t.requestArriveTime = int64(binary.BigEndian.Uint64(buffer.Next(8)))
}

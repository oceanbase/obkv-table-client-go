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
	"strconv"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableMoveReplicaInfo struct {
	ObUniVersionHeader
	tableId       uint64
	schemaVersion uint64
	partitionId   uint64
	server        *ObAddr
	role          ObRole
	replicaType   ObReplicaType
	partRenewTime int64
	reserved      uint64
}

func (i *ObTableMoveReplicaInfo) TableId() uint64 {
	return i.tableId
}

func (i *ObTableMoveReplicaInfo) SetTableId(tableId uint64) {
	i.tableId = tableId
}

func (i *ObTableMoveReplicaInfo) SchemaVersion() uint64 {
	return i.schemaVersion
}

func (i *ObTableMoveReplicaInfo) SetSchemaVersion(schemaVersion uint64) {
	i.schemaVersion = schemaVersion
}

func (i *ObTableMoveReplicaInfo) PartitionId() uint64 {
	return i.partitionId
}

func (i *ObTableMoveReplicaInfo) SetPartitionId(partitionId uint64) {
	i.partitionId = partitionId
}

func (i *ObTableMoveReplicaInfo) Server() *ObAddr {
	return i.server
}

func (i *ObTableMoveReplicaInfo) SetServer(server *ObAddr) {
	i.server = server
}

func (i *ObTableMoveReplicaInfo) Role() ObRole {
	return i.role
}

func (i *ObTableMoveReplicaInfo) SetRole(role ObRole) {
	i.role = role
}

func (i *ObTableMoveReplicaInfo) ReplicaType() ObReplicaType {
	return i.replicaType
}

func (i *ObTableMoveReplicaInfo) SetReplicaType(replicaType ObReplicaType) {
	i.replicaType = replicaType
}

func (i *ObTableMoveReplicaInfo) PartRenewTime() int64 {
	return i.partRenewTime
}

func (i *ObTableMoveReplicaInfo) SetPartRenewTime(partRenewTime int64) {
	i.partRenewTime = partRenewTime
}

func (i *ObTableMoveReplicaInfo) Reserved() uint64 {
	return i.reserved
}

func (i *ObTableMoveReplicaInfo) SetReserved(reserved uint64) {
	i.reserved = reserved
}

func NewObTableMoveReplicaInfo() *ObTableMoveReplicaInfo {
	return &ObTableMoveReplicaInfo{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		tableId:       0,
		schemaVersion: 0,
		partitionId:   0,
		server:        NewObAddr(),
		role:          ObRoleInvalid,
		replicaType:   ReplicaTypeInvalid,
		partRenewTime: 0,
		reserved:      0,
	}
}

func (i *ObTableMoveReplicaInfo) PayloadLen() int {
	return i.PayloadContentLen() + i.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (i *ObTableMoveReplicaInfo) PayloadContentLen() int {
	totalLen := 0

	totalLen += util.EncodedLengthByVi64(int64(i.tableId))
	totalLen += util.EncodedLengthByVi64(int64(i.schemaVersion))
	totalLen += util.EncodedLengthByVi64(int64(i.partitionId))
	totalLen += i.server.PayloadLen()
	totalLen += 1 // role
	totalLen += util.EncodedLengthByVi32(int32(i.replicaType))
	totalLen += util.EncodedLengthByVi64(i.partRenewTime)
	totalLen += util.EncodedLengthByVi64(int64(i.reserved))

	i.ObUniVersionHeader.SetContentLength(totalLen)
	return i.ObUniVersionHeader.ContentLength()
}

func (i *ObTableMoveReplicaInfo) Encode(buffer *bytes.Buffer) {
	i.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(int32(i.tableId)))
	util.EncodeVi64(buffer, int64(int32(i.schemaVersion)))
	util.EncodeVi64(buffer, int64(int32(i.partitionId)))
	i.server.Encode(buffer)
	util.PutUint8(buffer, uint8(i.role))
	util.EncodeVi32(buffer, int32(i.replicaType))
	util.EncodeVi64(buffer, int64(int32(i.partRenewTime)))
	util.EncodeVi64(buffer, int64(int32(i.reserved)))
}

func (i *ObTableMoveReplicaInfo) Decode(buffer *bytes.Buffer) {
	i.ObUniVersionHeader.Decode(buffer)
	i.tableId = uint64(util.DecodeVi64(buffer))
	i.schemaVersion = uint64(util.DecodeVi64(buffer))
	i.partitionId = uint64(util.DecodeVi64(buffer))
	i.server.Decode(buffer)
	i.role = ObRole(util.Uint8(buffer))
	i.replicaType = ObReplicaType(util.DecodeVi32(buffer))
	i.partRenewTime = util.DecodeVi64(buffer)
	i.reserved = uint64(util.DecodeVi64(buffer))
}

func (i *ObTableMoveReplicaInfo) String() string {
	var ObUniVersionHeaderStr = "nil"
	if i.ObUniVersionHeader != (ObUniVersionHeader{}) {
		ObUniVersionHeaderStr = i.ObUniVersionHeader.String()
	}

	return "ObAddr{" +
		"ObUniVersionHeader:" + ObUniVersionHeaderStr + ", " +
		"tableId:" + strconv.Itoa(int(i.tableId)) + ", " +
		"schemaVersion:" + strconv.Itoa(int(i.schemaVersion)) + ", " +
		"partitionId:" + strconv.Itoa(int(i.partitionId)) + ", " +
		"server:" + i.server.String() + ", " +
		"role:" + strconv.Itoa(int(i.role)) + ", " +
		"replicaType:" + strconv.Itoa(int(i.replicaType)) + ", " +
		"partRenewTime:" + strconv.Itoa(int(i.partRenewTime)) + ", " +
		"reserved:" + strconv.Itoa(int(i.reserved)) +
		"}"
}

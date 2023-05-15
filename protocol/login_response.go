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

	"github.com/oceanbase/obkv-table-client-go/util"
)

type LoginResponse struct {
	*UniVersionHeader
	serverCapabilities int32
	reserved1          int32
	reserved2          int64

	serverVersion string
	credential    []byte
	tenantId      uint64
	userId        int64
	databaseId    int64
}

func NewLoginResponse() *LoginResponse {
	return &LoginResponse{
		UniVersionHeader:   NewUniVersionHeader(),
		serverCapabilities: 0,
		reserved1:          0,
		reserved2:          0,
		serverVersion:      "",
		credential:         nil,
		tenantId:           0,
		userId:             0,
		databaseId:         0,
	}
}

func (r *LoginResponse) ServerCapabilities() int32 {
	return r.serverCapabilities
}

func (r *LoginResponse) SetServerCapabilities(serverCapabilities int32) {
	r.serverCapabilities = serverCapabilities
}

func (r *LoginResponse) Reserved1() int32 {
	return r.reserved1
}

func (r *LoginResponse) SetReserved1(reserved1 int32) {
	r.reserved1 = reserved1
}

func (r *LoginResponse) Reserved2() int64 {
	return r.reserved2
}

func (r *LoginResponse) SetReserved2(reserved2 int64) {
	r.reserved2 = reserved2
}

func (r *LoginResponse) ServerVersion() string {
	return r.serverVersion
}

func (r *LoginResponse) SetServerVersion(serverVersion string) {
	r.serverVersion = serverVersion
}

func (r *LoginResponse) UserId() int64 {
	return r.userId
}

func (r *LoginResponse) SetUserId(userId int64) {
	r.userId = userId
}

func (r *LoginResponse) DatabaseId() int64 {
	return r.databaseId
}

func (r *LoginResponse) SetDatabaseId(databaseId int64) {
	r.databaseId = databaseId
}

func (r *LoginResponse) TenantId() uint64 {
	return r.tenantId
}

func (r *LoginResponse) SetTenantId(tenantId uint64) {
	r.tenantId = tenantId
}

func (r *LoginResponse) PCode() TablePacketCode {
	return TableApiLogin
}

func (r *LoginResponse) PayloadLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *LoginResponse) PayloadContentLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *LoginResponse) Credential() []byte {
	return r.credential
}

func (r *LoginResponse) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *LoginResponse) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *LoginResponse) Decode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Decode(buffer)

	r.serverCapabilities = util.DecodeVi32(buffer)
	r.reserved1 = util.DecodeVi32(buffer)
	r.reserved2 = util.DecodeVi64(buffer)

	r.serverVersion = util.DecodeVString(buffer)
	r.credential = util.DecodeBytesString(buffer)
	r.tenantId = uint64(util.DecodeVi64(buffer))
	r.userId = util.DecodeVi64(buffer)
	r.databaseId = util.DecodeVi64(buffer)
}

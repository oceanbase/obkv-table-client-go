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

type ObLoginRequest struct {
	*ObUniVersionHeader
	authMethod    uint8
	clientType    uint8
	clientVersion uint8
	reversed1     uint8

	clientCapabilities int32
	maxPacketSize      int32
	reversed2          int32
	reversed3          int64

	tenantName   string
	userName     string
	passSecret   string
	passScramble string
	databaseName string
	ttlUs        int64
}

const passScrambleLen = 20

func NewObLoginRequest(tenantName string, databaseName string, userName string, password string) *ObLoginRequest {
	passScramble := util.GetPasswordScramble(passScrambleLen)
	passSecret := util.ScramblePassword(password, passScramble)

	return &ObLoginRequest{
		ObUniVersionHeader: NewObUniVersionHeader(),
		authMethod:         0x01,
		clientType:         0x02,
		clientVersion:      0x01,
		reversed1:          0,
		clientCapabilities: 0,
		maxPacketSize:      0,
		reversed2:          0,
		reversed3:          0,
		tenantName:         tenantName,
		userName:           userName,
		passSecret:         passSecret,
		passScramble:       passScramble,
		databaseName:       databaseName,
		ttlUs:              0,
	}
}

func (r *ObLoginRequest) AuthMethod() uint8 {
	return r.authMethod
}

func (r *ObLoginRequest) SetAuthMethod(authMethod uint8) {
	r.authMethod = authMethod
}

func (r *ObLoginRequest) ClientType() uint8 {
	return r.clientType
}

func (r *ObLoginRequest) SetClientType(clientType uint8) {
	r.clientType = clientType
}

func (r *ObLoginRequest) ClientVersion() uint8 {
	return r.clientVersion
}

func (r *ObLoginRequest) SetClientVersion(clientVersion uint8) {
	r.clientVersion = clientVersion
}

func (r *ObLoginRequest) Reversed1() uint8 {
	return r.reversed1
}

func (r *ObLoginRequest) SetReversed1(reversed1 uint8) {
	r.reversed1 = reversed1
}

func (r *ObLoginRequest) ClientCapabilities() int32 {
	return r.clientCapabilities
}

func (r *ObLoginRequest) SetClientCapabilities(clientCapabilities int32) {
	r.clientCapabilities = clientCapabilities
}

func (r *ObLoginRequest) MaxPacketSize() int32 {
	return r.maxPacketSize
}

func (r *ObLoginRequest) SetMaxPacketSize(maxPacketSize int32) {
	r.maxPacketSize = maxPacketSize
}

func (r *ObLoginRequest) Reversed2() int32 {
	return r.reversed2
}

func (r *ObLoginRequest) SetReversed2(reversed2 int32) {
	r.reversed2 = reversed2
}

func (r *ObLoginRequest) Reversed3() int64 {
	return r.reversed3
}

func (r *ObLoginRequest) SetReversed3(reversed3 int64) {
	r.reversed3 = reversed3
}

func (r *ObLoginRequest) TenantName() string {
	return r.tenantName
}

func (r *ObLoginRequest) SetTenantName(tenantName string) {
	r.tenantName = tenantName
}

func (r *ObLoginRequest) UserName() string {
	return r.userName
}

func (r *ObLoginRequest) SetUserName(userName string) {
	r.userName = userName
}

func (r *ObLoginRequest) PassSecret() string {
	return r.passSecret
}

func (r *ObLoginRequest) SetPassSecret(passSecret string) {
	r.passSecret = passSecret
}

func (r *ObLoginRequest) PassScramble() string {
	return r.passScramble
}

func (r *ObLoginRequest) SetPassScramble(passScramble string) {
	r.passScramble = passScramble
}

func (r *ObLoginRequest) DatabaseName() string {
	return r.databaseName
}

func (r *ObLoginRequest) SetDatabaseName(databaseName string) {
	r.databaseName = databaseName
}

func (r *ObLoginRequest) TtlUs() int64 {
	return r.ttlUs
}

func (r *ObLoginRequest) SetTtlUs(ttlUs int64) {
	r.ttlUs = ttlUs
}

func (r *ObLoginRequest) PCode() ObTablePacketCode {
	return ObTableApiLogin
}

func (r *ObLoginRequest) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObLoginRequest) PayloadContentLen() int {
	totalLen := 4 + // authMethod clientType clientVersion reversed1
		util.EncodedLengthByVi32(r.clientCapabilities) +
		util.EncodedLengthByVi32(r.maxPacketSize) +
		util.EncodedLengthByVi32(r.reversed2) +
		util.EncodedLengthByVi64(r.reversed3) +
		util.EncodedLengthByVString(r.tenantName) +
		util.EncodedLengthByVString(r.userName) +
		util.EncodedLengthByVString(r.passSecret) +
		util.EncodedLengthByVString(r.passScramble) +
		util.EncodedLengthByVString(r.databaseName) +
		util.EncodedLengthByVi64(r.ttlUs)

	r.ObUniVersionHeader.SetContentLength(totalLen) // Set on first acquisition
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObLoginRequest) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, r.authMethod)
	util.PutUint8(buffer, r.clientType)
	util.PutUint8(buffer, r.clientVersion)
	util.PutUint8(buffer, r.reversed1)

	util.EncodeVi32(buffer, r.clientCapabilities)
	util.EncodeVi32(buffer, r.maxPacketSize)
	util.EncodeVi32(buffer, r.reversed2)

	util.EncodeVi64(buffer, r.reversed3)

	// some VString convert to bytesString
	util.EncodeVString(buffer, r.tenantName)
	util.EncodeVString(buffer, r.userName)
	util.EncodeVString(buffer, r.passSecret)
	util.EncodeVString(buffer, r.passScramble)
	util.EncodeVString(buffer, r.databaseName)

	util.EncodeVi64(buffer, r.ttlUs)
}

func (r *ObLoginRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

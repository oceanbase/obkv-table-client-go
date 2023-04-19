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

func (r *LoginResponse) PCode() TablePacketCode {
	return Login
}

func (r *LoginResponse) PayloadLen() int64 {
	// TODO implement me
	panic("implement me")
}

func (r *LoginResponse) PayloadContentLen() int64 {
	// TODO implement me
	panic("implement me")
}

func (r *LoginResponse) SessionId() uint64 {
	return 0
}

func (r *LoginResponse) SetSessionId(sessionId uint64) {
}

func (r *LoginResponse) Credential() []byte {
	return r.credential
}

func (r *LoginResponse) SetCredential(credential []byte) {
	r.credential = credential
}

func (r *LoginResponse) Encode() []byte {
	// TODO implement me
	panic("implement me")
}

func (r *LoginResponse) Decode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Decode(buffer)

	r.serverCapabilities = util.DecodeVi32(buffer)
	_ = util.DecodeVi32(buffer) // reserved1
	_ = util.DecodeVi64(buffer) // reserved2

	r.serverVersion = util.DecodeVString(buffer)
	r.credential = util.DecodeBytesString(buffer)

	r.tenantId = uint64(util.DecodeVi64(buffer))
	r.userId = util.DecodeVi64(buffer)
	r.databaseId = util.DecodeVi64(buffer)
}

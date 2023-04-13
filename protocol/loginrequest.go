package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type LoginRequest struct {
	*UniVersionHeader
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

func NewLoginRequest(tenantName string, databaseName string, userName string, password string) *LoginRequest {
	passScramble := util.GetPasswordScramble(passScrambleLen)
	passSecret := util.ScramblePassword(password, passScramble)

	return &LoginRequest{
		UniVersionHeader:   NewUniVersionHeader(),
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

func (r *LoginRequest) AuthMethod() uint8 {
	return r.authMethod
}

func (r *LoginRequest) SetAuthMethod(authMethod uint8) {
	r.authMethod = authMethod
}

func (r *LoginRequest) ClientType() uint8 {
	return r.clientType
}

func (r *LoginRequest) SetClientType(clientType uint8) {
	r.clientType = clientType
}

func (r *LoginRequest) ClientVersion() uint8 {
	return r.clientVersion
}

func (r *LoginRequest) SetClientVersion(clientVersion uint8) {
	r.clientVersion = clientVersion
}

func (r *LoginRequest) Reversed1() uint8 {
	return r.reversed1
}

func (r *LoginRequest) SetReversed1(reversed1 uint8) {
	r.reversed1 = reversed1
}

func (r *LoginRequest) ClientCapabilities() int32 {
	return r.clientCapabilities
}

func (r *LoginRequest) SetClientCapabilities(clientCapabilities int32) {
	r.clientCapabilities = clientCapabilities
}

func (r *LoginRequest) MaxPacketSize() int32 {
	return r.maxPacketSize
}

func (r *LoginRequest) SetMaxPacketSize(maxPacketSize int32) {
	r.maxPacketSize = maxPacketSize
}

func (r *LoginRequest) Reversed2() int32 {
	return r.reversed2
}

func (r *LoginRequest) SetReversed2(reversed2 int32) {
	r.reversed2 = reversed2
}

func (r *LoginRequest) Reversed3() int64 {
	return r.reversed3
}

func (r *LoginRequest) SetReversed3(reversed3 int64) {
	r.reversed3 = reversed3
}

func (r *LoginRequest) TenantName() string {
	return r.tenantName
}

func (r *LoginRequest) SetTenantName(tenantName string) {
	r.tenantName = tenantName
}

func (r *LoginRequest) UserName() string {
	return r.userName
}

func (r *LoginRequest) SetUserName(userName string) {
	r.userName = userName
}

func (r *LoginRequest) PassSecret() string {
	return r.passSecret
}

func (r *LoginRequest) SetPassSecret(passSecret string) {
	r.passSecret = passSecret
}

func (r *LoginRequest) PassScramble() string {
	return r.passScramble
}

func (r *LoginRequest) SetPassScramble(passScramble string) {
	r.passScramble = passScramble
}

func (r *LoginRequest) DatabaseName() string {
	return r.databaseName
}

func (r *LoginRequest) SetDatabaseName(databaseName string) {
	r.databaseName = databaseName
}

func (r *LoginRequest) TtlUs() int64 {
	return r.ttlUs
}

func (r *LoginRequest) SetTtlUs(ttlUs int64) {
	r.ttlUs = ttlUs
}

func (r *LoginRequest) PCode() TablePacketCode {
	return Login
}

func (r *LoginRequest) PayloadLen() int64 {
	return r.PayloadContentLen() + int64(r.UniVersionHeader.UniVersionHeaderLen()) // Do not change the order
}

func (r *LoginRequest) PayloadContentLen() int64 {
	r.UniVersionHeader.calculateLengthOnce.Do(func() {
		totalLen := 0
		totalLen += 4 // authMethod clientType clientVersion reversed1
		totalLen = totalLen +
			util.NeedLengthByVi32(r.clientCapabilities) +
			util.NeedLengthByVi32(r.maxPacketSize) +
			util.NeedLengthByVi32(r.reversed2) +
			util.NeedLengthByVi64(r.reversed3) +
			util.NeedLengthByVString(r.tenantName) +
			util.NeedLengthByVString(r.userName) +
			util.NeedLengthByVString(r.passSecret) +
			util.NeedLengthByVString(r.passScramble) +
			util.NeedLengthByVString(r.databaseName) +
			util.NeedLengthByVi64(r.ttlUs)
		r.UniVersionHeader.SetContentLength(int64(totalLen)) // Set on first acquisition
	})
	return r.UniVersionHeader.ContentLength()
}

func (r *LoginRequest) SessionId() uint64 {
	return 0
}

func (r *LoginRequest) SetSessionId(sessionId uint64) {
	return
}

func (r *LoginRequest) Credential() string {
	return ""
}

func (r *LoginRequest) SetCredential(credential string) {
	return
}

func (r *LoginRequest) Encode() []byte {
	var index = 0

	payloadLength := r.PayloadLen()
	requestBuf := make([]byte, payloadLength)

	headerLen := r.UniVersionHeader.UniVersionHeaderLen()
	r.UniVersionHeader.Encode(requestBuf[:headerLen])
	index += headerLen

	util.PutUint8(requestBuf[index:index+1], r.authMethod)
	index++
	util.PutUint8(requestBuf[index:index+1], r.clientType)
	index++
	util.PutUint8(requestBuf[index:index+1], r.clientVersion)
	index++
	util.PutUint8(requestBuf[index:index+1], r.reversed1)
	index++

	needLength := util.NeedLengthByVi32(r.clientCapabilities)
	util.EncodeVi32(requestBuf[index:index+needLength], r.clientCapabilities)
	index += needLength

	needLength = util.NeedLengthByVi32(r.maxPacketSize)
	util.EncodeVi32(requestBuf[index:index+needLength], r.maxPacketSize)
	index += needLength

	needLength = util.NeedLengthByVi32(r.reversed2)
	util.EncodeVi32(requestBuf[index:index+needLength], r.reversed2)
	index += needLength

	needLength = util.NeedLengthByVi64(r.reversed3)
	util.EncodeVi64(requestBuf[index:index+needLength], r.reversed3)
	index += needLength

	needLength = util.NeedLengthByVString(r.tenantName)
	util.EncodeVString(requestBuf[index:index+needLength], r.tenantName)
	index += needLength

	needLength = util.NeedLengthByVString(r.userName)
	util.EncodeVString(requestBuf[index:index+needLength], r.userName)
	index += needLength

	needLength = util.NeedLengthByVString(r.passSecret)
	util.EncodeVString(requestBuf[index:index+needLength], r.passSecret)
	index += needLength

	needLength = util.NeedLengthByVString(r.passScramble)
	util.EncodeVString(requestBuf[index:index+needLength], r.passScramble)
	index += needLength

	needLength = util.NeedLengthByVString(r.databaseName)
	util.EncodeVString(requestBuf[index:index+needLength], r.databaseName)
	index += needLength

	needLength = util.NeedLengthByVi64(r.ttlUs)
	util.EncodeVi64(requestBuf[index:index+needLength], r.ttlUs)
	index += needLength

	return requestBuf
}

func (r *LoginRequest) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

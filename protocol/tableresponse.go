package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type TableResponse struct {
	*UniVersionHeader
	errorNo  int32
	sqlState []byte
	msg      []byte
}

func NewTableResponse() *TableResponse {
	return &TableResponse{
		UniVersionHeader: NewUniVersionHeader(),
		errorNo:          0,
		sqlState:         nil,
		msg:              nil,
	}
}

func (r *TableResponse) ErrorNo() int32 {
	return r.errorNo
}

func (r *TableResponse) SetErrorNo(errorNo int32) {
	r.errorNo = errorNo
}

func (r *TableResponse) SqlState() []byte {
	return r.sqlState
}

func (r *TableResponse) SetSqlState(sqlState []byte) {
	r.sqlState = sqlState
}

func (r *TableResponse) Msg() []byte {
	return r.msg
}

func (r *TableResponse) SetMsg(msg []byte) {
	r.msg = msg
}

func (r *TableResponse) PayloadLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *TableResponse) PayloadContentLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *TableResponse) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

func (r *TableResponse) Decode(buffer *bytes.Buffer) {
	r.UniVersionHeader.Decode(buffer)

	r.errorNo = util.DecodeVi32(buffer)

	r.sqlState = util.DecodeBytes(buffer)

	r.msg = util.DecodeBytes(buffer)
}

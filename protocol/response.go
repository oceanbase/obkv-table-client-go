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

type ObTableResponse struct {
	ObUniVersionHeader
	errorNo  int32
	sqlState []byte
	msg      []byte
}

func NewObTableResponse() *ObTableResponse {
	return &ObTableResponse{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		errorNo:  0,
		sqlState: nil,
		msg:      nil,
	}
}

func (r *ObTableResponse) ErrorNo() int32 {
	return r.errorNo
}

func (r *ObTableResponse) SetErrorNo(errorNo int32) {
	r.errorNo = errorNo
}

func (r *ObTableResponse) SqlState() []byte {
	return r.sqlState
}

func (r *ObTableResponse) SetSqlState(sqlState []byte) {
	r.sqlState = sqlState
}

func (r *ObTableResponse) Msg() []byte {
	return r.msg
}

func (r *ObTableResponse) SetMsg(msg []byte) {
	r.msg = msg
}

func (r *ObTableResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableResponse) PayloadContentLen() int {
	totalLen := util.EncodedLengthByVi32(r.errorNo) +
		util.EncodedLengthByBytes(r.sqlState) +
		util.EncodedLengthByBytes(r.msg)

	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi32(buffer, r.errorNo)

	util.EncodeBytes(buffer, r.sqlState)

	util.EncodeBytes(buffer, r.msg)
}

func (r *ObTableResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	r.errorNo = util.DecodeVi32(buffer)

	r.sqlState = util.DecodeBytes(buffer)

	r.msg = util.DecodeBytes(buffer)
}

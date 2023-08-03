/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
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

type ObTableQueryResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	propertiesNames []string
	rowCount        int64
	propertiesRows  [][]*ObObject
}

func NewObTableQueryResponse() *ObTableQueryResponse {
	return &ObTableQueryResponse{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		ObPayloadBase: ObPayloadBase{
			uniqueId:  0,
			sequence:  0,
			tenantId:  1,
			sessionId: 0,
			flag:      0,
			timeout:   10 * 1000 * time.Millisecond,
		},
		propertiesNames: nil,
		rowCount:        0,
		propertiesRows:  nil,
	}
}

func (r *ObTableQueryResponse) PropertiesNames() []string {
	return r.propertiesNames
}

func (r *ObTableQueryResponse) SetPropertiesNames(propertiesNames []string) {
	r.propertiesNames = propertiesNames
}

func (r *ObTableQueryResponse) RowCount() int64 {
	return r.rowCount
}

func (r *ObTableQueryResponse) SetRowCount(rowCount int64) {
	r.rowCount = rowCount
}

func (r *ObTableQueryResponse) PropertiesRows() [][]*ObObject {
	return r.propertiesRows
}

func (r *ObTableQueryResponse) SetPropertiesRows(propertiesRows [][]*ObObject) {
	r.propertiesRows = propertiesRows
}

func (r *ObTableQueryResponse) PCode() ObTablePacketCode {
	return ObTableApiExecuteQuery
}

func (r *ObTableQueryResponse) PayloadLen() int {
	return r.PayloadContentLen() + r.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (r *ObTableQueryResponse) PayloadContentLen() int {
	totalLen := 0
	totalLen += util.EncodedLengthByVi64(int64(len(r.propertiesNames)))
	for _, propertiesName := range r.propertiesNames {
		totalLen += util.EncodedLengthByVString(propertiesName)
	}

	totalLen += util.EncodedLengthByVi64(r.rowCount)
	totalLen += util.EncodedLengthByVi64(int64(len(r.propertiesRows)))
	for _, propertiesRow := range r.propertiesRows {
		for _, obObject := range propertiesRow {
			totalLen += obObject.EncodedLength()
		}
	}
	r.ObUniVersionHeader.SetContentLength(totalLen)
	return r.ObUniVersionHeader.ContentLength()
}

func (r *ObTableQueryResponse) Credential() []byte {
	return nil
}

func (r *ObTableQueryResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableQueryResponse) Encode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(r.propertiesNames)))

	for _, propertiesName := range r.propertiesNames {
		util.EncodeVString(buffer, propertiesName)
	}

	util.EncodeVi64(buffer, r.rowCount)

	util.EncodeVi64(buffer, int64(len(r.propertiesRows)))

	for _, propertiesRow := range r.propertiesRows {
		for _, obObject := range propertiesRow {
			obObject.Encode(buffer)
		}
	}
}

func (r *ObTableQueryResponse) Decode(buffer *bytes.Buffer) {
	r.ObUniVersionHeader.Decode(buffer)

	propertiesNamesLen := util.DecodeVi64(buffer)

	r.propertiesNames = make([]string, 0, propertiesNamesLen)
	var i int64
	for i = 0; i < propertiesNamesLen; i++ {
		r.propertiesNames = append(r.propertiesNames, util.DecodeVString(buffer))
	}

	r.rowCount = util.DecodeVi64(buffer)

	_ = util.DecodeVi64(buffer) // ObDataBuffer

	r.propertiesRows = make([][]*ObObject, 0, r.rowCount)
	for i = 0; i < r.rowCount; i++ {
		row := make([]*ObObject, 0, len(r.propertiesNames))
		for j := 0; j < len(r.propertiesNames); j++ {
			obObject := NewObObject()
			obObject.Decode(buffer)
			row = append(row, obObject)
		}
		r.propertiesRows = append(r.propertiesRows, row)
	}
}

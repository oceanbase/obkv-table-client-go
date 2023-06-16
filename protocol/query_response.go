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

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableQueryResponse struct {
	ObUniVersionHeader
	ObPayloadBase
	propertiesNames []string
	rowCount        int64
	propertiesRows  [][]*ObObject
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
	// TODO implement me
	panic("implement me")
}

func (r *ObTableQueryResponse) PayloadContentLen() int {
	// TODO implement me
	panic("implement me")
}

func (r *ObTableQueryResponse) Credential() []byte {
	return nil
}

func (r *ObTableQueryResponse) SetCredential(credential []byte) {
	return
}

func (r *ObTableQueryResponse) Encode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
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

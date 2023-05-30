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

type ObTableBatchOperation struct {
	ObUniVersionHeader
	obTableOperations   []*ObTableOperation
	readOnly            bool
	sameType            bool
	samePropertiesNames bool
}

func NewObTableBatchOperation() *ObTableBatchOperation {
	return &ObTableBatchOperation{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		obTableOperations:   nil,
		readOnly:            true,
		sameType:            true,
		samePropertiesNames: false,
	}
}

func (o *ObTableBatchOperation) ObTableOperations() []*ObTableOperation {
	return o.obTableOperations
}

func (o *ObTableBatchOperation) SetObTableOperations(obTableOperations []*ObTableOperation) {
	o.obTableOperations = obTableOperations
}

func (o *ObTableBatchOperation) AppendObTableOperation(obTableOperation *ObTableOperation) {
	o.obTableOperations = append(o.obTableOperations, obTableOperation)
	if o.readOnly && obTableOperation.opType != ObTableOperationGet {
		o.readOnly = false
	}

	length := len(o.obTableOperations)
	if o.sameType && length > 1 && o.obTableOperations[length-1].opType != o.obTableOperations[length-2].opType {
		o.sameType = false
	}
}

func (o *ObTableBatchOperation) ReadOnly() bool {
	return o.readOnly
}

func (o *ObTableBatchOperation) SetReadOnly(readOnly bool) {
	o.readOnly = readOnly
}

func (o *ObTableBatchOperation) SameType() bool {
	return o.sameType
}

func (o *ObTableBatchOperation) SetSameType(sameType bool) {
	o.sameType = sameType
}

func (o *ObTableBatchOperation) SamePropertiesNames() bool {
	return o.samePropertiesNames
}

func (o *ObTableBatchOperation) SetSamePropertiesNames(samePropertiesNames bool) {
	o.samePropertiesNames = samePropertiesNames
}

func (o *ObTableBatchOperation) PayloadLen() int {
	return o.PayloadContentLen() + o.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (o *ObTableBatchOperation) PayloadContentLen() int {
	totalLen := 0

	totalLen += util.EncodedLengthByVi64(int64(len(o.obTableOperations)))

	for _, tableOperation := range o.obTableOperations {
		totalLen += tableOperation.PayloadLen()
	}

	totalLen += 3 // readOnly sameType samePropertiesNames

	o.ObUniVersionHeader.SetContentLength(totalLen)
	return o.ObUniVersionHeader.ContentLength()
}

func (o *ObTableBatchOperation) Encode(buffer *bytes.Buffer) {
	o.ObUniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(o.obTableOperations)))

	for _, tableOperation := range o.obTableOperations {
		tableOperation.Encode(buffer)
	}

	util.PutUint8(buffer, util.BoolToByte(o.readOnly))

	util.PutUint8(buffer, util.BoolToByte(o.sameType))

	util.PutUint8(buffer, util.BoolToByte(o.samePropertiesNames))
}

func (o *ObTableBatchOperation) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

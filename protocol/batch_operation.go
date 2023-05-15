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

type TableBatchOperation struct {
	*UniVersionHeader
	tableOperations     []*TableOperation
	readOnly            bool
	sameType            bool
	samePropertiesNames bool
}

func NewTableBatchOperation() *TableBatchOperation {
	return &TableBatchOperation{
		UniVersionHeader:    NewUniVersionHeader(),
		tableOperations:     nil,
		readOnly:            false,
		sameType:            false,
		samePropertiesNames: false,
	}
}

func (o *TableBatchOperation) TableOperations() []*TableOperation {
	return o.tableOperations
}

func (o *TableBatchOperation) SetTableOperations(tableOperations []*TableOperation) {
	o.tableOperations = tableOperations
}

func (o *TableBatchOperation) AppendTableOperation(tableOperation *TableOperation) {
	o.tableOperations = append(o.tableOperations, tableOperation)
}

func (o *TableBatchOperation) ReadOnly() bool {
	return o.readOnly
}

func (o *TableBatchOperation) SetReadOnly(readOnly bool) {
	o.readOnly = readOnly
}

func (o *TableBatchOperation) SameType() bool {
	return o.sameType
}

func (o *TableBatchOperation) SetSameType(sameType bool) {
	o.sameType = sameType
}

func (o *TableBatchOperation) SamePropertiesNames() bool {
	return o.samePropertiesNames
}

func (o *TableBatchOperation) SetSamePropertiesNames(samePropertiesNames bool) {
	o.samePropertiesNames = samePropertiesNames
}

func (o *TableBatchOperation) PayloadLen() int {
	return o.PayloadContentLen() + o.UniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (o *TableBatchOperation) PayloadContentLen() int {
	totalLen := 0

	totalLen += util.EncodedLengthByVi64(int64(len(o.tableOperations)))

	for _, tableOperation := range o.tableOperations {
		totalLen += tableOperation.PayloadLen()
	}

	totalLen += 3 // readOnly sameType samePropertiesNames

	o.UniVersionHeader.SetContentLength(totalLen)
	return o.UniVersionHeader.ContentLength()
}

func (o *TableBatchOperation) Encode(buffer *bytes.Buffer) {
	o.UniVersionHeader.Encode(buffer)

	util.EncodeVi64(buffer, int64(len(o.tableOperations)))

	for _, tableOperation := range o.tableOperations {
		tableOperation.Encode(buffer)
	}

	util.PutUint8(buffer, util.BoolToByte(o.readOnly))

	util.PutUint8(buffer, util.BoolToByte(o.sameType))

	util.PutUint8(buffer, util.BoolToByte(o.samePropertiesNames))
}

func (o *TableBatchOperation) Decode(buffer *bytes.Buffer) {
	// TODO implement me
	panic("implement me")
}

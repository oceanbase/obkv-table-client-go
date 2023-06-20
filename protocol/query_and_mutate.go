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

type ObTableQueryAndMutate struct {
	ObUniVersionHeader
	tableQuery           *ObTableQuery
	mutations            *ObTableBatchOperation
	returnAffectedEntity bool
}

func NewObTableQueryAndMutate() *ObTableQueryAndMutate {
	return &ObTableQueryAndMutate{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		tableQuery:           NewObTableQuery(),
		mutations:            NewObTableBatchOperation(),
		returnAffectedEntity: false,
	}
}

func (q *ObTableQueryAndMutate) TableQuery() *ObTableQuery {
	return q.tableQuery
}

func (q *ObTableQueryAndMutate) SetTableQuery(tableQuery *ObTableQuery) {
	q.tableQuery = tableQuery
}

func (q *ObTableQueryAndMutate) Mutations() *ObTableBatchOperation {
	return q.mutations
}

func (q *ObTableQueryAndMutate) SetMutations(mutations *ObTableBatchOperation) {
	q.mutations = mutations
}

func (q *ObTableQueryAndMutate) ReturnAffectedEntity() bool {
	return q.returnAffectedEntity
}

func (q *ObTableQueryAndMutate) SetReturnAffectedEntity(returnAffectedEntity bool) {
	q.returnAffectedEntity = returnAffectedEntity
}

func (q *ObTableQueryAndMutate) PayloadLen() int {
	return q.PayloadContentLen() + q.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (q *ObTableQueryAndMutate) PayloadContentLen() int {
	totalLen := 0
	totalLen += q.tableQuery.PayloadLen() +
		q.mutations.PayloadLen() +
		1 // returnAffectedEntity
	q.ObUniVersionHeader.SetContentLength(totalLen)
	return q.ObUniVersionHeader.ContentLength()
}

func (q *ObTableQueryAndMutate) Encode(buffer *bytes.Buffer) {
	q.ObUniVersionHeader.Encode(buffer)

	q.tableQuery.Encode(buffer)

	q.mutations.Encode(buffer)

	util.PutUint8(buffer, util.BoolToByte(q.returnAffectedEntity))
}

func (q *ObTableQueryAndMutate) Decode(buffer *bytes.Buffer) {
	q.ObUniVersionHeader.Decode(buffer)

	q.tableQuery.Decode(buffer)

	q.mutations.Decode(buffer)

	q.returnAffectedEntity = util.ByteToBool(util.Uint8(buffer))
}

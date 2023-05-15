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

package client

import "strconv"

type ObTableParam struct {
	table       *ObTable
	tableId     uint64
	partitionId int64 // partition id in 3.x aka tablet id in 4.x
}

func NewObTableParam(table *ObTable, tableId uint64, partitionId int64) *ObTableParam {
	return &ObTableParam{table, tableId, partitionId}
}

func (p *ObTableParam) String() string {
	tableStr := "nil"
	if p.table != nil {
		tableStr = p.table.String()
	}
	return "ObTableParam{" +
		"table:" + tableStr + ", " +
		"tableId:" + strconv.Itoa(int(p.tableId)) + ", " +
		"partitionId:" + strconv.Itoa(int(p.partitionId)) +
		"}"
}

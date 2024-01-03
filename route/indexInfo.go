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

package route

import (
	"github.com/oceanbase/obkv-table-client-go/protocol"
)

type ObIndexInfo struct {
	dataTableId    uint64
	indexTableId   uint64
	indexTableName string
	indexType      protocol.ObIndexType
}

func (i *ObIndexInfo) IndexTableName() string {
	return i.indexTableName
}

func (i *ObIndexInfo) IndexTableId() uint64 {
	return i.indexTableId
}

func (i *ObIndexInfo) DataTableId() uint64 {
	return i.dataTableId
}

func (i *ObIndexInfo) IndexType() protocol.ObIndexType {
	return i.indexType
}

func NewObIndexInfo() *ObIndexInfo {
	return &ObIndexInfo{
		dataTableId:    0,
		indexTableId:   0,
		indexTableName: "",
		indexType:      protocol.IndexTypeIsNot,
	}
}

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
	"github.com/oceanbase/obkv-table-client-go/table"
)

// obColumn represents a column in the table and contains the column meta-information.
// obColumn represents obSimpleColumn or obGeneratedColumn.
type obColumn interface {
	ColumnName() string
	ObjType() protocol.ObObjType
	CollationType() protocol.ObCollationType
	// EvalValue calculate the value of the partition column.
	eval(rowKey []*table.Column) (interface{}, error)
	extractColumn(rowKey []*table.Column) (*table.Column, error)
	String() string
}

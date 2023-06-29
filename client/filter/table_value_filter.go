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

package filter

import (
	"strings"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObTableValueFilter struct {
	op         ObCompareOperator
	columnName string
	value      interface{}
}

func NewObTableValueFilter(op ObCompareOperator, columnName string, value interface{}) *ObTableValueFilter {
	return &ObTableValueFilter{
		op:         op,
		columnName: columnName,
		value:      value,
	}
}

func (f *ObTableValueFilter) Set(op ObCompareOperator, columnName string, value interface{}) {
	f.op = op
	f.columnName = columnName
	f.value = value
}

func (f *ObTableValueFilter) ColumnName() string {
	return f.columnName
}

func (f *ObTableValueFilter) String() string {
	var builder strings.Builder
	if f.columnName == "" {
		return ""
	}
	builder.WriteString(tableCompareFilter)
	builder.WriteString("(")
	builder.WriteString(f.op.String())
	builder.WriteString(", '")
	builder.WriteString(f.columnName)
	builder.WriteString(":")
	if f.value != nil {
		builder.WriteString(util.InterfaceToString(f.value))
	}
	builder.WriteString("')")
	return builder.String()
}

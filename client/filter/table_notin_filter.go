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

type ObTableNotInFilter struct {
	columnName string
	values     []interface{}
}

func NewObTableNotInFilter(columnName string, values ...interface{}) *ObTableNotInFilter {
	return &ObTableNotInFilter{
		columnName: columnName,
		values:     values,
	}
}

func (f *ObTableNotInFilter) ColumnName() string {
	return f.columnName
}

func (f *ObTableNotInFilter) String() string {
	filterList := AndList()
	for _, value := range f.values {
		filterList.AddFilter(CompareVal(NotEqual, f.columnName, value))
	}
	return filterList.String()
}

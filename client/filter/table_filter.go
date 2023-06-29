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

type ObTableFilter interface {
	String() string
}

const tableCompareFilter = "TableCompareFilter"

func AndList(filters ...ObTableFilter) *ObTableFilterList {
	return NewObTableFilterListWithOperatorAndTableFilter(OperatorAnd, filters...)
}

func OrList(filters ...ObTableFilter) *ObTableFilterList {
	return NewObTableFilterListWithOperatorAndTableFilter(OperatorOr, filters...)
}

func CompareVal(op ObCompareOperator, columnName string, value interface{}) *ObTableValueFilter {
	return NewObTableValueFilter(op, columnName, value)
}

func In(columnName string, values ...interface{}) *ObTableInFilter {
	return NewObTableInFilter(columnName, values...)
}

func NotIn(columnName string, values ...interface{}) *ObTableNotInFilter {
	return NewObTableNotInFilter(columnName, values...)
}

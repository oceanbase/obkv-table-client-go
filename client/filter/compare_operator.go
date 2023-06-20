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

type ObCompareOperator uint8

const (
	LessThan ObCompareOperator = iota
	GreaterThan
	LessOrEqualThan
	GreaterOrEqualThan
	NotEqual
	Equal
	IsNull
	IsNotNull
)

const (
	LessThanStr           string = "<"
	GreaterThanStr        string = ">"
	LessOrEqualThanStr    string = "<="
	GreaterOrEqualThanStr string = ">="
	NotEqualStr           string = "!="
	EqualStr              string = "="
	IsNullStr             string = "IS"
	IsNotNullStr          string = "IS_NOT"
)

var ObCompareOperatorStrings = []string{
	LessThan:           LessThanStr,
	GreaterThan:        GreaterThanStr,
	LessOrEqualThan:    LessOrEqualThanStr,
	GreaterOrEqualThan: GreaterOrEqualThanStr,
	NotEqual:           NotEqualStr,
	Equal:              EqualStr,
	IsNull:             IsNullStr,
	IsNotNull:          IsNotNullStr,
}

func (o ObCompareOperator) String() string {
	if o > IsNotNull {
		return ""
	}
	return ObCompareOperatorStrings[o]
}

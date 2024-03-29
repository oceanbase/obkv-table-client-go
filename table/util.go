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

package table

// ColumnsToString converts a column to a string
func ColumnsToString(columns []*Column) string {
	var str string
	str = str + "["
	for i := 0; i < len(columns); i++ {
		if i > 0 {
			str += ", "
		}
		str += columns[i].String()
	}
	str += "]"
	return str
}

// RangePairsToString converts a ranges pair to a string
func RangePairsToString(rangesPairs []*RangePair) string {
	var str string
	str = str + "["
	for i := 0; i < len(rangesPairs); i++ {
		if i > 0 {
			str += ", "
		}
		str += rangesPairs[i].String()
	}
	str += "]"
	return str
}

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

type AggregateResult interface {
	Value(columnName string) interface{}
	Values() []interface{}
}

func newObAggregateResult(aggResultSet QueryResult) *obAggregateResult {
	return &obAggregateResult{aggResultSet}
}

type obAggregateResult struct {
	aggResult QueryResult
}

// Value returns the value of the specified column.
func (r *obAggregateResult) Value(columnName string) interface{} {
	if r.aggResult == nil {
		return nil
	}
	return r.aggResult.Value(columnName)
}

// Values returns all values in the query result.
func (r *obAggregateResult) Values() []interface{} {
	if r.aggResult == nil {
		return nil
	}
	return r.aggResult.Values()
}

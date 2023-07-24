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
	// Value returns the value of the specified aggregate name, such as Value("min(c1)").
	Value(columnName string) interface{}
	// Values returns all aggregation values.
	Values() []interface{}
}

func newObAggregateResult(aggResultSet QueryResult) *obAggregateResult {
	return &obAggregateResult{aggResultSet}
}

type obAggregateResult struct {
	result QueryResult
}

// Value returns the value of the specified aggregate name, such as Value("min(c1)").
func (r *obAggregateResult) Value(aggregateName string) interface{} {
	if r.result == nil {
		return nil
	}
	return r.result.Value(aggregateName)
}

// Values returns all aggregation values.
func (r *obAggregateResult) Values() []interface{} {
	if r.result == nil {
		return nil
	}
	return r.result.Values()
}

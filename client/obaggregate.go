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

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/pkg/errors"
)

func newObAggExecutor(queryExecutor *obQueryExecutor) *obAggExecutor {
	return &obAggExecutor{
		queryExecutor:     queryExecutor,
		aggOperations:     make([]*protocol.ObTableAggregation, 0),
		aggOperationNames: make([]string, 0),
	}
}

type obAggExecutor struct {
	queryExecutor     *obQueryExecutor
	aggOperations     []*protocol.ObTableAggregation
	aggOperationNames []string
}

// Min add a min operation to the agg executor.
func (q *obAggExecutor) Min(aggColumn string) AggExecutor {
	q.aggOperations = append(q.aggOperations, protocol.NewObTableAggregationWithParams(protocol.ObTableAggregationTypeMin, aggColumn))
	q.aggOperationNames = append(q.aggOperationNames, "min("+aggColumn+")")
	return q
}

// Max add a max operation to the agg executor.
func (q *obAggExecutor) Max(aggColumn string) AggExecutor {
	q.aggOperations = append(q.aggOperations, protocol.NewObTableAggregationWithParams(protocol.ObTableAggregationTypeMax, aggColumn))
	q.aggOperationNames = append(q.aggOperationNames, "max("+aggColumn+")")
	return q
}

// Count add a count operation to the agg executor.
func (q *obAggExecutor) Count() AggExecutor {
	q.aggOperations = append(q.aggOperations, protocol.NewObTableAggregationWithParams(protocol.ObTableAggregationTypeCount, "*"))
	q.aggOperationNames = append(q.aggOperationNames, "count(*)")
	return q
}

// Sum add a sum operation to the agg executor.
func (q *obAggExecutor) Sum(aggColumn string) AggExecutor {
	q.aggOperations = append(q.aggOperations, protocol.NewObTableAggregationWithParams(protocol.ObTableAggregationTypeSum, aggColumn))
	q.aggOperationNames = append(q.aggOperationNames, "sum("+aggColumn+")")
	return q
}

// Avg add an avg operation to the agg executor.
func (q *obAggExecutor) Avg(aggColumn string) AggExecutor {
	q.aggOperations = append(q.aggOperations, protocol.NewObTableAggregationWithParams(protocol.ObTableAggregationTypeAvg, aggColumn))
	q.aggOperationNames = append(q.aggOperationNames, "avg("+aggColumn+")")
	return q
}

// setParams set param to the agg executor.
func (q *obAggExecutor) setParams() {
	q.queryExecutor.tableQuery.SetSelectColumns(q.aggOperationNames)
	q.queryExecutor.tableQuery.SetAggregations(q.aggOperations)
	return
}

// Execute an agg operation.
// AggregateResult contains the results of all operations.
func (q *obAggExecutor) Execute(ctx context.Context) (AggregateResult, error) {
	if len(q.aggOperations) == 0 {
		return nil, errors.New("empty aggregation operations")
	}
	q.setParams()
	resSet, err := q.queryExecutor.init(ctx)
	if err != nil {
		return nil, errors.WithMessagef(err, "init query executor failed")
	}
	res, err := resSet.Next()
	if err != nil {
		return nil, errors.WithMessagef(err, "get aggregate result failed")
	}
	return newObAggregateResult(res), nil
}

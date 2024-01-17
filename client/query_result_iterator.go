/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package client

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/log"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/route"
)

type QueryResultIterator interface {
	IsClosed() bool
	Close()
	Next() (QueryResult, error)
	NextBatch() ([]QueryResult, error)
}

type ObQueryResultIterator struct {
	ctx                   context.Context
	queryExecutor         *obQueryExecutor
	lock                  sync.Mutex
	cachedPropertiesRows  [][]*protocol.ObObject
	cachedPropertiesNames []string
	targetParts           []*route.ObTableParam
	readConsistency       protocol.ObTableConsistencyLevel
	prevSessionId         int64
	rowIndex              int
	closed                bool
	hasNext               bool
}

// newObQueryResultWithParam creates a new ObQueryResultIterator.
func newObQueryResultIteratorWithParams(
	ctx context.Context,
	queryExecutor *obQueryExecutor,
	targetParts []*route.ObTableParam) *ObQueryResultIterator {
	return &ObQueryResultIterator{
		ctx:                   ctx,
		queryExecutor:         queryExecutor,
		cachedPropertiesRows:  nil,
		cachedPropertiesNames: nil,
		targetParts:           targetParts,
		readConsistency:       protocol.ObTableConsistencyLevelStrong,
		prevSessionId:         0,
		rowIndex:              0,
		closed:                false,
		hasNext:               true,
	}
}

// IsClosed returns true if the query result is closed.
func (q *ObQueryResultIterator) IsClosed() bool {
	return q.closed
}

// Close closes the query result iterator.
func (q *ObQueryResultIterator) Close() {
	// TODO: Send a close request to the server
	q.closed = true
}

// Next returns nil if there is no more row.
// Notes: Next() could not be called with NextBatch() at the same time.
// We strongly recommend you to call Next().
func (q *ObQueryResultIterator) Next() (QueryResult, error) {
	// check status
	err := q.checkStatus()
	if err != nil {
		return nil, err
	}
	var startTime int64 = time.Now().UnixMilli()
	// lock
	q.lock.Lock()
	defer q.lock.Unlock()

	// get next row from cache
	if q.cachedPropertiesRows != nil && q.rowIndex < len(q.cachedPropertiesRows) {
		row := q.cachedPropertiesRows[q.rowIndex]
		q.rowIndex++
		return newObQueryResult(q.cachedPropertiesNames, row), nil
	}

	// get next row from previous server
	// if prevSessionId != 0, it means that the previous server has not been read completely
	if q.prevSessionId != 0 {
		err = q.fetchNextWithRetry(true)
		if err != nil {
			return nil, err
		}
		if !q.hasNext {
			return nil, errors.WithMessage(err, "fetch next row from previous server failed")
		}
		row := q.cachedPropertiesRows[q.rowIndex]
		q.rowIndex++
		return newObQueryResult(q.cachedPropertiesNames, row), nil
	}

	// get next row from next server
	err = q.fetchNextWithRetry(false)
	endTime := time.Now().UnixMilli()
	var duration int64 = endTime - startTime
	MonitorSlowQuery(duration, log.SlowQueryThreshold, q.queryExecutor.tableName, q.ctx.Value(log.ObkvTraceIdName))
	if err != nil {
		return nil, err
	}
	if !q.hasNext {
		// no more row
		return nil, nil
	}
	row := q.cachedPropertiesRows[q.rowIndex]
	q.rowIndex++
	return newObQueryResult(q.cachedPropertiesNames, row), nil
}

// NextBatch returns nil if there is no more row.
// Notes: NextBatch() could not be called with Next() at the same time.
// Number of returned rows is less than or equal to the batch size.
// Batch size, limit, amount of data will affect the number of returned rows.
func (q *ObQueryResultIterator) NextBatch() ([]QueryResult, error) {
	// check status
	err := q.checkStatus()
	if err != nil {
		return nil, err
	}

	// lock
	q.lock.Lock()
	defer q.lock.Unlock()

	// get next row from previous server
	// if prevSessionId != 0, it means that the previous server has not been read completely
	err = q.fetchNextWithRetry(q.prevSessionId != 0)
	if err != nil {
		return nil, err
	}
	if !q.hasNext {
		// no more row
		return nil, nil
	}
	result := make([]QueryResult, 0, len(q.cachedPropertiesRows))
	for _, row := range q.cachedPropertiesRows {
		result = append(result, newObQueryResult(q.cachedPropertiesNames, row))
	}
	return result, nil
}

// fetchNextWithRetry fetches the next batch from the server, retry when route table change
func (q *ObQueryResultIterator) fetchNextWithRetry(hasPrev bool) error {
	err, needRetry := q.fetchNext(hasPrev)
	for err != nil && needRetry {
		select {
		case <-q.ctx.Done():
			return errors.WithMessage(err, "retry and timeout")
		default:
			// get and set route table again
			targetParts, err := q.queryExecutor.getTableParams(q.ctx, q.queryExecutor.tableName, q.queryExecutor.keyRanges)
			if err != nil {
				return err
			}
			q.targetParts = targetParts

			err, needRetry = q.fetchNext(hasPrev)
			time.Sleep(1 * time.Millisecond)
		}
	}

	if err != nil {
		return err
	}

	return nil
}

// fetchNext fetches the next batch from the server.
func (q *ObQueryResultIterator) fetchNext(hasPrev bool) (error, bool) {
	needRetry := false
	// check status
	err := q.checkStatus()
	if err != nil {
		return err, needRetry
	}

	// loop to get batch from servers
	cacheRows := int64(0)
	result := protocol.NewObTableAsyncQueryResponse()
	for cacheRows == 0 {
		if len(q.targetParts) == 0 {
			break
		}
		nextParam := q.targetParts[0]

		// prepare request
		queryRequest := protocol.NewObTableQueryRequestWithParams(
			q.queryExecutor.tableName,
			nextParam.TableId(),
			nextParam.PartitionId(),
			q.queryExecutor.entityType,
			q.queryExecutor.tableQuery,
		)
		asyncQueryRequest := protocol.NewObTableAsyncQueryRequestWithParams(
			queryRequest,
			q.queryExecutor.cli.config.OperationTimeOut,
			q.queryExecutor.cli.GetRpcFlag(),
		)
		if hasPrev {
			asyncQueryRequest.SetQueryType(protocol.QueryNext)
			asyncQueryRequest.SetQuerySessionId(q.prevSessionId)
		} else {
			asyncQueryRequest.SetQueryType(protocol.QueryStart)
		}

		// execute
		err, needRetry = q.queryExecutor.cli.executeInternal(q.ctx, q.queryExecutor.tableName, nextParam.Table(), asyncQueryRequest, result)
		if err != nil {
			return err, needRetry
		}
		// deal with result, update status
		cacheRows = result.ResultRowCount()
		if cacheRows > 0 {
			// cache result
			q.cachedPropertiesNames = result.PropertiesNames()
			q.cachedPropertiesRows = result.PropertiesRows()
			// update status
			q.rowIndex = 0
			if result.IsEnd() {
				// remove current server that has been read completely
				q.prevSessionId = 0
				q.targetParts = q.targetParts[1:]
			} else {
				// current server has not been read completely
				q.prevSessionId = result.QuerySessionId()
			}
		} else {
			// remove current server
			q.targetParts = q.targetParts[1:]
			q.prevSessionId = 0
			hasPrev = false
		}
	}

	if cacheRows == 0 {
		q.hasNext = false
		return nil, needRetry
	}

	return nil, needRetry
}

// checkStatus checks the status of the query result.
func (q *ObQueryResultIterator) checkStatus() error {
	if q.closed {
		return errors.Errorf("ObQueryResult is closed")
	}
	return nil
}

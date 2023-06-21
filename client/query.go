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
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
	"github.com/pkg/errors"
)

type ObQueryExecutor struct {
	tableName  string
	cli        *obClient
	keyRanges  []*table.RangePair
	tableQuery *protocol.ObTableQuery
	entityType protocol.ObTableEntityType
}

// NewObQueryExecutorWithParams creates a new ObQueryExecutor.
func NewObQueryExecutorWithParams(tableName string, cli *obClient) *ObQueryExecutor {
	return &ObQueryExecutor{
		tableName:  tableName,
		cli:        cli,
		keyRanges:  nil,
		tableQuery: protocol.NewObTableQueryWithParams(-1),
		entityType: protocol.ObTableEntityTypeDynamic,
	}
}

// setTableName returns the table name.
func (q *ObQueryExecutor) setTableName() string {
	return q.tableName
}

// setClient returns the client.
func (q *ObQueryExecutor) setClient() *obClient {
	return q.cli
}

// AddKeyRanges adds key ranges.
func (q *ObQueryExecutor) AddKeyRanges(keyRanges []*table.RangePair) {
	if q.keyRanges == nil {
		q.keyRanges = make([]*table.RangePair, 0, len(keyRanges))
	}
	q.keyRanges = append(q.keyRanges, keyRanges...)
}

// SetEntityTypes sets the entity type.
func (q *ObQueryExecutor) SetEntityTypes(entityType protocol.ObTableEntityType) {
	q.entityType = entityType
}

// setQueryOptions sets the query option.
func (q *ObQueryExecutor) setQueryOptions(queryOptions *ObkvQueryOptions) {
	// TODO: modify queryFilter and set it to tableQuery.
	q.tableQuery.SetFilterString("")
	// TODO: modify hTableFilter and set it to tableQuery.
	q.tableQuery.SetHTableFilter(nil)
	q.tableQuery.SetSelectColumns(queryOptions.selectColumns)
	q.tableQuery.SetIndexName(queryOptions.indexName)
	q.tableQuery.SetBatchSize(queryOptions.batchSize)
	q.tableQuery.SetMaxResultSize(queryOptions.maxResultSize)
	q.tableQuery.SetLimit(queryOptions.limit)
	q.tableQuery.SetOffset(queryOptions.offset)
	q.tableQuery.SetScanOrder(protocol.ObScanOrder(queryOptions.scanOrder))
	q.tableQuery.SetIsHbaseQuery(queryOptions.isHbaseQuery)
	if queryOptions.isHbaseQuery {
		q.entityType = protocol.ObTableEntityTypeHKV
	}
}

// getTableParams returns the table params.
func (q *ObQueryExecutor) getTableParams(
	ctx context.Context,
	tableName string,
	keyRanges []*table.RangePair,
	refresh bool) ([]*ObTableParam, error) {
	entry, err := q.cli.getOrRefreshTableEntry(ctx, tableName, refresh, false)
	if err != nil {
		return nil, errors.WithMessagef(err, "get or refresh table entry, tableName:%s", tableName)
	}

	// get partition ids from key ranges
	partIdList := make([]uint64, 0)
	for _, keyRange := range keyRanges {
		partIds, err := q.cli.getPartitionIds(entry, keyRange)
		if err != nil {
			return nil, errors.WithMessagef(err, "get partition id, tableName:%s, tableEntry:%s", tableName, entry.String())
		}
		partIdList = append(partIdList, partIds...)
	}

	// remove duplicate partIds
	partIds := removeDuplicates(partIdList)

	// construct table params
	tableParams := make([]*ObTableParam, 0, len(partIds))
	for _, partId := range partIds {
		t, err := q.cli.getTable(entry, partId)
		if err != nil {
			return nil, errors.WithMessagef(err, "get table, tableName:%s, tableEntry:%s, partId:%d",
				tableName, entry.String(), partId)
		}

		if util.ObVersion() >= 4 && entry.IsPartitionTable() {
			partId, err = entry.PartitionInfo().GetTabletId(partId)
			if err != nil {
				return nil, errors.WithMessagef(err, "get tablet id, tableName:%s, tableEntry:%s, partId:%d",
					tableName, entry.String(), partId)
			}
		}
		tableParams = append(tableParams, NewObTableParam(t, entry.TableId(), partId))
	}

	return tableParams, nil
}

// removeDuplicates removes duplicates id in partIdList.
func removeDuplicates(nums []uint64) []uint64 {
	set := make(map[uint64]bool)
	var result []uint64
	for _, num := range nums {
		if !set[num] {
			set[num] = true
			result = append(result, num)
		}
	}
	return result
}

// checkQueryParams checks the query params.
func (q *ObQueryExecutor) checkQueryParams() error {
	if q.tableName == "" {
		return errors.New("table name is empty")
	}
	if q.cli == nil {
		return errors.New("client is nil")
	}
	if q.keyRanges == nil {
		return errors.New("key ranges is nil")
	}
	if q.tableQuery == nil {
		return errors.New("table query is nil")
	}
	return nil
}

// transferQueryRange sets the query range into tableQuery.
func (q *ObQueryExecutor) transferQueryRange() error {
	queryRanges := make([]*protocol.ObNewRange, 0, len(q.keyRanges))
	for _, rangePair := range q.keyRanges {
		if len(rangePair.Start()) != len(rangePair.End()) {
			return errors.New("startRange and endRange key length is not equal")
		}
		startObjs := make([]*protocol.ObObject, 0, len(rangePair.Start()))
		endObjs := make([]*protocol.ObObject, 0, len(rangePair.End()))
		for i := 0; i < len(rangePair.Start()); i++ {
			// append start obj
			objMeta, err := protocol.DefaultObjMeta(rangePair.Start()[i].Value())
			if err != nil {
				return errors.WithMessage(err, "create obj meta by Range key")
			}
			startObjs = append(startObjs, protocol.NewObObjectWithParams(objMeta, rangePair.Start()[i].Value()))

			// append end obj
			objMeta, err = protocol.DefaultObjMeta(rangePair.End()[i].Value())
			if err != nil {
				return errors.WithMessage(err, "create obj meta by Range key")
			}
			endObjs = append(endObjs, protocol.NewObObjectWithParams(objMeta, rangePair.End()[i].Value()))
		}
		borderFlag := protocol.NewObBorderFlag()
		if rangePair.IncludeStart() {
			borderFlag.SetInclusiveStart()
		}
		if rangePair.IncludeEnd() {
			borderFlag.SetInclusiveEnd()
		}
		queryRanges = append(queryRanges, protocol.NewObNewRangeWithParams(startObjs, endObjs, borderFlag))
	}
	q.tableQuery.SetKeyRanges(queryRanges)
	return nil
}

// init calculate the expectant and construct the query result.
func (q *ObQueryExecutor) init(ctx context.Context) (*ObQueryResultIterator, error) {
	err := q.checkQueryParams()
	if err != nil {
		return nil, errors.WithMessage(err, "check query params")
	}

	// get table params
	expectant, err := q.getTableParams(ctx, q.tableName, q.keyRanges, false)
	if err != nil {
		return nil, errors.WithMessage(err, "get table params")
	}

	// set query range into table query
	if nil != q.transferQueryRange() {
		return nil, errors.WithMessage(err, "transfer query range")
	}

	return newObQueryResultIteratorWithParams(&ctx, q.cli, q.tableQuery, expectant, q.entityType, q.tableName), nil
}

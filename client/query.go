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
	"fmt"

	"github.com/oceanbase/obkv-table-client-go/client/option"
	"github.com/oceanbase/obkv-table-client-go/route"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type obQueryExecutor struct {
	tableName  string
	cli        *obClient
	keyRanges  []*table.RangePair
	tableQuery *protocol.ObTableQuery
	entityType protocol.ObTableEntityType
}

// newObQueryExecutorWithParams creates a new obQueryExecutor.
func newObQueryExecutorWithParams(tableName string, cli *obClient) *obQueryExecutor {
	return &obQueryExecutor{
		tableName:  tableName,
		cli:        cli,
		keyRanges:  nil,
		tableQuery: protocol.NewObTableQueryWithParams(-1),
		entityType: protocol.ObTableEntityTypeDynamic,
	}
}

// setTableName sets the table name.
func (q *obQueryExecutor) setTableName(tableName string) {
	q.tableName = tableName
}

// setClient sets the client.
func (q *obQueryExecutor) setClient(cli *obClient) {
	q.cli = cli
}

// addKeyRanges adds key ranges.
func (q *obQueryExecutor) addKeyRanges(keyRanges []*table.RangePair) {
	if q.keyRanges == nil {
		q.keyRanges = make([]*table.RangePair, 0, len(keyRanges))
	}
	q.keyRanges = append(q.keyRanges, keyRanges...)
}

// setEntityType sets the entity type.
func (q *obQueryExecutor) setEntityType(entityType protocol.ObTableEntityType) {
	q.entityType = entityType
}

// setQueryOptions sets the query option.
func (q *obQueryExecutor) setQueryOptions(queryOptions *option.ObQueryOptions) {
	if queryOptions.QueryFilter != nil {
		q.tableQuery.SetFilterString(queryOptions.QueryFilter.String())
	}
	q.tableQuery.SetSelectColumns(queryOptions.SelectColumns)
	q.tableQuery.SetIndexName(queryOptions.IndexName)
	q.tableQuery.SetBatchSize(queryOptions.BatchSize)
	q.tableQuery.SetMaxResultSize(queryOptions.MaxResultSize)
	q.tableQuery.SetLimit(queryOptions.Limit)
	q.tableQuery.SetOffset(queryOptions.Offset)
	q.tableQuery.SetScanOrder(protocol.ObScanOrder(queryOptions.ScanOrder))

	if queryOptions.HTableFilter != nil {
		q.tableQuery.SetHTableFilter(queryOptions.HTableFilter.Transfrom2Proto())
		q.tableQuery.SetIsHbaseQuery(true)
		q.entityType = protocol.ObTableEntityTypeHKV
	} else {
		q.tableQuery.SetIsHbaseQuery(queryOptions.IsHbaseQuery)
		switch queryOptions.KeyValueMode {
		case table.DynamicMode:
			q.entityType = protocol.ObTableEntityTypeDynamic
		case table.ObTableMode:
			q.entityType = protocol.ObTableEntityTypeKV
		case table.ObHBaseMode:
			q.entityType = protocol.ObTableEntityTypeHKV
		default:
			q.entityType = protocol.ObTableEntityTypeDynamic
		}
	}
}

// getTableParams returns the table params.
func (q *obQueryExecutor) getTableParams(
	ctx context.Context,
	tableName string,
	keyRanges []*table.RangePair,
	refresh bool) ([]*ObTableParam, error) {
	// odp table
	if q.cli.odpTable != nil {
		return []*ObTableParam{NewObTableParam(q.cli.odpTable, 0, 0)}, nil
	}

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

	// defense for aggregate of multiple parts
	if len(partIds) > 1 && q.tableQuery.IsAggregations() {
		if q.tableQuery.IsAggregations() {
			return nil, errors.New("aggregate multiple partitions")
		}
	}

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
func (q *obQueryExecutor) checkQueryParams() error {
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

// init calculate the targetParts and construct the query result.
func (q *obQueryExecutor) init(ctx context.Context) (*ObQueryResultIterator, error) {
	err := q.checkQueryParams()
	if err != nil {
		return nil, errors.WithMessage(err, "check query params")
	}

	// construct index table name if do index scan
	tableName := q.tableName
	if "" != q.tableQuery.IndexName() {
		indexTableName, err := q.constructIndexTableName(ctx)
		if err != nil {
			return nil, errors.WithMessage(err, "construct index table name")
		}

		info, err := q.cli.getOrRefreshIndexInfo(ctx, indexTableName)
		if err != nil {
			return nil, errors.WithMessage(err, "get index info fail")
		}

		if info.IndexType().IsGlobalIndex() {
			tableName = indexTableName
		}
	}

	// get table params
	targetParts, err := q.getTableParams(ctx, tableName, q.keyRanges, false)
	if err != nil {
		return nil, errors.WithMessage(err, "get table params")
	}

	// set query range into table query
	keyRanges, err := TransferQueryRange(q.keyRanges)
	if err != nil {
		return nil, errors.WithMessage(err, "transfer query range")
	}
	q.tableQuery.SetKeyRanges(keyRanges)

	return newObQueryResultIteratorWithParams(ctx, q.cli, q.tableQuery, targetParts, q.entityType, q.tableName), nil
}

// constructIndexTableName construct index table name, get main table id from cache or remote.
func (q *obQueryExecutor) constructIndexTableName(ctx context.Context) (string, error) {
	var tableId uint64
	var err error
	entry := q.cli.getTableEntryFromCache(q.tableName)
	if entry == nil { // do dml in sql, do query in obkv, entry is null
		addr := q.cli.serverRoster.GetServer()
		tableId, err = route.GetTableIdFromRemote(ctx, addr, q.cli.sysUA, q.tableName)
		if err != nil {
			return "", errors.WithMessagef(err, "get table id from remote, tableName:%s", q.tableName)
		}
	} else {
		tableId = entry.TableId()
	}

	// [__idx_][data_table_id][_index_name]
	return fmt.Sprintf("__idx_%d_%s", tableId, q.tableQuery.IndexName()), nil
}

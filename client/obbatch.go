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
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type obBatchExecutor struct {
	tableName string
	batchOps  *protocol.ObTableBatchOperation
	cli       *ObClient
}

func newObBatchExecutor(tableName string, cli *ObClient) *obBatchExecutor {
	return &obBatchExecutor{
		tableName: tableName,
		batchOps:  protocol.NewObTableBatchOperation(),
		cli:       cli,
	}
}

// addDmlOp add dml operation witch include insert/update/insertOrUpdate/replace/increment/append
// operation to batch executor
func (b *obBatchExecutor) addDmlOp(
	opType protocol.ObTableOperationType,
	rowKey []*table.Column,
	mutateValues []*table.Column,
	opts ...ObkvOption) error {
	op, err := protocol.NewObTableOperation(opType, rowKey, mutateValues)
	if err != nil {
		return errors.WithMessagef(err, "new table operation, opType:%d, tableName:%s, rowKey:%s, mutateValues:%s",
			opType, b.tableName, table.ColumnsToString(rowKey), table.ColumnsToString(mutateValues))
	}
	b.batchOps.AppendObTableOperation(op)
	return nil
}

// AddInsertOp add an insert operation to the batch executor.
func (b *obBatchExecutor) AddInsertOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.addDmlOp(protocol.ObTableOperationInsert, rowKey, mutateValues, opts...)
}

// AddUpdateOp add an update operation to the batch executor.
func (b *obBatchExecutor) AddUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.addDmlOp(protocol.ObTableOperationUpdate, rowKey, mutateValues, opts...)
}

// AddInsertOrUpdateOp add an insertOrUpdate operation to the batch executor
func (b *obBatchExecutor) AddInsertOrUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.addDmlOp(protocol.ObTableOperationInsertOrUpdate, rowKey, mutateValues, opts...)
}

// AddReplaceOp add a replace operation to the batch executor
func (b *obBatchExecutor) AddReplaceOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.addDmlOp(protocol.ObTableOperationReplace, rowKey, mutateValues, opts...)
}

// AddIncrementOp add an increment operation to the batch executor
func (b *obBatchExecutor) AddIncrementOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.addDmlOp(protocol.ObTableOperationIncrement, rowKey, mutateValues, opts...)
}

// AddAppendOp add an append operation to the batch executor
func (b *obBatchExecutor) AddAppendOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.addDmlOp(protocol.ObTableOperationAppend, rowKey, mutateValues, opts...)
}

// AddDeleteOp add a delete operation to the batch executor
func (b *obBatchExecutor) AddDeleteOp(rowKey []*table.Column, opts ...ObkvOption) error {
	op, err := protocol.NewObTableOperation(protocol.ObTableOperationDel, rowKey, nil)
	if err != nil {
		return errors.WithMessagef(err, "new delete table operation, tableName:%s, rowKey:%s",
			b.tableName, table.ColumnsToString(rowKey))
	}
	b.batchOps.AppendObTableOperation(op)
	return nil
}

// AddGetOp add a get operation to the batch executor
func (b *obBatchExecutor) AddGetOp(rowKey []*table.Column, getColumns []string, opts ...ObkvOption) error {
	var columns []*table.Column
	for _, columnName := range getColumns {
		columns = append(columns, table.NewColumn(columnName, nil))
	}
	op, err := protocol.NewObTableOperation(protocol.ObTableOperationGet, rowKey, columns)
	if err != nil {
		return errors.WithMessagef(err, "new get table operation, tableName:%s, rowKey:%s",
			b.tableName, table.ColumnsToString(rowKey))
	}
	b.batchOps.AppendObTableOperation(op)
	return nil
}

// constructPartOpMap classify all operations by the dimension of the partition.
func (b *obBatchExecutor) constructPartOpMap(ctx context.Context) (map[int64]*obPartOp, error) {
	partOpMap := make(map[int64]*obPartOp)
	for i, op := range b.batchOps.ObTableOperations() {
		rowKey := op.Entity().RowKey().GetRowKeyValue()
		tableParam, err := b.cli.getTableParam(ctx, b.tableName, rowKey, false)
		if err != nil {
			return nil, errors.WithMessagef(err, "get table param, tableName:%s, rowKey:%s",
				b.tableName, util.InterfacesToString(rowKey))
		}
		singleOp := newSingleOp(i, op)
		partOp, exist := partOpMap[tableParam.partitionId]
		if !exist {
			partOp = newPartOp(tableParam)
			partOpMap[tableParam.partitionId] = partOp
		}
		partOp.addOperation(singleOp)
	}
	return partOpMap, nil
}

// partitionExecute execute operation on a single partition.
func (b *obBatchExecutor) partitionExecute(
	ctx context.Context,
	partOp *obPartOp,
	res []*protocol.ObTableOperationResponse) error {
	// 1. Construct batch operation request
	// 1.1 Construct batch operation
	batchOp := protocol.NewObTableBatchOperation()
	ops := make([]*protocol.ObTableOperation, 0, len(partOp.ops))
	for _, op := range partOp.ops {
		ops = append(ops, op.op)
	}
	batchOp.SetObTableOperations(ops)
	// 1.2 Construct batch operation request
	request := protocol.NewObTableBatchOperationRequest(
		b.tableName,
		partOp.tableParam.tableId,
		partOp.tableParam.partitionId,
		batchOp,
		b.cli.config.OperationTimeOut,
		b.cli.config.LogLevel,
	)

	// 2. Execute
	partRes := protocol.NewObTableBatchOperationResponse()
	err := partOp.tableParam.table.execute(ctx, request, partRes)
	if err != nil {
		return errors.WithMessagef(err, "table execute, request:%s", request.String())
	}

	// 3. Handle result
	subResSize := len(partRes.ObTableOperationResponses())
	subOpSize := len(partOp.ops)
	if subResSize < subOpSize {
		// only one result when it across failed
		// only one result when hkv puts
		if len(partRes.ObTableOperationResponses()) == 1 {
			for _, op := range partOp.ops {
				res[op.indexOfBatch] = partRes.ObTableOperationResponses()[0]
			}
		} else {
			return errors.Errorf("unexpected batch result size, subResSize:%d", subResSize)
		}
	} else {
		if subResSize != subOpSize {
			return errors.Errorf("unexpected batch result size, subResSize:%d, subOpSize:%d", subResSize, subOpSize)
		}
		for i, op := range partOp.ops {
			res[op.indexOfBatch] = partRes.ObTableOperationResponses()[i]
		}
	}

	return nil
}

// Execute a batch operation.
// batch operation only ensures atomicity of a single partition.
// BatchOperationResult contains the results of all operations.
func (b *obBatchExecutor) Execute(ctx context.Context) (BatchOperationResult, error) {
	if b.cli == nil {
		return nil, errors.New("client handle is nil")
	}
	if len(b.batchOps.ObTableOperations()) == 0 {
		return nil, errors.New("operation is empty")
	}

	if _, ok := ctx.Deadline(); !ok {
		ctx, _ = context.WithTimeout(ctx, b.cli.config.OperationTimeOut) // default timeout operation timeout
	}

	res := make([]*protocol.ObTableOperationResponse, len(b.batchOps.ObTableOperations()))
	// 1. construct partition operation map
	partOpMap, err := b.constructPartOpMap(ctx)
	if err != nil {
		return nil, errors.WithMessagef(err, "construct partition operation map")
	}

	// 2. Loop map, execute per partition operations in goroutine
	if len(partOpMap) > 1 {
		errArr := make([]error, 0, 1)
		var errArrLock sync.Mutex
		var wg sync.WaitGroup
		for _, partOp := range partOpMap {
			wg.Add(1)
			go func(ctx context.Context, partOp *obPartOp) {
				defer wg.Done()
				err := b.partitionExecute(ctx, partOp, res)
				if err != nil {
					log.Warn("failed to execute partition operations", log.String("partOp", partOp.String()))
					errArrLock.Lock()
					errArr = append(errArr, err)
					errArrLock.Unlock()
				}
			}(ctx, partOp)
		}
		wg.Wait()
		if len(errArr) != 0 {
			log.Warn("error occur when execute partition operations")
			return nil, errArr[0]
		}
	} else {
		for _, partOp := range partOpMap {
			err := b.partitionExecute(ctx, partOp, res)
			if err != nil {
				return nil, errors.WithMessagef(err, "execute partition operations, partOp:%s", partOp.String())
			}
		}
	}

	return newBatchOperationResult(res), nil
}

type obPartOp struct {
	tableParam *ObTableParam
	ops        []*obSingleOp
}

func newPartOp(tableParam *ObTableParam) *obPartOp {
	ops := make([]*obSingleOp, 0)
	return &obPartOp{tableParam, ops}
}

func (p *obPartOp) addOperation(op *obSingleOp) {
	p.ops = append(p.ops, op)
}

func (p *obPartOp) String() string {
	var opsStr string
	opsStr = opsStr + "["
	for i := 0; i < len(p.ops); i++ {
		if i > 0 {
			opsStr += ", "
		}
		opsStr += p.ops[i].String()
	}
	opsStr += "]"
	return "obPartOp{" +
		"tableParam:" + p.tableParam.String() + ", " +
		"ops:" + opsStr +
		"}"
}

type obSingleOp struct {
	indexOfBatch int
	op           *protocol.ObTableOperation
}

func newSingleOp(index int, op *protocol.ObTableOperation) *obSingleOp {
	return &obSingleOp{index, op}
}

func (s *obSingleOp) String() string {
	return "obSingleOp{" +
		"indexOfBatch:" + strconv.Itoa(s.indexOfBatch) + ", " +
		"op:" + s.op.String() +
		"}"
}

type BatchOperationResult interface {
	GetResults() []*protocol.ObTableOperationResponse
}

type obBatchOperationResult struct {
	results []*protocol.ObTableOperationResponse
}

func newBatchOperationResult(results []*protocol.ObTableOperationResponse) *obBatchOperationResult {
	return &obBatchOperationResult{results}
}

func (r *obBatchOperationResult) GetResults() []*protocol.ObTableOperationResponse {
	return r.results
}

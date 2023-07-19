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
	"github.com/oceanbase/obkv-table-client-go/client/option"
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

// newObBatchExecutor create a batch executor and bind a client.
func newObBatchExecutor(tableName string, cli *obClient) *obBatchExecutor {
	return &obBatchExecutor{
		tableName:           tableName,
		batchOps:            protocol.NewObTableBatchOperation(),
		cli:                 cli,
		rowKeyName:          nil,
		readOnly:            true,
		sameType:            true,
		samePropertiesNames: false,
	}
}

type obBatchExecutor struct {
	tableName           string
	batchOps            *protocol.ObTableBatchOperation
	cli                 *obClient
	rowKeyName          []string
	readOnly            bool
	sameType            bool
	samePropertiesNames bool
}

func (c *obBatchExecutor) getOperationOptions(opts ...option.ObOperationOption) *option.ObOperationOptions {
	operationOptions := option.NewOperationOptions()
	for _, opt := range opts {
		opt.Apply(operationOptions)
	}
	return operationOptions
}

func (c *obBatchExecutor) setBatchOptions(batchOptions *option.ObBatchOptions) {
	c.readOnly = batchOptions.ReadOnly
	c.sameType = batchOptions.SameType
	c.samePropertiesNames = batchOptions.SamePropertiesNames
}

func (c *obBatchExecutor) setReadonly(readOnly bool) {
	c.readOnly = readOnly
}

func (c *obBatchExecutor) setSameType(sameType bool) {
	c.sameType = sameType
}

func (c *obBatchExecutor) setSamePropertiesNames(samePropertiesNames bool) {
	c.samePropertiesNames = samePropertiesNames
}

func (b *obBatchExecutor) String() string {
	var rowKeyNameStr string
	rowKeyNameStr = rowKeyNameStr + "["
	for i := 0; i < len(b.rowKeyName); i++ {
		if i > 0 {
			rowKeyNameStr += ", "
		}
		rowKeyNameStr += b.rowKeyName[i]
	}
	rowKeyNameStr += "]"
	return "obBatchExecutor{" +
		"tableName:" + b.tableName + ", " +
		"rowKeyName:" + rowKeyNameStr +
		"}"
}

// addDmlOp add dml operation witch include insert/update/insertOrUpdate/replace/increment/append
// operation to batch executor
func (b *obBatchExecutor) addDmlOp(
	opType protocol.ObTableOperationType,
	rowKey []*table.Column,
	mutateValues []*table.Column,
	opts ...option.ObOperationOption) error {

	if rowKey == nil {
		return errors.New("rowKey is nil")
	}
	if mutateValues == nil {
		return errors.New("mutateValues is nil")
	}

	// 1. Add rowKey name firstly
	if b.rowKeyName == nil {
		b.rowKeyName = make([]string, 0, len(rowKey))
		for _, column := range rowKey {
			b.rowKeyName = append(b.rowKeyName, column.Name())
		}
	}

	// 2. Create new operation
	op, err := protocol.NewObTableOperationWithParams(opType, rowKey, mutateValues)
	if err != nil {
		return errors.WithMessagef(err, "new table operation, opType:%d, tableName:%s, rowKey:%s, mutateValues:%s",
			opType, b.tableName, table.ColumnsToString(rowKey), table.ColumnsToString(mutateValues))
	}

	// 3. Append operation
	b.batchOps.AppendObTableOperation(op)
	return nil
}

// AddInsertOp add an insert operation to the batch executor.
func (b *obBatchExecutor) AddInsertOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObOperationOption) error {
	return b.addDmlOp(protocol.ObTableOperationInsert, rowKey, mutateValues, opts...)
}

// AddUpdateOp add an update operation to the batch executor.
func (b *obBatchExecutor) AddUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObOperationOption) error {
	return b.addDmlOp(protocol.ObTableOperationUpdate, rowKey, mutateValues, opts...)
}

// AddInsertOrUpdateOp add an insertOrUpdate operation to the batch executor
func (b *obBatchExecutor) AddInsertOrUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObOperationOption) error {
	return b.addDmlOp(protocol.ObTableOperationInsertOrUpdate, rowKey, mutateValues, opts...)
}

// AddReplaceOp add a replace operation to the batch executor
func (b *obBatchExecutor) AddReplaceOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObOperationOption) error {
	return b.addDmlOp(protocol.ObTableOperationReplace, rowKey, mutateValues, opts...)
}

// AddIncrementOp add an increment operation to the batch executor
func (b *obBatchExecutor) AddIncrementOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObOperationOption) error {
	return b.addDmlOp(protocol.ObTableOperationIncrement, rowKey, mutateValues, opts...)
}

// AddAppendOp add an append operation to the batch executor
func (b *obBatchExecutor) AddAppendOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObOperationOption) error {
	return b.addDmlOp(protocol.ObTableOperationAppend, rowKey, mutateValues, opts...)
}

// AddDeleteOp add a delete operation to the batch executor
func (b *obBatchExecutor) AddDeleteOp(rowKey []*table.Column, opts ...option.ObOperationOption) error {
	if rowKey == nil {
		return errors.New("rowKey is nil")
	}

	// 1. Add rowKey name firstly
	if b.rowKeyName == nil {
		b.rowKeyName = make([]string, 0, len(rowKey))
		for _, column := range rowKey {
			b.rowKeyName = append(b.rowKeyName, column.Name())
		}
	}

	// 2. Create new operation
	op, err := protocol.NewObTableOperationWithParams(protocol.ObTableOperationDel, rowKey, nil)
	if err != nil {
		return errors.WithMessagef(err, "new delete table operation, tableName:%s, rowKey:%s",
			b.tableName, table.ColumnsToString(rowKey))
	}

	// 3. Append operation
	b.batchOps.AppendObTableOperation(op)
	return nil
}

// AddGetOp add a get operation to the batch executor
func (b *obBatchExecutor) AddGetOp(rowKey []*table.Column, getColumns []string, opts ...option.ObOperationOption) error {
	if rowKey == nil {
		return errors.New("rowKey is nil")
	}

	// 1. Add rowKey name firstly
	b.rowKeyName = make([]string, 0, len(rowKey))
	for _, column := range rowKey {
		b.rowKeyName = append(b.rowKeyName, column.Name())
	}

	// 2. Create new operation
	var columns []*table.Column
	for _, columnName := range getColumns {
		columns = append(columns, table.NewColumn(columnName, nil))
	}
	op, err := protocol.NewObTableOperationWithParams(protocol.ObTableOperationGet, rowKey, columns)
	if err != nil {
		return errors.WithMessagef(err, "new get table operation, tableName:%s, rowKey:%s",
			b.tableName, table.ColumnsToString(rowKey))
	}

	// 3. Append operation
	b.batchOps.AppendObTableOperation(op)
	return nil
}

// constructPartOpMap classify all operations by the dimension of the partition.
func (b *obBatchExecutor) constructPartOpMap(ctx context.Context) (map[uint64]*obPartOp, error) {
	partOpMap := make(map[uint64]*obPartOp)
	for i, op := range b.batchOps.ObTableOperations() {
		rowKey := make([]*table.Column, 0, len(b.rowKeyName))
		rowKeyValue := op.Entity().GetRowKeyValue()
		for i, v := range rowKeyValue {
			rowKey = append(rowKey, table.NewColumn(b.rowKeyName[i], v))
		}
		tableParam, err := b.cli.getTableParam(ctx, b.tableName, rowKey, false)
		if err != nil {
			return nil, errors.WithMessagef(err, "get table param, tableName:%s, rowKeyValue:%s",
				b.tableName, util.InterfacesToString(rowKeyValue))
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
	res []SingleResult) error {
	// 1. Construct batch operation request
	// 1.1 Construct batch operation
	batchOp := protocol.NewObTableBatchOperation()
	ops := make([]*protocol.ObTableOperation, 0, len(partOp.ops))
	batchOp.SetObTableOperations(ops)
	batchOp.SetReadOnly(b.readOnly)
	batchOp.SetSameType(b.sameType)
	batchOp.SetSamePropertiesNames(b.samePropertiesNames)
	for _, op := range partOp.ops {
		batchOp.AppendObTableOperation(op.op)
	}
	// 1.2 Construct batch operation request
	request := protocol.NewObTableBatchOperationRequestWithParams(
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
				res[op.indexOfBatch] = operationResponse2SingleResult(partRes.ObTableOperationResponses()[0])
			}
		} else {
			return errors.Errorf("unexpected batch result size, subResSize:%d", subResSize)
		}
	} else {
		if subResSize != subOpSize {
			return errors.Errorf("unexpected batch result size, subResSize:%d, subOpSize:%d", subResSize, subOpSize)
		}
		for i, op := range partOp.ops {
			res[op.indexOfBatch] = operationResponse2SingleResult(partRes.ObTableOperationResponses()[i])
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

	res := make([]SingleResult, len(b.batchOps.ObTableOperations()))
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
					log.Warn("failed to execute partition operations", log.String("partOp", partOp.String()), log.String("err", err.Error()))
					errArrLock.Lock()
					errArr = append(errArr, err)
					errArrLock.Unlock()
				}
			}(ctx, partOp)
		}
		wg.Wait()
		if len(errArr) != 0 {
			log.Warn("error occur when execute partition operations")
			return newObBatchOperationResult(res), errArr[0]
		}
	} else {
		for _, partOp := range partOpMap {
			err := b.partitionExecute(ctx, partOp, res)
			if err != nil {
				return newObBatchOperationResult(res), errors.WithMessagef(err, "execute partition operations, partOp:%s", partOp.String())
			}
		}
	}

	return newObBatchOperationResult(res), nil
}

// operationResponse2SingleResult convert operation response to single result.
func operationResponse2SingleResult(res *protocol.ObTableOperationResponse) SingleResult {
	return newObSingleResult(res.AffectedRows(), res.Entity())
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
	var tableParamStr = "nil"
	if p.tableParam != nil {
		tableParamStr = p.tableParam.String()
	}

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
		"tableParam:" + tableParamStr + ", " +
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
	var opStr = "nil"
	if s.op != nil {
		opStr = s.op.String()
	}
	return "obSingleOp{" +
		"indexOfBatch:" + strconv.Itoa(s.indexOfBatch) + ", " +
		"op:" + opStr +
		"}"
}

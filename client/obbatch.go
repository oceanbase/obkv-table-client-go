package client

import (
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObBatchExecutor struct {
	tableName string
	batchOps  *protocol.TableBatchOperation
	cli       *ObClient
}

func newObBatchExecutor(tableName string, cli *ObClient) *ObBatchExecutor {
	return &ObBatchExecutor{
		tableName: tableName,
		batchOps:  protocol.NewTableBatchOperation(),
		cli:       cli,
	}
}

func (b *ObBatchExecutor) AddDmlOp(
	opType protocol.TableOperationType,
	rowKey []*table.Column,
	mutateValues []*table.Column,
	opts ...ObkvOption) error {
	var rowKeyValues []interface{}
	var columnNames []string
	var properties []interface{}
	for _, column := range rowKey {
		rowKeyValues = append(rowKeyValues, column.Value())
	}
	for _, column := range mutateValues {
		columnNames = append(columnNames, column.Name())
		properties = append(properties, column.Value())
	}
	op, err := protocol.NewTableOperation(opType, rowKeyValues, columnNames, properties)
	if err != nil {
		log.Warn("failed to new table operation",
			log.Int("type", int(opType)),
			log.String("tableName", b.tableName),
			log.String("rowKey", columnsToString(rowKey)),
			log.String("mutateValues", columnsToString(mutateValues)))
		return err
	}
	b.batchOps.AppendTableOperation(op)
	return nil
}

func (b *ObBatchExecutor) AddInsertOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.AddDmlOp(protocol.Insert, rowKey, mutateValues, opts...)
}

func (b *ObBatchExecutor) AddUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.AddDmlOp(protocol.Update, rowKey, mutateValues, opts...)
}

func (b *ObBatchExecutor) AddInsertOrUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.AddDmlOp(protocol.InsertOrUpdate, rowKey, mutateValues, opts...)
}

func (b *ObBatchExecutor) AddReplaceOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.AddDmlOp(protocol.Replace, rowKey, mutateValues, opts...)
}

func (b *ObBatchExecutor) AddIncrementOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.AddDmlOp(protocol.Increment, rowKey, mutateValues, opts...)
}

func (b *ObBatchExecutor) AddAppendOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error {
	return b.AddDmlOp(protocol.Increment, rowKey, mutateValues, opts...)
}

func (b *ObBatchExecutor) AddDeleteOp(rowKey []*table.Column, opts ...ObkvOption) error {
	var rowKeyValues []interface{}
	for _, column := range rowKey {
		rowKeyValues = append(rowKeyValues, column.Value())
	}
	op, err := protocol.NewTableOperation(protocol.Del, rowKeyValues, nil, nil)
	if err != nil {
		log.Warn("failed to new table operation",
			log.Int("type", int(protocol.Del)),
			log.String("tableName", b.tableName),
			log.String("rowKey", columnsToString(rowKey)))
		return err
	}
	b.batchOps.AppendTableOperation(op)
	return nil
}

func (b *ObBatchExecutor) AddGetOp(rowKey []*table.Column, getColumns []string, opts ...ObkvOption) error {
	var rowKeyValues []interface{}
	for _, column := range rowKey {
		rowKeyValues = append(rowKeyValues, column.Value())
	}
	op, err := protocol.NewTableOperation(protocol.Get, rowKeyValues, getColumns, nil)
	if err != nil {
		log.Warn("failed to new table operation",
			log.Int("type", int(protocol.Get)),
			log.String("tableName", b.tableName),
			log.String("rowKey", columnsToString(rowKey)),
			log.String("getColumns", util.StringArrayToString(getColumns)))
		return err
	}
	b.batchOps.AppendTableOperation(op)
	return nil
}

func (b *ObBatchExecutor) constructPartOpMap() (map[int64]*ObPartOp, error) {
	partOpMap := make(map[int64]*ObPartOp)
	for i, op := range b.batchOps.TableOperations() {
		rowKey := op.Entity().RowKey().GetRowKeyValue()
		tableParam, err := b.cli.getTableParam(b.tableName, rowKey, false)
		if err != nil {
			log.Warn("failed to get table param",
				log.String("tableName", b.tableName),
				log.String("rowKey", util.InterfacesToString(rowKey)))
			return nil, err
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

func (b *ObBatchExecutor) partitionExecute(
	partOp *ObPartOp,
	res []*protocol.TableOperationResponse) error {
	// 1. Construct batch operation request
	// 1.1 Construct batch operation
	batchOp := protocol.NewTableBatchOperation()
	ops := make([]*protocol.TableOperation, 0, len(partOp.ops))
	for _, op := range partOp.ops {
		ops = append(ops, op.op)
	}
	batchOp.SetTableOperations(ops)
	// 1.2 Construct batch operation request
	request := protocol.NewTableBatchOperationRequest(
		b.tableName,
		partOp.tableParam.tableId,
		partOp.tableParam.partitionId,
		batchOp,
		b.cli.config.OperationTimeOut,
		b.cli.config.LogLevel,
	)

	// 2. Execute
	partRes := protocol.NewTableBatchOperationResponse()
	err := partOp.tableParam.table.execute(request, partRes)
	if err != nil {
		log.Warn("failed to execute batch request", log.String("request", request.String()))
		return errors.WithMessagef(err, "[%s]", request.String())
	}

	// 3. Handle result
	subResSize := len(partRes.TableOperationResponses())
	subOpSize := len(partOp.ops)
	if subResSize < subOpSize {
		// only one result when it across failed
		// only one result when hkv puts
		if len(partRes.TableOperationResponses()) == 1 {
			for _, op := range partOp.ops {
				res[op.indexOfBatch] = partRes.TableOperationResponses()[0]
			}
		} else {
			log.Warn("unexpected batch result size", log.Int("subResSize", subResSize))
			return errors.New("unexpected batch result size")
		}
	} else {
		if subResSize != subOpSize {
			log.Warn("unexpected batch result size",
				log.Int("subResSize", subResSize),
				log.Int("subOpSize", subOpSize))
			return errors.New("unexpected batch result size")
		}
		for i, op := range partOp.ops {
			res[op.indexOfBatch] = partRes.TableOperationResponses()[i]
		}
	}

	return nil
}

func (b *ObBatchExecutor) Execute() (BatchOperationResult, error) {
	if b.cli == nil {
		log.Warn("client handle is nil")
		return nil, errors.New("client handle is nil")
	}
	if len(b.batchOps.TableOperations()) == 0 {
		log.Warn("operation is empty")
		return nil, errors.New("operation is empty")
	}
	res := make([]*protocol.TableOperationResponse, len(b.batchOps.TableOperations()))
	// 1. construct partition operation map
	partOpMap, err := b.constructPartOpMap()
	if err != nil {
		log.Warn("failed to construct partition operation map")
		return nil, err
	}

	// 2. Loop map, execute per partition operations in goroutine
	if len(partOpMap) > 1 {
		errArr := make([]error, 0, 1)
		var errArrLock sync.Mutex
		var wg sync.WaitGroup
		for _, partOp := range partOpMap {
			wg.Add(1)
			go func(partOp *ObPartOp) {
				defer wg.Done()
				err := b.partitionExecute(partOp, res)
				if err != nil {
					log.Warn("failed to execute partition operations", log.String("partOp", partOp.String()))
					errArrLock.Lock()
					errArr = append(errArr, err)
					errArrLock.Unlock()
				}
			}(partOp)
		}
		wg.Wait()
		if len(errArr) != 0 {
			log.Warn("error occur when execute partition operations")
			return nil, errArr[0]
		}
	} else {
		for _, partOp := range partOpMap {
			err := b.partitionExecute(partOp, res)
			if err != nil {
				log.Warn("failed to execute partition operations", log.String("partOp", partOp.String()))
				return nil, err
			}
		}
	}

	return newBatchOperationResult(res), nil
}

type ObPartOp struct {
	tableParam *ObTableParam
	ops        []*ObSingleOp
}

func newPartOp(tableParam *ObTableParam) *ObPartOp {
	ops := make([]*ObSingleOp, 0)
	return &ObPartOp{tableParam, ops}
}

func (p *ObPartOp) addOperation(op *ObSingleOp) {
	p.ops = append(p.ops, op)
}

func (p *ObPartOp) String() string {
	var opsStr string
	opsStr = opsStr + "["
	for i := 0; i < len(p.ops); i++ {
		if i > 0 {
			opsStr += ", "
		}
		opsStr += p.ops[i].String()
	}
	opsStr += "]"
	return "ObPartOp{" +
		"tableParam:" + p.tableParam.String() + ", " +
		"ops:" + opsStr +
		"}"
}

type ObSingleOp struct {
	indexOfBatch int
	op           *protocol.TableOperation
}

func newSingleOp(index int, op *protocol.TableOperation) *ObSingleOp {
	return &ObSingleOp{index, op}
}

func (s *ObSingleOp) String() string {
	return "ObSingleOp{" +
		"indexOfBatch:" + strconv.Itoa(s.indexOfBatch) + ", " +
		"op:" + s.op.String() +
		"}"
}

type BatchOperationResult interface {
	GetResults() []*protocol.TableOperationResponse
}

type ObBatchOperationResult struct {
	results []*protocol.TableOperationResponse
}

func newBatchOperationResult(results []*protocol.TableOperationResponse) *ObBatchOperationResult {
	return &ObBatchOperationResult{results}
}

func (r *ObBatchOperationResult) GetResults() []*protocol.TableOperationResponse {
	return r.results
}

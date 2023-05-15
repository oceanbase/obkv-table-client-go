package client

import (
	"context"

	"github.com/oceanbase/obkv-table-client-go/table"
)

type BatchExecutor interface {
	// AddInsertOp add an insert operation to the batch executor.
	AddInsertOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	// AddUpdateOp add an update operation to the batch executor.
	AddUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	// AddInsertOrUpdateOp add an insertOrUpdate operation to the batch executor
	AddInsertOrUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	// AddReplaceOp add a replace operation to the batch executor
	AddReplaceOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	// AddIncrementOp add an increment operation to the batch executor
	AddIncrementOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	// AddAppendOp add an append operation to the batch executor
	AddAppendOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	// AddDeleteOp add a delete operation to the batch executor
	AddDeleteOp(rowKey []*table.Column, opts ...ObkvOption) error
	// AddGetOp add a get operation to the batch executor
	AddGetOp(rowKey []*table.Column, getColumns []string, opts ...ObkvOption) error
	// Execute a batch operation.
	// batch operation only ensures atomicity of a single partition.
	// BatchOperationResult contains the results of all operations.
	Execute(ctx context.Context) (BatchOperationResult, error)
}

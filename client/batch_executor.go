package client

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type BatchExecutor interface {
	AddInsertOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddInsertOrUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddReplaceOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddIncrementOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddAppendOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddDeleteOp(rowKey []*table.Column, opts ...ObkvOption) error
	AddGetOp(rowKey []*table.Column, getColumns []string, opts ...ObkvOption) error
	Execute(ctx context.Context) (BatchOperationResult, error)
}

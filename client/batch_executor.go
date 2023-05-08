package client

import "github.com/oceanbase/obkv-table-client-go/table"

type BatchExecutor interface {
	AddInsertOp(rowkey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddUpdateOp(rowkey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddInsertOrUpdateOp(rowkey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddReplaceOp(rowkey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddIncrementOp(rowkey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddAppendOp(rowkey []*table.Column, mutateValues []*table.Column, opts ...ObkvOption) error
	AddDeleteOp(rowkey []*table.Column, opts ...ObkvOption) error
	AddGetOp(rowkey []*table.Column, getColumns []string, opts ...ObkvOption) error
	Execute() (BatchOperationResult, error)
}

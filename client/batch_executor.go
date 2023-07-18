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

	"github.com/oceanbase/obkv-table-client-go/table"
)

// BatchExecutor is for batch operation.
type BatchExecutor interface {
	// AddInsertOp add an insert operation to the batch executor.
	AddInsertOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObkvOperationOption) error
	// AddUpdateOp add an update operation to the batch executor.
	AddUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObkvOperationOption) error
	// AddInsertOrUpdateOp add an insertOrUpdate operation to the batch executor
	AddInsertOrUpdateOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObkvOperationOption) error
	// AddReplaceOp add a replace operation to the batch executor
	AddReplaceOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObkvOperationOption) error
	// AddIncrementOp add an increment operation to the batch executor
	AddIncrementOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObkvOperationOption) error
	// AddAppendOp add an append operation to the batch executor
	AddAppendOp(rowKey []*table.Column, mutateValues []*table.Column, opts ...option.ObkvOperationOption) error
	// AddDeleteOp add a delete operation to the batch executor
	AddDeleteOp(rowKey []*table.Column, opts ...option.ObkvOperationOption) error
	// AddGetOp add a get operation to the batch executor.
	AddGetOp(rowKey []*table.Column, getColumns []string, opts ...option.ObkvOperationOption) error
	// Execute a batch operation.
	// batch operation only ensures atomicity of a single partition.
	// BatchOperationResult contains the results of all operations.
	Execute(ctx context.Context) (BatchOperationResult, error)
}

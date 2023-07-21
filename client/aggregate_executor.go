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
)

type AggExecutor interface {
	// Min add a min operation to the agg executor.
	Min(aggColumn string) AggExecutor
	// Max add a max operation to the agg executor.
	Max(aggColumn string) AggExecutor
	// Count add a count operation to the agg executor.
	Count() AggExecutor
	// Sum add a sum operation to the agg executor.
	Sum(aggColumn string) AggExecutor
	// Avg add an avg operation to the agg executor.
	Avg(aggColumn string) AggExecutor
	// Execute an agg operation.
	// AggregateResult contains the results of all operations.
	Execute(ctx context.Context) (AggregateResult, error)
}

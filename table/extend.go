/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
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

package table

// Entend to represent the minimum and maximum values of the key now.
type Entend int64

const (
	// Min -> -3 = -1 - 2 = MaxUint64 - 2
	Min Entend = -3
	// Max -> -2 = -1 - 1 = MaxUint64 - 1
	Max Entend = -2
)

func (e Entend) String() string {
	switch e {
	case Min:
		return "Min"
	case Max:
		return "Max"
	default:
		return "unknown"
	}
}

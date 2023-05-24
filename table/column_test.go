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

package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn_String(t *testing.T) {
	col := &Column{}
	assert.Equal(t, "column{name:, value:<nil>}", col.String())

	col.SetName("c1")
	col.SetValue(1)
	assert.Equal(t, "column{name:c1, value:1}", col.String())
	assert.EqualValues(t, "c1", col.Name())
	assert.EqualValues(t, 1, col.Value())

	col = NewColumn("c1", 1)
	assert.Equal(t, "column{name:c1, value:1}", col.String())
}

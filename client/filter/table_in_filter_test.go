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

package filter

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObTableInFilterString(t *testing.T) {
	testColumnName := "testColumnName"
	value := rand.Intn(100)
	obTableInFilter := In(testColumnName, value)
	newObTableInFilter := NewObTableInFilter(testColumnName, value)
	assert.EqualValues(t, obTableInFilter.ColumnName(), newObTableInFilter.ColumnName())
	assert.EqualValues(t, obTableInFilter.String(), newObTableInFilter.String())
	assert.EqualValues(t, obTableInFilter, newObTableInFilter)
}

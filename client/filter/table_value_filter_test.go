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

func TestObTableValueFilterString(t *testing.T) {
	testColumnName := "testColumnName"
	value := rand.Intn(100)

	obTableValueFilter := CompareVal(LessThan, testColumnName, value)
	newObTableValueFilter := NewObTableValueFilter(LessThan, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(GreaterThan, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(GreaterThan, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(LessOrEqualThan, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(LessOrEqualThan, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(GreaterOrEqualThan, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(GreaterOrEqualThan, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(NotEqual, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(NotEqual, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(Equal, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(Equal, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(IsNull, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(IsNull, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(IsNotNull, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(IsNotNull, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)

	obTableValueFilter = CompareVal(8, testColumnName, value)
	newObTableValueFilter = NewObTableValueFilter(8, testColumnName, value)
	assert.EqualValues(t, obTableValueFilter.ColumnName(), newObTableValueFilter.ColumnName())
	assert.EqualValues(t, obTableValueFilter.String(), newObTableValueFilter.String())
	assert.EqualValues(t, obTableValueFilter, newObTableValueFilter)
}

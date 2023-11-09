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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentMap_Add(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.Add("key1", "value1")
	cmap.Add("key2", "value2")

	assert.Equal(t, 2, cmap.Size(), "Size() failed")
	assert.True(t, cmap.Contains("key1"), "Contains() failed")
}

func TestConcurrentMap_AddIfAbsent(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.AddIfAbsent("key1", "value1")
	cmap.AddIfAbsent("key2", "value2")
	cmap.AddIfAbsent("key2", "new value2")

	assert.Equal(t, 2, cmap.Size(), "Size() failed")
	assert.True(t, cmap.Contains("key2"), "Contains() failed")
}

func TestConcurrentMap_Remove(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.Add("key1", "value1")
	cmap.Add("key2", "value2")

	cmap.Remove("key1")

	assert.Equal(t, 1, cmap.Size(), "Size() failed")
	assert.False(t, cmap.Contains("key1"), "Contains() failed")
}

func TestConcurrentMap_Contains(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.Add("key1", "value1")
	cmap.Add("key2", "value2")

	assert.True(t, cmap.Contains("key1"), "Contains() failed")
	assert.False(t, cmap.Contains("key3"), "Contains() failed")
}

func TestConcurrentMap_Size(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.Add("key1", "value1")
	cmap.Add("key2", "value2")

	assert.Equal(t, 2, cmap.Size(), "Size() failed")
}

func TestConcurrentMap_Update(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.Add("key1", "value1")
	cmap.Add("key2", "value2")

	assert.EqualValues(t, true, cmap.Update("key1", "new value1"))
	assert.EqualValues(t, false, cmap.Update("key3", "new value1"))

}

func TestConcurrentMap_Get(t *testing.T) {
	cmap := NewConcurrentMap()

	cmap.Add("key1", "value1")
	cmap.Add("key2", "value2")

	assert.EqualValues(t, "value1", cmap.Get("key1"))
	assert.EqualValues(t, nil, cmap.Get("key3"))
}

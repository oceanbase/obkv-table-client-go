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
	"sync"
)

type ConcurrentMap struct {
	items map[interface{}]interface{}
	mu    sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[interface{}]interface{}),
	}
}

func (c *ConcurrentMap) Add(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}

func (c *ConcurrentMap) AddIfAbsent(key interface{}, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[key]; !found {
		c.items[key] = value
	}
}

func (c *ConcurrentMap) Update(key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, found := c.items[key]; found {
		c.items[key] = value
		return true
	}

	return false
}

func (c *ConcurrentMap) Remove(key interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *ConcurrentMap) Get(key interface{}) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.items[key]
}

func (c *ConcurrentMap) Contains(key interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.items[key]
	return found
}

func (c *ConcurrentMap) Range(fn func(key, value interface{})) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.items {
		fn(key, value)
	}
}

func (c *ConcurrentMap) Size() int {
	return len(c.items)
}

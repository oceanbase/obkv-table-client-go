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

import "github.com/oceanbase/obkv-table-client-go/util"

// Column represents a column in a table,
// consisting of a column name and a column value.
type Column struct {
	name  string
	value interface{}
}

func NewColumn(name string, value interface{}) *Column {
	return &Column{name: name, value: value}
}

func (c *Column) Name() string {
	return c.name
}

func (c *Column) SetName(name string) {
	c.name = name
}

func (c *Column) Value() interface{} {
	return c.value
}

func (c *Column) SetValue(value interface{}) {
	c.value = value
}

func (c *Column) String() string {
	return "column{" +
		"name:" + c.name + ", " +
		"value:" + util.InterfaceToString(c.value) +
		"}"
}

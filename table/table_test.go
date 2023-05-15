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

func TestColumn_ToString(t *testing.T) {
	col := Column{}
	assert.Equal(t, col.String(), "column{name: , value: <nil>}")
	col = Column{"c1", 123}
	assert.Equal(t, col.String(), "column{name: c1, value: 123}")
}

func TestObRowKeyElement_ToString(t *testing.T) {
	v := ObRowKeyElement{}
	assert.Equal(t, v.String(), "ObRowKeyElement{nameIdxMap:{}}")
	m := make(map[string]int, 3)
	m["c1"] = 0
	m["c2"] = 1
	m["c3"] = 2
	v = ObRowKeyElement{m}
	assert.Equal(t, v.String(), "ObRowKeyElement{nameIdxMap:{m[c1]=0, m[c2]=1, m[c3]=2}}")
}

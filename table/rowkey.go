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
	"strconv"
)

// ObRowKeyElement store each primary key column name
// and its location index within the primary key.
type ObRowKeyElement struct {
	nameIdxMap map[string]int
}

func NewObRowKeyElement(nameIdxMap map[string]int) *ObRowKeyElement {
	return &ObRowKeyElement{nameIdxMap}
}

func (e *ObRowKeyElement) NameIdxMap() map[string]int {
	return e.nameIdxMap
}

func (e *ObRowKeyElement) String() string {
	var nameIdxMapStr string
	var i = 0
	nameIdxMapStr = nameIdxMapStr + "{"
	for k, v := range e.nameIdxMap {
		if i > 0 {
			nameIdxMapStr += ", "
		}
		i++
		nameIdxMapStr += "m[" + k + "]=" + strconv.Itoa(v)
	}
	nameIdxMapStr += "}"
	return "ObRowKeyElement{" +
		"nameIdxMap:" + nameIdxMapStr +
		"}"
}

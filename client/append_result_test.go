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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/protocol"
)

func TestAppendResult(t *testing.T) {
	res := newObAppendResult(1, nil)
	assert.EqualValues(t, 1, res.AffectedRows())
	assert.Equal(t, nil, res.Value("c1"))

	obj := protocol.NewObObject()
	obj.SetValue(1)
	m := make(map[string]*protocol.ObObject, 1)
	m["c1"] = obj
	entity := protocol.NewObTableEntity()
	entity.SetProperties(m)
	res.affectedEntity = entity
	assert.EqualValues(t, 1, res.Value("c1"))
	assert.EqualValues(t, 1, res.Value("C1"))
	assert.EqualValues(t, nil, res.Value("C2"))

	assert.EqualValues(t, []interface{}(nil), res.RowKey())
	keys := []*protocol.ObObject{obj}
	rowKey := protocol.NewRowKey()
	rowKey.SetKeys(keys)
	entity.SetRowKey(rowKey)
	r := res.RowKey()
	assert.EqualValues(t, 1, r[0])
}

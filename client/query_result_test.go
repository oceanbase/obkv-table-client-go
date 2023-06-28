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

func TestQueryResult(t *testing.T) {
	columnNames := []string{"c1", "c2"}
	res := newObQueryResult(columnNames, nil)
	assert.Equal(t, nil, res.Value("c1"))

	obj1 := protocol.NewObObject()
	obj1.SetValue(1)
	obj2 := protocol.NewObObject()
	obj2.SetValue("str")
	v := []*protocol.ObObject{obj1, obj2}
	res = newObQueryResult(columnNames, v)
	assert.EqualValues(t, 1, res.Value("c1"))
	assert.EqualValues(t, 1, res.Value("C1"))
	assert.EqualValues(t, "str", res.Value("C2"))

}

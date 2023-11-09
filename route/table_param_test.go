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

package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testTableIdV3 uint64 = 1099511677791
	testTableIdV4 uint64 = 500039
	testPartIdV3  uint64 = 0
	testPartIdV4  uint64 = 500040
)

func TestObTableParam_String(t *testing.T) {
	tableParam := NewObTableParam(&ObTable{}, testTableIdV3, testPartIdV3)
	assert.Equal(
		t,
		"ObTableParam{table:ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}, tableId:1099511677791, partitionId:0}",
		tableParam.String(),
	)
	tableParam = NewObTableParam(&ObTable{}, testTableIdV4, testPartIdV4)
	assert.Equal(
		t,
		"ObTableParam{table:ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}, tableId:500039, partitionId:500040}",
		tableParam.String(),
	)
	tb := NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	tableParam = NewObTableParam(tb, testTableIdV4, testPartIdV4)
	assert.Equal(
		t,
		"ObTableParam{table:ObTable{ip:127.0.0.1, port:8080, tenantName:sys, userName:root, password:, database:test, isClosed:false}, tableId:500039, partitionId:500040}",
		tableParam.String(),
	)
}

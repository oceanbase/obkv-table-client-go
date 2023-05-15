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
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testIp         string = "127.0.0.1"
	testPort       int    = 8080
	testTenantName string = "sys"
	testUserName   string = "root"
	testPassword   string = ""
	testDatabase   string = "test"
)

func TestObTable_init(t *testing.T) {
	tb := NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	err := tb.init(1, time.Duration(1000)*time.Millisecond)
	assert.NotEqual(t, nil, err)
}

func TestObTable_close(t *testing.T) {
	tb := NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	tb.close()
	assert.EqualValues(t, true, tb.isClosed)
}

func TestObTable_String(t *testing.T) {
	tb := &ObTable{}
	assert.Equal(
		t,
		"ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}",
		tb.String(),
	)
	tb = NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	assert.Equal(
		t,
		"ObTable{ip:127.0.0.1, port:8080, tenantName:sys, userName:root, password:, database:test, isClosed:false}",
		tb.String(),
	)
}

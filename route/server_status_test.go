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

func TestObServerStatus_IsActive(t *testing.T) {
	status := &obServerStatus{}
	assert.False(t, status.IsActive())
	status = newServerStatus(0, "Active")
	assert.True(t, status.IsActive())
	status = newServerStatus(0, "active")
	assert.True(t, status.IsActive())
	status = newServerStatus(1, "active")
	assert.False(t, status.IsActive())
	status = newServerStatus(0, "InActive")
	assert.False(t, status.IsActive())
	status = newServerStatus(1, "InActive")
	assert.False(t, status.IsActive())
}

func TestObServerStatus_String(t *testing.T) {
	status := &obServerStatus{}
	assert.Equal(t, "obServerStatus{stopTime:0, status:}", status.String())
	status = newServerStatus(0, "Active")
	assert.Equal(t, "obServerStatus{stopTime:0, status:Active}", status.String())
}

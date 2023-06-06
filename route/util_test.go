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

func TestUtil_createInStatement(t *testing.T) {
	inStr := createInStatement(nil)
	assert.Equal(t, "();", inStr)
	inStr = createInStatement([]uint64{1})
	assert.Equal(t, "(1);", inStr)
	inStr = createInStatement([]uint64{1, 2})
	assert.Equal(t, "(1, 2);", inStr)
}

func TestUtil_murmurHash64A(t *testing.T) {
	result := murmurHash64A([]byte{1}, len([]byte{1}), int64(0))
	assert.Equal(t, int64(-5720937396023583481), result)
	result = murmurHash64A([]byte{1}, len([]byte{1}), int64(1))
	assert.Equal(t, int64(6351753276682545529), result)
	result = murmurHash64A([]byte{1, 2, 3}, len([]byte{1, 2, 3}), int64(123456789))
	assert.Equal(t, int64(-4356950700900923028), result)
}

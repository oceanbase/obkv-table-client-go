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

func Test_hashSortUtf8Mb4(t *testing.T) {
	data := []byte{1}
	hashCode := 0
	res := hashSortUtf8Mb4(data, int64(hashCode), 10, true)
	assert.Equal(t, int64(-7030129012826305577), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 10, false)
	assert.Equal(t, int64(2570), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, true)
	assert.Equal(t, int64(-7030129012826305577), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, false)
	assert.Equal(t, int64(7062546676564130965), res)
	data = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res = hashSortUtf8Mb4(data, int64(hashCode), 10, true)
	assert.Equal(t, int64(-1916273894318764036), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 10, false)
	assert.Equal(t, int64(1013208367030014238), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, true)
	assert.Equal(t, int64(-1916273894318764036), res)
	res = hashSortUtf8Mb4(data, int64(hashCode), 0xc6a4a7935bd1e995, false)
	assert.Equal(t, int64(7452881443849355883), res)
}

func Test_hashSortMbBin(t *testing.T) {
	data := []byte{1}
	hashCode := 0
	res := hashSortMbBin(data, int64(hashCode), 10)
	assert.Equal(t, int64(10), res)
}

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

package util

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinary(t *testing.T) {
	var (
		u8  uint8  = 1<<8 - 1
		u16 uint16 = 1<<16 - 1
		u32 uint32 = 1<<32 - 1
		u64 uint64 = 1<<64 - 1
	)
	buf := make([]byte, 15)
	buffer := bytes.NewBuffer(buf)
	PutUint8(buffer, u8)
	PutUint16(buffer, u16)
	PutUint32(buffer, u32)
	PutUint64(buffer, u64)

	buffer = bytes.NewBuffer(buf)
	assert.EqualValues(t, Uint8(buffer), u8)
	assert.EqualValues(t, Uint16(buffer), u16)
	assert.EqualValues(t, Uint32(buffer), u32)
	assert.EqualValues(t, Uint64(buffer), u64)

	u8 = uint8(rand.Uint64())
	u16 = uint16(rand.Uint64())
	u32 = uint32(rand.Uint64())
	u64 = rand.Uint64()
	buf = make([]byte, 15)
	buffer = bytes.NewBuffer(buf)
	PutUint8(buffer, u8)
	PutUint16(buffer, u16)
	PutUint32(buffer, u32)
	PutUint64(buffer, u64)

	buffer = bytes.NewBuffer(buf)
	assert.EqualValues(t, Uint8(buffer), u8)
	assert.EqualValues(t, Uint16(buffer), u16)
	assert.EqualValues(t, Uint32(buffer), u32)
	assert.EqualValues(t, Uint64(buffer), u64)
}

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

package protocol

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObUniVersionHeaderEncodeDecode(t *testing.T) {
	obUniVersionHeader := ObUniVersionHeader{}
	obUniVersionHeader.SetVersion(int64(rand.Uint64()))
	obUniVersionHeader.SetContentLength(int(rand.Uint64()))

	uniVersionHeaderLen := obUniVersionHeader.UniVersionHeaderLen()
	buf := make([]byte, uniVersionHeaderLen)
	buffer := bytes.NewBuffer(buf)
	obUniVersionHeader.Encode(buffer)

	newObUniVersionHeader := ObUniVersionHeader{}
	buffer = bytes.NewBuffer(buf)
	newObUniVersionHeader.Decode(buffer)

	assert.EqualValues(t, obUniVersionHeader.Version(), newObUniVersionHeader.Version())
	assert.EqualValues(t, obUniVersionHeader.ContentLength(), newObUniVersionHeader.ContentLength())
	assert.EqualValues(t, obUniVersionHeader.UniVersionHeaderLen(), newObUniVersionHeader.UniVersionHeaderLen())
	assert.EqualValues(t, obUniVersionHeader.String(), newObUniVersionHeader.String())
	assert.EqualValues(t, obUniVersionHeader, newObUniVersionHeader)
}

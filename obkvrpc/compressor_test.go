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

package obkvrpc

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/golang/snappy"
	"github.com/klauspost/compress/zstd"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/pierrec/lz4/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestNoneDecompress(t *testing.T) {
	s := "hello world"
	decompressor, err := getDecompressor(protocol.ObCompressTypeNone)
	assert.Equal(t, nil, err)
	dstBuffer, err := decompressor.Decompress(bytes.NewBufferString(s), int32(len(s)))
	assert.Equal(t, nil, err)
	assert.EqualValues(t, s, dstBuffer.String())
}

func TestLz4Decompress(t *testing.T) {
	s := "helo world"
	data := []byte(s)
	// compress data
	compressBuf := make([]byte, lz4.CompressBlockBound(len(data)))
	var c lz4.Compressor
	n, err := c.CompressBlock(data, compressBuf)
	assert.Equal(t, nil, err)
	compressBuf = compressBuf[:n]
	fmt.Println(n, len(data))
	// decompress
	decompressor, err := getDecompressor(protocol.ObCompressTypeLZ4)
	assert.Equal(t, nil, err)
	dstBuffer, err := decompressor.Decompress(bytes.NewBuffer(compressBuf), int32(len(s)))
	assert.Equal(t, nil, err)
	assert.EqualValues(t, s, dstBuffer.String())
}

func TestSnappyDecompress(t *testing.T) {
	s := "hello world"
	// compress data
	encodeBytes := snappy.Encode(nil, []byte(s))
	assert.NotEqual(t, 0, len(encodeBytes))

	// decompress
	decompressor, err := getDecompressor(protocol.ObCompressTypeSnappy)
	assert.Equal(t, nil, err)
	dstBuffer, err := decompressor.Decompress(bytes.NewBuffer(encodeBytes), int32(len(s)))
	assert.Equal(t, nil, err)
	assert.EqualValues(t, s, dstBuffer.String())
}

func TestZlibDecompress(t *testing.T) {
	s := "hello world"
	r := strings.NewReader(s)
	// compress data
	srcBuffer := new(bytes.Buffer)
	zw := zlib.NewWriter(srcBuffer)
	_, err := io.Copy(zw, r)
	assert.Equal(t, nil, err)
	err = zw.Close()

	// decompress
	decompressor, err := getDecompressor(protocol.ObCompressTypeZlib)
	assert.Equal(t, nil, err)
	dstBuffer, err := decompressor.Decompress(srcBuffer, int32(len(s)))
	assert.Equal(t, nil, err)
	assert.EqualValues(t, s, dstBuffer.String())
}

func TestZStdDecompress(t *testing.T) {
	s := "hello world"
	r := strings.NewReader(s)
	// compress data
	srcBuffer := new(bytes.Buffer)
	zw, err := zstd.NewWriter(srcBuffer)
	assert.Equal(t, nil, err)
	_, err = io.Copy(zw, r)
	assert.Equal(t, nil, err)
	err = zw.Close()

	// decompress data
	decompressor, err := getDecompressor(protocol.ObCompressTypeZstd)
	assert.Equal(t, nil, err)
	dstBuffer, err := decompressor.Decompress(srcBuffer, int32(len(s)))
	assert.Equal(t, nil, err)
	assert.EqualValues(t, s, dstBuffer.String())
}

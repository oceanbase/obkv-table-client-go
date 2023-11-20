/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
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
	"github.com/pierrec/lz4"
	"github.com/pkg/errors"
	"io"
)

type DeCompressor interface {
	Decompress(src *bytes.Buffer, originLen int32) (*bytes.Buffer, error)
}

type NoneDecompressor struct{}
type Lz4Decompressor struct{}
type SnappyDecompressor struct{}
type ZlibDecompressor struct{}
type ZStdDecompressor struct{}

func (d *NoneDecompressor) Decompress(src *bytes.Buffer, originLen int32) (*bytes.Buffer, error) {
	return src, nil
}

func (d *Lz4Decompressor) Decompress(src *bytes.Buffer, originLen int32) (*bytes.Buffer, error) {
	srcBytes := src.Bytes()
	dstBytes := make([]byte, originLen)
	n, err := lz4.UncompressBlock(srcBytes, dstBytes)
	if err != nil {
		return nil, err
	}
	if int32(n) != originLen {
		return nil, errors.New(fmt.Sprintf("length of decompress content is not expected: actual-%d, expected-%d", n, originLen))
	}
	return bytes.NewBuffer(dstBytes), nil
}

func (d *SnappyDecompressor) Decompress(src *bytes.Buffer, originLen int32) (*bytes.Buffer, error) {
	srcBytes := src.Bytes()
	dstBytes, err := snappy.Decode(nil, srcBytes)
	if err != nil {
		return nil, err
	}
	if int32(len(dstBytes)) != originLen {
		return nil, errors.New(fmt.Sprintf("length of decompress content is not expected: actual-%d, expected-%d", int32(len(dstBytes)), originLen))
	}
	return bytes.NewBuffer(dstBytes), nil
}

func (d *ZlibDecompressor) Decompress(src *bytes.Buffer, originLen int32) (*bytes.Buffer, error) {
	// 1. create zlib reader
	zlibReader, err := zlib.NewReader(src)
	if err != nil {
		return nil, err
	}

	defer func(zlibReader io.ReadCloser) {
		err := zlibReader.Close()
		if err != nil {
			panic(err.Error())
		}
	}(zlibReader)

	// 2. create buffer
	dstBuffer := new(bytes.Buffer)
	// 3. write the decompress data to dstBuffer
	n, err := io.Copy(dstBuffer, zlibReader)
	if err != nil {
		return nil, err
	}
	if int32(n) != originLen {
		return nil, errors.New(fmt.Sprintf("length of decompress content is not expected: actual-%d, expected-%d", n, originLen))
	}
	return dstBuffer, nil
}

func (d *ZStdDecompressor) Decompress(src *bytes.Buffer, originLen int32) (*bytes.Buffer, error) {
	zstdReader, err := zstd.NewReader(src)
	if err != nil {
		return nil, err
	}
	defer zstdReader.Close()
	dstBuffer := new(bytes.Buffer)
	n, err := io.Copy(dstBuffer, zstdReader)
	if err != nil {
		return nil, err
	}
	if int32(n) != originLen {
		return nil, errors.New(fmt.Sprintf("length of decompress content is not expected: actual-%d, expected-%d", n, originLen))
	}
	return dstBuffer, nil
}

func getDecompressor(compressType protocol.ObCompressType) (DeCompressor, error) {
	switch compressType {
	case protocol.ObCompressTypeLZ4:
		return &Lz4Decompressor{}, nil
	case protocol.ObCompressTypeSnappy:
		return &SnappyDecompressor{}, nil
	case protocol.ObCompressTypeZlib:
		return &ZlibDecompressor{}, nil
	case protocol.ObCompressTypeZstd:
		return &ZStdDecompressor{}, nil
	case protocol.ObCompressTypeZstd_138:
		return &ZStdDecompressor{}, nil
	case protocol.ObCompressTypeLZ4_191:
		return &Lz4Decompressor{}, nil
	case protocol.ObCompressTypeNone:
		return &NoneDecompressor{}, nil
	default:
		return nil, errors.New("unknown compression algorithm")
	}
}

func convertCompressTypeToString(compressType protocol.ObCompressType) string {
	switch compressType {
	case protocol.ObCompressTypeLZ4:
		return "lz4"
	case protocol.ObCompressTypeSnappy:
		return "snappy"
	case protocol.ObCompressTypeZlib:
		return "zlib"
	case protocol.ObCompressTypeZstd:
		return "zStd"
	case protocol.ObCompressTypeNone:
		return "none"
	case protocol.ObCompressTypeZstd_138:
		return "zstd_1.3.8"
	case protocol.ObCompressTypeLZ4_191:
		return "lz4_1.9.1"
	default:
		return "invalid compress type"
	}
}

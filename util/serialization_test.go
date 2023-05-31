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
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodedLengthByVi32(t *testing.T) {
	type args struct {
		value int32
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{value: 0}, wantEncodedLength: 1},
		{name: "test", args: args{value: -1}, wantEncodedLength: 5},
		{name: "test", args: args{value: 1<<31 - 1}, wantEncodedLength: 5},
		{name: "test", args: args{value: -(1 << 31)}, wantEncodedLength: 5},
		{name: "test", args: args{value: int32(obMaxV1b)}, wantEncodedLength: 1},
		{name: "test", args: args{value: int32(obMaxV2b)}, wantEncodedLength: 2},
		{name: "test", args: args{value: int32(obMaxV3b)}, wantEncodedLength: 3},
		{name: "test", args: args{value: int32(obMaxV4b)}, wantEncodedLength: 4},
		{name: "test", args: args{value: int32(obMaxV4b + 1)}, wantEncodedLength: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByVi32(tt.args.value), "EncodedLengthByVi32(%v)", tt.args.value)
		})
	}
}

func TestEncodedLengthByVi64(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{value: 0}, wantEncodedLength: 1},
		{name: "test", args: args{value: -1}, wantEncodedLength: 10},
		{name: "test", args: args{value: -(1 << 63)}, wantEncodedLength: 10},
		{name: "test", args: args{value: int64(obMaxV1b)}, wantEncodedLength: 1},
		{name: "test", args: args{value: int64(obMaxV2b)}, wantEncodedLength: 2},
		{name: "test", args: args{value: int64(obMaxV3b)}, wantEncodedLength: 3},
		{name: "test", args: args{value: int64(obMaxV4b)}, wantEncodedLength: 4},
		{name: "test", args: args{value: int64(obMaxV5b)}, wantEncodedLength: 5},
		{name: "test", args: args{value: int64(obMaxV6b)}, wantEncodedLength: 6},
		{name: "test", args: args{value: int64(obMaxV7b)}, wantEncodedLength: 7},
		{name: "test", args: args{value: int64(obMaxV8b)}, wantEncodedLength: 8},
		{name: "test", args: args{value: int64(obMaxV9b)}, wantEncodedLength: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByVi64(tt.args.value), "EncodedLengthByVi64(%v)", tt.args.value)
		})
	}
}

func TestEncodedLengthByVf32(t *testing.T) {
	type args struct {
		value float32
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{value: math.Float32frombits(0)}, wantEncodedLength: 1},
		{name: "test", args: args{value: math.Float32frombits(1<<31 - 1)}, wantEncodedLength: 5},
		{name: "test", args: args{value: math.Float32frombits(1<<32 - 1)}, wantEncodedLength: 5},
		{name: "test", args: args{value: math.Float32frombits(uint32(obMaxV1b))}, wantEncodedLength: 1},
		{name: "test", args: args{value: math.Float32frombits(uint32(obMaxV2b))}, wantEncodedLength: 2},
		{name: "test", args: args{value: math.Float32frombits(uint32(obMaxV3b))}, wantEncodedLength: 3},
		{name: "test", args: args{value: math.Float32frombits(uint32(obMaxV4b))}, wantEncodedLength: 4},
		{name: "test", args: args{value: math.Float32frombits(uint32(obMaxV4b + 1))}, wantEncodedLength: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByVf32(tt.args.value), "EncodedLengthByVf32(%v)", tt.args.value)
		})
	}
}

func TestEncodedLengthByVf64(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{value: math.Float64frombits(0)}, wantEncodedLength: 1},
		{name: "test", args: args{value: math.Float64frombits(obMaxV1b)}, wantEncodedLength: 1},
		{name: "test", args: args{value: math.Float64frombits(obMaxV2b)}, wantEncodedLength: 2},
		{name: "test", args: args{value: math.Float64frombits(obMaxV3b)}, wantEncodedLength: 3},
		{name: "test", args: args{value: math.Float64frombits(obMaxV4b)}, wantEncodedLength: 4},
		{name: "test", args: args{value: math.Float64frombits(obMaxV5b)}, wantEncodedLength: 5},
		{name: "test", args: args{value: math.Float64frombits(obMaxV6b)}, wantEncodedLength: 6},
		{name: "test", args: args{value: math.Float64frombits(obMaxV7b)}, wantEncodedLength: 7},
		{name: "test", args: args{value: math.Float64frombits(obMaxV8b)}, wantEncodedLength: 8},
		{name: "test", args: args{value: math.Float64frombits(obMaxV9b)}, wantEncodedLength: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByVf64(tt.args.value), "EncodedLengthByVf64(%v)", tt.args.value)
		})
	}
}

func TestEncodedLengthByVString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{str: ""}, wantEncodedLength: 2},
		{name: "test", args: args{str: "this is test."}, wantEncodedLength: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByVString(tt.args.str), "EncodedLengthByVString(%v)", tt.args.str)
		})
	}
}

func TestEncodedLengthByBytesString(t *testing.T) {
	type args struct {
		bys []byte
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{bys: []byte("")}, wantEncodedLength: 2},
		{name: "test", args: args{bys: []byte("this is test.")}, wantEncodedLength: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByBytesString(tt.args.bys), "EncodedLengthByBytesString(%v)", tt.args.bys)
		})
	}
}

func TestEncodedLengthByBytes(t *testing.T) {
	type args struct {
		bys []byte
	}
	tests := []struct {
		name              string
		args              args
		wantEncodedLength int
	}{
		{name: "test", args: args{bys: []byte("")}, wantEncodedLength: 1},
		{name: "test", args: args{bys: []byte("this is test.")}, wantEncodedLength: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantEncodedLength, EncodedLengthByBytes(tt.args.bys), "EncodedLengthByBytes(%v)", tt.args.bys)
		})
	}
}

func TestEncodeDecodeVi32(t *testing.T) {
	var (
		u1       = int32(rand.Uint32())
		u2       = int32(rand.Uint32())
		u3       = int32(rand.Uint32())
		u4       = int32(rand.Uint32())
		u5       = int32(rand.Uint32())
		u6       = int32(rand.Uint32())
		u7       = int32(rand.Uint32())
		u8       = int32(rand.Uint32())
		totalLen = 0
	)
	totalLen += EncodedLengthByVi32(u1)
	totalLen += EncodedLengthByVi32(u2)
	totalLen += EncodedLengthByVi32(u3)
	totalLen += EncodedLengthByVi32(u4)
	totalLen += EncodedLengthByVi32(u5)
	totalLen += EncodedLengthByVi32(u6)
	totalLen += EncodedLengthByVi32(u7)
	totalLen += EncodedLengthByVi32(u8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeVi32(buffer, u1)
	EncodeVi32(buffer, u2)
	EncodeVi32(buffer, u3)
	EncodeVi32(buffer, u4)
	EncodeVi32(buffer, u5)
	EncodeVi32(buffer, u6)
	EncodeVi32(buffer, u7)
	EncodeVi32(buffer, u8)

	buffer = bytes.NewBuffer(buf)
	var (
		i1 = DecodeVi32(buffer)
		i2 = DecodeVi32(buffer)
		i3 = DecodeVi32(buffer)
		i4 = DecodeVi32(buffer)
		i5 = DecodeVi32(buffer)
		i6 = DecodeVi32(buffer)
		i7 = DecodeVi32(buffer)
		i8 = DecodeVi32(buffer)
	)
	assert.EqualValues(t, u1, i1)
	assert.EqualValues(t, u2, i2)
	assert.EqualValues(t, u3, i3)
	assert.EqualValues(t, u4, i4)
	assert.EqualValues(t, u5, i5)
	assert.EqualValues(t, u6, i6)
	assert.EqualValues(t, u7, i7)
	assert.EqualValues(t, u8, i8)
}

func TestEncodeDecodeVi64(t *testing.T) {
	var (
		u1       = int64(rand.Uint64())
		u2       = int64(rand.Uint64())
		u3       = int64(rand.Uint64())
		u4       = int64(rand.Uint64())
		u5       = int64(rand.Uint64())
		u6       = int64(rand.Uint64())
		u7       = int64(rand.Uint64())
		u8       = int64(rand.Uint64())
		totalLen = 0
	)
	totalLen += EncodedLengthByVi64(u1)
	totalLen += EncodedLengthByVi64(u2)
	totalLen += EncodedLengthByVi64(u3)
	totalLen += EncodedLengthByVi64(u4)
	totalLen += EncodedLengthByVi64(u5)
	totalLen += EncodedLengthByVi64(u6)
	totalLen += EncodedLengthByVi64(u7)
	totalLen += EncodedLengthByVi64(u8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeVi64(buffer, u1)
	EncodeVi64(buffer, u2)
	EncodeVi64(buffer, u3)
	EncodeVi64(buffer, u4)
	EncodeVi64(buffer, u5)
	EncodeVi64(buffer, u6)
	EncodeVi64(buffer, u7)
	EncodeVi64(buffer, u8)

	buffer = bytes.NewBuffer(buf)
	var (
		i1 = DecodeVi64(buffer)
		i2 = DecodeVi64(buffer)
		i3 = DecodeVi64(buffer)
		i4 = DecodeVi64(buffer)
		i5 = DecodeVi64(buffer)
		i6 = DecodeVi64(buffer)
		i7 = DecodeVi64(buffer)
		i8 = DecodeVi64(buffer)
	)
	assert.EqualValues(t, u1, i1)
	assert.EqualValues(t, u2, i2)
	assert.EqualValues(t, u3, i3)
	assert.EqualValues(t, u4, i4)
	assert.EqualValues(t, u5, i5)
	assert.EqualValues(t, u6, i6)
	assert.EqualValues(t, u7, i7)
	assert.EqualValues(t, u8, i8)
}

func TestEncodeDecodeVf32(t *testing.T) {
	var (
		u1       = rand.Float32()
		u2       = rand.Float32()
		u3       = rand.Float32()
		u4       = rand.Float32()
		u5       = rand.Float32()
		u6       = rand.Float32()
		u7       = rand.Float32()
		u8       = rand.Float32()
		totalLen = 0
	)
	totalLen += EncodedLengthByVf32(u1)
	totalLen += EncodedLengthByVf32(u2)
	totalLen += EncodedLengthByVf32(u3)
	totalLen += EncodedLengthByVf32(u4)
	totalLen += EncodedLengthByVf32(u5)
	totalLen += EncodedLengthByVf32(u6)
	totalLen += EncodedLengthByVf32(u7)
	totalLen += EncodedLengthByVf32(u8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeVf32(buffer, u1)
	EncodeVf32(buffer, u2)
	EncodeVf32(buffer, u3)
	EncodeVf32(buffer, u4)
	EncodeVf32(buffer, u5)
	EncodeVf32(buffer, u6)
	EncodeVf32(buffer, u7)
	EncodeVf32(buffer, u8)

	buffer = bytes.NewBuffer(buf)
	var (
		i1 = DecodeVf32(buffer)
		i2 = DecodeVf32(buffer)
		i3 = DecodeVf32(buffer)
		i4 = DecodeVf32(buffer)
		i5 = DecodeVf32(buffer)
		i6 = DecodeVf32(buffer)
		i7 = DecodeVf32(buffer)
		i8 = DecodeVf32(buffer)
	)
	assert.EqualValues(t, u1, i1)
	assert.EqualValues(t, u2, i2)
	assert.EqualValues(t, u3, i3)
	assert.EqualValues(t, u4, i4)
	assert.EqualValues(t, u5, i5)
	assert.EqualValues(t, u6, i6)
	assert.EqualValues(t, u7, i7)
	assert.EqualValues(t, u8, i8)
}

func TestEncodeDecodeVf64(t *testing.T) {
	var (
		u1       = rand.Float64()
		u2       = rand.Float64()
		u3       = rand.Float64()
		u4       = rand.Float64()
		u5       = rand.Float64()
		u6       = rand.Float64()
		u7       = rand.Float64()
		u8       = rand.Float64()
		totalLen = 0
	)
	totalLen += EncodedLengthByVf64(u1)
	totalLen += EncodedLengthByVf64(u2)
	totalLen += EncodedLengthByVf64(u3)
	totalLen += EncodedLengthByVf64(u4)
	totalLen += EncodedLengthByVf64(u5)
	totalLen += EncodedLengthByVf64(u6)
	totalLen += EncodedLengthByVf64(u7)
	totalLen += EncodedLengthByVf64(u8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeVf64(buffer, u1)
	EncodeVf64(buffer, u2)
	EncodeVf64(buffer, u3)
	EncodeVf64(buffer, u4)
	EncodeVf64(buffer, u5)
	EncodeVf64(buffer, u6)
	EncodeVf64(buffer, u7)
	EncodeVf64(buffer, u8)

	buffer = bytes.NewBuffer(buf)
	var (
		i1 = DecodeVf64(buffer)
		i2 = DecodeVf64(buffer)
		i3 = DecodeVf64(buffer)
		i4 = DecodeVf64(buffer)
		i5 = DecodeVf64(buffer)
		i6 = DecodeVf64(buffer)
		i7 = DecodeVf64(buffer)
		i8 = DecodeVf64(buffer)
	)
	assert.EqualValues(t, u1, i1)
	assert.EqualValues(t, u2, i2)
	assert.EqualValues(t, u3, i3)
	assert.EqualValues(t, u4, i4)
	assert.EqualValues(t, u5, i5)
	assert.EqualValues(t, u6, i6)
	assert.EqualValues(t, u7, i7)
	assert.EqualValues(t, u8, i8)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func TestEncodeDecodeVString(t *testing.T) {
	var (
		s1       = String(rand.Intn(100))
		s2       = String(rand.Intn(100))
		s3       = String(rand.Intn(100))
		s4       = String(rand.Intn(100))
		s5       = String(rand.Intn(100))
		s6       = String(rand.Intn(100))
		s7       = String(rand.Intn(100))
		s8       = String(rand.Intn(100))
		totalLen = 0
	)
	totalLen += EncodedLengthByVString(s1)
	totalLen += EncodedLengthByVString(s2)
	totalLen += EncodedLengthByVString(s3)
	totalLen += EncodedLengthByVString(s4)
	totalLen += EncodedLengthByVString(s5)
	totalLen += EncodedLengthByVString(s6)
	totalLen += EncodedLengthByVString(s7)
	totalLen += EncodedLengthByVString(s8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeVString(buffer, s1)
	EncodeVString(buffer, s2)
	EncodeVString(buffer, s3)
	EncodeVString(buffer, s4)
	EncodeVString(buffer, s5)
	EncodeVString(buffer, s6)
	EncodeVString(buffer, s7)
	EncodeVString(buffer, s8)

	buffer = bytes.NewBuffer(buf)
	var (
		str1 = DecodeVString(buffer)
		str2 = DecodeVString(buffer)
		str3 = DecodeVString(buffer)
		str4 = DecodeVString(buffer)
		str5 = DecodeVString(buffer)
		str6 = DecodeVString(buffer)
		str7 = DecodeVString(buffer)
		str8 = DecodeVString(buffer)
	)
	assert.EqualValues(t, s1, str1)
	assert.EqualValues(t, s2, str2)
	assert.EqualValues(t, s3, str3)
	assert.EqualValues(t, s4, str4)
	assert.EqualValues(t, s5, str5)
	assert.EqualValues(t, s6, str6)
	assert.EqualValues(t, s7, str7)
	assert.EqualValues(t, s8, str8)
}

func TestEncodeDecodeBytesString(t *testing.T) {
	var (
		bys1     = []byte(String(rand.Intn(100)))
		bys2     = []byte(String(rand.Intn(100)))
		bys3     = []byte(String(rand.Intn(100)))
		bys4     = []byte(String(rand.Intn(100)))
		bys5     = []byte(String(rand.Intn(100)))
		bys6     = []byte(String(rand.Intn(100)))
		bys7     = []byte(String(rand.Intn(100)))
		bys8     = []byte(String(rand.Intn(100)))
		totalLen = 0
	)
	totalLen += EncodedLengthByBytesString(bys1)
	totalLen += EncodedLengthByBytesString(bys2)
	totalLen += EncodedLengthByBytesString(bys3)
	totalLen += EncodedLengthByBytesString(bys4)
	totalLen += EncodedLengthByBytesString(bys5)
	totalLen += EncodedLengthByBytesString(bys6)
	totalLen += EncodedLengthByBytesString(bys7)
	totalLen += EncodedLengthByBytesString(bys8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeBytesString(buffer, bys1)
	EncodeBytesString(buffer, bys2)
	EncodeBytesString(buffer, bys3)
	EncodeBytesString(buffer, bys4)
	EncodeBytesString(buffer, bys5)
	EncodeBytesString(buffer, bys6)
	EncodeBytesString(buffer, bys7)
	EncodeBytesString(buffer, bys8)

	buffer = bytes.NewBuffer(buf)
	var (
		bytes1 = DecodeBytesString(buffer)
		bytes2 = DecodeBytesString(buffer)
		bytes3 = DecodeBytesString(buffer)
		bytes4 = DecodeBytesString(buffer)
		bytes5 = DecodeBytesString(buffer)
		bytes6 = DecodeBytesString(buffer)
		bytes7 = DecodeBytesString(buffer)
		bytes8 = DecodeBytesString(buffer)
	)
	assert.EqualValues(t, bys1, bytes1)
	assert.EqualValues(t, bys2, bytes2)
	assert.EqualValues(t, bys3, bytes3)
	assert.EqualValues(t, bys4, bytes4)
	assert.EqualValues(t, bys5, bytes5)
	assert.EqualValues(t, bys6, bytes6)
	assert.EqualValues(t, bys7, bytes7)
	assert.EqualValues(t, bys8, bytes8)
}

func TestEncodeDecodeBytes(t *testing.T) {
	var (
		bys1     = []byte(String(rand.Intn(100)))
		bys2     = []byte(String(rand.Intn(100)))
		bys3     = []byte(String(rand.Intn(100)))
		bys4     = []byte(String(rand.Intn(100)))
		bys5     = []byte(String(rand.Intn(100)))
		bys6     = []byte(String(rand.Intn(100)))
		bys7     = []byte(String(rand.Intn(100)))
		bys8     = []byte(String(rand.Intn(100)))
		totalLen = 0
	)
	totalLen += EncodedLengthByBytes(bys1)
	totalLen += EncodedLengthByBytes(bys2)
	totalLen += EncodedLengthByBytes(bys3)
	totalLen += EncodedLengthByBytes(bys4)
	totalLen += EncodedLengthByBytes(bys5)
	totalLen += EncodedLengthByBytes(bys6)
	totalLen += EncodedLengthByBytes(bys7)
	totalLen += EncodedLengthByBytes(bys8)
	buf := make([]byte, totalLen)
	buffer := bytes.NewBuffer(buf)
	EncodeBytes(buffer, bys1)
	EncodeBytes(buffer, bys2)
	EncodeBytes(buffer, bys3)
	EncodeBytes(buffer, bys4)
	EncodeBytes(buffer, bys5)
	EncodeBytes(buffer, bys6)
	EncodeBytes(buffer, bys7)
	EncodeBytes(buffer, bys8)

	buffer = bytes.NewBuffer(buf)
	var (
		bytes1 = DecodeBytes(buffer)
		bytes2 = DecodeBytes(buffer)
		bytes3 = DecodeBytes(buffer)
		bytes4 = DecodeBytes(buffer)
		bytes5 = DecodeBytes(buffer)
		bytes6 = DecodeBytes(buffer)
		bytes7 = DecodeBytes(buffer)
		bytes8 = DecodeBytes(buffer)
	)
	assert.EqualValues(t, bys1, bytes1)
	assert.EqualValues(t, bys2, bytes2)
	assert.EqualValues(t, bys3, bytes3)
	assert.EqualValues(t, bys4, bytes4)
	assert.EqualValues(t, bys5, bytes5)
	assert.EqualValues(t, bys6, bytes6)
	assert.EqualValues(t, bys7, bytes7)
	assert.EqualValues(t, bys8, bytes8)
}

/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at
 *          http//license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS WITHOUT WARRANTIES OF ANY KIND
 * EITHER EXPRESS OR IMPLIED INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObHTableFilter struct {
	ObUniVersionHeader
	isValid               bool
	selectColumnQualifier [][]byte
	minStamp              int64
	maxStamp              int64
	maxVersions           int32
	limitPerRowPerCf      int32
	offsetPerRowPerCf     int32
	filterString          string
}

func (f *ObHTableFilter) IsValid() bool {
	return f.isValid
}

func (f *ObHTableFilter) SetIsValid(isValid bool) {
	f.isValid = isValid
}

func (f *ObHTableFilter) SelectColumnQualifier() [][]byte {
	return f.selectColumnQualifier
}

func (f *ObHTableFilter) SetSelectColumnQualifier(selectColumnQualifier [][]byte) {
	f.selectColumnQualifier = selectColumnQualifier
}

func (f *ObHTableFilter) MinStamp() int64 {
	return f.minStamp
}

func (f *ObHTableFilter) SetMinStamp(minStamp int64) {
	f.minStamp = minStamp
}

func (f *ObHTableFilter) MaxStamp() int64 {
	return f.maxStamp
}

func (f *ObHTableFilter) SetMaxStamp(maxStamp int64) {
	f.maxStamp = maxStamp
}

func (f *ObHTableFilter) MaxVersions() int32 {
	return f.maxVersions
}

func (f *ObHTableFilter) SetMaxVersions(maxVersions int32) {
	f.maxVersions = maxVersions
}

func (f *ObHTableFilter) LimitPerRowPerCf() int32 {
	return f.limitPerRowPerCf
}

func (f *ObHTableFilter) SetLimitPerRowPerCf(limitPerRowPerCf int32) {
	f.limitPerRowPerCf = limitPerRowPerCf
}

func (f *ObHTableFilter) OffsetPerRowPerCf() int32 {
	return f.offsetPerRowPerCf
}

func (f *ObHTableFilter) SetOffsetPerRowPerCf(offsetPerRowPerCf int32) {
	f.offsetPerRowPerCf = offsetPerRowPerCf
}

func (f *ObHTableFilter) FilterString() string {
	return f.filterString
}

func (f *ObHTableFilter) SetFilterString(filterString string) {
	f.filterString = filterString
}

func (f *ObHTableFilter) PayloadLen() int {
	return f.PayloadContentLen() + f.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (f *ObHTableFilter) PayloadContentLen() int {
	totalLen := 0
	totalLen += 1 // isValid
	totalLen += util.EncodedLengthByVi64(int64(len(f.selectColumnQualifier)))
	for _, bys := range f.selectColumnQualifier {
		totalLen += util.EncodedLengthByBytesString(bys)
	}

	totalLen += util.EncodedLengthByVi64(f.minStamp) +
		util.EncodedLengthByVi64(f.maxStamp) +
		util.EncodedLengthByVi32(f.maxVersions) +
		util.EncodedLengthByVi32(f.limitPerRowPerCf) +
		util.EncodedLengthByVi32(f.offsetPerRowPerCf) +
		util.EncodedLengthByVString(f.filterString)
	return totalLen
}

func (f *ObHTableFilter) Encode(buffer *bytes.Buffer) {
	f.ObUniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, util.BoolToByte(f.isValid))

	util.EncodeVi64(buffer, int64(len(f.selectColumnQualifier)))
	for _, bys := range f.selectColumnQualifier {
		util.EncodeBytesString(buffer, bys)
	}

	util.EncodeVi64(buffer, f.minStamp)
	util.EncodeVi64(buffer, f.maxStamp)

	util.EncodeVi32(buffer, f.maxVersions)
	util.EncodeVi32(buffer, f.limitPerRowPerCf)
	util.EncodeVi32(buffer, f.offsetPerRowPerCf)

	util.EncodeVString(buffer, f.filterString)
}

func (f *ObHTableFilter) Decode(buffer *bytes.Buffer) {
}

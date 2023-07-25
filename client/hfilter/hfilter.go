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

package hfilter

import (
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"math"
)

type ObHTableFilter interface {
	SetIsValid(isValid bool)
	IsValid() bool
	SetSelectColumnQualifier(selectColumnQualifier [][]byte)
	SelectColumnQualifier() [][]byte
	SetMinStamp(minStamp int64)
	MinStamp() int64
	SetMaxStamp(maxStamp int64)
	MaxStamp() int64
	SetMaxVersions(maxVersions int32)
	MaxVersions() int32
	SetLimitPerRowPerCf(limitPerRowPerCf int32)
	LimitPerRowPerCf() int32
	SetOffsetPerRowPerCf(offsetPerRowPerCf int32)
	OffsetPerRowPerCf() int32
	SetFilterString(filterString string)
	FilterString() string
	Transfrom2Proto() *protocol.ObHTableFilter
}

type obHTableFilter struct {
	isValid               bool
	selectColumnQualifier [][]byte
	minStamp              int64
	maxStamp              int64
	maxVersions           int32
	limitPerRowPerCf      int32
	offsetPerRowPerCf     int32
	filterString          string
}

func NewObHTableFilter() ObHTableFilter {
	return &obHTableFilter{
		isValid:               true,
		selectColumnQualifier: nil,
		minStamp:              0,
		maxStamp:              math.MaxInt64,
		maxVersions:           1,
		limitPerRowPerCf:      -1,
		offsetPerRowPerCf:     0,
		filterString:          "",
	}
}

// SetIsValid set isValid
func (f *obHTableFilter) SetIsValid(isValid bool) {
	f.isValid = isValid
}

// IsValid get isValid
func (f *obHTableFilter) IsValid() bool {
	return f.isValid
}

// SetSelectColumnQualifier set selectColumnQualifier
func (f *obHTableFilter) SetSelectColumnQualifier(selectColumnQualifier [][]byte) {
	f.selectColumnQualifier = selectColumnQualifier
}

// SelectColumnQualifier get selectColumnQualifier
func (f *obHTableFilter) SelectColumnQualifier() [][]byte {
	return f.selectColumnQualifier
}

// SetMinStamp set minStamp
func (f *obHTableFilter) SetMinStamp(minStamp int64) {
	f.minStamp = minStamp
}

// MinStamp get minStamp
func (f *obHTableFilter) MinStamp() int64 {
	return f.minStamp
}

// SetMaxStamp set maxStamp
func (f *obHTableFilter) SetMaxStamp(maxStamp int64) {
	f.maxStamp = maxStamp
}

// MaxStamp get maxStamp
func (f *obHTableFilter) MaxStamp() int64 {
	return f.maxStamp
}

// SetMaxVersions set maxVersions
func (f *obHTableFilter) SetMaxVersions(maxVersions int32) {
	f.maxVersions = maxVersions
}

// MaxVersions get maxVersions
func (f *obHTableFilter) MaxVersions() int32 {
	return f.maxVersions
}

// SetLimitPerRowPerCf set limitPerRowPerCf
func (f *obHTableFilter) SetLimitPerRowPerCf(limitPerRowPerCf int32) {
	f.limitPerRowPerCf = limitPerRowPerCf
}

// LimitPerRowPerCf get limitPerRowPerCf
func (f *obHTableFilter) LimitPerRowPerCf() int32 {
	return f.limitPerRowPerCf
}

// SetOffsetPerRowPerCf set offsetPerRowPerCf
func (f *obHTableFilter) SetOffsetPerRowPerCf(offsetPerRowPerCf int32) {
	f.offsetPerRowPerCf = offsetPerRowPerCf
}

// OffsetPerRowPerCf get offsetPerRowPerCf
func (f *obHTableFilter) OffsetPerRowPerCf() int32 {
	return f.offsetPerRowPerCf
}

// SetFilterString set filterString
func (f *obHTableFilter) SetFilterString(filterString string) {
	f.filterString = filterString
}

// FilterString get filterString
func (f *obHTableFilter) FilterString() string {
	return f.filterString
}

// Transfrom2Proto transform to protocol ObHTableFilter
func (f *obHTableFilter) Transfrom2Proto() *protocol.ObHTableFilter {
	proto := protocol.NewObHTableFilter()
	proto.SetIsValid(f.isValid)
	proto.SetSelectColumnQualifier(f.selectColumnQualifier)
	proto.SetMinStamp(f.minStamp)
	proto.SetMaxStamp(f.maxStamp)
	proto.SetMaxVersions(f.maxVersions)
	proto.SetLimitPerRowPerCf(f.limitPerRowPerCf)
	proto.SetOffsetPerRowPerCf(f.offsetPerRowPerCf)
	proto.SetFilterString(f.filterString)
	return proto
}

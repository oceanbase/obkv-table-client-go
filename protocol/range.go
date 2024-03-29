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

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObNewRange struct {
	tableId    uint64
	borderFlag ObBorderFlag
	startKey   []*ObObject
	endKey     []*ObObject
	flag       int64
}

func NewObNewRange() *ObNewRange {
	return &ObNewRange{
		tableId:    0,
		borderFlag: 0,
		startKey:   nil,
		endKey:     nil,
		flag:       0,
	}
}

func NewObNewRangeWithColumns(startKeyColumns []*table.Column, endKeyColumns []*table.Column) (*ObNewRange, error) {
	startKey := make([]*ObObject, 0, len(startKeyColumns))
	endKey := make([]*ObObject, 0, len(endKeyColumns))

	// add start key columns
	for _, column := range startKeyColumns {
		objMeta, err := DefaultObjMeta(column.Value())
		if err != nil {
			return nil, errors.WithMessage(err, "create obj meta by start key")
		}
		object := NewObObjectWithParams(objMeta, column.Value())
		startKey = append(startKey, object)
	}

	// add end key columns
	for _, column := range endKeyColumns {
		objMeta, err := DefaultObjMeta(column.Value())
		if err != nil {
			return nil, errors.WithMessage(err, "create obj meta by end key")
		}
		object := NewObObjectWithParams(objMeta, column.Value())
		endKey = append(endKey, object)
	}

	obBorderFlag := NewObBorderFlag()

	return &ObNewRange{
		tableId:    0,
		borderFlag: obBorderFlag,
		startKey:   startKey,
		endKey:     endKey,
		flag:       0,
	}, nil
}

// NewObNewRangeWithParams creates a new ObNewRange.
func NewObNewRangeWithParams(startKey []*ObObject, endKey []*ObObject, borderFlag ObBorderFlag) *ObNewRange {
	return &ObNewRange{
		tableId:    0,
		borderFlag: borderFlag,
		startKey:   startKey,
		endKey:     endKey,
		flag:       0,
	}
}

func (r *ObNewRange) TableId() uint64 {
	return r.tableId
}

func (r *ObNewRange) SetTableId(tableId uint64) {
	r.tableId = tableId
}

func (r *ObNewRange) BorderFlag() ObBorderFlag {
	return r.borderFlag
}

func (r *ObNewRange) SetBorderFlag(borderFlag ObBorderFlag) {
	r.borderFlag = borderFlag
}

func (r *ObNewRange) StartKey() []*ObObject {
	return r.startKey
}

func (r *ObNewRange) SetStartKey(startKey []*ObObject) {
	r.startKey = startKey
}

func (r *ObNewRange) EndKey() []*ObObject {
	return r.endKey
}

func (r *ObNewRange) SetEndKey(endKey []*ObObject) {
	r.endKey = endKey
}

func (r *ObNewRange) Flag() int64 {
	return r.flag
}

func (r *ObNewRange) SetFlag(flag int64) {
	r.flag = flag
}

func (r *ObNewRange) EncodedLength() int {
	totalLen := 0
	totalLen += util.EncodedLengthByVi64(int64(r.tableId)) +
		1 // borderFlag
	totalLen += util.EncodedLengthByVi64(int64(len(r.startKey)))
	for _, obObject := range r.startKey {
		totalLen += obObject.EncodedLength()
	}

	totalLen += util.EncodedLengthByVi64(int64(len(r.endKey)))
	for _, obObject := range r.endKey {
		totalLen += obObject.EncodedLength()
	}

	if util.ObVersion() >= 4 {
		totalLen += util.EncodedLengthByVi64(r.flag)
	}

	return totalLen
}

func (r *ObNewRange) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, int64(r.tableId))

	util.PutUint8(buffer, uint8(r.borderFlag))

	util.EncodeVi64(buffer, int64(len(r.startKey)))
	for _, obObject := range r.startKey {
		obObject.Encode(buffer)
	}

	util.EncodeVi64(buffer, int64(len(r.endKey)))
	for _, obObject := range r.endKey {
		obObject.Encode(buffer)
	}

	if util.ObVersion() >= 4 {
		util.EncodeVi64(buffer, r.flag)
	}
}

func (r *ObNewRange) Decode(buffer *bytes.Buffer) {
	r.tableId = uint64(util.DecodeVi64(buffer))

	r.borderFlag = ObBorderFlag(util.Uint8(buffer))

	startKeyLen := util.DecodeVi64(buffer)
	r.startKey = make([]*ObObject, 0, startKeyLen)
	var i int64
	for i = 0; i < startKeyLen; i++ {
		obObject := NewObObject()
		obObject.Decode(buffer)
		r.startKey = append(r.startKey, obObject)
	}

	endKeyLen := util.DecodeVi64(buffer)
	r.endKey = make([]*ObObject, 0, startKeyLen)
	for i = 0; i < endKeyLen; i++ {
		obObject := NewObObject()
		obObject.Decode(buffer)
		r.endKey = append(r.endKey, obObject)
	}

	if util.ObVersion() >= 4 {
		r.flag = util.DecodeVi64(buffer)
	}
}

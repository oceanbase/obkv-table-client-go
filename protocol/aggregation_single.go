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

type ObTableAggregation struct {
	ObUniVersionHeader
	aggType   ObTableAggregationType
	aggColumn string
}

func NewObTableAggregation() *ObTableAggregation {
	return &ObTableAggregation{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		aggType:   0,
		aggColumn: "",
	}
}

func NewObTableAggregationWithParams(aggregationType ObTableAggregationType, aggregationColumn string) *ObTableAggregation {
	return &ObTableAggregation{
		ObUniVersionHeader: ObUniVersionHeader{
			version:       1,
			contentLength: 0,
		},
		aggType:   aggregationType,
		aggColumn: aggregationColumn,
	}
}

type ObTableAggregationType uint8

const (
	ObTableAggregationTypeInvaild ObTableAggregationType = iota
	ObTableAggregationTypeMax
	ObTableAggregationTypeMin
	ObTableAggregationTypeCount
	ObTableAggregationTypeSum
	ObTableAggregationTypeAvg
)

func (s *ObTableAggregation) AggOperation() string {
	switch s.aggType {
	case ObTableAggregationTypeMax:
		return "max(" + s.aggColumn + ")"
	case ObTableAggregationTypeMin:
		return "min(" + s.aggColumn + ")"
	case ObTableAggregationTypeCount:
		return "count(" + s.aggColumn + ")"
	case ObTableAggregationTypeSum:
		return "sum(" + s.aggColumn + ")"
	case ObTableAggregationTypeAvg:
		return "avg(" + s.aggColumn + ")"
	}
	return "invalid"
}

func (s *ObTableAggregation) AggType() ObTableAggregationType {
	return s.aggType
}

func (s *ObTableAggregation) SetAggType(aggType ObTableAggregationType) {
	s.aggType = aggType
}

func (s *ObTableAggregation) AggColumn() string {
	return s.aggColumn
}

func (s *ObTableAggregation) SetAggColumn(aggColumn string) {
	s.aggColumn = aggColumn
}

func (s *ObTableAggregation) PayloadLen() int {
	return s.PayloadContentLen() + s.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (s *ObTableAggregation) PayloadContentLen() int {
	totalLen := 0
	totalLen += 1 // aggType
	totalLen += util.EncodedLengthByVString(s.aggColumn)

	s.ObUniVersionHeader.SetContentLength(totalLen)
	return s.ObUniVersionHeader.ContentLength()
}

func (s *ObTableAggregation) Encode(buffer *bytes.Buffer) {
	s.ObUniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, uint8(s.aggType))

	util.EncodeVString(buffer, s.aggColumn)
}

func (s *ObTableAggregation) Decode(buffer *bytes.Buffer) {
	s.ObUniVersionHeader.Decode(buffer)

	s.aggType = ObTableAggregationType(util.Uint8(buffer))

	s.aggColumn = util.DecodeVString(buffer)
}

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

type ObTableAggregationSingle struct {
	ObUniVersionHeader
	aggType   ObTableAggregationType
	aggColumn string
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

func (s *ObTableAggregationSingle) AggType() ObTableAggregationType {
	return s.aggType
}

func (s *ObTableAggregationSingle) SetAggType(aggType ObTableAggregationType) {
	s.aggType = aggType
}

func (s *ObTableAggregationSingle) AggColumn() string {
	return s.aggColumn
}

func (s *ObTableAggregationSingle) SetAggColumn(aggColumn string) {
	s.aggColumn = aggColumn
}

func (s *ObTableAggregationSingle) PayloadLen() int {
	return s.PayloadContentLen() + s.ObUniVersionHeader.UniVersionHeaderLen() // Do not change the order
}

func (s *ObTableAggregationSingle) PayloadContentLen() int {
	totalLen := 0
	totalLen += 1 // aggType
	totalLen += util.EncodedLengthByVString(s.aggColumn)
	return totalLen
}

func (s *ObTableAggregationSingle) Encode(buffer *bytes.Buffer) {
	s.ObUniVersionHeader.Encode(buffer)

	util.PutUint8(buffer, uint8(s.aggType))

	util.EncodeVString(buffer, s.aggColumn)
}

func (s *ObTableAggregationSingle) Decode(buffer *bytes.Buffer) {
}

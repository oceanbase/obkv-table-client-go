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

package table

type RangePair struct {
	start          []*Column
	end            []*Column
	inclusiveStart bool
	inclusiveEnd   bool
}

// NewRangePair creates a new RangePair.
// If inclusive is not specified, it defaults to true.
// If only one inclusive is specified, it is used for both.
// If both inclusive are specified, the first is used for the start and the second for the end.
func NewRangePair(start []*Column, end []*Column, inclusive ...bool) *RangePair {
	if len(inclusive) == 2 {
		return &RangePair{start: start, end: end, inclusiveStart: inclusive[0], inclusiveEnd: inclusive[1]}
	} else if len(inclusive) == 1 {
		return &RangePair{start: start, end: end, inclusiveStart: inclusive[0], inclusiveEnd: inclusive[0]}
	}
	return &RangePair{start: start, end: end, inclusiveStart: true, inclusiveEnd: true}
}

// Start returns the start of the range.
func (c *RangePair) Start() []*Column {
	return c.start
}

// SetStart sets the start of the range.
func (c *RangePair) SetStart(start []*Column) {
	c.start = start
}

// AddStart adds a column to the start of the range.
func (c *RangePair) AddStart(column *Column) {
	if c.start == nil {
		c.start = make([]*Column, 0)
	}
	c.start = append(c.start, column)
}

// End returns the end of the range.
func (c *RangePair) End() []*Column {
	return c.end
}

// SetEnd sets the end of the range.
func (c *RangePair) SetEnd(end []*Column) {
	c.end = end
}

// AddEnd adds a column to the end of the range.
func (c *RangePair) AddEnd(column *Column) {
	if c.end == nil {
		c.end = make([]*Column, 0)
	}
	c.end = append(c.end, column)
}

// IncludeStart returns true if the start is inclusive.
func (c *RangePair) IncludeStart() bool {
	return c.inclusiveStart
}

// SetIncludeStart sets the start to be inclusive.
func (c *RangePair) SetIncludeStart(inclusiveStart bool) {
	c.inclusiveStart = inclusiveStart
}

// IncludeEnd returns true if the end is inclusive.
func (c *RangePair) IncludeEnd() bool {
	return c.inclusiveEnd
}

// SetIncludeEnd sets the end to be inclusive.
func (c *RangePair) SetIncludeEnd(inclusiveEnd bool) {
	c.inclusiveEnd = inclusiveEnd
}

// IsStartEqEnd returns true if the start is equal to the end.
func (c *RangePair) IsStartEqEnd() bool {
	for i := 0; i < len(c.start); i++ {
		if !c.start[i].IsEqual(c.end[i]) {
			return false
		}
	}
	return true
}

// String returns a string representation of the RangePair.
func (c *RangePair) String() string {
	retString := "RangePair{"
	if c.start != nil {
		retString += "start:["
		for _, column := range c.start {
			retString += column.String()
		}
		retString += "]"
	}
	if c.inclusiveStart {
		retString += " inclusiveStart,"
	}
	if c.end != nil {
		retString += "end:["
		for _, column := range c.end {
			retString += column.String()
		}
		retString += "]"
	}
	if c.inclusiveEnd {
		retString += " inclusiveEnd"
	}
	retString += "}"
	return retString
}

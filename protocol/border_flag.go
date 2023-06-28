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

type ObBorderFlag int8

const (
	InclusiveStart     ObBorderFlag = 0x1
	InclusiveEnd       ObBorderFlag = 0x2
	BorderFlagMinValue ObBorderFlag = 0x3
	BorderFlagMaxValue ObBorderFlag = 0x8
)

// NewObBorderFlag creates a new ObBorderFlag.
func NewObBorderFlag() ObBorderFlag {
	borderFlag := ObBorderFlag(0)
	borderFlag.SetInclusiveStart()
	borderFlag.SetInclusiveEnd()
	return borderFlag
}

// SetInclusiveStart sets the InclusiveStart flag.
func (bf *ObBorderFlag) SetInclusiveStart() {
	*bf |= InclusiveStart
}

// UnSetInclusiveStart unsets the InclusiveStart flag.
func (bf *ObBorderFlag) UnSetInclusiveStart() {
	*bf &= ^InclusiveStart
}

// IsInclusiveStart returns true if the InclusiveStart flag is set.
func (bf *ObBorderFlag) IsInclusiveStart() bool {
	return *bf&InclusiveStart == InclusiveStart
}

// SetInclusiveEnd sets the InclusiveEnd flag.
func (bf *ObBorderFlag) SetInclusiveEnd() {
	*bf |= InclusiveEnd
}

// UnSetInclusiveEnd unsets the InclusiveEnd flag.
func (bf *ObBorderFlag) UnSetInclusiveEnd() {
	*bf &= ^InclusiveEnd
}

// IsInclusiveEnd returns true if the InclusiveEnd flag is set.
func (bf *ObBorderFlag) IsInclusiveEnd() bool {
	return *bf&InclusiveEnd == InclusiveEnd
}

// SetMinValue sets the BorderFlagMinValue flag.
func (bf *ObBorderFlag) SetMinValue() {
	*bf |= BorderFlagMinValue
}

// UnSetMinValue unsets the BorderFlagMinValue flag.
func (bf *ObBorderFlag) UnSetMinValue() {
	*bf &= ^BorderFlagMinValue
}

// IsMinValue returns true if the BorderFlagMinValue flag is set.
func (bf *ObBorderFlag) IsMinValue() bool {
	return *bf&BorderFlagMinValue == BorderFlagMinValue
}

// SetMaxValue sets the BorderFlagMaxValue flag.
func (bf *ObBorderFlag) SetMaxValue() {
	*bf |= BorderFlagMaxValue
}

// UnSetMaxValue unsets the BorderFlagMaxValue flag.
func (bf *ObBorderFlag) UnSetMaxValue() {
	*bf &= ^BorderFlagMaxValue
}

// IsMaxValue returns true if the BorderFlagMaxValue flag is set.
func (bf *ObBorderFlag) IsMaxValue() bool {
	return *bf&BorderFlagMaxValue == BorderFlagMaxValue
}

func (bf *ObBorderFlag) String() string {
	switch *bf {
	case InclusiveStart:
		return "InclusiveStart"
	case InclusiveEnd:
		return "InclusiveEnd"
	case BorderFlagMinValue:
		return "BorderFlagMinValue"
	case BorderFlagMaxValue:
		return "BorderFlagMaxValue"
	default:
		return "unknown"
	}
}

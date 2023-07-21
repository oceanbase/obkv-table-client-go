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
	inclusiveStart     ObBorderFlag = 0x1
	inclusiveEnd       ObBorderFlag = 0x2
	borderFlagMinValue ObBorderFlag = 0x3
	borderFlagMaxValue ObBorderFlag = 0x8
)

// NewObBorderFlag creates a new ObBorderFlag.
// The default value is inclusiveStart and inclusiveEnd.
func NewObBorderFlag() ObBorderFlag {
	borderFlag := ObBorderFlag(0)
	borderFlag.SetInclusiveStart()
	borderFlag.SetInclusiveEnd()
	return borderFlag
}

// SetInclusiveStart sets the inclusiveStart flag.
func (bf *ObBorderFlag) SetInclusiveStart() {
	*bf |= inclusiveStart
}

// UnSetInclusiveStart unsets the inclusiveStart flag.
func (bf *ObBorderFlag) UnSetInclusiveStart() {
	*bf &= ^inclusiveStart
}

// IsInclusiveStart returns true if the inclusiveStart flag is set.
func (bf *ObBorderFlag) IsInclusiveStart() bool {
	return *bf&inclusiveStart == inclusiveStart
}

// SetInclusiveEnd sets the inclusiveEnd flag.
func (bf *ObBorderFlag) SetInclusiveEnd() {
	*bf |= inclusiveEnd
}

// UnSetInclusiveEnd unsets the inclusiveEnd flag.
func (bf *ObBorderFlag) UnSetInclusiveEnd() {
	*bf &= ^inclusiveEnd
}

// IsInclusiveEnd returns true if the inclusiveEnd flag is set.
func (bf *ObBorderFlag) IsInclusiveEnd() bool {
	return *bf&inclusiveEnd == inclusiveEnd
}

// SetMinValue sets the borderFlagMinValue flag.
func (bf *ObBorderFlag) SetMinValue() {
	*bf |= borderFlagMinValue
}

// UnSetMinValue unsets the borderFlagMinValue flag.
func (bf *ObBorderFlag) UnSetMinValue() {
	*bf &= ^borderFlagMinValue
}

// IsMinValue returns true if the borderFlagMinValue flag is set.
func (bf *ObBorderFlag) IsMinValue() bool {
	return *bf&borderFlagMinValue == borderFlagMinValue
}

// SetMaxValue sets the borderFlagMaxValue flag.
func (bf *ObBorderFlag) SetMaxValue() {
	*bf |= borderFlagMaxValue
}

// UnSetMaxValue unsets the borderFlagMaxValue flag.
func (bf *ObBorderFlag) UnSetMaxValue() {
	*bf &= ^borderFlagMaxValue
}

// IsMaxValue returns true if the borderFlagMaxValue flag is set.
func (bf *ObBorderFlag) IsMaxValue() bool {
	return *bf&borderFlagMaxValue == borderFlagMaxValue
}

func (bf *ObBorderFlag) String() string {
	switch *bf {
	case inclusiveStart:
		return "inclusiveStart"
	case inclusiveEnd:
		return "inclusiveEnd"
	case borderFlagMinValue:
		return "borderFlagMinValue"
	case borderFlagMaxValue:
		return "borderFlagMaxValue"
	default:
		return "unknown"
	}
}

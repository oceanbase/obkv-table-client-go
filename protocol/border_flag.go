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
	INCLUSIVE_START       ObBorderFlag = 0x1
	INCLUSIVE_END         ObBorderFlag = 0x2
	BORDER_FLAG_MIN_VALUE ObBorderFlag = 0x3
	BORDER_FLAG_MAX_VALUE ObBorderFlag = 0x8
)

// NewObBorderFlag creates a new ObBorderFlag.
func NewObBorderFlag() ObBorderFlag {
	borderFlag := ObBorderFlag(0)
	borderFlag.SetInclusiveStart()
	borderFlag.SetInclusiveEnd()
	return borderFlag
}

// SetInclusiveStart sets the INCLUSIVE_START flag.
func (bf *ObBorderFlag) SetInclusiveStart() {
	*bf |= INCLUSIVE_START
}

// UnSetInclusiveStart unsets the INCLUSIVE_START flag.
func (bf *ObBorderFlag) UnSetInclusiveStart() {
	*bf &= ^INCLUSIVE_START
}

// IsInclusiveStart returns true if the INCLUSIVE_START flag is set.
func (bf ObBorderFlag) IsInclusiveStart() bool {
	return bf&INCLUSIVE_START == INCLUSIVE_START
}

// SetInclusiveEnd sets the INCLUSIVE_END flag.
func (bf *ObBorderFlag) SetInclusiveEnd() {
	*bf |= INCLUSIVE_END
}

// UnSetInclusiveEnd unsets the INCLUSIVE_END flag.
func (bf *ObBorderFlag) UnSetInclusiveEnd() {
	*bf &= ^INCLUSIVE_END
}

// IsInclusiveEnd returns true if the INCLUSIVE_END flag is set.
func (bf ObBorderFlag) IsInclusiveEnd() bool {
	return bf&INCLUSIVE_END == INCLUSIVE_END
}

// SetMinValue sets the BORDER_FLAG_MIN_VALUE flag.
func (bf ObBorderFlag) SetMinValue() {
	bf |= BORDER_FLAG_MIN_VALUE
}

// UnSetMinValue unsets the BORDER_FLAG_MIN_VALUE flag.
func (bf ObBorderFlag) UnSetMinValue() {
	bf &= ^BORDER_FLAG_MIN_VALUE
}

// IsMinValue returns true if the BORDER_FLAG_MIN_VALUE flag is set.
func (bf ObBorderFlag) IsMinValue() bool {
	return bf&BORDER_FLAG_MIN_VALUE == BORDER_FLAG_MIN_VALUE
}

// SetMaxValue sets the BORDER_FLAG_MAX_VALUE flag.
func (bf ObBorderFlag) SetMaxValue() {
	bf |= BORDER_FLAG_MAX_VALUE
}

// UnSetMaxValue unsets the BORDER_FLAG_MAX_VALUE flag.
func (bf ObBorderFlag) UnSetMaxValue() {
	bf &= ^BORDER_FLAG_MAX_VALUE
}

// IsMaxValue returns true if the BORDER_FLAG_MAX_VALUE flag is set.
func (bf ObBorderFlag) IsMaxValue() bool {
	return bf&BORDER_FLAG_MAX_VALUE == BORDER_FLAG_MAX_VALUE
}

func (bf ObBorderFlag) String() string {
	switch bf {
	case INCLUSIVE_START:
		return "INCLUSIVE_START"
	case INCLUSIVE_END:
		return "INCLUSIVE_END"
	case BORDER_FLAG_MIN_VALUE:
		return "BORDER_FLAG_MIN_VALUE"
	case BORDER_FLAG_MAX_VALUE:
		return "BORDER_FLAG_MAX_VALUE"
	default:
		return "UNKNOWN"
	}
}

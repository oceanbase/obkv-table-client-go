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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	s1 := "StringNeedToBeCalculatedCheckSumAndTest,Make String longer longer longer longer longer to test"
	assert.EqualValues(t, 1566783161, Calculate(0, []byte(s1)))
	s1 = "String Need To Be Calculated CheckSum And Test,Make String longer longer longer longer longer to test"
	assert.EqualValues(t, Calculate(0x05010927, []byte(s1)), Calculate(0x05010927, []byte(s1)))
}

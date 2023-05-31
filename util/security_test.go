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

func TestGetPasswordScramble(t *testing.T) {
	for i := 0; i < 20; i++ {
		s := GetPasswordScramble(i)
		assert.EqualValues(t, i, len(s))
	}
}

func TestScramblePassword(t *testing.T) {
	password := "hello"
	seed := "qXA4YhW5PwaWsARvj3KC"
	scramblePassword1 := ScramblePassword(password, seed)
	scramblePassword2 := ScramblePassword(password, seed)
	assert.EqualValues(t, 20, len(scramblePassword1))
	assert.EqualValues(t, 20, len(scramblePassword2))
	assert.EqualValues(t, scramblePassword1, scramblePassword2)
	assert.EqualValues(t, "", ScramblePassword("", seed))
}

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

package filter

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObTableFilterListString(t *testing.T) {
	testColumnName := "testColumnName"
	value := rand.Intn(100)

	var filterList []ObTableFilter
	filterList = append(filterList, In(testColumnName, value))
	filterList = append(filterList, NotIn(testColumnName, value))
	filterList = append(filterList, CompareVal(LessThan, testColumnName, value))
	filterList = append(filterList, CompareVal(GreaterThan, testColumnName, value))
	filterList = append(filterList, CompareVal(LessOrEqualThan, testColumnName, value))
	filterList = append(filterList, CompareVal(GreaterOrEqualThan, testColumnName, value))
	filterList = append(filterList, CompareVal(NotEqual, testColumnName, value))
	filterList = append(filterList, CompareVal(Equal, testColumnName, value))
	filterList = append(filterList, CompareVal(IsNull, testColumnName, value))

	obTableFilterList := NewObTableFilterListWithOperatorAndTableFilter(OperatorAnd)
	obTableFilterList.AddFilter(filterList...)

	tableFilterList := AndList()
	tableFilterList.AddFilter(filterList...)

	assert.EqualValues(t, tableFilterList.Size(), obTableFilterList.Size())
	assert.EqualValues(t, tableFilterList.String(), obTableFilterList.String())
	assert.EqualValues(t, tableFilterList, obTableFilterList)
}

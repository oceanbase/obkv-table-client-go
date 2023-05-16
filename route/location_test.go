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

package route

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/protocol"
)

func TestGetOrderedPartColumns(t *testing.T) {
	desc := newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2, "c1")
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeBinary)
	desc.PartColumns = []*obColumn{col}
	partitionKeyColumns := []*obColumn{col}
	columns := getOrderedPartColumns(partitionKeyColumns, desc)
	assert.Equal(t, 1, len(columns))
}

func TestSetPartDescProperty(t *testing.T) {
	// key partition
	descKey := newObKeyPartDesc(0, 10, partFuncTypeKeyImplV2, "c1")
	objType, _ := protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col := newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeBinary)
	partColumns := []*obColumn{col}
	orderedCompareColumns := []*obColumn{col}
	err := setPartDescProperty(descKey, partColumns, orderedCompareColumns)
	assert.Equal(t, nil, err)
	// range partition
	descRange := newObRangePartDesc(0, 10, partFuncTypeRange, "c1")
	objType, _ = protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col = newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeBinary)
	partColumns = []*obColumn{col}
	orderedCompareColumns = []*obColumn{col}
	err = setPartDescProperty(descRange, partColumns, orderedCompareColumns)
	assert.Equal(t, nil, err)
	// hash partition
	descHash := newObHashPartDesc(0, 10, partFuncTypeHash, "c1")
	objType, _ = protocol.NewObjType(protocol.ObObjTypeInt64TypeValue)
	col = newObSimpleColumn("c1", 1, objType, protocol.ObCollationTypeBinary)
	partColumns = []*obColumn{col}
	orderedCompareColumns = []*obColumn{col}
	err = setPartDescProperty(descHash, partColumns, orderedCompareColumns)
	assert.NotEqual(t, nil, err)
}

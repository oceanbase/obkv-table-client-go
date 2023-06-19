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
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

func TestObTableQueryAndMutateResponseEncodeDecode(t *testing.T) {
	obTableQueryAndMutateResponse := NewObTableQueryAndMutateResponse()
	obTableQueryAndMutateResponse.SetAffectedRows(int64(rand.Uint64()))
	obTableQueryResponse := NewObTableQueryResponse()
	randomLen := rand.Intn(20)
	propertiesNames := make([]string, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		propertiesNames = append(propertiesNames, util.String(10))
	}
	obTableQueryResponse.SetPropertiesNames(propertiesNames)
	obTableQueryResponse.SetRowCount(int64(randomLen))
	propertiesRows := make([][]*ObObject, 0, randomLen)
	for i := 0; i < randomLen; i++ {
		obObjects := make([]*ObObject, 0, randomLen)
		for j := 0; j < randomLen; j++ {
			column := table.NewColumn(util.String(10), int64(rand.Intn(10000)))
			objMeta, _ := DefaultObjMeta(column.Value())
			object := NewObObjectWithParams(objMeta, column.Value())
			obObjects = append(obObjects, object)
		}
		propertiesRows = append(propertiesRows, obObjects)
	}
	obTableQueryResponse.SetPropertiesRows(propertiesRows)
	obTableQueryAndMutateResponse.SetAffectedEntity(obTableQueryResponse)

	payloadLen := obTableQueryAndMutateResponse.PayloadLen()
	buf := make([]byte, payloadLen)
	buffer := bytes.NewBuffer(buf)
	obTableQueryAndMutateResponse.Encode(buffer)

	newObTableQueryAndMutateResponse := NewObTableQueryAndMutateResponse()
	newBuffer := bytes.NewBuffer(buf)
	newObTableQueryAndMutateResponse.Decode(newBuffer)

	assert.EqualValues(t, obTableQueryAndMutateResponse.AffectedRows(), newObTableQueryAndMutateResponse.AffectedRows())
	assert.EqualValues(t, obTableQueryAndMutateResponse.AffectedEntity(), newObTableQueryAndMutateResponse.AffectedEntity())
	assert.EqualValues(t, obTableQueryAndMutateResponse, newObTableQueryAndMutateResponse)
}

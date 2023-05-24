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
	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/table"
)

const (
	ObInvalidPartId = 0
	ObPartIdBitNum  = 28
	ObPartIdShift   = 32
	ObMask          = (1 << ObPartIdBitNum) | 1<<(ObPartIdBitNum+ObPartIdShift)
	ObSubPartIdMask = 0xffffffff & 0xfffffff
)

type obPartDesc interface {
	String() string
	PartFuncType() obPartFuncType
	SetPartColumns(partColumns []obColumn)
	PartColumns() []obColumn
	SetPartNum(partNum int)
	GetPartId(rowKey []*table.Column) (uint64, error)
}

// evalPartKeyValues calculate the value of the partition key
func evalPartKeyValues(desc obPartDesc, rowKey []*table.Column) ([]interface{}, error) {
	if desc == nil {
		return nil, errors.New("part desc is nil")
	}
	partValues := make([]interface{}, 0, len(desc.PartColumns()))
	for _, column := range desc.PartColumns() {
		value, err := column.eval(rowKey)
		if err != nil {
			return nil, errors.WithMessagef(err, "eval column, column:%s", column.String())
		}
		partValues = append(partValues, value)
	}
	return partValues, nil
}

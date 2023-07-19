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
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

func newObSimpleColumn(
	columnName string,
	objType protocol.ObObjType,
	collationType protocol.ObCollationType) *obSimpleColumn {
	return &obSimpleColumn{columnName, objType, collationType}
}

type obSimpleColumn struct {
	columnName    string
	objType       protocol.ObObjType
	collationType protocol.ObCollationType
}

func (c *obSimpleColumn) CollationType() protocol.ObCollationType {
	return c.collationType
}

func (c *obSimpleColumn) ObjType() protocol.ObObjType {
	return c.objType
}

func (c *obSimpleColumn) ColumnName() string {
	return c.columnName
}

// eval calculate simple column value
func (c *obSimpleColumn) eval(rowKey []*table.Column) (interface{}, error) {
	for _, column := range rowKey {
		if strings.EqualFold(column.Name(), c.columnName) {
			return column.Value(), nil
		}
	}

	return nil, errors.Errorf("partition column not match, column:%s", c.String())
}

// extractColumn extract the same column from obSimpleColumn
func (c *obSimpleColumn) extractColumn(rowKey []*table.Column) (*table.Column, error) {
	for _, column := range rowKey {
		if strings.EqualFold(column.Name(), c.columnName) {
			return column, nil
		}
	}

	return nil, nil
}

func (c *obSimpleColumn) String() string {
	var objTypeStr = "nil"
	if c.objType != nil {
		objTypeStr = c.objType.String()
	}
	return "obSimpleColumn{" +
		"columnName:" + c.columnName + ", " +
		"objType:" + objTypeStr + ", " +
		"collationType:" + strconv.Itoa(int(c.collationType)) +
		"}"
}

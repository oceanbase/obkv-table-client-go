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

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

// not support generate column now, support later
type obGeneratedColumnSimpleFunc interface {
	String() string
	getRefColumnNames() []string
}

func newObGeneratedColumn(
	columnName string,
	objType protocol.ObObjType,
	collationType protocol.ObCollationType,
	refColumnNames []string,
	columnExpress obGeneratedColumnSimpleFunc,
) *obGeneratedColumn {
	return &obGeneratedColumn{columnName, objType, collationType, refColumnNames, columnExpress}
}

// obGeneratedColumn is a special column in a table.
// We currently only support generated column with substring expressions.
type obGeneratedColumn struct {
	columnName    string
	objType       protocol.ObObjType
	collationType protocol.ObCollationType
	// refColumnNames: Represents which columns are referenced by the current column
	// generate column: key_prefix VARCHAR(1024) GENERATED ALWAYS AS (SUBSTRING(`col1`, 1, 4)
	// 		refColumnNames = ["col1"]
	refColumnNames []string
	columnExpress  obGeneratedColumnSimpleFunc // only support 'SUBSTRING' expr now
}

func (c *obGeneratedColumn) CollationType() protocol.ObCollationType {
	return c.collationType
}

func (c *obGeneratedColumn) ObjType() protocol.ObObjType {
	return c.objType
}

func (c *obGeneratedColumn) ColumnName() string {
	return c.columnName
}

func (c *obGeneratedColumn) eval(rowKey []*table.Column) (interface{}, error) {
	return nil, errors.New("not support generated column now")
}

func (c *obGeneratedColumn) extractColumn(rowKey []*table.Column) (*table.Column, error) {
	return nil, errors.New("not support generated column now")
}

func (c *obGeneratedColumn) String() string {
	var objTypeStr = "nil"
	if c.objType != nil {
		objTypeStr = c.objType.String()
	}

	var columnExpressStr = "nil"
	if c.columnExpress != nil {
		columnExpressStr = c.columnExpress.String()
	}

	var refColumnNamesStr string
	refColumnNamesStr = refColumnNamesStr + "["
	for i := 0; i < len(c.refColumnNames); i++ {
		if i > 0 {
			refColumnNamesStr += ", "
		}
		refColumnNamesStr += c.refColumnNames[i]
	}
	refColumnNamesStr += "]"

	return "obGeneratedColumn{" +
		"columnName:" + c.columnName + ", " +
		"objType:" + objTypeStr + ", " +
		"collationType:" + strconv.Itoa(int(c.collationType)) + ", " +
		"refColumnNames:" + refColumnNamesStr + ", " +
		"columnExpress:" + columnExpressStr +
		"}"
}

package protocol

import (
	"strconv"
	"strings"
)

func NewObGeneratedColumn(
	columnName string,
	index int,
	objType ObObjType,
	collType ObCollationType,
	columnExpress ObGeneratedColumnSimpleFunc) *ObColumn {
	return &ObColumn{
		columnName:     columnName,
		index:          index,
		objType:        objType,
		collationType:  collType,
		refColumnNames: columnExpress.getRefColumnNames(),
		isGenColumn:    true,
		columnExpress:  columnExpress,
	}
}

func NewObSimpleColumn(
	columnName string,
	index int,
	objType ObObjType,
	collType ObCollationType) *ObColumn {
	return &ObColumn{
		columnName:     columnName,
		index:          index,
		objType:        objType,
		collationType:  collType,
		refColumnNames: []string{columnName},
		isGenColumn:    false,
		columnExpress:  nil,
	}
}

type ObColumn struct {
	columnName     string
	index          int
	objType        ObObjType
	collationType  ObCollationType
	refColumnNames []string
	isGenColumn    bool
	columnExpress  ObGeneratedColumnSimpleFunc
}

func (c *ObColumn) ColumnName() string {
	return c.columnName
}

func (c *ObColumn) ToString() string {
	var isGenColumnStr string
	if c.isGenColumn {
		isGenColumnStr = "true"
	} else {
		isGenColumnStr = "false"
	}

	var columnExpressStr string
	if c.isGenColumn {
		columnExpressStr = c.columnExpress.ToString()
	} else {
		columnExpressStr = "nil"
	}
	return "ObColumn{" +
		"columnName:" + c.columnName + ", " +
		"index:" + strconv.Itoa(c.index) + ", " +
		"objType:" + c.objType.ToString() + ", " +
		"collationType:" + c.collationType.ToString() + ", " +
		"refColumnNames:" + "[" + strings.Join(c.refColumnNames, ",") + "]" + ", " +
		"isGenColumn:" + isGenColumnStr + ", " +
		"columnExpress:" + columnExpressStr +
		"}"
}

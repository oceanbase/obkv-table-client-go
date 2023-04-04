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
	// isGenColumn to string
	var isGenColumnStr string
	if c.isGenColumn {
		isGenColumnStr = "true"
	} else {
		isGenColumnStr = "false"
	}

	// objType to string
	var objTypeStr string
	if c.objType != nil {
		objTypeStr = c.objType.ToString()
	} else {
		objTypeStr = "nil"
	}

	// columnExpress to string
	var columnExpressStr string
	if c.isGenColumn {
		columnExpressStr = c.columnExpress.ToString()
	} else {
		columnExpressStr = "nil"
	}
	return "ObColumn{" +
		"columnName:" + c.columnName + ", " +
		"index:" + strconv.Itoa(c.index) + ", " +
		"objType:" + objTypeStr + ", " +
		"collationType:" + c.collationType.ToString() + ", " +
		"refColumnNames:" + "[" + strings.Join(c.refColumnNames, ",") + "]" + ", " +
		"isGenColumn:" + isGenColumnStr + ", " +
		"columnExpress:" + columnExpressStr +
		"}"
}

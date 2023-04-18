package protocol

import (
	"errors"
	"github.com/oceanbase/obkv-table-client-go/log"
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
	columnName    string
	index         int
	objType       ObObjType
	collationType ObCollationType
	// refColumnNames: Represents which columns are referenced by the current column
	// 1. generate column: key_prefix VARCHAR(1024) GENERATED ALWAYS AS (SUBSTRING(`col1`, 1, 4)
	// 					   refColumnNames = ["col1"]
	// 2. normal column: col_normal int
	//					   refColumnNames = ["col_normal"]
	refColumnNames []string
	isGenColumn    bool
	columnExpress  ObGeneratedColumnSimpleFunc // only support 'SUBSTRING' expr now
}

func (c *ObColumn) RefColumnNames() []string {
	return c.refColumnNames
}

func (c *ObColumn) ColumnName() string {
	return c.columnName
}

func (c *ObColumn) EvalValue(refs ...interface{}) (interface{}, error) {
	if !c.isGenColumn {
		if len(refs) == 0 || len(refs) > 1 {
			log.Warn("simple column is refer to itself so that the length of the refs must be 1",
				log.Int("refs len", len(refs)))
			return nil, errors.New("simple column is refer to itself so that the length of the refs must be 1")
		}
		return c.objType.parseToComparable(refs[0], c.collationType)
	} else {
		if len(refs) != len(c.refColumnNames) {
			log.Warn("input refs count is not equal to refColumnNames count when column is generate column",
				log.Int("refs len", len(refs)), log.Int("refColumnNames len", len(c.refColumnNames)))
			return nil, errors.New("input refs count is not equal to refColumnNames count when column is generate column")
		}
		// todo: impl generate column
		return nil, errors.New("not support generate column now")
	}
}

func (c *ObColumn) String() string {
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
		objTypeStr = c.objType.String()
	} else {
		objTypeStr = "nil"
	}

	// columnExpress to string
	var columnExpressStr string
	if c.isGenColumn {
		columnExpressStr = c.columnExpress.String()
	} else {
		columnExpressStr = "nil"
	}
	return "ObColumn{" +
		"columnName:" + c.columnName + ", " +
		"index:" + strconv.Itoa(c.index) + ", " +
		"objType:" + objTypeStr + ", " +
		"collationType:" + c.collationType.String() + ", " +
		"refColumnNames:" + "[" + strings.Join(c.refColumnNames, ",") + "]" + ", " +
		"isGenColumn:" + isGenColumnStr + ", " +
		"columnExpress:" + columnExpressStr +
		"}"
}

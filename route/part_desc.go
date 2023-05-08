package route

import (
	"errors"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"strconv"
	"strings"
)

const (
	ObInvalidPartId = -1
	ObPartIdBitNum  = 28
	ObPartIdShift   = 32
	ObMask          = (1 << ObPartIdBitNum) | 1<<(ObPartIdBitNum+ObPartIdShift)
	ObSubPartIdMask = 0xffffffff & 0xfffffff
)

type ObColumnIndexesPair struct {
	column  *protocol.ObColumn
	indexes []int
}

func NewObColumnIndexesPair(column *protocol.ObColumn, indexes []int) *ObColumnIndexesPair {
	return &ObColumnIndexesPair{column, indexes}
}

func (p *ObColumnIndexesPair) String() string {
	columnStr := "nil"
	if p.column != nil {
		columnStr = p.column.String()
	}
	var indexesStr string
	indexesStr = indexesStr + "["
	for i := 0; i < len(p.indexes); i++ {
		if i > 0 {
			indexesStr += ", "
		}
		indexesStr += strconv.Itoa(p.indexes[i])
	}
	indexesStr += "]"
	return "ObColumnIndexesPair{" +
		"column:" + columnStr + ", " +
		"indexes:" + indexesStr +
		"}"
}

type ObPartDescCommon struct {
	PartFuncType ObPartFuncType
	PartExpr     string
	// orderedPartColumnNames Represents all partitioned column names
	// eg:
	//    partition by range(c1, c2)
	//    orderedPartColumnNames = ["c1", "c2"]
	OrderedPartColumnNames              []string
	OrderedPartRefColumnRowKeyRelations []*ObColumnIndexesPair
	PartColumns                         []*protocol.ObColumn
	RowKeyElement                       *table.ObRowkeyElement
}

func (c *ObPartDescCommon) setCommRowKeyElement(rowKeyElement *table.ObRowkeyElement) {
	c.RowKeyElement = rowKeyElement
	if len(c.OrderedPartColumnNames) != 0 && len(c.PartColumns) != 0 {
		relations := make([]*ObColumnIndexesPair, 0, len(c.OrderedPartColumnNames))
		for _, partOrderColumnName := range c.OrderedPartColumnNames {
			for _, col := range c.PartColumns {
				if strings.EqualFold(col.ColumnName(), partOrderColumnName) {
					partRefColumnRowKeyIndexes := make([]int, 0, len(col.RefColumnNames()))
					for _, refColumnName := range col.RefColumnNames() {
						for rowKeyElementName, index := range c.RowKeyElement.NameIdxMap() {
							if strings.EqualFold(rowKeyElementName, refColumnName) {
								partRefColumnRowKeyIndexes = append(partRefColumnRowKeyIndexes, index)
							}
						}
					}
					pair := NewObColumnIndexesPair(col, partRefColumnRowKeyIndexes)
					relations = append(relations, pair)
				}
			}
		}
		c.OrderedPartRefColumnRowKeyRelations = relations
	}
}

func (c *ObPartDescCommon) CommString() string {
	// orderedPartRefColumnRowKeyRelations to string
	var relationsStr string
	relationsStr = relationsStr + "["
	for i := 0; i < len(c.OrderedPartRefColumnRowKeyRelations); i++ {
		if i > 0 {
			relationsStr += ", "
		}
		relationsStr += c.OrderedPartRefColumnRowKeyRelations[i].String()
	}
	relationsStr += "]"

	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(c.PartColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += c.PartColumns[i].String()
	}
	partColumnsStr += "]"

	// rowKeyElement to string
	rowKeyElementStr := "nil"
	if c.RowKeyElement != nil {
		rowKeyElementStr = c.RowKeyElement.String()
	}

	return "ObPartDescCommon{" +
		"partFuncType:" + c.PartFuncType.String() + ", " +
		"partExpr:" + c.PartExpr + ", " +
		"orderedPartColumnNames:" + strings.Join(c.OrderedPartColumnNames, ",") + ", " +
		"orderedPartRefColumnRowKeyRelations:" + relationsStr + ", " +
		"partColumns:" + partColumnsStr + ", " +
		"rowKeyElement:" + rowKeyElementStr +
		"}"
}

type ObPartDesc interface {
	String() string
	partFuncType() ObPartFuncType
	orderedPartColumnNames() []string
	setOrderedPartColumnNames(partExpr string)
	orderedPartRefColumnRowKeyRelations() []*ObColumnIndexesPair
	rowKeyElement() *table.ObRowkeyElement
	setRowKeyElement(rowKeyElement *table.ObRowkeyElement)
	setPartColumns(partColumns []*protocol.ObColumn)
	GetPartId(rowkey []interface{}) (int64, error)
}

func evalPartKeyValues(desc ObPartDesc, rowkey []interface{}) ([]interface{}, error) {
	if desc == nil {
		log.Warn("part desc is nil")
		return nil, errors.New("part desc is nil")
	}
	if desc.rowKeyElement() == nil {
		log.Warn("rowkey element is nil")
		return nil, errors.New("rowkey element is nil")
	}
	if len(rowkey) < len(desc.rowKeyElement().NameIdxMap()) {
		log.Warn("rowkey count not match",
			log.Int("rowkey len", len(rowkey)),
			log.Int("rowKeyElement len", len(desc.rowKeyElement().NameIdxMap())))
		return nil, errors.New("rowkey count not match")
	}
	partRefColumnSize := len(desc.orderedPartRefColumnRowKeyRelations())
	evalValues := make([]interface{}, 0, partRefColumnSize)
	for i := 0; i < partRefColumnSize; i++ {
		relation := desc.orderedPartRefColumnRowKeyRelations()[i]
		evalParams := make([]interface{}, len(relation.indexes))
		for i := 0; i < len(relation.indexes); i++ {
			evalParams[i] = rowkey[relation.indexes[i]]
		}
		val, err := relation.column.EvalValue(evalParams...)
		if err != nil {
			log.Warn("fail to eval column value", log.String("relations", relation.String()))
			return nil, err
		}
		evalValues = append(evalValues, val)
	}
	return evalValues, nil
}

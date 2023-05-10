package route

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type ObRangePartDesc struct {
	ObPartDescCommon
	partSpace                 int
	partNum                   int
	orderedCompareColumns     []*ObColumn
	orderedCompareColumnTypes []protocol.ObObjType
	// todo: List<ObComparableKV<ObPartitionKey, Long>> bounds
}

func newObRangePartDesc() *ObRangePartDesc {
	return &ObRangePartDesc{}
}

func (d *ObRangePartDesc) partFuncType() ObPartFuncType {
	return d.PartFuncType
}

func (d *ObRangePartDesc) orderedPartColumnNames() []string {
	return d.OrderedPartColumnNames
}

func (d *ObRangePartDesc) setOrderedPartColumnNames(partExpr string) {
	// eg:"c1, c2", need to remove ' '
	str := strings.ReplaceAll(partExpr, " ", "")
	d.OrderedPartColumnNames = strings.Split(str, ",")
}

func (d *ObRangePartDesc) orderedPartRefColumnRowKeyRelations() []*ObColumnIndexesPair {
	return d.OrderedPartRefColumnRowKeyRelations
}
func (d *ObRangePartDesc) rowKeyElement() *table.ObRowKeyElement {
	return d.RowKeyElement
}

func (d *ObRangePartDesc) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	d.setCommRowKeyElement(rowKeyElement)
}

func (d *ObRangePartDesc) setPartColumns(partColumns []*ObColumn) {
	d.PartColumns = partColumns
}

func (d *ObRangePartDesc) GetPartId(rowKey []interface{}) (int64, error) {
	// todo: impl
	return ObInvalidPartId, errors.New("not implement")
}

// func (d *ObRangePartDesc) setOrderedCompareColumns(orderedPartColumn []protocol.ObColumn) {
//	d.orderedCompareColumns = orderedPartColumn
// }

func (d *ObRangePartDesc) String() string {
	// orderedCompareColumns to string
	var orderedCompareColumnsStr string
	orderedCompareColumnsStr = orderedCompareColumnsStr + "["
	for i := 0; i < len(d.orderedCompareColumns); i++ {
		if i > 0 {
			orderedCompareColumnsStr += ", "
		}
		orderedCompareColumnsStr += d.orderedCompareColumns[i].String()
	}
	orderedCompareColumnsStr += "]"

	// orderedCompareColumnTypes to string
	var orderedCompareColumnTypesStr string
	orderedCompareColumnTypesStr = orderedCompareColumnTypesStr + "["
	for i := 0; i < len(d.orderedCompareColumns); i++ {
		if i > 0 {
			orderedCompareColumnTypesStr += ", "
		}
		orderedCompareColumnTypesStr += d.orderedCompareColumnTypes[i].String()
	}
	orderedCompareColumnTypesStr += "]"

	return "ObRangePartDesc{" +
		"comm:" + d.CommString() + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"orderedCompareColumns:" + orderedCompareColumnsStr + ", " +
		"orderedCompareColumnTypes:" + orderedCompareColumnTypesStr +
		"}"
}

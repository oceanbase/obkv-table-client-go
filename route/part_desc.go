package route

import (
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"strconv"
	"strings"
)

type ObColumnIndexesPair struct {
	column  protocol.ObColumn
	indexes []int
}

func (p *ObColumnIndexesPair) ToString() string {
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
		"column:" + p.column.ToString() + ", " +
		"indexes:" + indexesStr +
		"}"
}

type ObPartDescCommon struct {
	partFuncType                        ObPartFuncType
	partExpr                            string
	orderedPartColumnNames              []string
	orderedPartRefColumnRowKeyRelations []ObColumnIndexesPair
	partColumns                         []protocol.ObColumn
	rowKeyElement                       map[string]int
}

func (c *ObPartDescCommon) ToString() string {
	// orderedPartRefColumnRowKeyRelations to string
	var relationsStr string
	relationsStr = relationsStr + "["
	for i := 0; i < len(c.orderedPartRefColumnRowKeyRelations); i++ {
		if i > 0 {
			relationsStr += ", "
		}
		relationsStr += c.orderedPartRefColumnRowKeyRelations[i].ToString()
	}
	relationsStr += "]"

	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(c.partColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += c.partColumns[i].ToString()
	}
	partColumnsStr += "]"

	// rowKeyElement to string
	var rowKeyElementStr string
	rowKeyElementStr = rowKeyElementStr + "{"
	for k, v := range c.rowKeyElement {
		rowKeyElementStr += "m[" + k + "]=" + strconv.Itoa(v) + ", "
	}
	rowKeyElementStr += "}"

	return "ObPartDescCommon{" +
		"partFuncType:" + c.partFuncType.ToString() + ", " +
		"partExpr:" + c.partExpr + ", " +
		"orderedPartColumnNames:" + strings.Join(c.orderedPartColumnNames, ",") + ", " +
		"orderedPartRefColumnRowKeyRelations:" + relationsStr + ", " +
		"partColumns:" + partColumnsStr + ", " +
		"rowKeyElement:" + rowKeyElementStr +
		"}"
}

type ObPartDesc interface {
	ToString() string
	getPartFuncType() ObPartFuncType
	getOrderedPartColumnNames() []string
	setPartColumns(partColumns []protocol.ObColumn)
}

type ObRangePartDesc struct {
	comm                      ObPartDescCommon
	orderedCompareColumns     []protocol.ObColumn
	orderedCompareColumnTypes []protocol.ObObjType
	//todo: List<ObComparableKV<ObPartitionKey, Long>> bounds
}

func (d ObRangePartDesc) getPartFuncType() ObPartFuncType {
	return d.comm.partFuncType
}

func (d ObRangePartDesc) getOrderedPartColumnNames() []string {
	return d.comm.orderedPartColumnNames
}

func (d ObRangePartDesc) setPartColumns(partColumns []protocol.ObColumn) {
	d.comm.partColumns = partColumns
}

func (d *ObRangePartDesc) setOrderedCompareColumns(orderedPartColumn []protocol.ObColumn) {
	d.orderedCompareColumns = orderedPartColumn
}

func (d ObRangePartDesc) ToString() string {
	// orderedCompareColumns to string
	var orderedCompareColumnsStr string
	orderedCompareColumnsStr = orderedCompareColumnsStr + "["
	for i := 0; i < len(d.orderedCompareColumns); i++ {
		if i > 0 {
			orderedCompareColumnsStr += ", "
		}
		orderedCompareColumnsStr += d.orderedCompareColumns[i].ToString()
	}
	orderedCompareColumnsStr += "]"

	// orderedCompareColumnTypes to string
	var orderedCompareColumnTypesStr string
	orderedCompareColumnTypesStr = orderedCompareColumnTypesStr + "["
	for i := 0; i < len(d.orderedCompareColumns); i++ {
		if i > 0 {
			orderedCompareColumnTypesStr += ", "
		}
		orderedCompareColumnTypesStr += d.orderedCompareColumnTypes[i].ToString()
	}
	orderedCompareColumnTypesStr += "]"

	return "ObRangePartDesc{" +
		"comm:" + d.comm.ToString() + ", " +
		"orderedCompareColumns:" + orderedCompareColumnsStr + ", " +
		"orderedCompareColumnTypes:" + orderedCompareColumnTypesStr +
		"}"
}

type ObHashPartDesc struct {
	comm          ObPartDescCommon
	completeWorks []int64
	partSpace     int
	partNum       int
	partNameIdMap map[string]int64
}

func (d ObHashPartDesc) getPartFuncType() ObPartFuncType {
	return d.comm.partFuncType
}

func (d ObHashPartDesc) getOrderedPartColumnNames() []string {
	return d.comm.orderedPartColumnNames
}

func (d ObHashPartDesc) setPartColumns(partColumns []protocol.ObColumn) {
	d.comm.partColumns = partColumns
}

func (d ObHashPartDesc) ToString() string {
	// completeWorks to string
	var completeWorksStr string
	completeWorksStr = completeWorksStr + "["
	for i := 0; i < len(d.completeWorks); i++ {
		if i > 0 {
			completeWorksStr += ", "
		}
		completeWorksStr += strconv.Itoa(int(d.completeWorks[i]))
	}
	completeWorksStr += "]"

	// partNameIdMap to string
	var partNameIdMapStr string
	partNameIdMapStr = partNameIdMapStr + "{"
	for k, v := range d.partNameIdMap {
		partNameIdMapStr += "m[" + k + "]=" + strconv.Itoa(int(v)) + ", "
	}
	partNameIdMapStr += "}"

	return "ObHashPartDesc{" +
		"comm:" + d.comm.ToString() + ", " +
		"completeWorks:" + completeWorksStr + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

type ObKeyPartDesc struct {
	comm          ObPartDescCommon
	partSpace     int
	partNum       int
	partNameIdMap map[string]int64
}

func (d ObKeyPartDesc) getPartFuncType() ObPartFuncType {
	return d.comm.partFuncType
}

func (d ObKeyPartDesc) getOrderedPartColumnNames() []string {
	return d.comm.orderedPartColumnNames
}

func (d ObKeyPartDesc) setPartColumns(partColumns []protocol.ObColumn) {
	d.comm.partColumns = partColumns
}

func (d ObKeyPartDesc) ToString() string {
	// partNameIdMap to string
	var partNameIdMapStr string
	partNameIdMapStr = partNameIdMapStr + "{"
	for k, v := range d.partNameIdMap {
		partNameIdMapStr += "m[" + k + "]=" + strconv.Itoa(int(v)) + ", "
	}
	partNameIdMapStr += "}"
	return "ObKeyPartDesc{" +
		"comm:" + d.comm.ToString() + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

const (
	partLevelUnknown = "partLevelUnknown"
	partLevelZero    = "partLevelZero"
	partLevelOne     = "partLevelOne"
	partLevelTwo     = "partLevelTwo"
)

const (
	partLevelUnknownIndex = -1
	partLevelZeroIndex    = 0
	partLevelOneIndex     = 1
	partLevelTwoIndex     = 2
)

type ObPartitionLevel struct {
	name  string
	index int
}

func newObPartitionLevel(index int) ObPartitionLevel {
	if index == partLevelZeroIndex {
		return ObPartitionLevel{partLevelZero, partLevelZeroIndex}
	} else if index == partLevelOneIndex {
		return ObPartitionLevel{partLevelOne, partLevelOneIndex}
	} else if index == partLevelTwoIndex {
		return ObPartitionLevel{partLevelTwo, partLevelTwoIndex}
	} else {
		return ObPartitionLevel{partLevelUnknown, partLevelUnknownIndex}
	}
}

func (l *ObPartitionLevel) ToString() string {
	return "ObPartitionLevel{" +
		"name:" + l.name + ", " +
		"index:" + strconv.Itoa(l.index) +
		"}"
}

const (
	partFuncTypeUnknownIndex  = -1
	partFuncTypeHashIndex     = 0
	partFuncTypeKeyIndex      = 1
	partFuncTypeKeyImplIndex  = 2
	partFuncTypeRangeIndex    = 3
	partFuncTypeRangeColIndex = 4
	partFuncTypeListIndex     = 5
	partFuncTypeKeyV2Index    = 6
	partFuncTypeListColIndex  = 7
	partFuncTypeHashV2Index   = 8
	partFuncTypeKeyV3Index    = 9
)

const (
	partFuncTypeUnknown  = "UNKNOWN"
	partFuncTypeHash     = "HASH"
	partFuncTypeKey      = "KEY"
	partFuncTypeKeyImpl  = "KEY_IMPLICIT"
	partFuncTypeRange    = "RANGE"
	partFuncTypeRangeCol = "RANGE_COLUMNS"
	partFuncTypeList     = "LIST"
	partFuncTypeKeyV2    = "KEY_V2"
	partFuncTypeListCol  = "LIST_COLUMNS"
	partFuncTypeHashV2   = "HASH_V2"
	partFuncTypeKeyV3    = "KEY_V3"
)

type ObPartFuncType struct {
	name  string
	index int
}

func (t ObPartFuncType) ToString() string {
	return "ObPartFuncType{" +
		"name:" + t.name + ", " +
		"index:" + strconv.Itoa(t.index) +
		"}"
}

func (t ObPartFuncType) isRangePart() bool {
	return t.index == partFuncTypeRangeIndex || t.index == partFuncTypeRangeColIndex
}

func (t ObPartFuncType) isKeyPart() bool {
	return t.index == partFuncTypeKeyImplIndex ||
		t.index == partFuncTypeKeyV2Index ||
		t.index == partFuncTypeKeyV3Index
}

func (t ObPartFuncType) isHashPart() bool {
	return t.index == partFuncTypeHashIndex || t.index == partFuncTypeHashV2Index
}

func (t ObPartFuncType) isListPart() bool {
	return t.index == partFuncTypeListIndex || t.index == partFuncTypeListColIndex
}

func newObPartFuncType(index int) ObPartFuncType {
	switch index {
	case partFuncTypeHashIndex:
		return ObPartFuncType{partFuncTypeHash, partFuncTypeHashIndex}
	case partFuncTypeKeyIndex:
		return ObPartFuncType{partFuncTypeKey, partFuncTypeKeyIndex}
	case partFuncTypeKeyImplIndex:
		return ObPartFuncType{partFuncTypeKeyImpl, partFuncTypeKeyImplIndex}
	case partFuncTypeRangeIndex:
		return ObPartFuncType{partFuncTypeRange, partFuncTypeRangeIndex}
	case partFuncTypeRangeColIndex:
		return ObPartFuncType{partFuncTypeRangeCol, partFuncTypeRangeColIndex}
	case partFuncTypeListIndex:
		return ObPartFuncType{partFuncTypeList, partFuncTypeListIndex}
	case partFuncTypeKeyV2Index:
		return ObPartFuncType{partFuncTypeKeyV2, partFuncTypeKeyV2Index}
	case partFuncTypeListColIndex:
		return ObPartFuncType{partFuncTypeListCol, partFuncTypeListColIndex}
	case partFuncTypeHashV2Index:
		return ObPartFuncType{partFuncTypeHashV2, partFuncTypeHashV2Index}
	case partFuncTypeKeyV3Index:
		return ObPartFuncType{partFuncTypeKeyV3, partFuncTypeKeyV3Index}
	default:
		return ObPartFuncType{partFuncTypeUnknown, partFuncTypeUnknownIndex}
	}
}

package route

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type ObHashPartDesc struct {
	ObPartDescCommon
	completeWorks []int64
	partSpace     int
	partNum       int
	partNameIdMap map[string]int64
}

func newObHashPartDesc() *ObHashPartDesc {
	return &ObHashPartDesc{}
}

func (d *ObHashPartDesc) partFuncType() ObPartFuncType {
	return d.PartFuncType
}

func (d *ObHashPartDesc) orderedPartColumnNames() []string {
	return d.OrderedPartColumnNames
}

func (d *ObHashPartDesc) setOrderedPartColumnNames(partExpr string) {
	// eg:"c1, c2", need to remove ' '
	str := strings.ReplaceAll(partExpr, " ", "")
	d.OrderedPartColumnNames = strings.Split(str, ",")
}

func (d *ObHashPartDesc) orderedPartRefColumnRowKeyRelations() []*ObColumnIndexesPair {
	return d.OrderedPartRefColumnRowKeyRelations
}
func (d *ObHashPartDesc) rowKeyElement() *table.ObRowKeyElement {
	return d.RowKeyElement
}

func (d *ObHashPartDesc) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	d.setCommRowKeyElement(rowKeyElement)
}

func (d *ObHashPartDesc) setPartColumns(partColumns []*ObColumn) {
	d.PartColumns = partColumns
}

func (d *ObHashPartDesc) GetPartId(rowKey []interface{}) (int64, error) {
	if len(rowKey) == 0 {
		log.Warn("rowKey size is 0")
		return ObInvalidPartId, errors.New("rowKey size is 0")
	}
	evalValues, err := evalPartKeyValues(d, rowKey)
	if err != nil {
		log.Warn("failed to eval part key values", log.String("part desc", d.String()))
		return ObInvalidPartId, err
	}
	longValue, err := protocol.ParseToLong(evalValues[0]) // hash part has one param at most
	if err != nil {
		log.Warn("failed to parse to long", log.String("part desc", d.String()))
		return ObInvalidPartId, err
	}
	if v, ok := longValue.(int64); !ok {
		log.Warn("failed to convert to long")
		return ObInvalidPartId, errors.New("failed to convert to long")
	} else {
		return d.innerHash(v), nil
	}
}

func (d *ObHashPartDesc) innerHash(hashVal int64) int64 {
	// abs(hashVal)
	if hashVal < 0 {
		hashVal = -hashVal
	}
	return (int64(d.partSpace) << ObPartIdBitNum) | (hashVal % int64(d.partNum))
}

func (d *ObHashPartDesc) String() string {
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
	var i = 0
	for k, v := range d.partNameIdMap {
		if i > 0 {
			partNameIdMapStr += ", "
		}
		i++
		partNameIdMapStr += "m[" + k + "]=" + strconv.Itoa(int(v))
	}
	partNameIdMapStr += "}"

	return "ObHashPartDesc{" +
		"comm:" + d.CommString() + ", " +
		"completeWorks:" + completeWorksStr + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

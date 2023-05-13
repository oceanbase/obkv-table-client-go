package route

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type obHashPartDesc struct {
	obPartDescCommon
	completeWorks []int64
	partSpace     int
	partNum       int
	partNameIdMap map[string]int64
}

func newObHashPartDesc() *obHashPartDesc {
	return &obHashPartDesc{}
}

func (d *obHashPartDesc) partFuncType() obPartFuncType {
	return d.PartFuncType
}

func (d *obHashPartDesc) orderedPartColumnNames() []string {
	return d.OrderedPartColumnNames
}

func (d *obHashPartDesc) setOrderedPartColumnNames(partExpr string) {
	// eg:"c1, c2", need to remove ' '
	str := strings.ReplaceAll(partExpr, " ", "")
	d.OrderedPartColumnNames = strings.Split(str, ",")
}

func (d *obHashPartDesc) orderedPartRefColumnRowKeyRelations() []*obColumnIndexesPair {
	return d.OrderedPartRefColumnRowKeyRelations
}
func (d *obHashPartDesc) rowKeyElement() *table.ObRowKeyElement {
	return d.RowKeyElement
}

func (d *obHashPartDesc) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	d.setCommRowKeyElement(rowKeyElement)
}

func (d *obHashPartDesc) setPartColumns(partColumns []*obColumn) {
	d.PartColumns = partColumns
}

func (d *obHashPartDesc) GetPartId(rowKey []interface{}) (int64, error) {
	if len(rowKey) == 0 {
		return ObInvalidPartId, errors.New("rowKey size is 0")
	}
	evalValues, err := evalPartKeyValues(d, rowKey)
	if err != nil {
		return ObInvalidPartId, errors.WithMessagef(err, "eval partition key value, partDesc:%s", d.String())
	}
	longValue, err := protocol.ParseToLong(evalValues[0]) // hash part has one param at most
	if err != nil {
		return ObInvalidPartId, errors.WithMessagef(err, "parse long, partDesc:%s", d.String())
	}
	if v, ok := longValue.(int64); !ok {
		return ObInvalidPartId, errors.Errorf("failed to convert to long, value:%T", longValue)
	} else {
		return d.innerHash(v), nil
	}
}

func (d *obHashPartDesc) innerHash(hashVal int64) int64 {
	// abs(hashVal)
	if hashVal < 0 {
		hashVal = -hashVal
	}
	return (int64(d.partSpace) << ObPartIdBitNum) | (hashVal % int64(d.partNum))
}

func (d *obHashPartDesc) String() string {
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

	return "obHashPartDesc{" +
		"comm:" + d.CommString() + ", " +
		"completeWorks:" + completeWorksStr + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

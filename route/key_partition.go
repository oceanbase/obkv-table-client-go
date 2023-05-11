package route

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type obKeyPartDesc struct {
	obPartDescCommon
	partSpace     int
	partNum       int
	partNameIdMap map[string]int64
}

func newObKeyPartDesc() *obKeyPartDesc {
	return &obKeyPartDesc{}
}

func (d *obKeyPartDesc) partFuncType() obPartFuncType {
	return d.PartFuncType
}

func (d *obKeyPartDesc) orderedPartColumnNames() []string {
	return d.OrderedPartColumnNames
}

func (d *obKeyPartDesc) setOrderedPartColumnNames(partExpr string) {
	// eg:"c1, c2", need to remove ' '
	str := strings.ReplaceAll(partExpr, " ", "")
	d.OrderedPartColumnNames = strings.Split(str, ",")
}

func (d *obKeyPartDesc) orderedPartRefColumnRowKeyRelations() []*obColumnIndexesPair {
	return d.OrderedPartRefColumnRowKeyRelations
}

func (d *obKeyPartDesc) rowKeyElement() *table.ObRowKeyElement {
	return d.RowKeyElement
}

func (d *obKeyPartDesc) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	d.setCommRowKeyElement(rowKeyElement)
}

func (d *obKeyPartDesc) setPartColumns(partColumns []*obColumn) {
	d.PartColumns = partColumns
}

func (d *obKeyPartDesc) GetPartId(rowKey []interface{}) (int64, error) {
	if len(rowKey) == 0 {
		log.Warn("rowKey size is 0")
		return ObInvalidPartId, errors.New("rowKey size is 0")
	}
	evalValues, err := evalPartKeyValues(d, rowKey)
	if err != nil {
		log.Warn("failed to eval part key values", log.String("part desc", d.String()))
		return ObInvalidPartId, err
	}
	if len(evalValues) < len(d.OrderedPartRefColumnRowKeyRelations) {
		log.Warn("invalid eval values length",
			log.Int("evalValues length", len(evalValues)),
			log.Int("OrderedPartRefColumnRowKeyRelations length", len(d.OrderedPartRefColumnRowKeyRelations)))
	}
	var hashValue int64
	for i := 0; i < len(d.OrderedPartRefColumnRowKeyRelations); i++ {
		hashValue, err = d.toHashCode(
			evalValues[i],
			d.OrderedPartRefColumnRowKeyRelations[i].column,
			hashValue,
			d.PartFuncType,
		)
		if err != nil {
			log.Warn("failed to convert to hash code", log.String("part desc", d.String()))
			return ObInvalidPartId, err
		}
	}
	if hashValue < 0 {
		hashValue = -hashValue
	}
	return (int64(d.partSpace) << ObPartIdBitNum) | (hashValue % int64(d.partNum)), nil
}

func intToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case bool:
		if v {
			return int64(1), nil
		} else {
			return int64(0), nil
		}
	case int8:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case int:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		log.Warn("invalid type to convert to int64", log.String("value", util.InterfaceToString(value)))
		return -1, errors.New("invalid type to convert to int64")
	}
}

func (d *obKeyPartDesc) toHashCode(
	value interface{},
	refColumn *obColumn,
	hashCode int64,
	partFuncType obPartFuncType) (int64, error) {
	typeValue := refColumn.objType.GetValue()
	if typeValue >= protocol.ObTinyIntTypeValue && typeValue <= protocol.ObUInt64TypeValue {
		i64, err := intToInt64(value)
		if err != nil {
			log.Warn("failed to convert int to int64", log.Int("type", typeValue))
			return -1, err
		}
		arr := d.longToByteArray(i64)
		return murmurHash64A(arr, len(arr), hashCode), nil
	} else if typeValue == protocol.ObDateTimeTypeValue || typeValue == protocol.ObTimestampTypeValue {
		t, ok := value.(time.Time)
		if !ok {
			log.Warn("invalid timestamp type", log.String("value", util.InterfaceToString(value)))
			return -1, errors.New("invalid timestamp type")
		}
		return d.timeStampHash(t, hashCode), nil
	} else if typeValue == protocol.ObDateTypeValue {
		date, ok := value.(time.Time)
		if !ok {
			log.Warn("invalid date type", log.String("value", util.InterfaceToString(value)))
			return -1, errors.New("invalid date type")
		}
		return d.dateHash(date, hashCode), nil
	} else if typeValue == protocol.ObVarcharTypeValue || typeValue == protocol.ObCharTypeValue {
		return d.varcharHash(value, refColumn.collationType, hashCode, partFuncType)
	} else {
		log.Warn("unsupported type for key hash", log.String("objType", refColumn.objType.String()))
		return -1, errors.New("unsupported type for key hash")
	}
}

func (d *obKeyPartDesc) longToByteArray(l int64) []byte {
	return []byte{(byte)(l & 0xFF), (byte)((l >> 8) & 0xFF), (byte)((l >> 16) & 0xFF),
		(byte)((l >> 24) & 0xFF), (byte)((l >> 32) & 0xFF), (byte)((l >> 40) & 0xFF),
		(byte)((l >> 48) & 0xFF), (byte)((l >> 56) & 0xFF)}
}

func (d *obKeyPartDesc) longHash(value int64, hashCode int64) int64 {
	arr := d.longToByteArray(value)
	return murmurHash64A(arr, len(arr), hashCode)
}

func (d *obKeyPartDesc) timeStampHash(ts time.Time, hashCode int64) int64 {
	return d.longHash(ts.UnixMilli(), hashCode)
}

func (d *obKeyPartDesc) dateHash(ts time.Time, hashCode int64) int64 {
	return d.longHash(ts.UnixMilli(), hashCode)
}

func (d *obKeyPartDesc) varcharHash(
	value interface{},
	collType protocol.ObCollationType,
	hashCode int64,
	partFuncType obPartFuncType) (int64, error) {
	var seed uint64 = 0xc6a4a7935bd1e995
	var bytes []byte
	if v, ok := value.(string); ok {
		// Right Now, only UTF8 String is supported, aligned with the Serialization.
		// string and []byte is utf8 default in go language
		bytes = []byte(v)
	} else if v, ok := value.([]byte); ok {
		bytes = v
	} else if v, ok := value.(protocol.ObBytesString); ok {
		bytes = v.BytesVal()
	} else {
		log.Warn("invalid varchar", log.String("value", util.InterfaceToString(value)))
		return -1, errors.New("invalid varchar value for calc hash value")
	}
	switch collType.Value() {
	case protocol.CsTypeUtf8mb4GeneralCi:
		if partFuncType == partFuncTypeKeyV3 ||
			partFuncType == partFuncTypeKeyImplV2 ||
			util.ObVersion() >= 4 {
			hashCode = hashSortUtf8Mb4(bytes, hashCode, seed, true)
		} else {
			hashCode = hashSortUtf8Mb4(bytes, hashCode, seed, false)
		}
	case protocol.CsTypeUtf8mb4Bin:
	case protocol.CsTypeBinary:
		if partFuncType == partFuncTypeKeyV3 ||
			partFuncType == partFuncTypeKeyImplV2 ||
			util.ObVersion() >= 4 {
			hashCode = murmurHash64A(bytes, len(bytes), hashCode)
		} else {
			hashCode = hashSortMbBin(bytes, hashCode, seed)
		}
	case protocol.CsTypeInvalid:
	case protocol.CsTypeCollationFree:
	case protocol.CsTypeMax:
		log.Warn("not supported collation type", log.Int("coll type", collType.Value()))
		return -1, errors.New("not supported collation type")
	}
	return hashCode, nil
}

func (d *obKeyPartDesc) String() string {
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
	return "obKeyPartDesc{" +
		"comm:" + d.CommString() + ", " +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

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
	"time"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/oceanbase/obkv-table-client-go/util"
)

// newObHashPartDesc create a key partition description.
func newObKeyPartDesc(
	partSpace int,
	partNum int,
	partFuncType obPartFuncType) *obKeyPartDesc {
	return &obKeyPartDesc{
		partFuncType: partFuncType,
		partSpace:    partSpace,
		partNum:      partNum,
	}
}

// obKeyPartDesc description of the key partition.
type obKeyPartDesc struct {
	partFuncType obPartFuncType
	partSpace    int
	partNum      int
	partColumns  []obColumn
}

func (d *obKeyPartDesc) PartColumns() []obColumn {
	return d.partColumns
}

func (d *obKeyPartDesc) PartNum() int {
	return d.partNum
}

func (d *obKeyPartDesc) SetPartNum(partNum int) {
	d.partNum = partNum
}

func (d *obKeyPartDesc) PartFuncType() obPartFuncType {
	return d.partFuncType
}

func (d *obKeyPartDesc) SetPartColumns(partColumns []obColumn) {
	d.partColumns = partColumns
}

// GetPartId get partition id.
func (d *obKeyPartDesc) GetPartId(rowKey []*table.Column) (uint64, error) {
	if len(rowKey) == 0 {
		return ObInvalidPartId, errors.New("rowKey size is 0")
	}
	evalValues, err := evalPartKeyValues(d, rowKey)
	if err != nil {
		return ObInvalidPartId, errors.WithMessagef(err, "eval part key value, part desc:%s", d.String())
	}
	var hashValue int64
	for i := 0; i < len(d.partColumns); i++ {
		hashValue, err = d.toHashCode(
			evalValues[i],
			d.partColumns[i],
			hashValue,
			d.partFuncType,
		)
		if err != nil {
			return ObInvalidPartId, errors.WithMessagef(err, "convert to hash code, part desc:%s", d.String())
		}
	}
	if hashValue < 0 {
		hashValue = -hashValue
	}
	return uint64((int64(d.partSpace) << ObPartIdBitNum) | (hashValue % int64(d.partNum))), nil
}

// GetPartIds get partition ids.
func (d *obKeyPartDesc) GetPartIds(rowKeyPair *table.RangePair) ([]uint64, error) {
	if rowKeyPair == nil {
		return nil, errors.New("rowKeyPair is nil")
	}
	if rowKeyPair.Start() == nil || rowKeyPair.End() == nil {
		return nil, errors.New("startKeys or endKeys in rangePair is nil")
	}
	if rowKeyPair.IsStartEqEnd() {
		// check if startKey or endKey is extremum
		for i := 0; i < len(rowKeyPair.Start()); i++ {
			if _, ok := rowKeyPair.Start()[i].Value().(table.Entend); ok {
				return nil, errors.New("one of startKey or endKey is extremum")
			}
		}
		// startKey == endKey means that the range is equal to a column
		partId, err := d.GetPartId(rowKeyPair.Start())
		if err != nil {
			return []uint64{ObInvalidPartId}, errors.WithMessagef(err, "get part id, part desc:%s", d.String())
		}
		return []uint64{partId}, nil
	} else {
		// if startKey != endKey, add all partitions to the partition list
		partIds := make([]uint64, 0, d.partNum)
		for i := 0; i < d.partNum; i++ {
			partIds = append(partIds, uint64(i))
		}
		return partIds, nil
	}
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
		return -1, errors.Errorf("invalid type to convert to int64， value：%T", value)
	}
}

func (d *obKeyPartDesc) toHashCode(
	value interface{},
	column obColumn,
	hashCode int64,
	partFuncType obPartFuncType) (int64, error) {
	typeValue := column.ObjType().Value()
	if typeValue >= protocol.ObObjTypeTinyIntTypeValue && typeValue <= protocol.ObObjTypeUInt64TypeValue {
		i64, err := intToInt64(value)
		if err != nil {
			return -1, errors.WithMessagef(err, "convert int to int64, value:%T", typeValue)
		}
		arr := d.longToByteArray(i64)
		return murmurHash64A(arr, len(arr), hashCode), nil
	} else if typeValue == protocol.ObObjTypeDateTimeTypeValue || typeValue == protocol.ObObjTypeTimestampTypeValue {
		t, ok := value.(time.Time)
		if !ok {
			return -1, errors.Errorf("invalid timestamp type, value:%T", value)
		}
		return d.timeStampHash(t, hashCode), nil
	} else if typeValue == protocol.ObObjTypeDateTypeValue {
		date, ok := value.(time.Time)
		if !ok {
			return -1, errors.Errorf("invalid date type, value:%T", value)
		}
		return d.dateHash(date, hashCode), nil
	} else if typeValue == protocol.ObObjTypeVarcharTypeValue || typeValue == protocol.ObObjTypeCharTypeValue {
		return d.varcharHash(value, column.CollationType(), hashCode, partFuncType)
	} else {
		return -1, errors.Errorf("unsupported type for key hash, objType:%s", column.ObjType().String())
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
	} else {
		return -1, errors.Errorf("invalid varchar value for calc hash value, value:%T", value)
	}
	switch collType {
	case protocol.ObCollationTypeUtf8mb4GeneralCi:
		if partFuncType == partFuncTypeKeyV3 ||
			partFuncType == partFuncTypeKeyImplV2 ||
			util.ObVersion() >= 4 {
			hashCode = hashSortUtf8Mb4(bytes, hashCode, seed, true)
		} else {
			hashCode = hashSortUtf8Mb4(bytes, hashCode, seed, false)
		}
	case protocol.ObCollationTypeUtf8mb4Bin:
		fallthrough
	case protocol.ObCollationTypeBinary:
		if partFuncType == partFuncTypeKeyV3 ||
			partFuncType == partFuncTypeKeyImplV2 ||
			util.ObVersion() >= 4 {
			hashCode = murmurHash64A(bytes, len(bytes), hashCode)
		} else {
			hashCode = hashSortMbBin(bytes, hashCode, seed)
		}
	case protocol.ObCollationTypeInvalid:
		fallthrough
	case protocol.ObCollationTypeCollationFree:
		fallthrough
	case protocol.ObCollationTypeMax:
		fallthrough
	default:
		return -1, errors.Errorf("not supported collation type, collType:%d", collType)
	}
	return hashCode, nil
}

func (d *obKeyPartDesc) String() string {
	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(d.partColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += d.partColumns[i].String()
	}
	partColumnsStr += "]"

	return "obKeyPartDesc{" +
		"partSpace:" + strconv.Itoa(d.partSpace) + ", " +
		"partNum:" + strconv.Itoa(d.partNum) + ", " +
		"partColumns" + partColumnsStr +
		"}"
}

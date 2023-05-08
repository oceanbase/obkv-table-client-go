package route

import "strconv"

const (
	partFuncTypeUnknownIndex   = -1
	partFuncTypeHashIndex      = 0
	partFuncTypeKeyIndex       = 1
	partFuncTypeKeyImplIndex   = 2
	partFuncTypeRangeIndex     = 3
	partFuncTypeRangeColIndex  = 4
	partFuncTypeListIndex      = 5
	partFuncTypeKeyV2Index     = 6
	partFuncTypeListColIndex   = 7
	partFuncTypeHashV2Index    = 8
	partFuncTypeKeyV3Index     = 9
	partFuncTypeKeyImplV2Index = 10
)

const (
	partFuncTypeUnknown   = "UNKNOWN"
	partFuncTypeHash      = "HASH"
	partFuncTypeKey       = "KEY"
	partFuncTypeKeyImpl   = "KEY_IMPLICIT"
	partFuncTypeRange     = "RANGE"
	partFuncTypeRangeCol  = "RANGE_COLUMNS"
	partFuncTypeList      = "LIST"
	partFuncTypeKeyV2     = "KEY_V2"
	partFuncTypeListCol   = "LIST_COLUMNS"
	partFuncTypeHashV2    = "HASH_V2"
	partFuncTypeKeyV3     = "KEY_V3"
	partFuncTypeKeyImplV2 = "KEY_IMPLICIT_V2"
)

type ObPartFuncType struct {
	name  string
	index int
}

func (t ObPartFuncType) String() string {
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
		t.index == partFuncTypeKeyV3Index ||
		t.index == partFuncTypeKeyImplV2Index
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
	case partFuncTypeKeyImplV2Index:
		return ObPartFuncType{partFuncTypeKeyImplV2, partFuncTypeKeyImplV2Index}
	default:
		return ObPartFuncType{partFuncTypeUnknown, partFuncTypeUnknownIndex}
	}
}

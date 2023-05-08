package route

import "strconv"

const (
	PartLevelUnknown = "partLevelUnknown"
	PartLevelZero    = "partLevelZero"
	PartLevelOne     = "partLevelOne"
	PartLevelTwo     = "partLevelTwo"
)

const (
	PartLevelUnknownIndex = -1
	PartLevelZeroIndex    = 0
	PartLevelOneIndex     = 1
	PartLevelTwoIndex     = 2
)

type ObPartitionLevel struct {
	name  string
	index int
}

func (l ObPartitionLevel) Index() int {
	return l.index
}

func newObPartitionLevel(index int) ObPartitionLevel {
	switch index {
	case PartLevelZeroIndex:
		return ObPartitionLevel{PartLevelZero, PartLevelZeroIndex}
	case PartLevelOneIndex:
		return ObPartitionLevel{PartLevelOne, PartLevelOneIndex}
	case PartLevelTwoIndex:
		return ObPartitionLevel{PartLevelTwo, PartLevelTwoIndex}
	default:
		return ObPartitionLevel{PartLevelUnknown, PartLevelUnknownIndex}
	}
}

func (l ObPartitionLevel) String() string {
	return "ObPartitionLevel{" +
		"name:" + l.name + ", " +
		"index:" + strconv.Itoa(l.index) +
		"}"
}

package protocol

const (
	InvalidTableId     int64 = -1
	InvalidPartitionId int64 = 0
)

type TableConsistencyLevel uint8

const (
	Strong TableConsistencyLevel = iota
	Eventual
)

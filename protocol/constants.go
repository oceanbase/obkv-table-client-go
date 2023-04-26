package protocol

const (
	InvalidTableId     uint64 = 0
	InvalidPartitionId int64  = 0
)

type TableConsistencyLevel uint8

const (
	Strong TableConsistencyLevel = iota
	Eventual
)

package protocol

type TableOperationRequest struct {
	*UniVersionHeader
	credential           []uint8
	tableName            string
	tableId              int64
	partitionId          int64
	entityType           TableEntityType
	tableOperation       *TableOperation
	consistencyLevel     TableConsistencyLevel
	returnRowKey         bool
	returnAffectedEntity bool
	returnAffectedRows   bool
}

type TableEntityType int32

const (
	Dynamic TableEntityType = iota
	KV
	HKV
)

type TableConsistencyLevel int32

const (
	Strong TableConsistencyLevel = iota
	Eventual
)

package protocol

type TableOperation struct {
	*UniVersionHeader
	opType TableOperationType
	entity *TableEntity
}

type TableOperationType int32

const (
	Get TableOperationType = iota
	Insert
	Del
	Update
	InsertOrUpdate
	Replace
	Increment
	Append
)

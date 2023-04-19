package protocol

type TableOperationRequest struct {
	*UniVersionHeader
	credential           []byte
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

func (t *TableOperationRequest) Credential() []byte {
	return t.credential
}

func (t *TableOperationRequest) SetCredential(credential []byte) {
	t.credential = credential
}

func (t *TableOperationRequest) TableName() string {
	return t.tableName
}

func (t *TableOperationRequest) SetTableName(tableName string) {
	t.tableName = tableName
}

func (t *TableOperationRequest) TableId() int64 {
	return t.tableId
}

func (t *TableOperationRequest) SetTableId(tableId int64) {
	t.tableId = tableId
}

func (t *TableOperationRequest) PartitionId() int64 {
	return t.partitionId
}

func (t *TableOperationRequest) SetPartitionId(partitionId int64) {
	t.partitionId = partitionId
}

func (t *TableOperationRequest) EntityType() TableEntityType {
	return t.entityType
}

func (t *TableOperationRequest) SetEntityType(entityType TableEntityType) {
	t.entityType = entityType
}

func (t *TableOperationRequest) TableOperation() *TableOperation {
	return t.tableOperation
}

func (t *TableOperationRequest) SetTableOperation(tableOperation *TableOperation) {
	t.tableOperation = tableOperation
}

func (t *TableOperationRequest) ConsistencyLevel() TableConsistencyLevel {
	return t.consistencyLevel
}

func (t *TableOperationRequest) SetConsistencyLevel(consistencyLevel TableConsistencyLevel) {
	t.consistencyLevel = consistencyLevel
}

func (t *TableOperationRequest) ReturnRowKey() bool {
	return t.returnRowKey
}

func (t *TableOperationRequest) SetReturnRowKey(returnRowKey bool) {
	t.returnRowKey = returnRowKey
}

func (t *TableOperationRequest) ReturnAffectedEntity() bool {
	return t.returnAffectedEntity
}

func (t *TableOperationRequest) SetReturnAffectedEntity(returnAffectedEntity bool) {
	t.returnAffectedEntity = returnAffectedEntity
}

func (t *TableOperationRequest) ReturnAffectedRows() bool {
	return t.returnAffectedRows
}

func (t *TableOperationRequest) SetReturnAffectedRows(returnAffectedRows bool) {
	t.returnAffectedRows = returnAffectedRows
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

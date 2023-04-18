package protocol

import (
	"github.com/oceanbase/obkv-table-client-go/table"
	"time"
)

const (
	ObTableOperationTypeGet            = 0
	ObTableOperationTypeInsert         = 1
	ObTableOperationTypeDel            = 2
	ObTableOperationTypeUpdate         = 3
	ObTableOperationTypeInsertOrUpdate = 4
	ObTableOperationTypeReplace        = 5
	ObTableOperationTypeIncrement      = 6
	ObTableOperationTypeAppend         = 7
)

type ObTableOperationType int8

type ObRowkey struct {
	keys []interface{}
}

type ObTableEntity struct {
	rowkey     ObRowkey
	properties map[string]interface{}
}

func (o ObTableEntity) Rowkey() ObRowkey {
	return o.rowkey
}

func (o ObTableEntity) Properties() map[string]interface{} {
	return o.properties
}

type ObTableOperation struct {
}

func NewObTableOperation(
	opType ObTableOperationType,
	rowkeyValue []interface{},
	columns []table.Column) (*ObTableOperation, error) {
	// todo: new ObTableEntity
	return nil, nil
}

type ObTableOperationResult struct {
	affectedRows int64
	entity       ObTableEntity
}

func (o ObTableOperationResult) Entity() ObTableEntity {
	return o.entity
}

func (o ObTableOperationResult) AffectedRows() int64 {
	return o.affectedRows
}

type ObTableOperationRequest struct {
}

func NewObTableOperationRequest(
	tableName string,
	tableId uint64,
	partId int64,
	opType ObTableOperationType,
	rowkeyValue []interface{},
	columns []string,
	properties []interface{},
	opTimeOut time.Duration,
	logLevel uint16) (*ObTableOperationRequest, error) {
	// todo: impl
	return nil, nil
}

func (r *ObTableOperationRequest) String() string {
	// todo: impl
	return "ObTableOperationRequest{" +
		"}"
}

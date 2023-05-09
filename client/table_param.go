package client

import "strconv"

type ObTableParam struct {
	table       *ObTable
	tableId     uint64
	partitionId int64 // partition id in 3.x aka tablet id in 4.x
}

func NewObTableParam(table *ObTable, tableId uint64, partitionId int64) *ObTableParam {
	return &ObTableParam{table, tableId, partitionId}
}

func (p *ObTableParam) String() string {
	tableStr := "nil"
	if p.table != nil {
		tableStr = p.table.String()
	}
	return "ObTableParam{" +
		"table:" + tableStr + ", " +
		"tableId:" + strconv.Itoa(int(p.tableId)) + ", " +
		"partitionId:" + strconv.Itoa(int(p.partitionId)) +
		"}"
}

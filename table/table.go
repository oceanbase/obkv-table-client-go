package table

import (
	"strconv"
	"time"
)

type ObTable struct {
	ip         string
	port       int
	tenantName string
	userName   string
	password   string
	database   string
}

func NewObTable(
	ip string,
	port int,
	tenantName string,
	userName string,
	password string,
	database string,
	connSize int,
	connTimeOut time.Duration) (*ObTable, error) {
	t := &ObTable{
		ip:         ip,
		port:       port,
		tenantName: tenantName,
		userName:   userName,
		password:   password,
		database:   database,
	}
	return t, nil
}

func (t *ObTable) init() {

}

func (t *ObTable) Execute(request interface{}, result interface{}) error {
	// todo: impl
	return nil
}

func (t *ObTable) String() string {
	return "ObTable{" +
		"ip:" + t.ip + ", " +
		"port:" + strconv.Itoa(t.port) + ", " +
		"tenantName:" + t.tenantName + ", " +
		"userName:" + t.userName + ", " +
		"password:" + t.password + ", " +
		"database:" + t.database +
		"}"
}

type ObTableParam struct {
	table       *ObTable
	tableId     uint64
	partitionId int64 // partition id in 3.x aka tablet id in 4.x
}

func (p *ObTableParam) Table() *ObTable {
	return p.table
}

func (p *ObTableParam) PartitionId() int64 {
	return p.partitionId
}

func (p *ObTableParam) TableId() uint64 {
	return p.tableId
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

type ObRowkeyElement struct {
	nameIdxMap map[string]int
}

func NewObRowkeyElement(nameIdxMap map[string]int) *ObRowkeyElement {
	return &ObRowkeyElement{nameIdxMap}
}

func (e *ObRowkeyElement) NameIdxMap() map[string]int {
	return e.nameIdxMap
}

func (e *ObRowkeyElement) String() string {
	var nameIdxMapStr string
	var i = 0
	nameIdxMapStr = nameIdxMapStr + "{"
	for k, v := range e.nameIdxMap {
		if i > 0 {
			nameIdxMapStr += ", "
		}
		i++
		nameIdxMapStr += "m[" + k + "]=" + strconv.Itoa(v)
	}
	nameIdxMapStr += "}"
	return "ObRowkeyElement{" +
		"nameIdxMap:" + nameIdxMapStr +
		"}"
}

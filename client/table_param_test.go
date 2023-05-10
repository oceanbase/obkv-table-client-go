package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testTableIdV3 uint64 = 1099511677791
	testTableIdV4 uint64 = 500039
	testPartIdV3  int64  = 0
	testPartIdV4  int64  = 500040
)

func TestObTableParam_String(t *testing.T) {
	tableParam := NewObTableParam(&ObTable{}, testTableIdV3, testPartIdV3)
	assert.Equal(
		t,
		"ObTableParam{table:ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}, tableId:1099511677791, partitionId:0}",
		tableParam.String(),
	)
	tableParam = NewObTableParam(&ObTable{}, testTableIdV4, testPartIdV4)
	assert.Equal(
		t,
		"ObTableParam{table:ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}, tableId:500039, partitionId:500040}",
		tableParam.String(),
	)
	tb := NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	tableParam = NewObTableParam(tb, testTableIdV4, testPartIdV4)
	assert.Equal(
		t,
		"ObTableParam{table:ObTable{ip:127.0.0.1, port:8080, tenantName:sys, userName:root, password:, database:test, isClosed:false}, tableId:500039, partitionId:500040}",
		tableParam.String(),
	)
}

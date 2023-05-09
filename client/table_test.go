package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testIp         string = "127.0.0.1"
	testPort       int    = 8080
	testTenantName string = "sys"
	testUserName   string = "root"
	testPassword   string = ""
	testDatabase   string = "test"
)

func TestObTable_init(t *testing.T) {
	tb := NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	err := tb.init(1, time.Duration(1000)*time.Millisecond)
	assert.NotEqual(t, nil, err)
}

func TestObTable_close(t *testing.T) {
	tb := NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	tb.close()
	assert.EqualValues(t, true, tb.isClosed)
}

func TestObTable_String(t *testing.T) {
	tb := &ObTable{}
	assert.Equal(
		t,
		"ObTable{ip:, port:0, tenantName:, userName:, password:, database:, isClosed:false}",
		tb.String(),
	)
	tb = NewObTable(testIp, testPort, testTenantName, testUserName, testPassword, testDatabase)
	assert.Equal(
		t,
		"ObTable{ip:127.0.0.1, port:8080, tenantName:sys, userName:root, password:, database:test, isClosed:false}",
		tb.String(),
	)
}

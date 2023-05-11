package route

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewDB(t *testing.T) {
	const (
		testUserName = "root"
		testPassword = "::::"
		testIp       = "127.0.0.1"
		testSqlPort  = 41101
		testDatabase = "test"
	)

	_, err := NewDB(testUserName, testPassword, testIp, strconv.Itoa(testSqlPort), testDatabase)
	assert.NotEqual(t, nil, err)
}

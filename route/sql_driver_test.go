package route

import (
	"strconv"
	"testing"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB(testUserName, testPassword, testIp, strconv.Itoa(testSqlPort), testDatabase)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()
}

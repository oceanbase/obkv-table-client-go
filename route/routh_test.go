package route

import (
	"fmt"
	"testing"
)

// write your true config and create table by sql first
var (
	testClusterName = "test"
	testTenantName  = "mysql"
	testDatabase    = "test"
	testTableName   = "test"
	testUserName    = "root"
	testPassword    = ""
	testIp          = "127.0.0.1"
	testSqlPort     = 41101
	testServerPort  = 41100
	testServerAddr  = ObServerAddr{testIp, testSqlPort, testServerPort}
	testUserAuth    = ObUserAuth{testUserName, testPassword}
)

func TestGetTableEntryFromRemote(t *testing.T) {
	err := GetObVersionFromRemote(&testServerAddr, &testUserAuth)
	if err != nil {
		panic(err)
	}
	fmt.Println("ob cluster version is:", ObVersion)
	key := ObTableEntryKey{
		testClusterName,
		testTenantName,
		testDatabase,
		testTableName,
	}
	InitSql(ObVersion)
	entry, err := GetTableEntryFromRemote(&testServerAddr, &testUserAuth, &key)
	if err != nil {
		panic(err)
	}
	fmt.Println("entry is:", entry.ToString())
}

package single

import (
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()

	test.CreateDB()

	test.CreateTable(testTinyintCreateStatement)
	test.CreateTable(testUTinyintCreateStatement)
	test.CreateTable(testSmallintCreateStatement)
	test.CreateTable(testUSmallintCreateStatement)
	test.CreateTable(testInt32CreateStatement)
	test.CreateTable(testUInt32CreateStatement)
	test.CreateTable(testInt64CreateStatement)
	test.CreateTable(testUInt64CreateStatement)
	test.CreateTable(testFloatCreateStatement)
	test.CreateTable(testDoubleCreateStatement)
	test.CreateTable(testVarcharCreateStatement)
	test.CreateTable(testVarbinaryCreateStatement)
	test.CreateTable(testDatetimeCreateStatement)
}

func teardown() {
	cli.Close()

	test.DropTable(testTinyintTableName)
	test.DropTable(testUTinyintTableName)
	test.DropTable(testSmallintTableName)
	test.DropTable(testUSmallintTableName)
	test.DropTable(testInt32TableName)
	test.DropTable(testUInt32TableName)
	test.DropTable(testInt64TableName)
	test.DropTable(testUInt64TableName)
	test.DropTable(testFloatTableName)
	test.DropTable(testDoubleTableName)
	test.DropTable(testVarcharTableName)
	test.DropTable(testVarbinaryTableName)
	test.DropTable(testDatetimeTableName)

	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

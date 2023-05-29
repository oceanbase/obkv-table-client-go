package partition

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

	// hash
	test.CreateTable(hashBigIntL0CreateStatement)
	test.CreateTable(hashBigIntL1CreateStatement)
	test.CreateTable(hashBigIntL2CreateStatement)

	// key-bigint
	test.CreateTable(keyBigIntL0CreateStatement)
	test.CreateTable(keyBigIntL1CreateStatement)
	test.CreateTable(keyBigIntL2CreateStatement)

	// key-varchar
	test.CreateTable(keyVarcharL0CreateStatement)
	test.CreateTable(keyVarcharL1CreateStatement)
	test.CreateTable(keyVarcharL2CreateStatement)

	// key-varbinary
	test.CreateTable(keyVarBinaryL0CreateStatement)
	test.CreateTable(keyVarBinaryL1CreateStatement)
	test.CreateTable(keyVarBinaryL2CreateStatement)
}

func teardown() {
	cli.Close()

	// hash
	test.DropTable(hashBigIntL0TableName)
	test.DropTable(hashBigIntL1TableName)
	test.DropTable(hashBigIntL2TableName)

	// key-bigint
	test.DropTable(keyBigIntL0TableName)
	test.DropTable(keyBigIntL1TableName)
	test.DropTable(keyBigIntL2TableName)

	// key-varchar
	test.DropTable(keyVarcharL0TableName)
	test.DropTable(keyVarcharL1TableName)
	test.DropTable(keyVarcharL2TableName)

	// key-varbinary
	test.DropTable(keyVarBinaryL0TableName)
	test.DropTable(keyVarBinaryL1TableName)
	test.DropTable(keyVarBinaryL2TableName)

	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

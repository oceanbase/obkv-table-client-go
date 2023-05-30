package partition

import (
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

var cli client.Client

const partitionTestRecordCount = 10

func setup() {
	cli = test.CreateClient()

	test.CreateDB()

	// hash
	test.CreateTable(hashBigintL0CreateStatement)
	test.CreateTable(hashBigintL1CreateStatement)
	test.CreateTable(hashBigintL2CreateStatement)

	// key-Bigint
	test.CreateTable(keyBigintL1CreateStatement)
	test.CreateTable(keyBigintL2CreateStatement)

	// key-varchar
	test.CreateTable(keyVarcharL1CreateStatement)
	test.CreateTable(keyVarcharL2CreateStatement)

	// key-varbinary
	test.CreateTable(keyVarBinaryL1CreateStatement)
	test.CreateTable(keyVarBinaryL2CreateStatement)

	// key-MultiBigint
	test.CreateTable(keyMultiBigintL1CreateStatement)
	test.CreateTable(keyMultiBigintL2CreateStatement)

	// key-MultiVarchar
	test.CreateTable(keyMultiVarcharL1CreateStatement)
	test.CreateTable(keyMultiVarcharL2CreateStatement)

	// key-MultiVarBinary
	test.CreateTable(keyMultiVarBinaryL1CreateStatement)
	test.CreateTable(keyMultiVarBinaryL2CreateStatement)
}

func teardown() {
	cli.Close()

	// hash
	test.DropTable(hashBigintL0TableName)
	test.DropTable(hashBigintL1TableName)
	test.DropTable(hashBigintL2TableName)

	// key-Bigint
	test.DropTable(keyBigintL1TableName)
	test.DropTable(keyBigintL2TableName)

	// key-varchar
	test.DropTable(keyVarcharL1TableName)
	test.DropTable(keyVarcharL2TableName)

	// key-varbinary
	test.DropTable(keyVarBinaryL1TableName)
	test.DropTable(keyVarBinaryL2TableName)

	// key-MultiBigint
	test.DropTable(keyMultiBigintL1TableName)
	test.DropTable(keyMultiBigintL2TableName)

	// key-MultiVarchar
	test.DropTable(keyMultiVarcharL1TableName)
	test.DropTable(keyMultiVarcharL2TableName)

	// key-MultiVarBinary
	test.DropTable(keyMultiVarBinaryL1TableName)
	test.DropTable(keyMultiVarBinaryL2TableName)

	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

package batch

import (
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	batchOpTableTableName       = "batchOpTable"
	batchOpTableCreateStatement = "create table if not exists batchOpTable(`c1` bigint(20) not null, c2 bigint(20) not null, primary key (`c1`)) partition by hash(c1) partitions 2;"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()

	test.CreateDB()
	test.CreateTable(batchOpTableCreateStatement)
}

func teardown() {
	cli.Close()

	test.DropTable(batchOpTableTableName)
	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

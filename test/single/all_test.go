package single

import (
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	singleOpTableTableName       = "singleOpTable"
	singleOpTableCreateStatement = "create table if not exists singleOpTable(`c1` bigint(20) not null, c2 bigint(20) not null, primary key (`c1`)) partition by hash(c1) partitions 2;"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()

	test.CreateDB()
	test.CreateTable(singleOpTableCreateStatement)
}

func teardown() {
	cli.Close()

	test.DropTable(singleOpTableTableName)
	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

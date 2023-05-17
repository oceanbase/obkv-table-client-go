package partition

import (
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

const (
	hashPartitionL1TableName       = "hashPartitionL1"
	hashPartitionL1CreateStatement = "create table if not exists hashPartitionL1(`c1` bigint(20) not null, c2 bigint(20) not null, primary key (`c1`)) partition by hash(c1) partitions 2;"

	keyPartitionIntL1TableName       = "keyPartitionIntL1"
	keyPartitionIntL1CreateStatement = "create table if not exists keyPartitionIntL1(`c1` bigint(20) not null, c2 bigint(20) not null, primary key (`c1`)) partition by key(c1) partitions 15;"

	keyPartitionVarcharL1TableName       = "keyPartitionVarcharL1"
	keyPartitionVarcharL1CreateStatement = "create table if not exists keyPartitionVarcharL1(`c1` varchar(20) not null, c2 bigint(20) not null, primary key (`c1`)) partition by key(c1) partitions 15;"

	hashPartitionL2TableName       = "hashPartitionL2"
	hashPartitionL2CreateStatement = "create table if not exists hashPartitionL2(`c1` bigint(20) not null, `c2` bigint(20) not null, `c3` bigint(20) not null, primary key (`c1`, `c2`)) partition by hash(`c1`) subpartition by hash(`c2`) subpartitions 4 partitions 16;"

	keyPartitionIntL2TableName       = "keyPartitionIntL2"
	keyPartitionIntL2CreateStatement = "create table if not exists keyPartitionIntL2(`c1` bigint(20) not null, c2 bigint(20) not null, `c3` bigint(20) not null, primary key (`c1`, `c2`)) partition by key(`c1`) subpartition by key(`c2`) subpartitions 4 partitions 16;"

	keyPartitionVarcharL2TableName       = "keyPartitionVarcharL2"
	keyPartitionVarcharL2CreateStatement = "create table if not exists keyPartitionVarcharL2(`c1` varchar(20) not null, `c2` varchar(20) not null, c3 bigint(20) not null, primary key (`c1`, `c2`)) partition by key(`c1`) subpartition by key(`c2`) subpartitions 4 partitions 16;"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()

	test.CreateDB()
	test.CreateTable(hashPartitionL1CreateStatement)
	test.CreateTable(keyPartitionIntL1CreateStatement)
	test.CreateTable(keyPartitionVarcharL1CreateStatement)
	test.CreateTable(hashPartitionL2CreateStatement)
	test.CreateTable(keyPartitionIntL2CreateStatement)
	test.CreateTable(keyPartitionVarcharL2CreateStatement)
}

func teardown() {
	cli.Close()

	test.DropTable(hashPartitionL1TableName)
	test.DropTable(keyPartitionIntL1TableName)
	test.DropTable(keyPartitionVarcharL1TableName)
	test.DropTable(hashPartitionL2TableName)
	test.DropTable(keyPartitionIntL2TableName)
	test.DropTable(keyPartitionVarcharL2TableName)
	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

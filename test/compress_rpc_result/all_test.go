/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2023 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package compress_rpc_result

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"testing"

	"github.com/oceanbase/obkv-table-client-go/client"
	"github.com/oceanbase/obkv-table-client-go/test"
)

var cli client.Client

func setup() {
	cli = test.CreateClient()
	test.CreateDB()
	test.CreateTable(createTableStat)
}

func setCompressType(compressTyepe string) error {
	if test.GlobalDB == nil {
		return errors.New("test GlobalDb is nil")
	}
	CompressFuncSql := "alter system set kv_transport_compress_func= '%s'"
	CompressThresholdSql := "alter system set kv_transport_compress_threshold= '%s'"
	setCompressFunc := fmt.Sprintf(CompressFuncSql, compressTyepe)
	setCompressThreshold := fmt.Sprintf(CompressThresholdSql, "0K")
	_, err := test.GlobalDB.Exec(setCompressFunc)
	if err != nil {
		return errors.WithMessagef(err, "fail to set compress func %s", compressTyepe)
	}
	_, err = test.GlobalDB.Exec(setCompressThreshold)
	if err != nil {
		return errors.WithMessagef(err, "fail to set compress threshold %s", "0K")
	}
	return nil
}

func teardown() {
	cli.Close()
	test.DropTable(tableName)
	test.CloseDB()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

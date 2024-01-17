/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
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

package logger

import (
	"context"
	"github.com/oceanbase/obkv-table-client-go/config"
	"github.com/oceanbase/obkv-table-client-go/log"
	"os"
	"testing"
)

func TestInitTraceId(t *testing.T) {
	file := "obclient-table-go.log"
	filePath := "../../configurations/obkv-table-default.toml"
	clientConfig, _ := config.GetClientConfigurationFromTOML(filePath)
	var logcon log.LogConfig
	logcon.LogFileName = clientConfig.LogConfig.LogFileName
	logcon.MaxAgeFileRem = clientConfig.LogConfig.MaxAgeFileRem
	logcon.MaxBackupFileSize = clientConfig.LogConfig.MaxBackupFileSize
	logcon.SingleFileMaxSize = clientConfig.LogConfig.SingleFileMaxSize
	logcon.Compress = clientConfig.LogConfig.Compress
	logcon.SlowQueryThreshold = clientConfig.LogConfig.SlowQueryThreshold
	ctx := context.TODO()

	log.InitLoggerWithConfig(logcon)
	ctx = context.TODO()
	log.InitTraceId(&ctx)
	log.Info("Default", ctx.Value(log.ObkvTraceIdName), "Info msg", log.String("Info", "second_Info"))
	for i := 0; i < 10; i++ {
		log.Error("BOOT", ctx.Value(log.ObkvTraceIdName), "Error msg", log.String("Error", "fourth_Error"))
		log.Debug("", ctx.Value(log.ObkvTraceIdName), "Debug msg", log.String("Debug", "first_Debug"))
		log.Warn("Default", nil, "Warn msg", log.String("Warn", "third_Warn"))
	}
	_ = log.Sync()
	_ = os.Remove(file)
}

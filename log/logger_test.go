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

package log

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	myLogger := NewLogger(os.Stderr, InfoLevel, AddCaller())
	ResetDefaultLogger(myLogger)
	myLogger.Debug("Debug msg", String("Debug", "Debug"))
	myLogger.Info("Info msg", String("Info", "Info"))
	myLogger.Warn("Warn msg", String("Warn", "Warn"))
	myLogger.Error("Error msg", String("Error", "Error"))
}

func TestFileLogger(t *testing.T) {
	var defWriteSync = getLogWriter("obclient-table-go.log")
	var Logger1 = NewLogger(defWriteSync, InfoLevel, AddCaller())
	ResetDefaultLogger(Logger1)
	Logger1.Debug("Debug msg", String("Debug", "Debug"))
	Logger1.Info("Info msg", String("Info", "Info"))
	Logger1.Warn("Warn msg", String("Warn", "Warn"))
	Logger1.Error("Error msg", String("Error", "Error"))
	_ = Logger1.Sync()
	_ = os.Remove("obclient-table-go.log")
}

func TestFileRotationLogger(t *testing.T) {
	var logc LogConfig
	logc.MaxBackupFileSize = 20
	logc.Compress = false
	logc.SingleFileMaxSize = 1
	logc.LogFileName = "./"
	logc.MaxAgeFileRem = 30
	ctx := context.TODO()
	var defWriteSync = getLogRotationWriter("obclient-table-go.log", logc)
	var Logger1 = NewLogger(defWriteSync, InfoLevel, AddCaller())
	ResetDefaultLogger(Logger1)
	Logger1.Debug("Debug msg", String("Debug", "first_Debug"))
	Logger1.Info("Info msg", String("Info", "second_Info"))
	Logger1.Warn("Warn msg", String("Warn", "thrid_Warn"))
	Logger1.Error("Error msg", String("Error", "fourth_Error"))

	// test traceId
	InitTraceId(&ctx)
	Debug("Default", ctx.Value(ObkvTraceIdName).(string), "Debug msg", String("Debug", "first_Debug"))
	Info("", ctx.Value(ObkvTraceIdName).(string), "Info msg", String("Info", "second_Info"))
	Warn("Boot", ctx.Value(ObkvTraceIdName).(string), "Warn msg", String("Warn", "thrid_Warn"))
	Error("", ctx.Value(ObkvTraceIdName).(string), "Error msg", String("Error", "fourth_Error"))
	DPanic("routine", ctx.Value(ObkvTraceIdName).(string), "DPanic msg", String("DPanic", "fifth_DPanic"))
	_ = Logger1.Sync()
	_ = os.Remove("obclient-table-go.log")
}

func TestFmtPrint(t *testing.T) {
	intArr := []int{1, 2, 3, 4}
	str := fmt.Sprintf("Y%X-%016X-%x-%x \n", intArr[0], intArr[1], intArr[2], intArr[3])
	fmt.Printf(str)
}

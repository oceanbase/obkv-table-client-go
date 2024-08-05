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

package config

import (
	"strconv"
	"time"

	"github.com/oceanbase/obkv-table-client-go/log"
)

type ClientConfig struct {
	ConnPoolMaxConnSize int
	ConnConnectTimeOut  time.Duration
	ConnLoginTimeout    time.Duration
	OperationTimeOut    time.Duration
	LogLevel            log.Level

	TableEntryRefreshLockTimeout     time.Duration
	TableEntryRefreshTryTimes        int
	TableEntryRefreshIntervalBase    time.Duration
	TableEntryRefreshIntervalCeiling time.Duration

	MetadataRefreshInterval    time.Duration
	MetadataRefreshLockTimeout time.Duration

	RsListLocalFileLocation    string
	RsListHttpGetTimeout       time.Duration
	RsListHttpGetRetryTimes    int
	RsListHttpGetRetryInterval time.Duration

	EnableRerouting bool

	// connection rebalance in ODP mode
	MaxConnectionAge     time.Duration
	EnableSLBLoadBalance bool

	// log config
	LogFileName        string // log file dir
	SingleFileMaxSize  int    // log file size（MB）
	MaxBackupFileSize  int    // Maximum number of old files to keep
	MaxAgeFileRem      int    // Maximum number of days to keep old files
	Compress           bool   // Whether to compress/archive old files
	SlowQueryThreshold int64  // Slow query threshold

	NeedCalculateGenerateColumn bool
}

func NewDefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		ConnPoolMaxConnSize:              1,
		ConnConnectTimeOut:               time.Duration(1000) * time.Millisecond,  // 1s
		ConnLoginTimeout:                 time.Duration(5000) * time.Millisecond,  // 5s
		OperationTimeOut:                 time.Duration(10000) * time.Millisecond, // 10s
		LogLevel:                         log.WarnLevel,
		TableEntryRefreshLockTimeout:     time.Duration(4000) * time.Millisecond, // 4s
		TableEntryRefreshTryTimes:        3,
		TableEntryRefreshIntervalBase:    time.Duration(100) * time.Millisecond,   // 100ms
		TableEntryRefreshIntervalCeiling: time.Duration(1600) * time.Millisecond,  // 1.6s
		MetadataRefreshInterval:          time.Duration(60000) * time.Millisecond, // 60s
		MetadataRefreshLockTimeout:       time.Duration(8000) * time.Millisecond,  // 8s
		RsListLocalFileLocation:          "",
		RsListHttpGetTimeout:             time.Duration(1000) * time.Millisecond, // 1s
		RsListHttpGetRetryTimes:          3,
		RsListHttpGetRetryInterval:       time.Duration(100) * time.Millisecond, // 100ms,
		EnableRerouting:                  true,
		MaxConnectionAge:                 time.Duration(0) * time.Second, // valid iff > 0
		EnableSLBLoadBalance:             false,
		LogFileName:                      "./",
		SingleFileMaxSize:                256, // MB
		MaxBackupFileSize:                10,
		MaxAgeFileRem:                    30,
		Compress:                         false,
		SlowQueryThreshold:               40, // ms
		NeedCalculateGenerateColumn:      true,
	}
}

func (c *ClientConfig) String() string {
	return "ClientConfig{" +
		"ConnPoolMaxConnSize:" + strconv.Itoa(c.ConnPoolMaxConnSize) + ", " +
		"ConnConnectTimeOut:" + c.ConnConnectTimeOut.String() + ", " +
		"ConnLoginTimeout:" + c.ConnLoginTimeout.String() + ", " +
		"OperationTimeOut:" + c.OperationTimeOut.String() + ", " +
		"LogLevel:" + strconv.Itoa(int(c.LogLevel)) + ", " +
		"TableEntryRefreshLockTimeout:" + c.TableEntryRefreshLockTimeout.String() + ", " +
		"TableEntryRefreshTryTimes:" + strconv.Itoa(c.TableEntryRefreshTryTimes) + ", " +
		"TableEntryRefreshIntervalBase:" + c.TableEntryRefreshIntervalBase.String() + ", " +
		"TableEntryRefreshIntervalCeiling:" + c.TableEntryRefreshIntervalCeiling.String() + ", " +
		"MetadataRefreshInterval:" + c.MetadataRefreshInterval.String() + ", " +
		"MetadataRefreshLockTimeout:" + c.MetadataRefreshLockTimeout.String() + ", " +
		"RsListLocalFileLocation:" + c.RsListLocalFileLocation + ", " +
		"RsListHttpGetTimeout:" + c.RsListHttpGetTimeout.String() + ", " +
		"RsListHttpGetRetryTimes:" + strconv.Itoa(c.RsListHttpGetRetryTimes) + ", " +
		"RsListHttpGetRetryInterval:" + c.RsListHttpGetRetryInterval.String() + ", " +
		"EnableRerouting:" + strconv.FormatBool(c.EnableRerouting) + ", " +
		"MaxConnectionAge:" + c.MaxConnectionAge.String() + ", " +
		"EnableSLBLoadBalance:" + strconv.FormatBool(c.EnableSLBLoadBalance) +
		"NeedCalculateGenerateColumn:" + strconv.FormatBool(c.NeedCalculateGenerateColumn) +
		"}"
}

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

package config

import (
	"github.com/naoina/toml"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/pkg/errors"
	"net"
	"os"
	"strconv"
	"time"
)

type ClientConfiguration struct {
	DirectClientConfig  DirectClientConfig
	OdpClientConfig     ODPClientConfig
	ConnConfig          ConnConfig
	TableEntryConfig    TableEntryConfig
	RouteMetaDataConfig RouteMetaDataConfig
	RsListConfig        RsListConfig
	ExtraConfig         ExtraConfig
	LogConfig           LogConfig
	Mode                string
}

type DirectClientConfig struct {
	ConfigUrl    string
	FullUserName string
	Password     string
	SysUserName  string
	SysPassword  string
}

type ODPClientConfig struct {
	OdpIp        string
	OdpRpcPort   int
	OdpSqlPort   int
	FullUserName string
	Password     string
	Database     string
}

type ConnConfig struct {
	PoolMaxConnSize int
	ConnectTimeOut  int
	LoginTimeout    int
}

type TableEntryConfig struct {
	RefreshLockTimeout     int
	RefreshTryTimes        int
	RefreshIntervalBase    int
	RefreshIntervalCeiling int
}

type RouteMetaDataConfig struct {
	RefreshInterval    int
	RefreshLockTimeout int
}

type RsListConfig struct {
	LocalFileLocation    string
	HttpGetTimeout       int
	HttpGetRetryTimes    int
	HttpGetRetryInterval int
}

type ExtraConfig struct {
	OperationTimeOut     int
	LogLevel             string
	EnableRerouting      bool
	MaxConnectionAge     int
	EnableSLBLoadBalance bool
}

type LogConfig struct {
	LogFileName        string // log file dir
	SingleFileMaxSize  int    // log file size（MB）
	MaxBackupFileSize  int    // Maximum number of old files to keep
	MaxAgeFileRem      int    // Maximum number of days to keep old files
	Compress           bool   // Whether to compress/archive old files
	SlowQueryThreshold int64  // Slow query threshold
}

func (c *ClientConfiguration) checkClientConfiguration() error {
	if c.Mode == "direct" {
		if c.DirectClientConfig.ConfigUrl == "" {
			return errors.New("config url is empty")
		} else if c.DirectClientConfig.FullUserName == "" {
			return errors.New("full user name is empty")
		} else if c.DirectClientConfig.SysUserName == "" {
			return errors.New("sys user name is empty")
		}
	} else if c.Mode == "proxy" {
		if net.ParseIP(c.OdpClientConfig.OdpIp) == nil {
			return errors.New("odp ip is empty")
		} else if c.OdpClientConfig.OdpRpcPort == 0 {
			return errors.New("odp rpc port is empty")
		} else if c.OdpClientConfig.OdpSqlPort == 0 {
			return errors.New("odp sql port is empty")
		} else if c.OdpClientConfig.Database == "" {
			return errors.New("database name is empty")
		} else if c.OdpClientConfig.FullUserName == "" {
			return errors.New("full user name is empty")
		}
	} else if c.Mode == "log" {
		if c.LogConfig.LogFileName == "" {
			return errors.New("should set log file name in toml config")
		} else if c.LogConfig.SingleFileMaxSize == 0 {
			return errors.New("single file maxSize is invalid")
		} else if c.LogConfig.SlowQueryThreshold == 0 {
			return errors.New("slow query threshold is invalid")
		}
	} else {
		return errors.New("mode is invalid")
	}
	return nil
}

func GetClientConfigurationFromTOML(filepath string) (*ClientConfiguration, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config := new(ClientConfiguration)
	if err = toml.NewDecoder(f).Decode(config); err != nil {
		return nil, err
	}
	if err = config.checkClientConfiguration(); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *ClientConfiguration) GetClientConfig() *ClientConfig {
	return &ClientConfig{
		ConnPoolMaxConnSize:              c.ConnConfig.PoolMaxConnSize,
		ConnConnectTimeOut:               time.Duration(c.ConnConfig.ConnectTimeOut) * time.Millisecond,
		ConnLoginTimeout:                 time.Duration(c.ConnConfig.LoginTimeout) * time.Millisecond,
		OperationTimeOut:                 time.Duration(c.ExtraConfig.OperationTimeOut) * time.Millisecond,
		LogLevel:                         log.MatchStr2LogLevel(c.ExtraConfig.LogLevel),
		TableEntryRefreshLockTimeout:     time.Duration(c.TableEntryConfig.RefreshLockTimeout) * time.Millisecond,
		TableEntryRefreshTryTimes:        c.TableEntryConfig.RefreshTryTimes,
		TableEntryRefreshIntervalBase:    time.Duration(c.TableEntryConfig.RefreshIntervalBase) * time.Millisecond,
		TableEntryRefreshIntervalCeiling: time.Duration(c.TableEntryConfig.RefreshIntervalCeiling) * time.Millisecond,
		MetadataRefreshInterval:          time.Duration(c.RouteMetaDataConfig.RefreshInterval) * time.Millisecond,
		MetadataRefreshLockTimeout:       time.Duration(c.RouteMetaDataConfig.RefreshLockTimeout) * time.Millisecond,
		RsListLocalFileLocation:          c.RsListConfig.LocalFileLocation,
		RsListHttpGetTimeout:             time.Duration(c.RsListConfig.HttpGetTimeout) * time.Millisecond,
		RsListHttpGetRetryTimes:          c.RsListConfig.HttpGetRetryTimes,
		RsListHttpGetRetryInterval:       time.Duration(c.RsListConfig.HttpGetRetryInterval) * time.Millisecond,
		EnableRerouting:                  c.ExtraConfig.EnableRerouting,
		MaxConnectionAge:                 time.Duration(c.ExtraConfig.MaxConnectionAge) * time.Millisecond,
		EnableSLBLoadBalance:             c.ExtraConfig.EnableSLBLoadBalance,
		LogFileName:                      c.LogConfig.LogFileName,
		SingleFileMaxSize:                c.LogConfig.SingleFileMaxSize, // MB
		MaxBackupFileSize:                c.LogConfig.MaxBackupFileSize,
		MaxAgeFileRem:                    c.LogConfig.MaxAgeFileRem,
		Compress:                         c.LogConfig.Compress,
		SlowQueryThreshold:               c.LogConfig.SlowQueryThreshold, // ms
	}
}

func (c *ClientConfiguration) String() string {
	if c.Mode == "odp" {
		return "Configuration{" + c.Mode + ", " + c.OdpClientConfig.FullUserName + ", " + c.OdpClientConfig.OdpIp + ", " + strconv.Itoa(c.OdpClientConfig.OdpRpcPort) + ", " + strconv.Itoa(c.OdpClientConfig.OdpSqlPort) + ", " + c.OdpClientConfig.Database + "}"
	}
	return "Configuration{" + c.Mode + ", " + c.DirectClientConfig.FullUserName + ", " + c.DirectClientConfig.ConfigUrl + "}"
}

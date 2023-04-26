package config

import (
	"strconv"
	"time"
)

type ClientConfig struct {
	ConnPoolMaxConnSize int
	RpcConnectTimeOut   time.Duration
	OperationTimeOut    time.Duration
	LogLevel            uint16

	TableEntryRefreshLockTimeout     time.Duration
	TableEntryRefreshTryTimes        int
	TableEntryRefreshIntervalBase    time.Duration
	TableEntryRefreshIntervalCeiling time.Duration

	MetadataRefreshInterval    time.Duration
	MetadataRefreshLockTimeout time.Duration
}

func NewDefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		ConnPoolMaxConnSize:              1,
		RpcConnectTimeOut:                time.Duration(1000) * time.Millisecond,  // 1s
		OperationTimeOut:                 time.Duration(10000) * time.Millisecond, // 10s
		LogLevel:                         7,
		TableEntryRefreshLockTimeout:     time.Duration(4000) * time.Millisecond, // 4s
		TableEntryRefreshTryTimes:        3,
		TableEntryRefreshIntervalBase:    time.Duration(100) * time.Millisecond,   // 100ms
		TableEntryRefreshIntervalCeiling: time.Duration(1600) * time.Millisecond,  // 1.6s
		MetadataRefreshInterval:          time.Duration(60000) * time.Millisecond, // 60s
		MetadataRefreshLockTimeout:       time.Duration(8000) * time.Millisecond,  // 8s
	}
}

func (c *ClientConfig) String() string {
	return "ClientConfig{" +
		"ConnPoolMaxConnSize:" + strconv.Itoa(c.ConnPoolMaxConnSize) + ", " +
		"OperationTimeOut:" + strconv.Itoa(int(c.OperationTimeOut)) + ", " +
		"LogLevel:" + strconv.Itoa(int(c.LogLevel)) +
		"}"
}

package config

import (
	"strconv"
	"time"
)

type ClientConfig struct {
	ConnPoolMaxConnSize int
	ConnTimeOut         time.Duration
	OperationTimeOut    time.Duration
	LogLevel            uint16
}

func NewDefaultClientConfig() *ClientConfig {
	return &ClientConfig{}
}

func (c *ClientConfig) String() string {
	return "ClientConfig{" +
		"ConnPoolMaxConnSize:" + strconv.Itoa(c.ConnPoolMaxConnSize) + ", " +
		"OperationTimeOut:" + strconv.Itoa(int(c.OperationTimeOut)) + ", " +
		"LogLevel:" + strconv.Itoa(int(c.LogLevel)) +
		"}"
}

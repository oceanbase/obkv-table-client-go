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
	"encoding/xml"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/pkg/errors"
	"net"
	"os"
	"strconv"
	"time"
)

type XmlLoginConfiguration struct {
	Mode         string
	ConfigUrl    string
	FullUserName string
	Password     string
	SysUserName  string
	SysPassword  string
	OdpIp        string
	OdpRpcPort   int
	OdpSqlPort   int
	Database     string
	TableConfig  *ClientConfig
}

type xmlConfiguration struct {
	Properties []xmlProperty `xml:"property"`
}

type xmlProperty struct {
	Name        string `xml:"name"`
	Value       string `xml:"value"`
	Description string `xml:"description"`
}

// NewDefaultXmlConfiguration creates a new default Configuration.
func NewDefaultXmlConfiguration() *XmlLoginConfiguration {
	return &XmlLoginConfiguration{
		Mode:        "direct",
		TableConfig: NewDefaultClientConfig(),
	}
}

// NewConfigurationWithXML creates a new Configuration with xml config file.
func NewConfigurationWithXML(configFilePath string) (*XmlLoginConfiguration, error) {
	config := NewDefaultXmlConfiguration()
	err := config.setConfigWithXml(configFilePath)
	if err != nil {
		return nil, errors.WithMessagef(err, "set config with xml failed")
	}
	err = checkXmlConfig(config)
	if err != nil {
		return nil, errors.WithMessagef(err, "check xml config failed")
	}
	return config, err
}

func (c *XmlLoginConfiguration) setConfigWithXml(configFileName string) error {
	wd, _ := os.Getwd()
	_ = wd
	file, err := os.Open(configFileName)
	if err != nil {
		return errors.WithMessagef(err, "open config file %s failed", configFileName)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// TODO: log
		}
	}(file)

	// decode xml
	decoder := xml.NewDecoder(file)
	config := xmlConfiguration{}
	err = decoder.Decode(&config)
	if err != nil {
		return errors.WithMessagef(err, "decode config file %s failed", configFileName)
	}

	err = c.setConfigWithXmlConfig(config)
	if err != nil {
		return errors.WithMessagef(err, "set config with xml config failed")
	}

	return nil
}

func (c *XmlLoginConfiguration) setConfigWithXmlConfig(xmlConfig xmlConfiguration) error {
	for _, property := range xmlConfig.Properties {
		var err error
		var value int
		switch property.Name {
		// OBKV client login configurations
		case "obkv.table.client.mode":
			c.Mode = property.Value
		case "obkv.table.client.config.url":
			c.ConfigUrl = property.Value
		case "obkv.table.client.full.username":
			c.FullUserName = property.Value
		case "obkv.table.client.password":
			c.Password = property.Value
		case "obkv.table.client.sys.username":
			c.SysUserName = property.Value
		case "obkv.table.client.sys.password":
			c.SysPassword = property.Value
		case "obkv.table.client.odp.ip":
			c.OdpIp = property.Value
		case "obkv.table.client.odp.rpc.port":
			value, err = strconv.Atoi(property.Value)
			c.OdpRpcPort = value
		case "obkv.table.client.odp.sql.port":
			value, err = strconv.Atoi(property.Value)
			c.OdpSqlPort = value
		case "obkv.table.client.odp.database.name":
			c.Database = property.Value

		// OBKV Table client configurations
		case "obkv.table.client.conn.pool.max.conn.size":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.ConnPoolMaxConnSize = value
		case "obkv.table.client.conn.Connect.timeout":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.ConnConnectTimeOut = time.Duration(value) * time.Millisecond
		case "obkv.table.client.conn.login.timeout":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.ConnLoginTimeout = time.Duration(value) * time.Millisecond
		case "obkv.table.client.conn.operation.timeout":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.OperationTimeOut = time.Duration(value) * time.Millisecond
		case "obkv.table.client.log.level":
			c.TableConfig.LogLevel = log.MatchStr2LogLevel(property.Value)
		case "obkv.table.client.table.entry.refresh.lock.timeout":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.MetadataRefreshLockTimeout = time.Duration(value) * time.Millisecond
		case "obkv.table.client.table.entry.refresh.try.times":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.TableEntryRefreshTryTimes = value
		case "obkv.table.client.table.entry.refresh.interval.base":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.TableEntryRefreshIntervalBase = time.Duration(value) * time.Millisecond
		case "obkv.table.client.table.entry.refresh.interval.ceiling":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.TableEntryRefreshIntervalCeiling = time.Duration(value) * time.Millisecond
		case "obkv.table.client.metadata.refresh.interval":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.MetadataRefreshInterval = time.Duration(value) * time.Millisecond
		case "obkv.table.client.metadata.refresh.lock.timeout":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.MetadataRefreshLockTimeout = time.Duration(value) * time.Millisecond
		case "obkv.table.client.rslist.local.file.location":
			c.TableConfig.RsListLocalFileLocation = property.Value
		case "obkv.table.client.rslist.http.get.timeout":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.RsListHttpGetTimeout = time.Duration(value) * time.Millisecond
		case "obkv.table.client.rslist.http.get.retry.times":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.RsListHttpGetRetryTimes = value
		case "obkv.table.client.rslist.http.get.retry.interval":
			value, err = strconv.Atoi(property.Value)
			c.TableConfig.RsListHttpGetRetryInterval = time.Duration(value) * time.Millisecond
		}
		if err != nil {
			return errors.WithMessagef(err, "set config with xml config %s failed", property.Name)
		}
	}
	return nil
}

func checkXmlConfig(config *XmlLoginConfiguration) error {
	if config.Mode == "direct" {
		if config.ConfigUrl == "" || config.ConfigUrl == "..." {
			return errors.New("config url is empty")
		} else if config.FullUserName == "" || config.FullUserName == "..." {
			return errors.New("full user name is empty")
		} else if config.SysUserName == "" || config.SysUserName == "..." {
			return errors.New("sys user name is empty")
		}
	} else if config.Mode == "proxy" {
		if net.ParseIP(config.OdpIp) == nil {
			return errors.New("odp ip is empty")
		} else if config.OdpRpcPort == 0 {
			return errors.New("odp rpc port is empty")
		} else if config.OdpSqlPort == 0 {
			return errors.New("odp sql port is empty")
		} else if config.Database == "" || config.Database == "..." {
			return errors.New("database name is empty")
		} else if config.FullUserName == "" || config.FullUserName == "..." {
			return errors.New("full user name is empty")
		}
	} else {
		return errors.New("mode is invali")
	}
	return nil
}

func (c *XmlLoginConfiguration) String() string {
	if c.Mode == "odp" {
		return "Configuration{" + c.Mode + ", " + c.FullUserName + ", " + c.OdpIp + ", " + strconv.Itoa(c.OdpRpcPort) + ", " + strconv.Itoa(c.OdpSqlPort) + ", " + c.Database + ", " + c.TableConfig.String() + "}"
	}
	return "Configuration{" + c.Mode + ", " + c.FullUserName + ", " + c.ConfigUrl + ", " + c.TableConfig.String() + ", " + c.TableConfig.String() + "}"
}

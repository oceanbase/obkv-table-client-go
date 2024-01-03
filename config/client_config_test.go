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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientConfig_String(t *testing.T) {
	cfg := &ClientConfig{}
	assert.Equal(t, "ClientConfig{"+
		"ConnPoolMaxConnSize:0, "+
		"ConnConnectTimeOut:0s, "+
		"ConnLoginTimeout:0s, "+
		"OperationTimeOut:0s, "+
		"LogLevel:0, "+
		"TableEntryRefreshLockTimeout:0s, "+
		"TableEntryRefreshTryTimes:0, "+
		"TableEntryRefreshIntervalBase:0s, "+
		"TableEntryRefreshIntervalCeiling:0s, "+
		"MetadataRefreshInterval:0s, "+
		"MetadataRefreshLockTimeout:0s, "+
		"RsListLocalFileLocation:, "+
		"RsListHttpGetTimeout:0s, "+
		"RsListHttpGetRetryTimes:0, "+
		"RsListHttpGetRetryInterval:0s, "+
		"EnableRerouting:false, "+
		"MaxConnectionAge:0s, "+
		"EnableSLBLoadBalance:false"+
		"}", cfg.String())
	cfg = NewDefaultClientConfig()
	assert.Equal(t, "ClientConfig{"+
		"ConnPoolMaxConnSize:1, "+
		"ConnConnectTimeOut:1s, "+
		"ConnLoginTimeout:5s, "+
		"OperationTimeOut:10s, "+
		"LogLevel:1, "+
		"TableEntryRefreshLockTimeout:4s, "+
		"TableEntryRefreshTryTimes:3, "+
		"TableEntryRefreshIntervalBase:100ms, "+
		"TableEntryRefreshIntervalCeiling:1.6s, "+
		"MetadataRefreshInterval:1m0s, "+
		"MetadataRefreshLockTimeout:8s, "+
		"RsListLocalFileLocation:, "+
		"RsListHttpGetTimeout:1s, "+
		"RsListHttpGetRetryTimes:3, "+
		"RsListHttpGetRetryInterval:100ms, "+
		"EnableRerouting:true, "+
		"MaxConnectionAge:0s, "+
		"EnableSLBLoadBalance:false"+
		"}", cfg.String())
}

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

package route

import (
	"time"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type ObRouteTaskInfo struct {
	refreshTableChan      chan bool
	configServerChan      chan bool
	createConnPoolChan    chan bool
	dropConnPoolChan      chan bool
	tables                *util.ConcurrentMap // The tables need to be refreshed
	createConnPoolServers *util.ConcurrentMap // The servers need to create connection pool
	dropConnPoolServers   *util.ConcurrentMap // The servers need to drop connection pool
	checkRslistTicker     *time.Ticker
}

func NewRouteTaskInfo() *ObRouteTaskInfo {
	return &ObRouteTaskInfo{
		refreshTableChan:      make(chan bool),
		configServerChan:      make(chan bool),
		createConnPoolChan:    make(chan bool),
		dropConnPoolChan:      make(chan bool),
		tables:                util.NewConcurrentMap(),
		createConnPoolServers: util.NewConcurrentMap(),
		dropConnPoolServers:   util.NewConcurrentMap(),
		checkRslistTicker:     time.NewTicker(rslistCheckInterval),
	}
}

func (i *ObRouteTaskInfo) TriggerRefreshTable() {
	if len(i.refreshTableChan) == 0 {
		i.refreshTableChan <- true
	}
}

func (i *ObRouteTaskInfo) TriggerCheckRslist() {
	if len(i.configServerChan) == 0 {
		i.configServerChan <- true
	}
}

func (i *ObRouteTaskInfo) TriggerCreateConnectionPool() {
	if len(i.createConnPoolChan) == 0 {
		i.createConnPoolChan <- true
	}
}

func (i *ObRouteTaskInfo) TriggerDropConnectionPool() {
	if len(i.dropConnPoolChan) == 0 {
		i.dropConnPoolChan <- true
	}
}

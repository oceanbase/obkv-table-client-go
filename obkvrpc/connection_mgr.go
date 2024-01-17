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

package obkvrpc

import (
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/log"
	"sync"
	"time"
)

const DefaultConnectWaitTime = time.Duration(3) * time.Second
const ConnectionMgrTaskInterval = time.Duration(3) * time.Second

type ConnectionMgr struct {
	connLifeCycleMgr *ConnectionLifeCycleMgr
	slbLoader        *SLBLoader
	needStop         chan bool
	wg               sync.WaitGroup
}

func NewConnectionMgr(p *ConnectionPool) *ConnectionMgr {
	connMgr := &ConnectionMgr{
		needStop: make(chan bool),
	}
	if p.option.enableSLBLoadBalance {
		connMgr.slbLoader = NewSLBLoader(p)
		connMgr.slbLoader.refreshSLBList()
	}

	if p.option.maxConnectionAge > 0 || p.option.enableSLBLoadBalance {
		connMgr.connLifeCycleMgr = NewConnectionLifeCycleMgr(p, p.option.maxConnectionAge)
	}
	return connMgr
}

func (c *ConnectionMgr) start() {
	ticker := time.NewTicker(ConnectionMgrTaskInterval)
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-c.needStop:
				ticker.Stop()
				log.Info("Monitor", nil, "Stop ConnectionMgr")
				return
			case <-ticker.C:
				c.run()
			}
		}
	}()
	log.Info("Monitor", nil, "start ConnectionMgr, "+c.String())
}

func (c *ConnectionMgr) run() {
	if c.slbLoader != nil {
		c.slbLoader.run()
	}
	if c.connLifeCycleMgr != nil {
		c.connLifeCycleMgr.run()
	}
}

func (c *ConnectionMgr) close() {
	c.needStop <- true
	c.wg.Wait()
}

func (c *ConnectionMgr) String() string {
	return fmt.Sprintf("ConnectionMgr{connLifeCycleMgr: %s, slbLoader: %s, needStop: %d, wg: %v}",
		c.connLifeCycleMgr, c.slbLoader, c.needStop, c.wg)
}

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
	"context"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/pkg/errors"
	"github.com/scylladb/go-set/strset"
	"go.uber.org/zap"
	"math"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type ConnectionLifeCycleMgr struct {
	connPool         *ConnectionPool
	maxConnectionAge time.Duration
	lastExpireIdx    int
}

const DefaultConnectWaitTime = time.Duration(3) * time.Second
const ConnectionMgrTaskInterval = time.Duration(3) * time.Second

type SLBLoader struct {
	connPool   *ConnectionPool
	dnsAddress string
	round      atomic.Int64
	mutex      sync.RWMutex
	slbAddress []string
}

type ConnectionMgr struct {
	connLifeCycleMgr *ConnectionLifeCycleMgr
	slbLoader        *SLBLoader
	needStop         chan bool
	wg               sync.WaitGroup
}

func NewSLBLoader(p *ConnectionPool) *SLBLoader {
	slbLoader := &SLBLoader{
		slbAddress: make([]string, 0, 10),
		dnsAddress: p.option.ip,
		connPool:   p,
	}
	slbLoader.round.Store(-1)
	return slbLoader
}

// refresh SLB list from DNS address
func (s *SLBLoader) refreshSLBList() (bool, error) {
	ips, err := net.LookupIP(s.dnsAddress)
	if err != nil {
		return false, errors.WithMessagef(err, "fail to look up slb address, dns addr: %s", s.dnsAddress)
	}
	slbAddress := strset.NewWithSize(len(ips))
	for _, ip := range ips {
		slbAddress.Add(ip.String())
	}
	changed := !slbAddress.IsEqual(strset.New(s.slbAddress...))
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if changed {
		log.Info(fmt.Sprint("SLB address changed, before: ", s.slbAddress, ", after: ", slbAddress))
		s.slbAddress = slbAddress.List()
	}
	return changed, nil
}

// round-robin get next slb address from slb list
func (s *SLBLoader) getNextSLBAddress() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	slbNum := len(s.slbAddress)
	if slbNum > 0 {
		slbAddr := s.slbAddress[(s.round.Add(1))%(int64(slbNum))]
		return slbAddr
	}
	return s.dnsAddress
}

// refresh SLBList and refresh connection expire time if SLBList changed
func (s *SLBLoader) run() {
	changed, err := s.refreshSLBList()
	if err != nil {
		log.Warn("reconnect failed", zap.Error(err))
		return
	}
	if changed {
		s.refreshConnectionLife()
	}
}

func (s *SLBLoader) refreshConnectionLife() {
	for _, conn := range s.connPool.connections {
		conn.expireTime = time.Now()
	}
}

func (s *SLBLoader) toString() string {
	return fmt.Sprintf("%#v", s)
}

// check and reconnect timeout connections
func (c *ConnectionLifeCycleMgr) run() {
	if c.connPool == nil {
		log.Error("connection pool is null")
		return
	}

	// 1. get all timeout connections
	expiredConnIds := make([]int, 0, len(c.connPool.connections))
	for i := 1; i <= len(c.connPool.connections); i++ {
		connection := c.connPool.connections[(i+c.lastExpireIdx)%(len(c.connPool.connections))]
		if !connection.expireTime.IsZero() && connection.expireTime.Before(time.Now()) {
			expiredConnIds = append(expiredConnIds, (i+c.lastExpireIdx)%(len(c.connPool.connections)))
		}
	}

	if len(expiredConnIds) > 0 {
		log.Info(fmt.Sprintf("Find %d expired connections", len(expiredConnIds)))
		for idx, connIdx := range expiredConnIds {
			log.Info(fmt.Sprintf("%d: ip=%s, port=%d", idx, c.connPool.connections[connIdx].option.ip, c.connPool.connections[connIdx].option.port))
		}
	}

	// 2. mark 30% expired connections as expired
	maxReconnIdx := int(math.Ceil(float64(len(expiredConnIds)) / 3))
	if maxReconnIdx > 0 {
		c.lastExpireIdx = expiredConnIds[maxReconnIdx-1]
		log.Info(fmt.Sprintf("Begin to refresh expired connections which idx less than %d", maxReconnIdx))
	}
	for i := 0; i < maxReconnIdx; i++ {
		// no one can get expired connection
		c.connPool.connections[expiredConnIds[i]].isExpired.Store(true)
	}
	defer func() {
		for i := 0; i < maxReconnIdx; i++ {
			c.connPool.connections[expiredConnIds[i]].isExpired.Store(false)
		}
	}()

	// 3. wait all expired connection finished
	time.Sleep(DefaultConnectWaitTime)
	for i := 0; i < maxReconnIdx; i++ {
		pool := c.connPool.connections
		idx := expiredConnIds[i]
		for j := 0; len(pool[idx].pending) > 0; j++ {
			time.Sleep(time.Duration(10) * time.Millisecond)
			if j > 0 && j%100 == 0 {
				log.Info(fmt.Sprintf("Wait too long time for the connection to end,"+
					"connection idx: %d, ip:%s, port:%d, current connection pending size: %d",
					idx, pool[idx].option.ip, pool[idx].option.port, len(pool[idx].pending)))
			}

			if j > 3000 {
				log.Warn("Wait too much time for the connection to end, stop ConnectionLifeCycleMgr")
				return
			}
		}
	}

	// 4. close and reconnect all expired connections
	ctx, _ := context.WithTimeout(context.Background(), c.connPool.option.connectTimeout)
	for i := 0; i < maxReconnIdx; i++ {
		// close and reconnect
		c.connPool.connections[expiredConnIds[i]].Close()
		_, err := c.connPool.RecreateConnection(ctx, expiredConnIds[i])
		if err != nil {
			log.Warn("reconnect failed", zap.Error(err))
			return
		}
	}
	if maxReconnIdx > 0 {
		log.Info(fmt.Sprintf("Finish to refresh expired connections which idx less than %d", maxReconnIdx))
	}
}

func (s *ConnectionLifeCycleMgr) toString() string {
	return fmt.Sprintf("%#v", s)
}

func NewConnectionLifeCycleMgr(connPool *ConnectionPool, maxConnectionAge time.Duration) *ConnectionLifeCycleMgr {
	connLifeCycleMgr := &ConnectionLifeCycleMgr{
		connPool:         connPool,
		maxConnectionAge: maxConnectionAge,
		lastExpireIdx:    0,
	}
	return connLifeCycleMgr
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
	log.Info("start ConnectionMgr, " + c.slbLoader.toString() + c.connLifeCycleMgr.toString())
	go func() {
		c.wg.Add(1)
		defer c.wg.Done()
		for {
			select {
			case <-c.needStop:
				ticker.Stop()
				log.Info("Stop ConnectionMgr")
				return
			case <-ticker.C:
				c.run()
			}
		}
	}()
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

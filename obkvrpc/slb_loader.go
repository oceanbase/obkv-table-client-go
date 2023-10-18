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
	"github.com/pkg/errors"
	"github.com/scylladb/go-set/strset"
	"go.uber.org/zap"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type SLBLoader struct {
	connPool   *ConnectionPool
	dnsAddress string
	round      atomic.Int64
	mutex      sync.RWMutex
	slbAddress []string
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

func (s *SLBLoader) String() string {
	return fmt.Sprintf("SLBLoader{connPool: %s, dnsAddress: %s, round: %d, slbAddress: %v}",
		s.connPool, s.dnsAddress, s.round, s.slbAddress)
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

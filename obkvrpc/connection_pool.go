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

package obkvrpc

import (
	"context"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/scylladb/go-set/strset"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type PoolOption struct {
	ip                  string
	port                int
	connPoolMaxConnSize int
	connectTimeout      time.Duration
	loginTimeout        time.Duration

	tenantName   string
	databaseName string
	userName     string
	password     string

	maxConnectionAge     time.Duration
	enableSLBLoadBalance bool
}

type ConnectionPool struct {
	option *PoolOption

	connections []*Connection
	rwMutexes   []sync.RWMutex
	connMgr     *ConnectionMgr
}

type ConnectionLifeCycleMgr struct {
	connPool         *ConnectionPool
	maxConnectionAge time.Duration
	lastExpireIdx    int
}

const DefaultConnectWaitTime = time.Duration(3) * time.Second

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
}

func NewPoolOption(ip string, port int, connPoolMaxConnSize int, connectTimeout time.Duration, loginTimeout time.Duration,
	tenantName string, databaseName string, userName string, password string, maxConnectionAge time.Duration, enableSLBLoadBalance bool) *PoolOption {
	return &PoolOption{
		ip:                   ip,
		port:                 port,
		connPoolMaxConnSize:  connPoolMaxConnSize,
		connectTimeout:       connectTimeout,
		loginTimeout:         loginTimeout,
		tenantName:           tenantName,
		databaseName:         databaseName,
		userName:             userName,
		password:             password,
		maxConnectionAge:     maxConnectionAge,
		enableSLBLoadBalance: enableSLBLoadBalance,
	}
}

func NewConnectionPool(option *PoolOption) (*ConnectionPool, error) {
	pool := &ConnectionPool{
		option:      option,
		connections: make([]*Connection, 0, option.connPoolMaxConnSize),
		rwMutexes:   make([]sync.RWMutex, 0, option.connPoolMaxConnSize),
	}

	if option.maxConnectionAge > 0 || option.enableSLBLoadBalance {
		pool.connMgr = NewConnectionMgr(pool)
	}

	for i := 0; i < pool.option.connPoolMaxConnSize; i++ {
		ctx, _ := context.WithTimeout(context.Background(), pool.option.connectTimeout)
		connection, err := pool.CreateConnection(ctx)
		if err != nil {
			return nil, errors.WithMessage(err, "create connection")
		}

		pool.connections = append(pool.connections, connection)
		pool.rwMutexes = append(pool.rwMutexes, sync.RWMutex{})

	}

	if pool.connMgr != nil {
		pool.connMgr.start()
	}

	return pool, nil
}

// GetConnection Find an unexpired and active connection to use
// In theory, all connection won't expire at the same time
func (p *ConnectionPool) GetConnection() (*Connection, int) {
	index := rand.Intn(p.option.connPoolMaxConnSize)
	for i := 0; i < p.option.connPoolMaxConnSize; i++ {
		if p.connections[(index+i)%p.option.connPoolMaxConnSize].isExpired.Load() == false {
			index = (index + i) % p.option.connPoolMaxConnSize
		}
	}

	p.rwMutexes[index].RLock()
	defer p.rwMutexes[index].RUnlock()

	if p.connections[index].active.Load() {
		return p.connections[index], 0
	}

	return nil, index
}

// RecreateConnection recreate the connection and login
func (p *ConnectionPool) RecreateConnection(ctx context.Context, connectionIdx int) (*Connection, error) {
	p.rwMutexes[connectionIdx].Lock()
	defer p.rwMutexes[connectionIdx].Unlock()

	if p.connections[connectionIdx].active.Load() {
		return p.connections[connectionIdx], nil
	}

	connection, err := p.CreateConnection(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "create connection")
	}

	p.connections[connectionIdx] = connection

	return p.connections[connectionIdx], nil
}

func (p *ConnectionPool) CreateConnection(ctx context.Context) (*Connection, error) {
	ip, port := p.getNextConnAddress()
	connectionOption := NewConnectionOption(ip, port, p.option.connectTimeout, p.option.loginTimeout,
		p.option.tenantName, p.option.databaseName, p.option.userName, p.option.password)
	connection := NewConnection(connectionOption)
	err := connection.Connect(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "connection connect")
	}
	err = connection.Login(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "connection login")
	}
	// put it to here to ensure connection should not expire during connect & login phase
	if p.option.maxConnectionAge > 0 {
		connection.expireTime = time.Now().Add(p.option.maxConnectionAge)
	}
	log.Info(fmt.Sprintf("connect success, remote addr:%s", connection.conn.RemoteAddr().String()))
	return connection, nil
}

func (p *ConnectionPool) Close() {
	for _, connection := range p.connections {
		connection.Close()
	}

	if p.connMgr != nil {
		p.connMgr.close()
	}
}

func (p *ConnectionPool) getNextConnAddress() (string, int) {
	ip := p.option.ip
	port := p.option.port
	if p.connMgr != nil && p.connMgr.slbLoader != nil {
		ip = p.connMgr.slbLoader.getNextSLBAddress()
		log.Info(fmt.Sprintf("Get a SLB address %s:%d", ip, port))
	}

	return ip, port
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
	//for idx, connection := range c.connPool.connections {
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
			if j > 0 && j%10000 == 0 {
				log.Info(fmt.Sprintf("Wait too long time for the connection to end,"+
					"connection idx: %d, ip:%s, port:%d, current connection pending size: %d",
					idx, pool[idx].option.ip, pool[idx].option.port, len(pool[idx].pending)))
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
	ticker := time.NewTicker(time.Duration(5) * time.Millisecond)
	log.Info("start ConnectionMgr, " + c.slbLoader.toString() + c.connLifeCycleMgr.toString())
	go func() {
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
}

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
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/oceanbase/obkv-table-client-go/log"

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

	connections       []*Connection
	rwMutexes         []sync.RWMutex
	connMgr           *ConnectionMgr
	disConnectedCount atomic.Int32
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
		option:            option,
		connections:       make([]*Connection, 0, option.connPoolMaxConnSize),
		rwMutexes:         make([]sync.RWMutex, 0, option.connPoolMaxConnSize),
		disConnectedCount: atomic.Int32{},
	}
	pool.disConnectedCount.Store(0)

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

func (p *ConnectionPool) IsDisconnected() bool {
	return p.disConnectedCount.Load() == int32(p.option.connPoolMaxConnSize)
}

// GetConnection Find an unexpired and active connection to use
// In theory, all connection won't expire at the same time
func (p *ConnectionPool) GetConnection() (*Connection, int) {
	index := rand.Intn(p.option.connPoolMaxConnSize)
	for i := 0; i < p.option.connPoolMaxConnSize; i++ {
		if p.connections[(index+i)%p.option.connPoolMaxConnSize].isExpired.Load() == false {
			index = (index + i) % p.option.connPoolMaxConnSize
			break
		} else if i == p.option.connPoolMaxConnSize-1 {
			log.Warn("All connections is expired, will pick a expired connection")
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

	p.disConnectedCount.Add(-1)

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

	// listen receive channel
	go func() {
		connection.receivePacket()
		p.disConnectedCount.Add(1) // receivePacket() break means disconnected
	}()

	// listen send channel
	go func() {
		connection.sendPacket()
	}()

	err = connection.Login(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "connection login")
	}
	// put it to here to ensure connection should not expire during connect & login phase
	if p.option.maxConnectionAge > 0 {
		connection.expireTime = time.Now().Add(p.option.maxConnectionAge)
	}
	log.Info(fmt.Sprintf("connect success, remote addr:%s, expire time: %s",
		connection.conn.RemoteAddr().String(), connection.expireTime.String()))
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

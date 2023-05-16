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
	"math/rand"
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
}

type ConnectionPool struct {
	option *PoolOption

	connections []*Connection
	rwMutexes   []sync.RWMutex
}

func NewPoolOption(ip string, port int, connPoolMaxConnSize int, connectTimeout time.Duration, loginTimeout time.Duration,
	tenantName string, databaseName string, userName string, password string) *PoolOption {
	return &PoolOption{
		ip:                  ip,
		port:                port,
		connPoolMaxConnSize: connPoolMaxConnSize,
		connectTimeout:      connectTimeout,
		loginTimeout:        loginTimeout,
		tenantName:          tenantName,
		databaseName:        databaseName,
		userName:            userName,
		password:            password,
	}
}

func NewConnectionPool(option *PoolOption) (*ConnectionPool, error) {
	pool := &ConnectionPool{
		option:      option,
		connections: make([]*Connection, 0, option.connPoolMaxConnSize),
		rwMutexes:   make([]sync.RWMutex, 0, option.connPoolMaxConnSize),
	}

	connectionOption := NewConnectionOption(pool.option.ip, pool.option.port, pool.option.connectTimeout, pool.option.loginTimeout,
		pool.option.tenantName, pool.option.databaseName, pool.option.userName, pool.option.password)

	for i := 0; i < pool.option.connPoolMaxConnSize; i++ {

		connection := NewConnection(connectionOption)

		ctx, _ := context.WithTimeout(context.Background(), pool.option.connectTimeout)
		err := connection.Connect(ctx)
		if err != nil {
			return nil, errors.WithMessage(err, "connection connect")
		}

		ctx, _ = context.WithTimeout(context.Background(), pool.option.loginTimeout)
		err = connection.Login(ctx)
		if err != nil {
			return nil, errors.WithMessage(err, "connection login")
		}

		pool.connections = append(pool.connections, connection)
		pool.rwMutexes = append(pool.rwMutexes, sync.RWMutex{})

	}

	return pool, nil
}

func (p *ConnectionPool) GetConnection() (*Connection, int) {
	index := rand.Intn(p.option.connPoolMaxConnSize)

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
	connectionOption := NewConnectionOption(p.option.ip, p.option.port, p.option.connectTimeout, p.option.loginTimeout,
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
	return connection, nil
}

func (p *ConnectionPool) Close() {
	for _, connection := range p.connections {
		connection.Close()
	}
}

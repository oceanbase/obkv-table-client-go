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

func (p *ConnectionPool) GetConnection(ctx context.Context) (*Connection, error) {
	randomIndex := rand.Intn(len(p.connections))

	p.rwMutexes[randomIndex].RLock()
	if p.connections[randomIndex].active.Load() {
		p.rwMutexes[randomIndex].RUnlock()
		return p.connections[randomIndex], nil
	}
	p.rwMutexes[randomIndex].RUnlock()

	p.rwMutexes[randomIndex].Lock()
	if p.connections[randomIndex].active.Load() {
		p.rwMutexes[randomIndex].Unlock()
		return p.connections[randomIndex], nil
	}
	// Recreate the connection and login
	connection, err := p.CreateConnection(ctx)
	if err != nil {
		p.rwMutexes[randomIndex].Unlock()
		return nil, errors.WithMessage(err, "recreate connection")
	}

	p.connections[randomIndex] = connection

	p.rwMutexes[randomIndex].Unlock()
	return p.connections[randomIndex], nil
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

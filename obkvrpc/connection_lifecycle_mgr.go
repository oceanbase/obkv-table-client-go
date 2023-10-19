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
	"go.uber.org/zap"
	"math"
	"time"
)

type ConnectionLifeCycleMgr struct {
	connPool         *ConnectionPool
	maxConnectionAge time.Duration
	lastExpireIdx    int
}

func (s *ConnectionLifeCycleMgr) String() string {
	return fmt.Sprintf("ConnectionLifeCycleMgr{connPool: %p, maxConnectionAge: %d,lastExpireIdx: %d}",
		s.connPool, s.maxConnectionAge, s.lastExpireIdx)
}

func NewConnectionLifeCycleMgr(connPool *ConnectionPool, maxConnectionAge time.Duration) *ConnectionLifeCycleMgr {
	connLifeCycleMgr := &ConnectionLifeCycleMgr{
		connPool:         connPool,
		maxConnectionAge: maxConnectionAge,
		lastExpireIdx:    0,
	}
	return connLifeCycleMgr
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

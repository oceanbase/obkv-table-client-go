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
	"sync"
	"time"
)

type ObTableRoster struct {
	tenantName     string
	userName       string
	password       string
	database       string
	connPoolSize   int
	connectTimeout time.Duration
	loginTimeout   time.Duration
	m              sync.Map
}

func (r *ObTableRoster) Add(addr tcpAddr, table *ObTable) {
	r.m.Store(addr, table)
}

func (r *ObTableRoster) Get(addr tcpAddr) (*ObTable, bool) {
	t, ok := r.m.Load(addr)
	if ok {
		return t.(*ObTable), ok
	}

	return nil, ok
}

func (r *ObTableRoster) Delete(addr tcpAddr) {
	r.m.Delete(addr)
}

func (r *ObTableRoster) LoadAndDelete(addr tcpAddr) (interface{}, bool) {
	return r.m.LoadAndDelete(addr)
}

func (r *ObTableRoster) Close() {
	r.m.Range(func(k, v interface{}) bool {
		r.Delete(k.(tcpAddr))
		t := v.(*ObTable)
		if t != nil {
			t.Close()
		}
		return true
	})
}

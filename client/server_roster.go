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

package client

import (
	"math/rand"
	"strconv"
	"sync/atomic"

	"github.com/oceanbase/obkv-table-client-go/route"
)

type obServerRoster struct {
	maxPriority atomic.Int32
	roster      []*route.ObServerAddr
}

func (r *obServerRoster) MaxPriority() int32 {
	return r.maxPriority.Load()
}

func (r *obServerRoster) Reset(servers []*route.ObServerAddr) {
	r.maxPriority.Store(0)
	r.roster = servers
}

func (r *obServerRoster) GetServer() *route.ObServerAddr {
	idx := rand.Intn(len(r.roster))
	return r.roster[idx]
}

func (r *obServerRoster) Size() int {
	return len(r.roster)
}

func (r *obServerRoster) String() string {
	var rostersStr string
	rostersStr = rostersStr + "["
	for i := 0; i < len(r.roster); i++ {
		if i > 0 {
			rostersStr += ", "
		}
		if r.roster[i] != nil {
			rostersStr += r.roster[i].String()
		} else {
			rostersStr += "nil"
		}
	}
	rostersStr += "]"
	return "obServerRoster{" +
		"maxPriority:" + strconv.Itoa(int(r.maxPriority.Load())) + ", " +
		"roster:" + rostersStr +
		"}"
}

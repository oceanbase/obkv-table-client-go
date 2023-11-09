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
	"math/rand"
)

type ObServerRoster struct {
	servers []*ObServerAddr
}

func (r *ObServerRoster) GetServer() *ObServerAddr {
	idx := rand.Intn(len(r.servers))
	return r.servers[idx]
}

func (r *ObServerRoster) Size() int {
	return len(r.servers)
}

func (r *ObServerRoster) String() string {
	var rostersStr string
	rostersStr = rostersStr + "["
	for i := 0; i < len(r.servers); i++ {
		if i > 0 {
			rostersStr += ", "
		}
		if r.servers[i] != nil {
			rostersStr += r.servers[i].String()
		} else {
			rostersStr += "nil"
		}
	}
	rostersStr += "]"
	return "ObServerRoster{" +
		"roster:" + rostersStr +
		"}"
}

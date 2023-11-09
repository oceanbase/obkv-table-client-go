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
	"testing"
)

func TestObConfigServerInfo_GetServerAddressRandomly(t *testing.T) {
	server1 := NewObServerAddr("127.0.0.1", 1, 1)
	server2 := NewObServerAddr("127.0.0.1", 1, 2)
	server3 := NewObServerAddr("127.0.0.1", 1, 3)
	server4 := NewObServerAddr("127.0.0.1", 1, 4)
	server5 := NewObServerAddr("127.0.0.1", 1, 5)
	servers := []*ObServerAddr{server1, server2, server3, server4, server5}
	info := NewConfigServerInfo()
	for _, server := range servers {
		info.rslist.Append(server)
	}
	for i := 0; i < 10; i++ {
		svr := info.GetServerAddressRandomly()
		println(svr.String())
	}
}

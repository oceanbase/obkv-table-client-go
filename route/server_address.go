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

import "strconv"

type ObServerAddr struct {
	ip      string
	sqlPort int
	svrPort int
}

func (a *ObServerAddr) SvrPort() int {
	return a.svrPort
}

func (a *ObServerAddr) Ip() string {
	return a.ip
}

func NewObServerAddr(ip string, sqlPort int, svrPort int) *ObServerAddr {
	return &ObServerAddr{ip, sqlPort, svrPort}
}

func (a *ObServerAddr) String() string {
	return "ObServerAddr{" +
		"ip:" + a.ip + ", " +
		"sqlPort:" + strconv.Itoa(a.sqlPort) + ", " +
		"svrPort:" + strconv.Itoa(a.svrPort) +
		"}"
}

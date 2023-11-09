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

type tcpAddr struct {
	ip   string
	port int
}

func (a *tcpAddr) Equal(other *tcpAddr) bool {
	return a.ip == other.ip && a.port == other.port
}

func (a *tcpAddr) String() string {
	return "tcpAddr{" +
		"ip:" + a.ip + ", " +
		"port:" + strconv.Itoa(a.port) +
		"}"
}

type ObServerAddr struct {
	tcpAddr
	sqlPort int
}

func (a *ObServerAddr) SvrPort() int {
	return a.port
}

func (a *ObServerAddr) Ip() string {
	return a.ip
}

func (a *ObServerAddr) Equal(other *ObServerAddr) bool {
	return a.tcpAddr.Equal(&other.tcpAddr) && a.sqlPort == other.sqlPort
}

func NewObServerAddr(ip string, sqlPort int, svrPort int) *ObServerAddr {
	return &ObServerAddr{
		tcpAddr{ip, svrPort},
		sqlPort}
}

func (a *ObServerAddr) String() string {
	return "ObServerAddr{" +
		"ip:" + a.ip + ", " +
		"sqlPort:" + strconv.Itoa(a.sqlPort) + ", " +
		"svrPort:" + strconv.Itoa(a.port) +
		"}"
}

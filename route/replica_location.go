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

type obServerRole int

const (
	serverRoleInvalid  obServerRole = -1
	serverRoleLeader   obServerRole = 1
	serverRoleFollower obServerRole = 2
)

type obReplicaType int

const (
	replicaTypeInvalid  obReplicaType = -1
	replicaTypeFull     obReplicaType = 0
	replicaTypeLogOnly  obReplicaType = 5
	replicaTypeReadOnly obReplicaType = 16
)

type obReplicaLocation struct {
	addr        *ObServerAddr
	svrStatus   *obServerStatus
	role        obServerRole
	replicaType obReplicaType
}

func (l *obReplicaLocation) SvrStatus() *obServerStatus {
	return l.svrStatus
}

func newReplicaLocation(addr *ObServerAddr, svrStatus *obServerStatus, role obServerRole, replicaType obReplicaType) *obReplicaLocation {
	return &obReplicaLocation{addr, svrStatus, role, replicaType}
}

func (l *obReplicaLocation) Addr() *ObServerAddr {
	return l.addr
}

func (l *obReplicaLocation) isValid() bool {
	return (l.addr != nil) && (l.role != serverRoleInvalid) && (l.svrStatus != nil) &&
		l.svrStatus.IsActive() && (l.replicaType != replicaTypeInvalid)
}

func (l *obReplicaLocation) isLeader() bool {
	return l.role == serverRoleLeader
}

func (l *obReplicaLocation) String() string {
	var addrStr string
	if l.addr == nil {
		addrStr = "nil"
	} else {
		addrStr = l.addr.String()
	}
	var statusStr string
	if l.svrStatus == nil {
		statusStr = "nil"
	} else {
		statusStr = l.svrStatus.String()
	}
	return "obReplicaLocation{" +
		"addr:" + addrStr + ", " +
		"info:" + statusStr + ", " +
		"role:" + strconv.Itoa(int(l.role)) + ", " +
		"replicaType:" + strconv.Itoa(int(l.replicaType)) +
		"}"
}

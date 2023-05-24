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

	"github.com/stretchr/testify/assert"
)

func TestObPartLocationEntry(t *testing.T) {
	e := newObPartLocationEntry(1)
	assert.Equal(t, "ObPartLocationEntry{partLocations:{}}", e.String())
	s1 := NewObServerAddr("127.0.0.1", 2001, 2000)
	s2 := NewObServerAddr("127.0.0.1", 2001, 2000)
	st := newServerStatus(0, "active")
	r1 := newReplicaLocation(s1, st, serverRoleLeader, replicaTypeFull)
	r2 := newReplicaLocation(s2, st, serverRoleFollower, replicaTypeFull)
	l := &obPartitionLocation{}
	l.addReplicaLocation(r1)
	l.addReplicaLocation(r2)
	e.partLocations[0] = l
	assert.Equal(t, "ObPartLocationEntry{partLocations:{m[0]=obPartitionLocation{leader:obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:2001, svrPort:2000}, info:obServerStatus{stopTime:0, status:active}, role:1, replicaType:0}, replicas:[obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:2001, svrPort:2000}, info:obServerStatus{stopTime:0, status:active}, role:1, replicaType:0}, obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:2001, svrPort:2000}, info:obServerStatus{stopTime:0, status:active}, role:2, replicaType:0}]}}}", e.String())
}

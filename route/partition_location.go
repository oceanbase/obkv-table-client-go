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

type ObConsistency int8

const (
	ConsistencyStrong ObConsistency = 0
	ConsistencyWeak   ObConsistency = 1
)

type obPartitionLocation struct {
	leader   obReplicaLocation
	replicas []obReplicaLocation
}

func (l *obPartitionLocation) addReplicaLocation(replica *obReplicaLocation) {
	if replica.isLeader() {
		l.leader = *replica
	}
	l.replicas = append(l.replicas, *replica)
}

func (l *obPartitionLocation) getReplica(consistency ObConsistency) *obReplicaLocation {
	if consistency == ConsistencyStrong {
		return &l.leader
	}
	return &l.leader
}

func (l *obPartitionLocation) String() string {
	var replicasStr string
	replicasStr = replicasStr + "["
	for i := 0; i < len(l.replicas); i++ {
		if i > 0 {
			replicasStr += ", "
		}
		replicasStr += l.replicas[i].String()
	}
	replicasStr += "]"
	return "obPartitionLocation{" +
		"leader:" + l.leader.String() + ", " +
		"replicas:" + replicasStr +
		"}"
}

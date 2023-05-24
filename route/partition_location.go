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

// obPartitionLocation indicates the location of a partition,
// including a leader replica and other read replicas
type obPartitionLocation struct {
	leader   *obReplicaLocation
	replicas []*obReplicaLocation
}

// addReplicaLocation add a replica to obPartitionLocation
func (l *obPartitionLocation) addReplicaLocation(replica *obReplicaLocation) {
	if replica.isLeader() {
		l.leader = replica
	}
	l.replicas = append(l.replicas, replica)
}

// getReplica get the copy according to the consistency requirements you need,
// strong consistency get the leader replica, weak consistency get the read replica (todo).
func (l *obPartitionLocation) getReplica(consistency ObConsistency) *obReplicaLocation {
	if consistency == ConsistencyStrong {
		return l.leader
	}
	// todo: get read replica
	return l.leader
}

func (l *obPartitionLocation) String() string {
	var leaderStr = "nil"
	if l.leader != nil {
		leaderStr = l.leader.String()
	}

	var replicasStr string
	replicasStr = replicasStr + "["
	for i := 0; i < len(l.replicas); i++ {
		if i > 0 {
			replicasStr += ", "
		}
		str := "nil"
		if l.replicas[i] != nil {
			str = l.replicas[i].String()
		}
		replicasStr += str
	}
	replicasStr += "]"
	return "obPartitionLocation{" +
		"leader:" + leaderStr + ", " +
		"replicas:" + replicasStr +
		"}"
}

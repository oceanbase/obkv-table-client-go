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

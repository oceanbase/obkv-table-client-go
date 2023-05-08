package route

type ObPartitionLocation struct {
	leader   ObReplicaLocation
	replicas []ObReplicaLocation
}

func (l *ObPartitionLocation) addReplicaLocation(replica *ObReplicaLocation) {
	if replica.isLeader() {
		l.leader = *replica
	}
	l.replicas = append(l.replicas, *replica)
}

func (l *ObPartitionLocation) getReplica(route *ObServerRoute) *ObReplicaLocation {
	if route.readConsistency == ObReadConsistencyStrong {
		return &l.leader
	}
	// todo:weak read by LDC
	return &l.leader
}

func (l *ObPartitionLocation) String() string {
	var replicasStr string
	replicasStr = replicasStr + "["
	for i := 0; i < len(l.replicas); i++ {
		if i > 0 {
			replicasStr += ", "
		}
		replicasStr += l.replicas[i].String()
	}
	replicasStr += "]"
	return "ObPartitionLocation{" +
		"leader:" + l.leader.String() + ", " +
		"replicas:" + replicasStr +
		"}"
}

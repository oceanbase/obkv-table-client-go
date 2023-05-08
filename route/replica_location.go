package route

type ObReplicaLocation struct {
	addr        ObServerAddr
	info        ObServerInfo
	role        ObServerRole
	replicaType ObReplicaType
}

func (l *ObReplicaLocation) Info() ObServerInfo {
	return l.info
}

func (l *ObReplicaLocation) Addr() *ObServerAddr {
	return &l.addr
}

func (l *ObReplicaLocation) isValid() bool {
	return !l.role.isInvalid() && l.info.IsActive()
}

func (l *ObReplicaLocation) isLeader() bool {
	return l.role.isLeader()
}

func (l *ObReplicaLocation) String() string {
	return "ObReplicaLocation{" +
		"addr:" + l.addr.String() + ", " +
		"info:" + l.info.String() + ", " +
		"role:" + l.role.String() + ", " +
		"replicaType:" + l.replicaType.String() +
		"}"
}

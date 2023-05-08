package route

type ObTableLocation struct {
	replicaLocations []ObReplicaLocation
}

func (l *ObTableLocation) ReplicaLocations() []ObReplicaLocation {
	return l.replicaLocations
}

func (l *ObTableLocation) String() string {
	var replicaLocationsStr string
	replicaLocationsStr = replicaLocationsStr + "["
	for i := 0; i < len(l.replicaLocations); i++ {
		if i > 0 {
			replicaLocationsStr += ", "
		}
		replicaLocationsStr += l.replicaLocations[i].String()
	}
	replicaLocationsStr += "]"
	return "ObTableLocation{" +
		"replicaLocations:" + replicaLocationsStr +
		"}"
}

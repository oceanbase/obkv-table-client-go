package route

import "strconv"

// ObReplicaType name
const (
	ReplicaTypeFull     = "FULL"
	ReplicaTypeLogOnly  = "LOGONLY"
	ReplicaTypeReadOnly = "READONLY"
	ReplicaTypeInvalid  = "INVALID"
)

// ObReplicaType index
const (
	ReplicaTypeInvalidIndex  = -1
	ReplicaTypeFullIndex     = 0
	ReplicaTypeLogOnlyIndex  = 5
	ReplicaTypeReadOnlyIndex = 16
)

type ObReplicaType struct {
	name  string
	index int
}

func newObReplicaType(index int) ObReplicaType {
	if index == ReplicaTypeFullIndex {
		return ObReplicaType{ReplicaTypeFull, ReplicaTypeFullIndex}
	} else if index == ReplicaTypeLogOnlyIndex {
		return ObReplicaType{ReplicaTypeLogOnly, ReplicaTypeLogOnlyIndex}
	} else if index == ReplicaTypeReadOnlyIndex {
		return ObReplicaType{ReplicaTypeReadOnly, ReplicaTypeReadOnlyIndex}
	} else {
		return ObReplicaType{ReplicaTypeInvalid, ReplicaTypeInvalidIndex}
	}
}

func (r *ObReplicaType) String() string {
	return "ObReplicaType{" +
		"name:" + r.name + ", " +
		"index:" + strconv.Itoa(r.index) +
		"}"
}

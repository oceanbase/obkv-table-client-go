package route

import (
	"strconv"
	"strings"
)

type ObServerAddr struct {
	ip      string
	sqlPort int
	svrPort int
}

func (a *ObServerAddr) ToString() string {
	return "ObServerAddr{" +
		"ip:" + a.ip + ", " +
		"sqlPort:" + strconv.Itoa(a.sqlPort) + ", " +
		"svrPort:" + strconv.Itoa(a.svrPort) +
		"}"
}

type ObServerInfo struct {
	stopTime int64
	status   string // Active/InActive/Deleting
}

func (i *ObServerInfo) isActive() bool {
	return i.stopTime == 0 && strings.EqualFold(i.status, "active") // ignore case
}

func (i *ObServerInfo) ToString() string {
	return "ObServerInfo{" +
		"stopTime:" + strconv.Itoa(int(i.stopTime)) + ", " +
		"status:" + i.status +
		"}"
}

// ObServerRole name
const (
	ServerRoleInvalid  = "INVALID_ROLE"
	ServerRoleLeader   = "LEADER"
	ServerRoleFollower = "FOLLOWER"
)

// ObServerRole index
const (
	ServerRoleInvalidIndex  = -1
	ServerRoleLeaderIndex   = 1
	ServerRoleFollowerIndex = 2
)

type ObServerRole struct {
	name  string
	index int
}

func (r *ObServerRole) isInvalid() bool {
	return r.name == ServerRoleInvalid && r.index == ServerRoleInvalidIndex
}

func (r *ObServerRole) isLeader() bool {
	return r.name == ServerRoleLeader && r.index == ServerRoleLeaderIndex
}

func newObServerRole(index int) ObServerRole {
	if index == ServerRoleLeaderIndex {
		return ObServerRole{ServerRoleLeader, ServerRoleLeaderIndex}
	} else if index == ServerRoleFollowerIndex {
		return ObServerRole{ServerRoleFollower, ServerRoleFollowerIndex}
	} else {
		return ObServerRole{ServerRoleInvalid, ServerRoleInvalidIndex}
	}
}

func (r *ObServerRole) ToString() string {
	return "ObServerRole{" +
		"name:" + r.name + ", " +
		"index:" + strconv.Itoa(r.index) +
		"}"
}

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

func (r *ObReplicaType) ToString() string {
	return "ObReplicaType{" +
		"name:" + r.name + ", " +
		"index:" + strconv.Itoa(r.index) +
		"}"
}

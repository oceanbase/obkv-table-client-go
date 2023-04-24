package route

import (
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
)

type ObServerAddr struct {
	ip      string
	sqlPort int
	svrPort int
}

func (a *ObServerAddr) SvrPort() int {
	return a.svrPort
}

func (a *ObServerAddr) Ip() string {
	return a.ip
}

func NewObServerAddr(ip string, sqlPort int, svrPort int) *ObServerAddr {
	return &ObServerAddr{ip, sqlPort, svrPort}
}

func (a *ObServerAddr) String() string {
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

func (i *ObServerInfo) IsActive() bool {
	return i.stopTime == 0 && strings.EqualFold(i.status, "active") // ignore case
}

func (i *ObServerInfo) String() string {
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

func (r *ObServerRole) String() string {
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

func (r *ObReplicaType) String() string {
	return "ObReplicaType{" +
		"name:" + r.name + ", " +
		"index:" + strconv.Itoa(r.index) +
		"}"
}

type ObServerRoster struct {
	maxPriority atomic.Int32
	roster      []*ObServerAddr
	// todo: serverLdc
}

func (r *ObServerRoster) MaxPriority() int32 {
	return r.maxPriority.Load()
}

func (r *ObServerRoster) Reset(servers []*ObServerAddr) {
	r.maxPriority.Store(0)
	r.roster = servers
}

func (r *ObServerRoster) GetServer() *ObServerAddr {
	idx := rand.Intn(len(r.roster))
	return r.roster[idx]
}

func (r *ObServerRoster) Size() int {
	return len(r.roster)
}

func (r *ObServerRoster) String() string {
	var rostersStr string
	rostersStr = rostersStr + "["
	for i := 0; i < len(r.roster); i++ {
		if i > 0 {
			rostersStr += ", "
		}
		if r.roster[i] != nil {
			rostersStr += r.roster[i].String()
		} else {
			rostersStr += "nil"
		}
	}
	rostersStr += "]"
	return "ObServerRoster{" +
		"maxPriority:" + strconv.Itoa(int(r.maxPriority.Load())) + ", " +
		"roster:" + rostersStr +
		"}"
}

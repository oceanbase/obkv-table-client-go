package route

import "strconv"

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

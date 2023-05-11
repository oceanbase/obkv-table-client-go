package route

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObReplicaLocation_isValid(t *testing.T) {
	addr := NewObServerAddr("127.0.0.1", 1222, 12223)
	activeStatus := newServerStatus(0, "Active")
	inactiveStatus := newServerStatus(1, "InActive")
	repLoc := &obReplicaLocation{}
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(nil, nil, serverRoleInvalid, replicaTypeInvalid)
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, nil, serverRoleInvalid, replicaTypeInvalid)
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleInvalid, replicaTypeInvalid)
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleLeader, replicaTypeInvalid)
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleLeader, replicaTypeInvalid)
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, inactiveStatus, serverRoleInvalid, replicaTypeFull)
	assert.False(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleLeader, replicaTypeFull)
	assert.True(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleFollower, replicaTypeFull)
	assert.True(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleLeader, replicaTypeLogOnly)
	assert.True(t, repLoc.isValid())
	repLoc = newReplicaLocation(addr, activeStatus, serverRoleLeader, replicaTypeReadOnly)
	assert.True(t, repLoc.isValid())
}

func TestObReplicaLocation_isLeader(t *testing.T) {
	repLoc := newReplicaLocation(nil, nil, serverRoleLeader, replicaTypeInvalid)
	assert.True(t, repLoc.isLeader())
	repLoc = newReplicaLocation(nil, nil, serverRoleInvalid, replicaTypeInvalid)
	assert.False(t, repLoc.isLeader())
	repLoc = newReplicaLocation(nil, nil, serverRoleFollower, replicaTypeInvalid)
	assert.False(t, repLoc.isLeader())
}

func TestObReplicaLocation_String(t *testing.T) {
	repLoc := &obReplicaLocation{}
	assert.Equal(t, "obReplicaLocation{addr:nil, info:nil, role:0, replicaType:0}", repLoc.String())
	repLoc = newReplicaLocation(
		NewObServerAddr("127.0.0.1", 1222, 12223),
		newServerStatus(0, "Active"),
		serverRoleInvalid,
		replicaTypeInvalid,
	)
	assert.Equal(t, "obReplicaLocation{addr:ObServerAddr{ip:127.0.0.1, sqlPort:1222, svrPort:12223}, info:obServerStatus{stopTime:0, status:Active}, role:-1, replicaType:-1}", repLoc.String())
}

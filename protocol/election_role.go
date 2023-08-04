package protocol

type ObRole uint8

const (
	ObRoleInvalid ObRole = iota
	ObRoleLeader
	ObRoleFollower
	ObRoleStandbyFollower
	ObRoleRestoreFollower
)

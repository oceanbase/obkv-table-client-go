package reroute

type Partition struct {
	tableId    uint64
	partId     uint64
	replicaNum uint64
	leader     *Replica
	follower   []*Replica
}

func NewPartition(partId uint64) *Partition {
	return &Partition{
		tableId:    0,
		partId:     partId,
		replicaNum: 0,
		leader:     nil,
		follower:   make([]*Replica, 0, 2),
	}
}

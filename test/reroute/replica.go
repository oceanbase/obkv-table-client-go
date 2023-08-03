package reroute

import (
	"github.com/oceanbase/obkv-table-client-go/protocol"
)

type Replica struct {
	TableId uint64
	PartId  uint64
	Ip      string
	Port    int
	Role    protocol.ObRole
}

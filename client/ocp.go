package client

import (
	"github.com/oceanbase/obkv-table-client-go/route"
	"math/rand"
)

type ObOcpModel struct {
	serverAddrs []route.ObServerAddr
	clusterId   int64
	// todo: idc2Region
}

func (o *ObOcpModel) getServerAddressRandomly() *route.ObServerAddr {
	idx := rand.Intn(len(o.serverAddrs) - 1)
	return &o.serverAddrs[idx]
}

package client

import (
	"github.com/oceanbase/obkv-table-client-go/route"
	"math/rand"
	"strconv"
	"sync/atomic"
)

type obServerRoster struct {
	maxPriority atomic.Int32
	roster      []*route.ObServerAddr
}

func (r *obServerRoster) MaxPriority() int32 {
	return r.maxPriority.Load()
}

func (r *obServerRoster) Reset(servers []*route.ObServerAddr) {
	r.maxPriority.Store(0)
	r.roster = servers
}

func (r *obServerRoster) GetServer() *route.ObServerAddr {
	idx := rand.Intn(len(r.roster))
	return r.roster[idx]
}

func (r *obServerRoster) Size() int {
	return len(r.roster)
}

func (r *obServerRoster) String() string {
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
	return "obServerRoster{" +
		"maxPriority:" + strconv.Itoa(int(r.maxPriority.Load())) + ", " +
		"roster:" + rostersStr +
		"}"
}

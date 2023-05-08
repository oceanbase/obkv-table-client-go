package route

import (
	"math/rand"
	"strconv"
	"sync/atomic"
)

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

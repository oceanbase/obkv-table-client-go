package route

import "strconv"

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

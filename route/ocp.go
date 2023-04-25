package route

import "math/rand"

type ObOcpModel struct {
	serverAddrs []ObServerAddr
	//clusterId   int64
	// todo: idc2Region
}

func (o *ObOcpModel) GetServerAddressRandomly() *ObServerAddr {
	idx := rand.Intn(len(o.serverAddrs) - 1)
	return &o.serverAddrs[idx]
}

func LoadOcpModel(configUrl string, fileName string) (*ObOcpModel, error) {
	// todo: impl
	return nil, nil
}

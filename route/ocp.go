package route

import "math/rand"

type ObOcpModel struct {
	ServerAddrs []ObServerAddr
	//clusterId   int64
	// todo: idc2Region
}

func (o *ObOcpModel) GetServerAddressRandomly() *ObServerAddr {
	idx := rand.Intn(len(o.ServerAddrs))
	return &o.ServerAddrs[idx]
}

func LoadOcpModel(configUrl string, fileName string) (*ObOcpModel, error) {
	// todo: impl
	return nil, nil
}

package model

type OcpModel struct {
	obServerAddrs []ObServerAddr
	clusterId     int64
	idc2Region    map[string]string
}

// GetObServerAddrs Get ob server addrs.
func (t *OcpModel) GetObServerAddrs() []ObServerAddr {
	return t.obServerAddrs
}

// SetObServerAddrs Set ob server addrs.
func (t *OcpModel) SetObServerAddrs(obServerAddrs []ObServerAddr) {
	t.obServerAddrs = obServerAddrs
}

// GetClusterId Get cluster id.
func (t *OcpModel) GetClusterId() int64 {
	return t.clusterId
}

// SetClusterId Set cluster id.
func (t *OcpModel) SetClusterId(clusterId int64) {
	t.clusterId = clusterId
}

// GetIdc2Region Get Region by IDC.
func (t *OcpModel) GetIdc2Region(idc string) string {
	return t.idc2Region[idc]
}

// GetIdc2Regions Get Region by IDC List.
func (t *OcpModel) GetIdc2Regions() map[string]string {
	return t.idc2Region
}

// AddIdc2Region Add Idc-Region pair.
func (t *OcpModel) AddIdc2Region(idc, region string) {
	t.idc2Region[idc] = region
}

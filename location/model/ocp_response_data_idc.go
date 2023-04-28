package model

import "fmt"

type OcpResponseDataIDC struct {
	idc    string
	region string
}

// GetIdc get idc
func (t *OcpResponseDataIDC) GetIdc() string {
	return t.idc
}

// SetIdc set idc
func (t *OcpResponseDataIDC) SetIdc(idc string) {
	t.idc = idc
}

// GetRegion get region
func (t *OcpResponseDataIDC) GetRegion() string {
	return t.region
}

// SetRegion set region
func (t *OcpResponseDataIDC) SetRegion(region string) {
	t.region = region
}

// ToString To string
func (t *OcpResponseDataIDC) ToString() string {
	return fmt.Sprintf("OcpResponseDataIDC{idc='%s',region='%s'}",
		t.idc, t.region)
}

package model

import (
	"fmt"
	"reflect"
)

type OcpResponseData struct {
	obRegion   string
	obRegionId int64
	rsList     []OcpResponseDataRs
	idcList    []OcpResponseDataIDC
}

// GetObRegion get obRegion
func (t *OcpResponseData) GetObRegion() string {
	return t.obRegion
}

// SetObRegion set obRegion
func (t *OcpResponseData) SetObRegion(obRegion string) {
	t.obRegion = obRegion
}

// GetObRegionId ob region id
func (t *OcpResponseData) GetObRegionId() int64 {
	return t.obRegionId
}

// SetObRegionId ob region id
func (t *OcpResponseData) SetObRegionId(obRegionId int64) {
	t.obRegionId = obRegionId
}

// GetRsList rs list
func (t *OcpResponseData) GetRsList() []OcpResponseDataRs {
	return t.rsList
}

// SetRsList rs list
func (t *OcpResponseData) SetRsList(rsList []OcpResponseDataRs) {
	t.rsList = rsList
}

// GetIDCList IDC list
func (t *OcpResponseData) GetIDCList() []OcpResponseDataIDC {
	return t.idcList
}

// SetIDCList IDC list
func (t *OcpResponseData) SetIDCList(IDCList []OcpResponseDataIDC) {
	t.idcList = IDCList
}

// Validate Validate
func (t *OcpResponseData) Validate() bool {
	return t.rsList != nil && len(t.rsList) > 0
}

func (t *OcpResponseData) IsEmpty() bool {
	return reflect.DeepEqual(t, OcpResponseData{})
}

// ToString To string
func (t *OcpResponseData) ToString() string {
	return fmt.Sprintf("OcpResponseData{ObRegion='%s',ObRegionId=%d,RsList=%v,IDCList=%v",
		t.obRegion, t.obRegionId, t.rsList, t.idcList)
}

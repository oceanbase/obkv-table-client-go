package route

import (
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/pkg/errors"
	"strconv"
)

type ObPartitionInfo struct {
	level           ObPartitionLevel
	firstPartDesc   ObPartDesc
	subPartDesc     ObPartDesc
	partColumns     []*ObColumn
	partTabletIdMap map[int64]int64
	partNameIdMap   map[string]int64
}

func (p *ObPartitionInfo) SubPartDesc() ObPartDesc {
	return p.subPartDesc
}

func (p *ObPartitionInfo) FirstPartDesc() ObPartDesc {
	return p.firstPartDesc
}

func (p *ObPartitionInfo) GetTabletId(partId int64) (int64, error) {
	if p.partTabletIdMap == nil {
		log.Warn("partTabletIdMap is nil")
		return 0, errors.New("partTabletIdMap is nil")
	}
	return p.partTabletIdMap[partId], nil
}

func (p *ObPartitionInfo) Level() ObPartitionLevel {
	return p.level
}

func (p *ObPartitionInfo) setRowKeyElement(rowKeyElement *table.ObRowkeyElement) {
	if p.firstPartDesc != nil {
		p.firstPartDesc.setRowKeyElement(rowKeyElement)
	}
	if p.subPartDesc != nil {
		p.subPartDesc.setRowKeyElement(rowKeyElement)
	}
}

func (p *ObPartitionInfo) String() string {
	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(p.partColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += p.partColumns[i].String()
	}
	partColumnsStr += "]"

	// partTabletIdMap to string
	var partTabletIdMapStr string
	var i = 0
	partTabletIdMapStr = partTabletIdMapStr + "{"
	for k, v := range p.partTabletIdMap {
		if i > 0 {
			partTabletIdMapStr += ", "
		}
		i++
		partTabletIdMapStr += "m[" + strconv.Itoa(int(k)) + "]=" + strconv.Itoa(int(v))
	}
	partTabletIdMapStr += "}"

	// partNameIdMap to string
	var partNameIdMapStr string
	i = 0
	partNameIdMapStr = partNameIdMapStr + "{"
	for k, v := range p.partNameIdMap {
		if i > 0 {
			partNameIdMapStr += ", "
		}
		i++
		partNameIdMapStr += "m[" + k + "]=" + strconv.Itoa(int(v))
	}
	partNameIdMapStr += "}"

	// firstPartDesc to string
	var firstPartDescStr string
	if p.firstPartDesc != nil {
		firstPartDescStr = p.firstPartDesc.String()
	} else {
		firstPartDescStr = "nil"
	}

	// subPartDesc to string
	var subPartDescStr string
	if p.subPartDesc != nil {
		subPartDescStr = p.firstPartDesc.String()
	} else {
		subPartDescStr = "nil"
	}

	return "ObPartitionInfo{" +
		"level:" + p.level.String() + ", " +
		"firstPartDesc:" + firstPartDescStr + ", " +
		"subPartDesc:" + subPartDescStr + ", " +
		"partColumns:" + partColumnsStr + ", " +
		"partTabletIdMap:" + partTabletIdMapStr + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}

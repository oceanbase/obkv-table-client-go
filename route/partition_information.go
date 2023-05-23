/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package route

import (
	"strconv"

	"github.com/pkg/errors"
)

type obPartLevel int

const (
	PartLevelUnknown obPartLevel = -1
	PartLevelZero    obPartLevel = 0
	PartLevelOne     obPartLevel = 1
	PartLevelTwo     obPartLevel = 2
)

// obPartitionInfo store meta information about partition keys
type obPartitionInfo struct {
	level           obPartLevel
	firstPartDesc   obPartDesc
	subPartDesc     obPartDesc
	partTabletIdMap map[int64]int64 // partId mean [0, partNum), tabletId For example, 500029
}

func newObPartitionInfo(level obPartLevel) *obPartitionInfo {
	return &obPartitionInfo{level: level}
}

func (p *obPartitionInfo) SubPartDesc() obPartDesc {
	return p.subPartDesc
}

func (p *obPartitionInfo) FirstPartDesc() obPartDesc {
	return p.firstPartDesc
}

// GetTabletId obtain tabletId through partId.
// In ob version 4.0 and above, tabletId needs to be serialized to the server
func (p *obPartitionInfo) GetTabletId(partId int64) (int64, error) {
	if p.partTabletIdMap == nil {
		return 0, errors.New("partTabletIdMap is nil")
	}
	return p.partTabletIdMap[partId], nil
}

func (p *obPartitionInfo) Level() obPartLevel {
	return p.level
}

func (p *obPartitionInfo) String() string {
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

	return "obPartitionInfo{" +
		"level:" + strconv.Itoa(int(p.level)) + ", " +
		"firstPartDesc:" + firstPartDescStr + ", " +
		"subPartDesc:" + subPartDescStr + ", " +
		"partTabletIdMap:" + partTabletIdMapStr +
		"}"
}

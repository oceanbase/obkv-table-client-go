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

import "strconv"

// ObPartLocationEntry store location information for all replicas.
type ObPartLocationEntry struct {
	partLocations map[int64]*obPartitionLocation
}

func newObPartLocationEntry(partNum int) *ObPartLocationEntry {
	entry := new(ObPartLocationEntry)
	entry.partLocations = make(map[int64]*obPartitionLocation, partNum)
	return entry
}

func (e *ObPartLocationEntry) String() string {
	var partitionLocationStr string
	var i = 0
	partitionLocationStr = partitionLocationStr + "{"
	for k, v := range e.partLocations {
		if i > 0 {
			partitionLocationStr += ", "
		}
		i++
		partitionLocationStr += "m[" + strconv.Itoa(int(k)) + "]="
		if v != nil {
			partitionLocationStr += v.String()
		} else {
			partitionLocationStr += "nil"
		}
	}
	partitionLocationStr += "}"
	return "ObPartLocationEntry{" +
		"partLocations:" + partitionLocationStr +
		"}"
}

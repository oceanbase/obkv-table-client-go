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

type ObTableLocation struct {
	replicaLocations []*obReplicaLocation
}

func (l *ObTableLocation) ReplicaLocations() []*obReplicaLocation {
	return l.replicaLocations
}

func (l *ObTableLocation) String() string {
	var replicaLocationsStr string
	replicaLocationsStr = replicaLocationsStr + "["
	for i := 0; i < len(l.replicaLocations); i++ {
		if i > 0 {
			replicaLocationsStr += ", "
		}
		replicaLocationsStr += l.replicaLocations[i].String()
	}
	replicaLocationsStr += "]"
	return "ObTableLocation{" +
		"replicaLocations:" + replicaLocationsStr +
		"}"
}

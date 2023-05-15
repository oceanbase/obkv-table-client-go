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
	"strings"
)

type obServerStatus struct {
	stopTime int64
	status   string // Active/InActive/Deleting
}

func newServerStatus(stopTime int64, status string) *obServerStatus {
	return &obServerStatus{stopTime, status}
}

func (i *obServerStatus) IsActive() bool {
	return i.stopTime == 0 && strings.EqualFold(i.status, "active") // ignore case
}

func (i *obServerStatus) String() string {
	return "obServerStatus{" +
		"stopTime:" + strconv.Itoa(int(i.stopTime)) + ", " +
		"status:" + i.status +
		"}"
}

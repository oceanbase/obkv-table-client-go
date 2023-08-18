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

package util

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"regexp"
	"strconv"
	"sync"
)

var globalObVersion float32 = 0.0
var obVersionGuard sync.Mutex

func ObVersion() float32 {
	return globalObVersion
}

func SetObVersion(version float32) {
	obVersionGuard.Lock()
	globalObVersion = version
	obVersionGuard.Unlock()
}

// ParseObVerionFromLogin may be used in ODP mode
func ParseObVerionFromLogin(serverVersion string) (float32, error) {
	// serverVersion is like "OceanBase 4.0.0.0"
	pattern := "^OceanBase\\s+(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)$"
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(serverVersion)
	if len(match) == 5 && match[0] == serverVersion {
		// transform version into 4.000
		subVersionStr := match[2] + match[3] + match[4]
		subVersion, err := strconv.Atoi(subVersionStr)
		if err != nil {
			return 0, errors.WithMessagef(err, "parse version %s failed", serverVersion)
		}
		mainVersion, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, errors.WithMessagef(err, "parse version %s failed", serverVersion)
		}
		return float32(mainVersion) + float32(subVersion)/float32(math.Pow10(len(subVersionStr))), nil
	}
	return 0, errors.New(fmt.Sprintf("parse version %s failed", serverVersion))
}

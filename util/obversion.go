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
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var globalObVersion float32 = 0.0
var globalOdpVersion float32 = 0.0
var obVersionGuard sync.Mutex

func ObVersion() float32 {
	return globalObVersion
}

func SetObVersion(version float32) {
	obVersionGuard.Lock()
	globalObVersion = version
	obVersionGuard.Unlock()
}

func SetOdpVersion(version float32) {
	obVersionGuard.Lock()
	globalOdpVersion = version
	obVersionGuard.Unlock()
}

func OdpVersion() float32 {
	return globalOdpVersion
}

// ParseObVerionFromLogin may be used in ODP mode
func ParseObVerionFromLogin(serverVersion string) (float32, float32, error) {
	pattern := ""
	if strings.HasPrefix(serverVersion, "OceanBase_CE") {
		// serverVersion in CE is like "OceanBase_CE 4.0.0.0 (+ Obproxy 4.3.6.0), content in () is optional and valid after Obproxy 4.3.5"
		pattern = "^OceanBase_CE\\s+(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)(\\s+\\+\\s+(Obproxy)\\s+(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+))?"
	} else {
		// serverVersion is like "OceanBase 4.0.0.0 (+ Obproxy 4.3.6.0), content in () is optional and valid after Obproxy 4.3.5"
		pattern = "^OceanBase\\s+(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+)(\\s+\\+\\s+(Obproxy)\\s+(\\d+)\\.(\\d+)\\.(\\d+)\\.(\\d+))?"
	}
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(serverVersion)
	// match length is always 11 when matched successfully (10 capture groups + match[0])
	// match[5] is empty when Obproxy part is not present
	if len(match) == 11 && match[0] == serverVersion {
		// transform ob version into 4.000
		subVersionStr := match[2] + match[3] + match[4]
		subVersion, err := strconv.Atoi(subVersionStr)
		if err != nil {
			return 0, 0, errors.WithMessagef(err, "parse version %s failed", serverVersion)
		}
		mainVersion, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, 0, errors.WithMessagef(err, "parse version %s failed", serverVersion)
		}
		obVersion := float32(mainVersion) + float32(subVersion)/float32(math.Pow10(len(subVersionStr)))

		// transform odp version if present
		var odpVersion float32 = 0
		if match[5] != "" && match[6] != "" {
			subOdpVersionStr := match[8] + match[9] + match[10]
			subOdpVersion, err := strconv.Atoi(subOdpVersionStr)
			if err != nil {
				return 0, 0, errors.WithMessagef(err, "parse version %s failed", serverVersion)
			}
			mainOdpVersion, err := strconv.Atoi(match[7])
			if err != nil {
				return 0, 0, errors.WithMessagef(err, "parse version %s failed", serverVersion)
			}
			odpVersion = float32(mainOdpVersion) + float32(subOdpVersion)/float32(math.Pow10(len(subOdpVersionStr)))
		}

		return obVersion, odpVersion, nil
	}
	return 0, 0, errors.New(fmt.Sprintf("parse version %s failed 5", serverVersion))
}

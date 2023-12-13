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
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
)

type ObRslist struct {
	list []*ObServerAddr
}

func NewRslist() *ObRslist {
	return &ObRslist{
		list: make([]*ObServerAddr, 0),
	}
}

// GetServerRandomly get one randomly server from all the servers
func (l *ObRslist) GetServerRandomly() (*ObServerAddr, error) {
	if len(l.list) == 0 {
		return nil, errors.New("list is empty")
	}

	idx := rand.Intn(len(l.list))
	addr := l.list[idx]
	if addr.IsConnectionFine() {
		return addr, nil
	}

	for _, serverAddr := range l.list {
		if serverAddr.IsConnectionFine() {
			return serverAddr, nil
		}
	}

	return nil, errors.New("has no valid server node")
}

func (l *ObRslist) Append(addr *ObServerAddr) {
	l.list = append(l.list, addr)
}

func (l *ObRslist) Size() int {
	return len(l.list)
}

func (l *ObRslist) Equal(other *ObRslist) bool {
	if l.Size() != other.Size() {
		return false
	}
	for i := 0; i < l.Size(); i++ {
		if !l.list[i].Equal(other.list[i]) {
			return false
		}
	}
	return true
}

func (l *ObRslist) FindMissingElements(other *ObRslist) []*ObServerAddr {
	missingElements := make([]*ObServerAddr, 0)

	bMap := make(map[string]bool)
	for _, element := range other.list {
		key := fmt.Sprintf("%s:%d", element.ip, element.port)
		bMap[key] = true
	}

	for _, element := range l.list {
		key := fmt.Sprintf("%s:%d", element.ip, element.port)
		if _, found := bMap[key]; !found {
			missingElements = append(missingElements, element)
		}
	}

	return missingElements
}

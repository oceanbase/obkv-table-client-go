/* Copyright  2017  Chao yuepan
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package obkvrpc

import (
	"math"
	"sync"
)

type levelPool struct {
	size int
	pool sync.Pool
}

func newLevelPool(size int) *levelPool {
	return &levelPool{
		size: size,
		pool: sync.Pool{
			New: func() interface{} {
				data := make([]byte, size)
				return &data
			},
		},
	}
}

type LimitedPool struct {
	minSize int
	maxSize int
	pools   []*levelPool
}

func NewLimitedPool(minSize, maxSize int) *LimitedPool {
	if maxSize < minSize {
		panic("maxSize can't be less than minSize")
	}
	const multiplier = 2
	var pools []*levelPool
	curSize := minSize
	for curSize < maxSize {
		pools = append(pools, newLevelPool(curSize))
		curSize *= multiplier
	}
	pools = append(pools, newLevelPool(maxSize))
	return &LimitedPool{
		minSize: minSize,
		maxSize: maxSize,
		pools:   pools,
	}
}

func (p *LimitedPool) findPool(size int) *levelPool {
	if size > p.maxSize {
		return nil
	}
	idx := int(math.Ceil(math.Log2(float64(size) / float64(p.minSize))))
	if idx < 0 {
		idx = 0
	}
	if idx > len(p.pools)-1 {
		return nil
	}
	return p.pools[idx]
}

func (p *LimitedPool) findPutPool(size int) *levelPool {
	if size > p.maxSize {
		return nil
	}
	if size < p.minSize {
		return nil
	}

	idx := int(math.Floor(math.Log2(float64(size) / float64(p.minSize))))
	if idx < 0 {
		idx = 0
	}
	if idx > len(p.pools)-1 {
		return nil
	}
	return p.pools[idx]
}

func (p *LimitedPool) Get(size int) *[]byte {
	sp := p.findPool(size)
	if sp == nil {
		data := make([]byte, size)
		return &data
	}
	buf := sp.pool.Get().(*[]byte)
	*buf = (*buf)[:size]
	return buf
}

func (p *LimitedPool) Put(b *[]byte) {
	sp := p.findPutPool(cap(*b))
	if sp == nil {
		return
	}
	*b = (*b)[:cap(*b)]
	sp.pool.Put(b)
}

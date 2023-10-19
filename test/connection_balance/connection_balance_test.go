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

package connection_balance

import (
	"context"
	"fmt"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/table"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func run(i int, done chan bool, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	executeNum := 0
	for {
		select {
		case <-done:
			log.Info(fmt.Sprintf("Finish %d worker, executeNum: %d", i, executeNum))
			return
		default:
			rowKey := []*table.Column{table.NewColumn("c1", fmt.Sprintf("key%d", i))}
			mutateColumns := []*table.Column{table.NewColumn("c2", int32(1))}
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10s
			affectRows, err := cli.InsertOrUpdate(ctx, testConnectionBalanceTableName, rowKey, mutateColumns)
			assert.Equal(t, nil, err)
			assert.EqualValues(t, 1, affectRows)
			executeNum++
		}
	}
}

func TestMaxConnectionAge(t *testing.T) {
	println("Test begin")
	done := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < concurrencyNum; i++ {
		wg.Add(1)
		go run(i, done, &wg, t)
	}
	time.Sleep(testDuration)
	close(done)
	println("Wait All Coroutine finish")
	wg.Wait()
	println("Test Finished")
}

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

package client

import (
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"strings"
)

type QueryResult interface {
	Value(columnName string) interface{}
}

func newObQueryResult(columnNames []string, values []*protocol.ObObject) *obQueryResult {
	return &obQueryResult{columnNames, values}
}

type obQueryResult struct {
	columnNames []string
	values      []*protocol.ObObject
}

func (r *obQueryResult) Value(columnName string) interface{} {
	if r.columnNames == nil {
		return nil
	}
	if r.values == nil {
		return nil
	}

	for i, resColumnName := range r.columnNames {
		if strings.EqualFold(columnName, resColumnName) {
			return r.values[i].Value()
		}
	}
	return nil
}

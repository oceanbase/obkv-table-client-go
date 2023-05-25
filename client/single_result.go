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
	"strings"

	"github.com/oceanbase/obkv-table-client-go/protocol"
)

type SingleResult interface {
	AffectedRows() int64
	Value(columnName string) interface{}
	RowKey() []interface{}
}

func newObSingleResult(affectedRows int64, affectedEntity *protocol.ObTableEntity) *obSingleResult {
	return &obSingleResult{affectedRows, affectedEntity}
}

type obSingleResult struct {
	affectedRows   int64
	affectedEntity *protocol.ObTableEntity
}

func (r *obSingleResult) AffectedRows() int64 {
	return r.affectedRows
}

func (r *obSingleResult) Value(columnName string) interface{} {
	if r.affectedEntity == nil {
		return nil
	}
	if r.affectedEntity.Properties() == nil {
		return nil
	}
	obj := r.affectedEntity.Properties()[strings.ToLower(columnName)]
	if obj == nil {
		return nil
	}
	return obj.Value()
}

func (r *obSingleResult) RowKey() []interface{} {
	if r.affectedEntity == nil {
		return nil
	}
	if r.affectedEntity.RowKey() == nil {
		return nil
	}
	if r.affectedEntity.RowKey().Keys() == nil || len(r.affectedEntity.RowKey().Keys()) == 0 {
		return nil
	}

	keys := r.affectedEntity.RowKey().Keys()
	res := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		if key != nil {
			res = append(res, key.Value())
		}
	}
	return res
}

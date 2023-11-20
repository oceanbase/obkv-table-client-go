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
	// IsEmptySet get empty row or not when do Get operation.
	IsEmptySet() bool
	// AffectedRows get affected row.
	AffectedRows() int64
	// Value get value by column name.
	Value(columnName string) interface{}
	// Values get key value map.
	Values() map[string]interface{}
	// RowKey get affected rowkey, only work in Increment and Append operation.
	RowKey() []interface{}
	// IsInsertOrUpdateDoInsert InsertOrUpdate impl maybe do insert or do update,
	// IsInsertOrUpdateDoInsert() to tell you which operation has done.
	IsInsertOrUpdateDoInsert() bool
	// IsInsertOrUpdateDoUpdate InsertOrUpdate impl maybe do insert or do update,
	// IsInsertOrUpdateDoUpdate() to tell you which operation has done.
	IsInsertOrUpdateDoUpdate() bool
	// IsInsertOrUpdateDoPut InsertOrUpdate impl maybe do put when all columns are filled.
	IsInsertOrUpdateDoPut() bool
}

func newObSingleResult(affectedRows int64, affectedEntity *protocol.ObTableEntity, flags uint64) *obSingleResult {
	return &obSingleResult{affectedRows, affectedEntity, flags}
}

type obSingleResult struct {
	affectedRows   int64
	affectedEntity *protocol.ObTableEntity
	flags          uint64
}

func (r *obSingleResult) IsEmptySet() bool {
	return len(r.affectedEntity.Properties()) == 0
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

	obj := r.affectedEntity.Properties()[columnName]
	if obj != nil {
		return obj.Value()
	}

	for name, obj := range r.affectedEntity.Properties() {
		if strings.EqualFold(columnName, name) {
			return obj.Value()
		}
	}
	return nil
}

func (r *obSingleResult) Values() map[string]interface{} {
	if len(r.affectedEntity.Properties()) == 0 {
		return nil
	}
	m := make(map[string]interface{}, len(r.affectedEntity.Properties()))
	for k, v := range r.affectedEntity.Properties() {
		m[k] = v.Value()
	}
	return m
}

func (r *obSingleResult) RowKey() []interface{} {
	if r.affectedEntity == nil {
		return nil
	}
	if r.affectedEntity.RowKey() == nil {
		return nil
	}
	if r.affectedEntity.RowKey() == nil || len(r.affectedEntity.RowKey()) == 0 {
		return nil
	}

	keys := r.affectedEntity.RowKey()
	res := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		if key != nil {
			res = append(res, key.Value())
		}
	}
	return res
}

func (r *obSingleResult) IsInsertOrUpdateDoInsert() bool {
	return r.flags&protocol.IsInsertUpDoInsertMask != 0
}

func (r *obSingleResult) IsInsertOrUpdateDoUpdate() bool {
	return !r.IsInsertOrUpdateDoInsert() && !r.IsInsertOrUpdateDoPut()
}

func (r *obSingleResult) IsInsertOrUpdateDoPut() bool {
	return r.flags&protocol.IsInsertUpDoPutMask != 0
}

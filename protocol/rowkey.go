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

package protocol

import (
	"bytes"

	"github.com/oceanbase/obkv-table-client-go/util"
)

type RowKey struct {
	keys []*ObObject
}

func NewRowKey() *RowKey {
	return &RowKey{
		keys: nil,
	}
}

func (k *RowKey) Keys() []*ObObject {
	return k.keys
}

func (k *RowKey) SetKeys(keys []*ObObject) {
	k.keys = keys
}

func (k *RowKey) AppendKey(key *ObObject) {
	k.keys = append(k.keys, key)
}

func (k *RowKey) Encode(buffer *bytes.Buffer) {
	util.EncodeVi64(buffer, int64(len(k.keys)))

	for _, key := range k.keys {
		key.Encode(buffer)
	}
}

func (k *RowKey) Decode(buffer *bytes.Buffer) {
	keysLen := util.DecodeVi64(buffer)

	var i int64
	for i = 0; i < keysLen; i++ {
		key := NewObObject()
		key.Decode(buffer)
		k.keys = append(k.keys, key)
	}
}

func (k *RowKey) EncodedLength() int {
	totalLen := util.EncodedLengthByVi64(int64(len(k.keys)))

	for _, key := range k.keys {
		totalLen += key.EncodedLength()
	}

	return totalLen
}

func (k *RowKey) GetRowKeyValue() []interface{} {
	rowKey := make([]interface{}, 0, len(k.keys))
	for _, key := range k.keys {
		rowKey = append(rowKey, key.value)
	}
	return rowKey
}

func (k *RowKey) String() string {
	var keysStr string
	keysStr = keysStr + "["
	for i := 0; i < len(k.keys); i++ {
		if i > 0 {
			keysStr += ", "
		}
		if k.keys[i] != nil {
			keysStr += k.keys[i].String()
		} else {
			keysStr += "nil"
		}
	}
	keysStr += "]"

	return "RowKey{" +
		"keys:" + keysStr +
		"}"
}

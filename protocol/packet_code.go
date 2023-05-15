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
	"github.com/pkg/errors"
)

type TablePacketCode int32

const (
	TableApiLogin TablePacketCode = iota
	TableApiExecute
	TableApiBatchExecute
	TableApiExecuteQuery
	TableApiQueryAndMute
	TableApiErrorPacket

	TableApiNoSuch = -1
)

const (
	tableApiLoginStr        string = "table login"
	tableApiExecuteStr      string = "table execute"
	tableApiBatchExecuteStr string = "table batch execute"
	tableApiExecuteQueryStr string = "table execute query"
	tableApiQueryAndMuteStr string = "table query and mute"
	tableApiErrorPacketStr  string = "table error"
)

const (
	tableApiPCodeLogin        uint32 = 0x1101
	tableApiPCodeExecute      uint32 = 0x1102
	tableApiPCodeBatchExecute uint32 = 0x1103
	tableApiPCodeExecuteQuery uint32 = 0x1104
	tableApiPCodeQueryAndMute uint32 = 0x1105
	tableApiPCodeErrorPacket  uint32 = 0x010
)

var tablePacketCodeStrings = []string{
	TableApiLogin:        tableApiLoginStr,
	TableApiExecute:      tableApiExecuteStr,
	TableApiBatchExecute: tableApiBatchExecuteStr,
	TableApiExecuteQuery: tableApiExecuteQueryStr,
	TableApiQueryAndMute: tableApiQueryAndMuteStr,
	TableApiErrorPacket:  tableApiErrorPacketStr,
}

var tablePacketCodePCodes = []uint32{
	TableApiLogin:        tableApiPCodeLogin,
	TableApiExecute:      tableApiPCodeExecute,
	TableApiBatchExecute: tableApiPCodeBatchExecute,
	TableApiExecuteQuery: tableApiPCodeExecuteQuery,
	TableApiQueryAndMute: tableApiPCodeQueryAndMute,
	TableApiErrorPacket:  tableApiPCodeErrorPacket,
}

func (c TablePacketCode) Value() uint32 {
	return tablePacketCodePCodes[c]
}

func (c TablePacketCode) ValueOf(pCode uint32) (TablePacketCode, error) { // TODO use map optimize
	switch pCode {
	case tableApiPCodeLogin:
		return TableApiLogin, nil
	case tableApiPCodeExecute:
		return TableApiExecute, nil
	case tableApiPCodeBatchExecute:
		return TableApiBatchExecute, nil
	case tableApiPCodeExecuteQuery:
		return TableApiExecuteQuery, nil
	case tableApiPCodeQueryAndMute:
		return TableApiQueryAndMute, nil
	case tableApiPCodeErrorPacket:
		return TableApiErrorPacket, nil
	}
	return TableApiNoSuch, errors.New("not match code")
}

func (c TablePacketCode) String() string {
	return tablePacketCodeStrings[c]
}

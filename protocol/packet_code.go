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

type ObTablePacketCode int32

const (
	ObTableApiLogin ObTablePacketCode = iota
	ObTableApiExecute
	ObTableApiBatchExecute
	ObTableApiExecuteQuery
	ObTableApiQueryAndMute
	ObTableApiExecuteAsyncQuery
	ObTableApiMove
	ObTableApiErrorPacket
	ObTableApiRedis

	ObTableApiNoSuch = -1
)

const (
	obTableApiLoginStr             string = "ob table login"
	obTableApiExecuteStr           string = "ob table execute"
	obTableApiBatchExecuteStr      string = "ob table batch execute"
	obTableApiExecuteQueryStr      string = "ob table execute query"
	obTableApiQueryAndMuteStr      string = "ob table query and mute"
	obTableApiExecuteAsyncQueryStr string = "ob table execute async query"
	obTableApiMoveStr              string = "ob table route"
	obTableApiErrorPacketStr       string = "ob table error"
	obTableApiRedisStr             string = "ob table redis"
)

const (
	obTableApiPCodeLogin             uint32 = 0x1101
	obTableApiPCodeExecute           uint32 = 0x1102
	obTableApiPCodeBatchExecute      uint32 = 0x1103
	obTableApiPCodeExecuteQuery      uint32 = 0x1104
	obTableApiPCodeQueryAndMute      uint32 = 0x1105
	obTableApiPCodeExecuteAsyncQuery uint32 = 0x1106
	obTableApiPCodeMove              uint32 = 0x1124
	obTableApiPCodeRedisExecute      uint32 = 0x1126
	obTableApiPCodeErrorPacket       uint32 = 0x010
)

var obTablePacketCodeStrings = []string{
	ObTableApiLogin:             obTableApiLoginStr,
	ObTableApiExecute:           obTableApiExecuteStr,
	ObTableApiBatchExecute:      obTableApiBatchExecuteStr,
	ObTableApiExecuteQuery:      obTableApiExecuteQueryStr,
	ObTableApiQueryAndMute:      obTableApiQueryAndMuteStr,
	ObTableApiExecuteAsyncQuery: obTableApiExecuteAsyncQueryStr,
	ObTableApiMove:              obTableApiMoveStr,
	ObTableApiErrorPacket:       obTableApiErrorPacketStr,
	ObTableApiRedis:             obTableApiRedisStr,
}

var obTablePacketCodePCodes = []uint32{
	ObTableApiLogin:             obTableApiPCodeLogin,
	ObTableApiExecute:           obTableApiPCodeExecute,
	ObTableApiBatchExecute:      obTableApiPCodeBatchExecute,
	ObTableApiExecuteQuery:      obTableApiPCodeExecuteQuery,
	ObTableApiQueryAndMute:      obTableApiPCodeQueryAndMute,
	ObTableApiExecuteAsyncQuery: obTableApiPCodeExecuteAsyncQuery,
	ObTableApiMove:              obTableApiPCodeMove,
	ObTableApiErrorPacket:       obTableApiPCodeErrorPacket,
	ObTableApiRedis:             obTableApiPCodeRedisExecute,
}

func (c ObTablePacketCode) Value() uint32 {
	return obTablePacketCodePCodes[c]
}

func (c ObTablePacketCode) ValueOf(pCode uint32) (ObTablePacketCode, error) { // use map optimize
	switch pCode {
	case obTableApiPCodeLogin:
		return ObTableApiLogin, nil
	case obTableApiPCodeExecute:
		return ObTableApiExecute, nil
	case obTableApiPCodeBatchExecute:
		return ObTableApiBatchExecute, nil
	case obTableApiPCodeExecuteQuery:
		return ObTableApiExecuteQuery, nil
	case obTableApiPCodeQueryAndMute:
		return ObTableApiQueryAndMute, nil
	case obTableApiPCodeExecuteAsyncQuery:
		return ObTableApiExecuteAsyncQuery, nil
	case obTableApiPCodeMove:
		return ObTableApiMove, nil
	case obTableApiPCodeErrorPacket:
		return ObTableApiErrorPacket, nil
	case obTableApiPCodeRedisExecute:
		return ObTableApiRedis, nil
	}
	return ObTableApiNoSuch, errors.New("not match code")
}

func (c ObTablePacketCode) String() string {
	return obTablePacketCodeStrings[c]
}

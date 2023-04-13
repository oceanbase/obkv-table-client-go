package protocol

import (
	"github.com/pkg/errors"
)

type TablePacketCode int

const (
	Login TablePacketCode = iota
	Execute
	BatchExecute
	ExecuteQuery
	QueryAndMute
	Error

	NoSuch = -1
)

const (
	loginStr        string = "login"
	executeStr      string = "execute"
	batchExecuteStr string = "batchExecute"
	executeQueryStr string = "executeQuery"
	queryAndMuteStr string = "queryAndMute"
	errorStr        string = "error"
)

const (
	pCodeLogin        uint32 = 0x1101
	pCodeExecute      uint32 = 0x1102
	pCodeBatchExecute uint32 = 0x1103
	pCodeExecuteQuery uint32 = 0x1104
	pCodeQueryAndMute uint32 = 0x1105
	pCodeErrorPacket  uint32 = 0x010
)

var tablePacketCodeStrings = [...]string{
	Login:        loginStr,
	Execute:      executeStr,
	BatchExecute: batchExecuteStr,
	ExecuteQuery: executeQueryStr,
	QueryAndMute: queryAndMuteStr,
	Error:        errorStr,
}

var tablePacketCodePCodes = [...]uint32{
	Login:        pCodeLogin,
	Execute:      pCodeExecute,
	BatchExecute: pCodeBatchExecute,
	ExecuteQuery: pCodeExecuteQuery,
	QueryAndMute: pCodeQueryAndMute,
	Error:        pCodeErrorPacket,
}

func (c TablePacketCode) Value() uint32 {
	return tablePacketCodePCodes[c]
}

func (c TablePacketCode) String() string {
	return tablePacketCodeStrings[c]
}

func (c TablePacketCode) FromPCode(pCode uint32) (TablePacketCode, error) {
	switch pCode {
	case pCodeLogin:
		return Login, nil
	case pCodeExecute:
		return Execute, nil
	case pCodeBatchExecute:
		return BatchExecute, nil
	case pCodeExecuteQuery:
		return ExecuteQuery, nil
	case pCodeQueryAndMute:
		return QueryAndMute, nil
	case pCodeErrorPacket:
		return Error, nil
	}
	return NoSuch, errors.New("no such this code")
}

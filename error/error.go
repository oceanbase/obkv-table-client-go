package error

import (
	"fmt"
)

type ObProtocolError struct {
	host      string
	port      int
	code      ObErrorCode
	trace     string
	tableName string
	message   string
}

// NewProtocolError
// error message format: [code][code name][error message][host:port][trace][table name]
func NewProtocolError(host string, port int, code ObErrorCode, sequence uint64, uniqueId uint64, tableName string) *ObProtocolError {
	trace := fmt.Sprintf("Y%X-%016X", uniqueId, sequence)
	server := fmt.Sprintf("%s:%d", host, port)
	var msg string
	switch code {
	case ObErrPrimaryKeyDuplicate:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrPrimaryKeyDuplicate",
			"Primary key duplicate.",
			server,
			trace,
			tableName,
		)
	case ObErrUnknownTable:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrUnknownTable",
			"Table not exist.",
			server,
			trace,
			tableName,
		)
	case ObErrColumnNotFound:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrColumnNotFound",
			"Column not found.",
			server,
			trace,
			tableName,
		)
	case ObObjTypeError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObObjTypeError",
			"Column type not match.",
			server,
			trace,
			tableName,
		)
	case ObBadNullError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObBadNullError",
			"Column doesn't have a default value.",
			server,
			trace,
			tableName,
		)
	case ObInvalidArgument:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObInvalidArgument",
			"Arguments invalid.",
			server,
			trace,
			tableName,
		)
	case ObDeserializeError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObDeserializeError",
			"Deserialize error.",
			server,
			trace,
			tableName,
		)
	case ObPasswordWrong:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObPasswordWrong",
			"Access deny, password is wrong.",
			server,
			trace,
			tableName,
		)
	case ObLocationLeaderNotExist:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObLocationLeaderNotExist",
			"Location leader not exist.",
			server,
			trace,
			tableName,
		)
	case ObNotMaster:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObNotMaster",
			"Current node is not master.",
			server,
			trace,
			tableName,
		)
	case ObRsNotMaster:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObRsNotMaster",
			"Current Rs is not master.",
			server,
			trace,
			tableName,
		)
	case ObRsShutdown:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObRsShutdown",
			"RS had been shut down.",
			server,
			trace,
			tableName,
		)
	case ObRpcConnectError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObRpcConnectError",
			"Connect error.",
			server,
			trace,
			tableName,
		)
	case ObPartitionNotExist:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObPartitionNotExist",
			"Partition not exist.",
			server,
			trace,
			tableName,
		)
	case ObPartitionIsStopped:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObPartitionIsStopped",
			"Partition is stopped.",
			server,
			trace,
			tableName,
		)
	case ObLocationNotExist:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObLocationNotExist",
			"location not exist.",
			server,
			trace,
			tableName,
		)
	case ObServerIsInit:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObServerIsInit",
			"Observer is initing, please wait.",
			server,
			trace,
			tableName,
		)
	case ObServerIsStopping:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObServerIsStopping",
			"Observer is stopping.",
			server,
			trace,
			tableName,
		)
	case ObTenantNotInServer:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTenantNotInServer",
			"Tenant not in server.",
			server,
			trace,
			tableName,
		)
	case ObTransRpcTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTransRpcTimeout",
			"Translation timeout.",
			server,
			trace,
			tableName,
		)
	case ObNoReadableReplica:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObNoReadableReplica",
			"No readable replica.",
			server,
			trace,
			tableName,
		)
	case ObReplicaNotReadable:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObReplicaNotReadable",
			"Current replica is not readable.",
			server,
			trace,
			tableName,
		)
	case ObTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTimeout",
			"Timeout.",
			server,
			trace,
			tableName,
		)
	case ObTransTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTransTimeout",
			"Translation timeout.",
			server,
			trace,
			tableName,
		)
	case ObWaitqueueTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObWaitqueueTimeout",
			"Wait queue timeout.",
			server,
			trace,
			tableName,
		)
	default:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrUnknown",
			"Get an error from observer.",
			server,
			trace,
			tableName,
		)
	}
	return &ObProtocolError{host, port, code, trace, tableName, msg}
}

func (e *ObProtocolError) Error() string {
	return e.message
}

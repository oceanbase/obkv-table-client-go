package error

import (
	"fmt"
)

type ObProtocolError struct {
	host    string
	port    int
	code    ObErrorCode
	trace   string
	message string
}

// NewProtocolError
// error message format: [code][code name][error message][host:port][trace]
func NewProtocolError(host string, port int, code ObErrorCode, sequence uint64, uniqueId uint64) *ObProtocolError {
	trace := fmt.Sprintf("Y%X-%016X", uniqueId, sequence)
	server := fmt.Sprintf("%s:%d", host, port)
	var msg string
	switch code {
	case ObErrPrimaryKeyDuplicate:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObErrPrimaryKeyDuplicate",
			"Primary key duplicate.",
			server,
			trace,
		)
	case ObErrUnknownTable:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObErrUnknownTable",
			"Table not exist.",
			server,
			trace,
		)
	case ObErrColumnNotFound:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObErrColumnNotFound",
			"Column not found.",
			server,
			trace,
		)
	case ObObjTypeError:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObObjTypeError",
			"Column type not match.",
			server,
			trace,
		)
	case ObBadNullError:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObBadNullError",
			"Column doesn't have a default value.",
			server,
			trace,
		)
	case ObInvalidArgument:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObInvalidArgument",
			"Arguments invalid.",
			server,
			trace,
		)
	case ObDeserializeError:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObDeserializeError",
			"Deserialize error.",
			server,
			trace,
		)
	case ObPasswordWrong:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObPasswordWrong",
			"Access deny, password is wrong.",
			server,
			trace,
		)
	case ObLocationLeaderNotExist:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObLocationLeaderNotExist",
			"Location leader not exist.",
			server,
			trace,
		)
	case ObNotMaster:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObNotMaster",
			"Current node is not master.",
			server,
			trace,
		)
	case ObRsNotMaster:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObRsNotMaster",
			"Current Rs is not master.",
			server,
			trace,
		)
	case ObRsShutdown:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObRsShutdown",
			"RS had been shut down.",
			server,
			trace,
		)
	case ObRpcConnectError:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObRpcConnectError",
			"Connect error.",
			server,
			trace,
		)
	case ObPartitionNotExist:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObPartitionNotExist",
			"Partition not exist.",
			server,
			trace,
		)
	case ObPartitionIsStopped:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObPartitionIsStopped",
			"Partition is stopped.",
			server,
			trace,
		)
	case ObLocationNotExist:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObLocationNotExist",
			"location not exist.",
			server,
			trace,
		)
	case ObServerIsInit:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObServerIsInit",
			"Observer is initing, please wait.",
			server,
			trace,
		)
	case ObServerIsStopping:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObServerIsStopping",
			"Observer is stopping.",
			server,
			trace,
		)
	case ObTenantNotInServer:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObTenantNotInServer",
			"Tenant not in server.",
			server,
			trace,
		)
	case ObTransRpcTimeout:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObTransRpcTimeout",
			"Translation timeout.",
			server,
			trace,
		)
	case ObNoReadableReplica:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObNoReadableReplica",
			"No readable replica.",
			server,
			trace,
		)
	case ObReplicaNotReadable:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObReplicaNotReadable",
			"Current replica is not readable.",
			server,
			trace,
		)
	case ObTimeout:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObTimeout",
			"Timeout.",
			server,
			trace,
		)
	case ObTransTimeout:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObTransTimeout",
			"Translation timeout.",
			server,
			trace,
		)
	case ObWaitqueueTimeout:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObWaitqueueTimeout",
			"Wait queue timeout.",
			server,
			trace,
		)
	default:
		msg += fmt.Sprintf(
			"[%d][%s][%s][%s][%s]",
			code,
			"ObErrUnknown",
			"Unknown error from observer.",
			server,
			trace,
		)
	}
	return &ObProtocolError{host, port, code, trace, msg}
}

func (e *ObProtocolError) Error() string {
	return fmt.Sprintf("Error code: %d, Error message: %s, Trace: %s", e.code, e.message, e.trace)
}

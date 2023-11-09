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
	"fmt"

	error2 "github.com/oceanbase/obkv-table-client-go/error"
)

type ObProtocolError struct {
	host      string
	port      int
	code      error2.ObErrorCode
	trace     string
	tableName string
	message   string
}

// NewProtocolError
// error message format: [code][code name][error message][host:port][trace][table name]
func NewProtocolError(
	host string,
	port int,
	code error2.ObErrorCode,
	sequence uint64,
	uniqueId uint64,
	tableName string) *ObProtocolError {

	trace := fmt.Sprintf("Y%X-%016X", uniqueId, sequence)
	server := fmt.Sprintf("%s:%d", host, port)
	var msg string
	switch code {
	case error2.ObErrPrimaryKeyDuplicate:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrPrimaryKeyDuplicate",
			"Primary key duplicate.",
			server,
			trace,
			tableName,
		)
	case error2.ObErrUnknownTable:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrUnknownTable",
			"Table not exist.",
			server,
			trace,
			tableName,
		)
	case error2.ObErrColumnNotFound:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObErrColumnNotFound",
			"Column not found.",
			server,
			trace,
			tableName,
		)
	case error2.ObObjTypeError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObObjTypeError",
			"Column type not match.",
			server,
			trace,
			tableName,
		)
	case error2.ObBadNullError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObBadNullError",
			"Column doesn't have a default value.",
			server,
			trace,
			tableName,
		)
	case error2.ObInvalidArgument:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObInvalidArgument",
			"Arguments invalid.",
			server,
			trace,
			tableName,
		)
	case error2.ObDeserializeError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObDeserializeError",
			"Deserialize error.",
			server,
			trace,
			tableName,
		)
	case error2.ObPasswordWrong:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObPasswordWrong",
			"Access deny, password is wrong.",
			server,
			trace,
			tableName,
		)
	case error2.ObLocationLeaderNotExist:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObLocationLeaderNotExist",
			"Location leader not exist.",
			server,
			trace,
			tableName,
		)
	case error2.ObNotMaster:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObNotMaster",
			"Current node is not master.",
			server,
			trace,
			tableName,
		)
	case error2.ObRsNotMaster:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObRsNotMaster",
			"Current Rs is not master.",
			server,
			trace,
			tableName,
		)
	case error2.ObRsShutdown:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObRsShutdown",
			"RS had been shut down.",
			server,
			trace,
			tableName,
		)
	case error2.ObRpcConnectError:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObRpcConnectError",
			"Connect error.",
			server,
			trace,
			tableName,
		)
	case error2.ObPartitionNotExist:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObPartitionNotExist",
			"Partition not exist.",
			server,
			trace,
			tableName,
		)
	case error2.ObPartitionIsStopped:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObPartitionIsStopped",
			"Partition is stopped.",
			server,
			trace,
			tableName,
		)
	case error2.ObLocationNotExist:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObLocationNotExist",
			"location not exist.",
			server,
			trace,
			tableName,
		)
	case error2.ObServerIsInit:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObServerIsInit",
			"Observer is initing, please wait.",
			server,
			trace,
			tableName,
		)
	case error2.ObServerIsStopping:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObServerIsStopping",
			"Observer is stopping.",
			server,
			trace,
			tableName,
		)
	case error2.ObTenantNotInServer:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTenantNotInServer",
			"Tenant not in server.",
			server,
			trace,
			tableName,
		)
	case error2.ObTransRpcTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTransRpcTimeout",
			"Translation timeout.",
			server,
			trace,
			tableName,
		)
	case error2.ObNoReadableReplica:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObNoReadableReplica",
			"No readable replica.",
			server,
			trace,
			tableName,
		)
	case error2.ObReplicaNotReadable:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObReplicaNotReadable",
			"Current replica is not readable.",
			server,
			trace,
			tableName,
		)
	case error2.ObTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTimeout",
			"Timeout.",
			server,
			trace,
			tableName,
		)
	case error2.ObTransTimeout:
		msg += fmt.Sprintf(
			"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s, tableName:%s",
			code,
			"ObTransTimeout",
			"Translation timeout.",
			server,
			trace,
			tableName,
		)
	case error2.ObWaitqueueTimeout:
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

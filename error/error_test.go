package error

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProtocolError(t *testing.T) {
	var sequence uint64 = 2
	var uniqueId uint64 = 3160802047
	err := NewProtocolError("127.0.0.1", 8080, ObErrPrimaryKeyDuplicate, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-5024, errCodeName:ObErrPrimaryKeyDuplicate, errMsg:Primary key duplicate., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObErrUnknownTable, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-5200, errCodeName:ObErrUnknownTable, errMsg:Table not exist., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObErrColumnNotFound, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-5031, errCodeName:ObErrColumnNotFound, errMsg:Column not found., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObObjTypeError, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4001, errCodeName:ObObjTypeError, errMsg:Column type not match., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObBadNullError, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4235, errCodeName:ObBadNullError, errMsg:Column doesn't have a default value., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObInvalidArgument, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4002, errCodeName:ObInvalidArgument, errMsg:Arguments invalid., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObDeserializeError, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4034, errCodeName:ObDeserializeError, errMsg:Deserialize error., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObPasswordWrong, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4043, errCodeName:ObPasswordWrong, errMsg:Access deny, password is wrong., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObLocationLeaderNotExist, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4654, errCodeName:ObLocationLeaderNotExist, errMsg:Location leader not exist., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObNotMaster, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4038, errCodeName:ObNotMaster, errMsg:Current node is not master., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObRsNotMaster, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4638, errCodeName:ObRsNotMaster, errMsg:Current Rs is not master., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObRsShutdown, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4550, errCodeName:ObRsShutdown, errMsg:RS had been shut down., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObRpcConnectError, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4265, errCodeName:ObRpcConnectError, errMsg:Connect error., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObPartitionNotExist, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4225, errCodeName:ObPartitionNotExist, errMsg:Partition not exist., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObPartitionIsStopped, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-6228, errCodeName:ObPartitionIsStopped, errMsg:Partition is stopped., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObLocationNotExist, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4653, errCodeName:ObLocationNotExist, errMsg:location not exist., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObServerIsInit, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-8001, errCodeName:ObServerIsInit, errMsg:Observer is initing, please wait., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObServerIsStopping, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-8002, errCodeName:ObServerIsStopping, errMsg:Observer is stopping., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTenantNotInServer, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-5150, errCodeName:ObTenantNotInServer, errMsg:Tenant not in server., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTransRpcTimeout, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-6230, errCodeName:ObTransRpcTimeout, errMsg:Translation timeout., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObNoReadableReplica, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-5296, errCodeName:ObNoReadableReplica, errMsg:No readable replica., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObReplicaNotReadable, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-6231, errCodeName:ObReplicaNotReadable, errMsg:Current replica is not readable., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTimeout, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4012, errCodeName:ObTimeout, errMsg:Timeout., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTransTimeout, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-6210, errCodeName:ObTransTimeout, errMsg:Translation timeout., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObWaitqueueTimeout, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4049, errCodeName:ObWaitqueueTimeout, errMsg:Wait queue timeout., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObErrUnexpected, sequence, uniqueId, "test")
	assert.EqualError(
		t,
		err,
		"errCode:-4016, errCodeName:ObErrUnknown, errMsg:Get an error from observer., server:127.0.0.1:8080, trace:YBC6602FF-0000000000000002, tableName:test",
	)

}

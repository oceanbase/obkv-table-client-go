package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProtocolError(t *testing.T) {
	var sequence uint64 = 2
	var uniqueId uint64 = 3160802047
	err := NewProtocolError("127.0.0.1", 8080, ObErrPrimaryKeyDuplicate, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-5024][ObErrPrimaryKeyDuplicate][Primary key duplicate.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObErrUnknownTable, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-5200][ObErrUnknownTable][Table not exist.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObErrColumnNotFound, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-5031][ObErrColumnNotFound][Column not found.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObObjTypeError, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4001][ObObjTypeError][Column type not match.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObBadNullError, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4235][ObBadNullError][Column doesn't have a default value.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObInvalidArgument, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4002][ObInvalidArgument][Arguments invalid.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObDeserializeError, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4034][ObDeserializeError][Deserialize error.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObPasswordWrong, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4043][ObPasswordWrong][Access deny, password is wrong.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObLocationLeaderNotExist, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4654][ObLocationLeaderNotExist][Location leader not exist.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObNotMaster, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4038][ObNotMaster][Current node is not master.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObRsNotMaster, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4638][ObRsNotMaster][Current Rs is not master.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObRsShutdown, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4550][ObRsShutdown][RS had been shut down.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObRpcConnectError, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4265][ObRpcConnectError][Connect error.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObPartitionNotExist, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4225][ObPartitionNotExist][Partition not exist.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObPartitionIsStopped, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-6228][ObPartitionIsStopped][Partition is stopped.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObLocationNotExist, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4653][ObLocationNotExist][location not exist.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObServerIsInit, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-8001][ObServerIsInit][Observer is initing, please wait.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObServerIsStopping, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-8002][ObServerIsStopping][Observer is stopping.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTenantNotInServer, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-5150][ObTenantNotInServer][Tenant not in server.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTransRpcTimeout, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-6230][ObTransRpcTimeout][Translation timeout.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObNoReadableReplica, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-5296][ObNoReadableReplica][No readable replica.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObReplicaNotReadable, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-6231][ObReplicaNotReadable][Current replica is not readable.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTimeout, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4012][ObTimeout][Timeout.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObTransTimeout, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-6210][ObTransTimeout][Translation timeout.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObWaitqueueTimeout, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4049][ObWaitqueueTimeout][Wait queue timeout.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)
	err = NewProtocolError("127.0.0.1", 8080, ObErrUnexpected, sequence, uniqueId)
	assert.EqualError(
		t,
		err,
		"[-4016][ObErrUnknown][Get an error from observer.][127.0.0.1:8080][YBC6602FF-0000000000000002]",
	)

}

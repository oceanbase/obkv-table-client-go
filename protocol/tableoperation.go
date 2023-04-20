package protocol

type TableOperation struct {
	*UniVersionHeader
	opType TableOperationType
	entity *TableEntity
}

func (o *TableOperation) Encode() []byte {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) PCode() TablePacketCode {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) PayloadLen() int64 {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) PayloadContentLen() int64 {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) SessionId() uint64 {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) SetSessionId(sessionId uint64) {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) Credential() []byte {
	// TODO implement me
	panic("implement me")
}

func (o *TableOperation) SetCredential(credential []byte) {
	// TODO implement me
	panic("implement me")
}

type TableOperationType int32

const (
	Get TableOperationType = iota
	Insert
	Del
	Update
	InsertOrUpdate
	Replace
	Increment
	Append
)

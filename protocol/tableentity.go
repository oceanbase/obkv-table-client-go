package protocol

type ObRowKey struct {
	// keys[]
}

type TableEntity struct {
	*UniVersionHeader
	rouKey *ObRowKey
	// properties map[string]*ObObject
}

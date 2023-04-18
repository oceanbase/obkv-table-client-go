package protocol

type TableEntity struct {
	*UniVersionHeader
	rouKey     *RowKey
	properties map[string]*Object
}

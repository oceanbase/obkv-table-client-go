package route

type ObTableEntryKey struct {
	clusterName  string
	tenantName   string
	databaseName string
	tableName    string
}

func NewObTableEntryKey(
	clusterName string,
	tenantName string,
	databaseName string,
	tableName string) *ObTableEntryKey {
	return &ObTableEntryKey{clusterName, tenantName, databaseName, tableName}
}

func (k *ObTableEntryKey) String() string {
	return "ObTableEntryKey{" +
		"clusterName:" + k.clusterName + ", " +
		"tenantNane:" + k.databaseName + ", " +
		"databaseName:" + k.databaseName + ", " +
		"tableName:" + k.tableName +
		"}"
}

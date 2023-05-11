package route

// not support generate column now, support later
type obGeneratedColumnSimpleFunc interface {
	String() string
	getRefColumnNames() []string
}

package table

import "github.com/oceanbase/obkv-table-client-go/util"

type Column struct {
	Name  string
	Value interface{}
}

func (c *Column) String() string {
	return "Column{" +
		"Name:" + c.Name + ", " +
		"Value:" + util.InterfaceToString(c.Value) +
		"}"
}

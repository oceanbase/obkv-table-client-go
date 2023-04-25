package table

import "github.com/oceanbase/obkv-table-client-go/util"

type Column struct {
	name  string
	value interface{}
}

func NewColumn(name string, value interface{}) *Column {
	return &Column{name: name, value: value}
}

func (c *Column) Name() string {
	return c.name
}

func (c *Column) SetName(name string) {
	c.name = name
}

func (c *Column) Value() interface{} {
	return c.value
}

func (c *Column) SetValue(value interface{}) {
	c.value = value
}

func (c *Column) String() string {
	return "column{" +
		"name: " + c.name + ", " +
		"value: " + util.InterfaceToString(c.value) +
		"}"
}

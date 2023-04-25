package client

import (
	"github.com/oceanbase/obkv-table-client-go/table"
)

func columnsToString(columns []*table.Column) string {
	var str string
	str = str + "["
	for i := 0; i < len(columns); i++ {
		if i > 0 {
			str += ", "
		}
		str += columns[i].String()
	}
	str += "]"
	return str
}

package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn_ToString(t *testing.T) {
	col := Column{}
	assert.Equal(t, col.String(), "column{name: , value: <nil>}")
	col = Column{"c1", 123}
	assert.Equal(t, col.String(), "column{name: c1, value: 123}")
}

func TestObTableParam_ToString(t *testing.T) {
	param := ObTableParam{}
	assert.Equal(t, param.String(), "ObTableParam{table:nil, tableId:0, partitionId:0}")
	param = ObTableParam{nil, 500023, 500012}
	assert.Equal(t, param.String(), "ObTableParam{table:nil, tableId:500023, partitionId:500012}")
}

func TestObRowkeyElement_ToString(t *testing.T) {
	v := ObRowkeyElement{}
	assert.Equal(t, v.String(), "ObRowkeyElement{nameIdxMap:{}}")
	m := make(map[string]int, 3)
	m["c1"] = 0
	m["c2"] = 1
	m["c3"] = 2
	v = ObRowkeyElement{m}
	assert.Equal(t, v.String(), "ObRowkeyElement{nameIdxMap:{m[c1]=0, m[c2]=1, m[c3]=2}}")
}

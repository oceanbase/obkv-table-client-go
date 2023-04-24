package table

import (
	"strconv"
)

type ObRowkeyElement struct {
	nameIdxMap map[string]int
}

func NewObRowkeyElement(nameIdxMap map[string]int) *ObRowkeyElement {
	return &ObRowkeyElement{nameIdxMap}
}

func (e *ObRowkeyElement) NameIdxMap() map[string]int {
	return e.nameIdxMap
}

func (e *ObRowkeyElement) String() string {
	var nameIdxMapStr string
	var i = 0
	nameIdxMapStr = nameIdxMapStr + "{"
	for k, v := range e.nameIdxMap {
		if i > 0 {
			nameIdxMapStr += ", "
		}
		i++
		nameIdxMapStr += "m[" + k + "]=" + strconv.Itoa(v)
	}
	nameIdxMapStr += "}"
	return "ObRowkeyElement{" +
		"nameIdxMap:" + nameIdxMapStr +
		"}"
}

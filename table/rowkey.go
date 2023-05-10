package table

import (
	"strconv"
)

type ObRowKeyElement struct {
	nameIdxMap map[string]int
}

func NewObRowKeyElement(nameIdxMap map[string]int) *ObRowKeyElement {
	return &ObRowKeyElement{nameIdxMap}
}

func (e *ObRowKeyElement) NameIdxMap() map[string]int {
	return e.nameIdxMap
}

func (e *ObRowKeyElement) String() string {
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
	return "ObRowKeyElement{" +
		"nameIdxMap:" + nameIdxMapStr +
		"}"
}

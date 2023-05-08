package route

import (
	"strconv"
	"strings"
)

type ObServerInfo struct {
	stopTime int64
	status   string // Active/InActive/Deleting
}

func (i *ObServerInfo) IsActive() bool {
	return i.stopTime == 0 && strings.EqualFold(i.status, "active") // ignore case
}

func (i *ObServerInfo) String() string {
	return "ObServerInfo{" +
		"stopTime:" + strconv.Itoa(int(i.stopTime)) + ", " +
		"status:" + i.status +
		"}"
}

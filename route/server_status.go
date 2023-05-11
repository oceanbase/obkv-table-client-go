package route

import (
	"strconv"
	"strings"
)

type obServerStatus struct {
	stopTime int64
	status   string // Active/InActive/Deleting
}

func newServerStatus(stopTime int64, status string) *obServerStatus {
	return &obServerStatus{stopTime, status}
}

func (i *obServerStatus) IsActive() bool {
	return i.stopTime == 0 && strings.EqualFold(i.status, "active") // ignore case
}

func (i *obServerStatus) String() string {
	return "obServerStatus{" +
		"stopTime:" + strconv.Itoa(int(i.stopTime)) + ", " +
		"status:" + i.status +
		"}"
}

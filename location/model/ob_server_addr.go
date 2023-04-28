package model

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
	"time"
)

type ObServerAddr struct {
	ip                string
	sqlPort           int
	svrPort           int
	priority          int
	grantPriorityTime int64
	lastAccessTime    int64
}

// IsExpired Whether the addr is expired given the timeout.
func (t *ObServerAddr) IsExpired(cachingTimeout int64) (bool, error) {
	return true, nil
}

// RecordAccess Record the access time.
func (t *ObServerAddr) RecordAccess() {
	t.lastAccessTime = time.Now().UnixMilli()
}

// GetIp Get ip.
func (t *ObServerAddr) GetIp() string {
	return t.ip
}

// SetIp Set ip.
func (t *ObServerAddr) SetIp(ip string) {
	t.ip = ip
}

// SetAddress Set address.
func (t *ObServerAddr) SetAddress(address string) {
	var err error
	if strings.Contains(address, ":") {
		t.ip = strings.Split(address, ":")[0]
		t.svrPort, err = strconv.Atoi(strings.Split(address, ":")[1])
		if err != nil {
			// TODO
			// log
		}
	} else {
		t.ip = address
	}
}

// GetSqlPort Get sql port.
func (t *ObServerAddr) GetSqlPort() int {
	return t.sqlPort
}

// SetSqlPort Set sql port.
func (t *ObServerAddr) SetSqlPort(sqlPort int) {
	t.sqlPort = sqlPort
}

// GetSvrPort Get svr port.
func (t *ObServerAddr) GetSvrPort() int {
	return t.svrPort
}

// SetSvrPort Set svr port.
func (t *ObServerAddr) SetSvrPort(svrPort int) {
	t.svrPort = svrPort
}

// GetPriority Get priority.
func (t *ObServerAddr) GetPriority() int {
	return t.priority
}

// SetGrantPriorityTime Set grant priority time.
func (t *ObServerAddr) SetGrantPriorityTime(grantPriorityTime int64) {
	t.grantPriorityTime = grantPriorityTime
}

// GetLastAccessTime Get Last Access Time.
func (t *ObServerAddr) GetLastAccessTime() int64 {
	return t.lastAccessTime
}

// Equals Equals
func (t *ObServerAddr) Equals(o interface{}) bool {
	if t == o {
		return true
	}
	return false
}

// HashCode Hash code.
func (t *ObServerAddr) HashCode() int {
	v := int(crc32.ChecksumIEEE([]byte(t.ip))) + t.sqlPort + t.svrPort
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}

// ToString To string.
func (t *ObServerAddr) ToString() string {
	return fmt.Sprintf("ObServerAddr{ip='%s', sqlPort=%d, svrPort=%d", t.ip, t.sqlPort, t.svrPort)
}

func (t *ObServerAddr) CompareTo(that ObServerAddr) int {
	thisValue := t.GetPriority()
	thatValue := that.GetPriority()
	if thisValue < thatValue {
		return 1
	}
	if thisValue > thatValue {
		return -1
	}
	return 0
}

package util

import (
	"sync"
)

var globalObVersion float32 = 0.0
var obVersionGuard sync.Mutex

func ObVersion() float32 {
	return globalObVersion
}

func SetObVersion(version float32) {
	obVersionGuard.Lock()
	globalObVersion = version
	obVersionGuard.Unlock()
}

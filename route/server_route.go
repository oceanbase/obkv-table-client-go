package route

import "strconv"

const (
	ObReadConsistencyStrong = 0
	ObReadConsistencyWeak   = 1
)

type ObReadConsistency int

type ObServerRoute struct {
	readConsistency ObReadConsistency
}

func (r *ObServerRoute) String() string {
	return "ObServerRoute{" +
		"readConsistency:" + strconv.Itoa(int(r.readConsistency)) +
		"}"
}

func NewObServerRoute(readOnly bool) *ObServerRoute {
	if readOnly {
		// todo: adapt java client
		return &ObServerRoute{ObReadConsistencyWeak}
	}
	return &ObServerRoute{ObReadConsistencyStrong}
}

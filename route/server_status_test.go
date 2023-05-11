package route

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObServerStatus_IsActive(t *testing.T) {
	status := &obServerStatus{}
	assert.False(t, status.IsActive())
	status = newServerStatus(0, "Active")
	assert.True(t, status.IsActive())
	status = newServerStatus(0, "active")
	assert.True(t, status.IsActive())
	status = newServerStatus(1, "active")
	assert.False(t, status.IsActive())
	status = newServerStatus(0, "InActive")
	assert.False(t, status.IsActive())
	status = newServerStatus(1, "InActive")
	assert.False(t, status.IsActive())
}

func TestObServerStatus_String(t *testing.T) {
	status := &obServerStatus{}
	assert.Equal(t, "obServerStatus{stopTime:0, status:}", status.String())
	status = newServerStatus(0, "Active")
	assert.Equal(t, "obServerStatus{stopTime:0, status:Active}", status.String())
}

package model

import "fmt"

type OcpResponseDataRs struct {
	address  string
	role     string
	sql_port int
}

// GetAddress get address.
func (t *OcpResponseDataRs) GetAddress() string {
	return t.address
}

// SetAddress set address.
func (t *OcpResponseDataRs) SetAddress(address string) {
	t.address = address
}

// GetRole get role.
func (t *OcpResponseDataRs) GetRole() string {
	return t.role
}

// SetRole set role.
func (t *OcpResponseDataRs) SetRole(role string) {
	t.role = role
}

// GetSQLPort get role.
func (t *OcpResponseDataRs) GetSQLPort() int {
	return t.sql_port
}

// SetSQLPort set role.
func (t *OcpResponseDataRs) SetSQLPort(sql_port int) {
	t.sql_port = sql_port
}

// ToString To string
func (t *OcpResponseDataRs) ToString() string {
	return fmt.Sprintf("OcpResponseDataRs{address='%s',role='%s',sql_port=%d}",
		t.address, t.role, t.sql_port)
}

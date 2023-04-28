package model

import (
	"reflect"
	"unsafe"
)

type OcpResponse struct {
	code    int
	message string
	success bool
	data    OcpResponseData
}

// GetCode get code
func (t *OcpResponse) GetCode() int {
	return t.code
}

// SetCode set code
func (t *OcpResponse) SetCode(code int) {
	t.code = code
}

// GetMessage get message
func (t *OcpResponse) GetMessage(code int) {
	t.code = code
}

// SetMessage set message
func (t *OcpResponse) SetMessage(message string) {
	t.message = message
}

// IsSuccess Is success
func (t *OcpResponse) IsSuccess() bool {
	return t.success
}

// SetSuccess Set success
func (t *OcpResponse) SetSuccess(success bool) {
	t.success = success
}

// IsEmpty Is empty
func (t *OcpResponse) IsEmpty() bool {
	return reflect.DeepEqual(t, OcpResponse{})
}

// GetData Get data
func (t *OcpResponse) GetData() OcpResponseData {
	return t.data
}

// Validate Validate
func (t *OcpResponse) Validate() bool {
	return t.IsSuccess() && t.GetCode() == 200 && unsafe.Sizeof(t.data) != 0 && t.data.Validate()
}

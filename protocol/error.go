/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package protocol

import (
	"fmt"

	error2 "github.com/oceanbase/obkv-table-client-go/error"
)

type ObProtocolError struct {
	server      string
	trace       string
	errCode     error2.ObErrorCode
	errCodeName string
	errMsg      string
}

// NewProtocolError construct ObProtocolError
func NewProtocolError(
	server string,
	code error2.ObErrorCode,
	errMsg string,
	sequence uint64,
	uniqueId uint64) *ObProtocolError {

	trace := fmt.Sprintf("Y%X-%016X", uniqueId, sequence)
	errCodeName := code.GetErrorCodeName()
	msg := errMsg
	if errMsg == "" {
		msg = "error occur in server"
	}
	return &ObProtocolError{server, trace, code, errCodeName, msg}
}

func (e *ObProtocolError) Error() string {
	return fmt.Sprintf(
		"errCode:%d, errCodeName:%s, errMsg:%s, server:%s, trace:%s",
		e.errCode,
		e.errCodeName,
		e.errMsg,
		e.server,
		e.trace,
	)
}

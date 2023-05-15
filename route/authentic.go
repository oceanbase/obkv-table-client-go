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

package route

type ObUserAuth struct {
	userName string
	password string
}

func (a *ObUserAuth) Password() string {
	return a.password
}

func (a *ObUserAuth) UserName() string {
	return a.userName
}

func NewObUserAuth(userName string, password string) *ObUserAuth {
	return &ObUserAuth{userName, password}
}

func (a *ObUserAuth) String() string {
	return "ObUserAuth{" +
		"userName:" + a.userName + ", " +
		"password:" + a.password +
		"}"
}

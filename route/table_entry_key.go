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

type ObTableEntryKey struct {
	clusterName  string
	tenantName   string
	databaseName string
	tableName    string
}

func NewObTableEntryKey(
	clusterName string,
	tenantName string,
	databaseName string,
	tableName string) *ObTableEntryKey {
	return &ObTableEntryKey{clusterName, tenantName, databaseName, tableName}
}

func (k *ObTableEntryKey) String() string {
	return "ObTableEntryKey{" +
		"clusterName:" + k.clusterName + ", " +
		"tenantNane:" + k.databaseName + ", " +
		"databaseName:" + k.databaseName + ", " +
		"tableName:" + k.tableName +
		"}"
}

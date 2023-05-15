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

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type DB = sql.DB
type Rows = sql.Rows

// NewDB create a DB, remember call DB.close when exit.
// Use the dsn connection mode.
func NewDB(
	userName string,
	password string,
	ip string,
	port string,
	database string) (*DB, error) {
	// "userName:password@tcp(ip:port)/database?charset=utf8"
	dsn := strings.Join([]string{userName, ":", password,
		"@tcp(", ip, ":", port, ")/", database, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithMessagef(err, "open db, dsn:%s", dsn)
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.WithMessagef(err, "ping db, dsn:%s", dsn)
	}
	return db, nil
}

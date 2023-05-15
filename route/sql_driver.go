package route

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type DB = sql.DB
type Rows = sql.Rows

// NewDB create a DB, remember call DB.close when exit
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

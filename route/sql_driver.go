package route

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oceanbase/obkv-table-client-go/log"
	"strings"
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
	path := strings.Join([]string{userName, ":", password,
		"@tcp(", ip, ":", port, ")/", database, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Warn("failed to open db", log.String("path", path))
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Warn("failed to ping db",
			log.String("ip", ip),
			log.String("port", port))
		return nil, err
	}
	return db, nil
}

package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	/*
	dbSource := fmt.Sprintf(
		"Fernando:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)
	*/
	d, err := sql.Open("mysql", "root:1234@tcp(schedule-db)/agendamiento")
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}

package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "lppmwira_golanguser"
	dbPass := "zQ0*WAuu5(kK"
	dbName := "lppmwira_golang"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(lppmwiraraja.ac.id)/"+dbName)
	return db, err
}

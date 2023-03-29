package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "admin"
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

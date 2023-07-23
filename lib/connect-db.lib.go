package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username  = GetEnv("DB_USERNAME")
	password  = GetEnv("DB_PASSWORD")
	host  = GetEnv("DB_HOST")
	port  = GetEnv("DB_PORT")
	db = GetEnv("DB_DBNAME")
)

var DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, db)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", DSN)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

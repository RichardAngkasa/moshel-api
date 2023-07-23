package lib

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username  = GetEnv("DB_USERNAME")
	password  = GetEnv("DB_PASSWORD")
	host  = GetEnv("DB_HOST")
	port  = GetEnv("DB_PORT")
	db = GetEnv("DB_DBNAME")
	db_uri = GetEnv("DB_URI")
)

// var DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, db)

func ConnectDB() *sql.DB {
	log.Println("db", db_uri)
	db, err := sql.Open("mysql", db_uri)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

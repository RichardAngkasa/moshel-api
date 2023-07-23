package lib

import (
	"database/sql"
	"fmt"
	"log"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
var connectStr = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", username, password, db, host, port)

func ConnectDB() *sql.DB {
	log.Println("DSN", connectStr)
	db, err := sql.Open("postgres", db_uri)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

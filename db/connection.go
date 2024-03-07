package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "C:/SQLite3/databases/Tasks_DB.db")
	if err != nil {
		log.Fatal(err)
	}

	if DB == nil {
		log.Fatal("Could not connect to database")
	}
}

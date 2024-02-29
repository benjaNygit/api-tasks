package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB
var port = 5432
var database = "Tasks_DB"
var DSN = "host=localhost user=benja password= database=Tasks_DB port=5432"
var connString = fmt.Sprintf("port=%d;database=%s;fedauth=ActiveDirectoryDefault;", port, database)

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}

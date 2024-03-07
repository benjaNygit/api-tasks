package db

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Migrate() {
	ConnectDB()
	defer DB.Close()

	filepath := "SQL/CreateDatabase.sql"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Convertir el []byte a string
	sqlScript := string(content)

	_, err = DB.Exec(sqlScript)
	if err != nil {
		log.Fatal(err)
	}

	var name string
	tables, _ := DB.Query("SELECT name FROM sqlite_master WHERE type = 'table' ORDER BY name")
	for tables.Next() {
		tables.Scan((&name))
		fmt.Printf("Tabla Creada: %s\n", name)
	}

	log.Print("Se migró la base de datos")
}

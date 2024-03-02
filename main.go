package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/benjaNygit/api-tasks/db"
)

func main() {
	// Comandos
	if len(os.Args) > 1 {
		// Migración de la base de datos
		if os.Args[1] == "migrate" {
			fmt.Println("Migrando la base de datos...")
			db.Migrate()
		}
		return
	}

	db.ConnectDB()
	defer db.DB.Close()

	query := "INSERT INTO category (Description) VALUES ('ff')"
	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// router := mux.NewRouter()

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World Welcome to API v1")
	// }).Methods("GET")

	// http.ListenAndServe(":3000", router)
}

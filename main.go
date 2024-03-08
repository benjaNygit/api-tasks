package main

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/benjaNygit/api-tasks/db"
	"github.com/benjaNygit/api-tasks/routes"
	mux "github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.HandleFunc("/", routes.IndexHandler)

	router.HandleFunc("/Category", routes.CategoryPostHandler).Methods("POST")
	router.HandleFunc("/Category", routes.CategoryGetHandler).Methods("GET")
	router.HandleFunc("/Category", routes.CategoryPatchHandler).Methods("PATCH")

	http.ListenAndServe(":3000", router)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	db "github.com/benjaNygit/api-tasks/db"
	"github.com/benjaNygit/api-tasks/models"
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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Welcome to API v1")
	}).Methods("GET")

	router.HandleFunc("/Category", func(w http.ResponseWriter, r *http.Request) {
		var category models.Category
		json.NewDecoder(r.Body).Decode(&category)

		err := db.Create(category)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(category)
	}).Methods("POST")

	http.ListenAndServe(":3000", router)
}

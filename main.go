package main

import (
	"log"

	db "github.com/benjaNygit/api-tasks/db"
)

func main() {
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

package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/benjaNygit/api-tasks/DB"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDB()
	defer db.DB.Close()

	x, err := db.DB.Prepare("insert into dbo.Category values (1,'test 1')")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Welcome to API v1")
	}).Methods("GET")

	http.ListenAndServe(":3000", router)
}

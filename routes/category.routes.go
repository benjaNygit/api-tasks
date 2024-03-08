package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benjaNygit/api-tasks/db"
	"github.com/benjaNygit/api-tasks/models"
	"github.com/gorilla/mux"
)

func CategoryPostHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	err := db.Create(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}

func CategoryGetHandler(w http.ResponseWriter, r *http.Request) {
	category, err := db.CategoryGetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&category)
}

func CategoryPatchHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	err := db.Update(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}

func CategoryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	params := mux.Vars(r)
	s := params["code"]

	fmt.Println(s)

	err := db.Delete(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}

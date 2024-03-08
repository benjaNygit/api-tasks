package routes

import (
	"encoding/json"
	"net/http"

	"github.com/benjaNygit/api-tasks/db"
	"github.com/benjaNygit/api-tasks/models"
)

func CategoryPostHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	err := db.Create(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(category)
}

func CategoryGetHandler(w http.ResponseWriter, r *http.Request) {
	category, err := db.CategoryGetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&category)
}

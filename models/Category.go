package models

import (
	"log"

	"github.com/benjaNygit/api-tasks/db"
)

type Category struct {
	Code        uint   `json:"code"`
	Description string `json:"description"`
}

func (c *Category) Create() bool {
	db.ConnectDB()
	defer db.DB.Close()

	query := db.Insert("Category", c.Description)

	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (c *Category) Read() Category {
	return Category{Code: c.Code}
}

func (c *Category) Update() Category {
	return Category{Code: c.Code}
}

func (c *Category) Delete() bool {
	db.ConnectDB()
	defer db.DB.Close()

	query := db.Delete("Category", string(rune(c.Code)))

	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

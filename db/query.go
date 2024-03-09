package db

import (
	"fmt"

	"github.com/benjaNygit/api-tasks/models"
)

// Category
func CategoryGetAll() ([]models.Category, error) {
	var category []models.Category
	var c models.Category

	rows, err := DB.Query(`SELECT Code, Description FROM Category`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			&c.Code,
			&c.Description,
		)

		category = append(category, c)
	}

	return category, nil
}

func CategoryGet(code uint64) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf("SELECT Code, Description FROM Category WHERE Code = %d", code)
	rows, err := DB.Query(query)
	if err != nil {
		return category, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(
			category.Code,
			category.Description,
		)
	}

	return category, nil
}

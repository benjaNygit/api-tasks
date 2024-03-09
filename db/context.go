package db

import (
	"fmt"

	"github.com/benjaNygit/api-tasks/models"
)

// Model Category
func Create(c models.Category) error {
	stmt, _ := DB.Prepare(`INSERT INTO Category (Description) VALUES (?)`)
	defer stmt.Close()

	_, err := stmt.Exec(c.Description)
	if err != nil {
		return err
	}

	return nil
}

func Update(c models.Category) error {
	stmt, _ := DB.Prepare(`UPDATE Category SET Description = ? WHERE Code = ?`)
	defer stmt.Close()

	_, err := stmt.Exec(c.Description, c.Code)
	if err != nil {
		return err
	}

	return nil
}

func Delete(c models.Category) error {
	stmt, _ := DB.Prepare(`DELETE Category WHERE Code = ?`)
	defer stmt.Close()

	fmt.Println(c.Code)
	_, err := stmt.Exec(c.Code)
	if err != nil {
		return err
	}

	return nil
}

package db

import "fmt"

func Insert(table string, values ...string) string {
	switch table {
	case "Category":
		return fmt.Sprintf("INSERT INTO Category (Code, Description) VALUES ('%s', '%s')", values[0], values[1])
	case "Priority":
		return fmt.Sprintf("INSERT INTO Priority (Code, Description) VALUES ('%s', '%s')", values[0], values[1])
	case "Task":
		return fmt.Sprintf("INSERT INTO Task (Id, Title, Description, CategoryCode, PriorityCode, CreatedAt, Done, UserId) VALUES (%s, '%s', '%s', '%s', '%s', '%s', '%s', '%s')", values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7])
	case "User":
		return fmt.Sprintf("INSERT INTO User (Id, Name, Email, Password, CreatedAt) VALUES ('%s', '%s', '%s', '%s', '%s')", values[0], values[1], values[2], values[3], values[4])
	}

	return ""
}

func Delete(table string, pk string) string {
	switch table {
	case "Category":
		return fmt.Sprintf("DELETE FROM Category WHERE Code = '%s'", pk)
	case "Priority":
		return fmt.Sprintf("DELETE FROM Priority WHERE Code = '%s'", pk)
	case "Task":
		return fmt.Sprintf("DELETE FROM Task WHERE Id = '%s'", pk)
	case "User":
		return fmt.Sprintf("DELETE FROM User WHERE Id = '%s'", pk)
	}

	return ""
}

func Update(table string, pk string, values ...string) string {
	switch table {
	case "Category":
		return fmt.Sprintf("UPDATE Category SET Description = '%s'", values[0])
	case "Priority":
		return fmt.Sprintf("UPDATE Priority SET Description = '%s'", values[0])
	case "Task":
		return fmt.Sprintf("UPDATE Task SET Title = '%s', Description = '%s', CategoryCode = '%s', PriorityCode, Done = '%s'", values[0], values[1], values[2], values[3])
	case "User":
		return fmt.Sprintf("UPDATE User SET Name = '%s', Email = '%s', Password = '%s', LastSession = '%s'", values[0], values[1], values[2], values[3])
	}

	return ""
}

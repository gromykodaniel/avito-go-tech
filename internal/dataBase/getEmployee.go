package dataBase

import (
	"context"
	"fmt"

	"github.com/azoma13/avito/models"
)

func GetEmployeeDB(username string) (models.Employee, error) {

	var employee models.Employee

	query := `SELECT id, username, password, balance FROM employee
		WHERE username = lower($1)`
	row := DB.QueryRow(context.Background(), query, username)

	err := row.Scan(&employee.ID, &employee.Username, &employee.Password, &employee.Balance)
	if err != nil {
		return models.Employee{}, fmt.Errorf("error scan employee for dataBase: %w", err)
	}

	return employee, nil
}

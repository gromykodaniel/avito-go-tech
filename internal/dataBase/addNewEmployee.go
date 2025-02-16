package dataBase

import (
	"context"
	"fmt"
)

func AddNewEmployeeDB(username, hashPassword string) error {

	exec := `INSERT INTO employee
		(username, password) 
		VALUES (lower($1), $2)`
	_, err := DB.Exec(context.Background(), exec, username, hashPassword)
	if err != nil {
		return fmt.Errorf("error exec new employee: %w", err)
	}

	return nil
}

package utils

import (
	"fmt"

	"github.com/azoma13/avito/internal/dataBase"
	"golang.org/x/crypto/bcrypt"
)

func RegisterEmployee(username, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = dataBase.AddNewEmployeeDB(username, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("error add new employee for dataBase: %w", err)
	}
	return err
}

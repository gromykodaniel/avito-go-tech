package dataBase

import (
	"context"
	"fmt"

	"github.com/azoma13/avito/models"
)

func SendCoinDB(employee models.Employee, employeeToUser models.Employee, sendCoin models.SentCoinRequest) error {

	tx, err := DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error to start transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	exec := `UPDATE employee
		SET balance = balance - $1
			WHERE username = $2`
	_, err = tx.Exec(context.Background(), exec, sendCoin.Amount, employee.Username)
	if err != nil {
		return fmt.Errorf("error exec update balance employee: %w", err)
	}

	exec = `UPDATE employee
	SET balance = balance + $1
		WHERE username = $2`
	_, err = tx.Exec(context.Background(), exec, sendCoin.Amount, employeeToUser.Username)
	if err != nil {
		return fmt.Errorf("error exec update balance toUser: %w", err)
	}

	exec = `INSERT INTO transaction
			(fromUser, toUser, amount) 
			VALUES ($1, $2, $3);`
	_, err = DB.Exec(context.Background(), exec, employee.Username, employeeToUser.Username, sendCoin.Amount)
	if err != nil {
		return fmt.Errorf("error exec insert transaction: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error commit transaction: %w", err)
	}

	return nil
}

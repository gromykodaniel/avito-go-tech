package dataBase

import (
	"context"
	"fmt"

	"github.com/azoma13/avito/models"
)

func PayBuyMerchDB(employee models.Employee, merch models.Merch) error {

	tx, err := DB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error to start transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	exec := `UPDATE employee
		SET balance = balance - $1
			WHERE username = $2`
	_, err = tx.Exec(context.Background(), exec, merch.Price, employee.Username)
	if err != nil {
		return fmt.Errorf("error exec update balance employee: %w", err)
	}

	quantity := 0

	query := `SELECT quantity FROM inventory 
		WHERE user_id = $1 AND item_id = $2`
	err = tx.QueryRow(context.Background(), query, employee.ID, merch.ID).Scan(&quantity)

	if err != nil {

		exec = `INSERT INTO inventory 
			(user_id, item_id, quantity) 
			VALUES ($1, $2, 1);`
		_, err = tx.Exec(context.Background(), exec, employee.ID, merch.ID)
		if err != nil {
			return fmt.Errorf("error exec insert inventory: %w", err)
		}

	} else {

		exec = `UPDATE inventory 
			SET quantity = quantity + 1 
				WHERE user_id = $1 AND item_id = $2`
		_, err = tx.Exec(context.Background(), exec, employee.ID, merch.ID)
		if err != nil {
			return fmt.Errorf("error exec update inventory: %w", err)
		}

	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("error commit transaction: %w", err)
	}

	return nil
}

package dataBase

import (
	"context"
	"fmt"

	"github.com/azoma13/avito/models"
)

func GetMerchDB(itemName string) (models.Merch, error) {

	var merch models.Merch

	query := `SELECT id, type, price FROM merch
		WHERE type = $1`
	row := DB.QueryRow(context.Background(), query, itemName)

	err := row.Scan(&merch.ID, &merch.Type, &merch.Price)
	if err != nil {
		return models.Merch{}, fmt.Errorf("error scan merch for dataBase: %w", err)
	}

	return merch, nil
}

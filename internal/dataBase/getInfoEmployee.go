package dataBase

import (
	"context"
	"fmt"

	"github.com/azoma13/avito/models"
)

func GetInfoEmployeeDB(employee models.Employee) (models.InfoResponse, error) {

	var infoRes models.InfoResponse

	infoRes.Coins = employee.Balance
	var err error
	infoRes.Inventory, err = getInventoryEmployeeDB(employee)
	if err != nil {
		return models.InfoResponse{}, fmt.Errorf("error func getInventoryEmployeeDB: %w", err)
	}

	infoRes.CoinHistory, err = getCoinHistoryEmployeeDB(employee)
	if err != nil {
		return models.InfoResponse{}, fmt.Errorf("error in func getCoinHistoryEmployeeDB: %w", err)
	}

	return infoRes, nil
}

func getCoinHistoryEmployeeDB(employee models.Employee) (models.CoinHistory, error) {

	var coinHistory models.CoinHistory
	var err error
	coinHistory.Received, err = getCoinHistoryReceivedDB(employee)
	if err != nil {
		return models.CoinHistory{}, fmt.Errorf("error in func getCoinHistoryReceivedDB: %w", err)
	}

	coinHistory.Sent, err = getCoinHistorySentDB(employee)
	if err != nil {
		return models.CoinHistory{}, fmt.Errorf("error in func getCoinHistorySentDB: %w", err)
	}

	return coinHistory, nil
}

func getCoinHistorySentDB(employee models.Employee) ([]models.Sent, error) {

	query := `SELECT toUser, amount FROM transaction
		WHERE fromUser = $1`
	rows, err := DB.Query(context.Background(), query, employee.Username)
	if err != nil {
		return nil, fmt.Errorf("error query sent: %w", err)
	}
	defer rows.Close()

	var sent []models.Sent
	for rows.Next() {

		var s models.Sent
		err := rows.Scan(&s.ToUser, &s.Amount)
		if err != nil {
			return nil, fmt.Errorf("error scan sent: %w", err)
		}
		sent = append(sent, s)

	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error cursor rows sent: %w", err)
	}

	if sent == nil {
		sent = []models.Sent{}
	}

	return sent, nil
}

func getCoinHistoryReceivedDB(employee models.Employee) ([]models.Received, error) {

	query := `SELECT fromUser, amount FROM transaction
		WHERE toUser = $1`
	rows, err := DB.Query(context.Background(), query, employee.Username)
	if err != nil {
		return nil, fmt.Errorf("error query received: %w", err)
	}
	defer rows.Close()

	var received []models.Received
	for rows.Next() {

		var receiv models.Received
		err := rows.Scan(&receiv.FromUser, &receiv.Amount)
		if err != nil {
			return nil, fmt.Errorf("error scan received: %w", err)
		}
		received = append(received, receiv)

	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error cursor rows received: %w", err)
	}

	if received == nil {
		received = []models.Received{}
	}

	return received, nil
}

func getInventoryEmployeeDB(employee models.Employee) ([]models.Item, error) {
	query := `SELECT item_id, quantity FROM inventory
		WHERE user_id = $1`
	rows, err := DB.Query(context.Background(), query, employee.ID)
	if err != nil {
		return nil, fmt.Errorf("error query in func getInventoryEmployeeDB: %w", err)
	}
	defer rows.Close()

	var inventory []models.Item
	for rows.Next() {
		var item models.Item
		var itemID int
		err := rows.Scan(&itemID, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error scan in func GetAllTasks: %w", err)
		}
		item.Type, err = GetItemNameDB(itemID)
		if err != nil {
			return nil, fmt.Errorf("error func GetItemNameDB: %w", err)
		}
		inventory = append(inventory, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error cursor rows in func getInventoryEmployeeDB: %w", err)
	}

	if inventory == nil {
		inventory = []models.Item{}
	}

	return inventory, nil
}

func GetItemNameDB(itemID int) (string, error) {
	query := `SELECT type FROM merch
		WHERE id = $1`
	row := DB.QueryRow(context.Background(), query, itemID)
	var itemName string
	err := row.Scan(&itemName)
	if err != nil {
		return "", fmt.Errorf("error scan item name for dataBase: %w", err)
	}
	return itemName, nil
}

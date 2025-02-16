package dataBase

import (
	"context"
	"log"
)

func CreateTableDB() error {
	query := `CREATE TABLE IF NOT EXISTS employee 
		(
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			balance INT DEFAULT 1000
		);
		CREATE TABLE IF NOT EXISTS merch 
		(
			id SERIAL PRIMARY KEY,
			type TEXT UNIQUE NOT NULL,
			price INT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS inventory 
		(
			id SERIAL PRIMARY KEY,
			user_id INT,
			item_id INT,
			quantity INT DEFAULT 0,
			FOREIGN KEY (user_id) REFERENCES employee(id),
			FOREIGN KEY (item_id) REFERENCES merch(id)
		);
		CREATE TABLE IF NOT EXISTS transaction
		(
			id SERIAL PRIMARY KEY,
			fromUser TEXT,
			toUser TEXT,
			amount INT DEFAULT 0,
			FOREIGN KEY (fromUser) REFERENCES employee(username),
			FOREIGN KEY (toUser) REFERENCES employee(username)
		);`
	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("error in compiling tables: %w", err)
	}
	return err
}

func CreateShopMerchDB() {
	itemsMerch := map[string]int{
		"t-shirt":    80,
		"cup":        20,
		"book":       50,
		"pen":        10,
		"powerbank":  200,
		"hoody":      300,
		"umbrella":   200,
		"socks":      10,
		"wallet":     50,
		"pink-hoody": 500,
	}

	exec := `INSERT INTO merch
			(id, type, price) VALUES
			(DEFAULT, $1, $2)
		ON CONFLICT (type)
			DO NOTHING;`
	for key, val := range itemsMerch {
		_, err := DB.Exec(context.Background(), exec, key, val)
		if err != nil {
			log.Fatal("Error adding merch: %w", err)
		}
	}
}

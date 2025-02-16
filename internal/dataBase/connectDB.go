package dataBase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/azoma13/avito/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB *pgxpool.Pool

func ConnectToDB() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", configs.UsernamePG, configs.PasswordPG, configs.HostPG, configs.PortPG, configs.DataBasePG))
	if err != nil {
		log.Fatalf("Unable to parse database config: %v", err)
	}

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	time.Sleep(5 * time.Second)

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	CreateTableDB()
	CreateShopMerchDB()
	return DB
}

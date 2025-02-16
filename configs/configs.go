package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	UsernamePG string
	PasswordPG string
	DataBasePG string
	HostPG     string
	PortPG     string
	PortAPI    string
	JwtKey     []byte
)

func Environment() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error load .env: %w", err)
	}

	UsernamePG = os.Getenv("PG_USERNAME")
	PasswordPG = os.Getenv("PG_PASSWORD")
	DataBasePG = os.Getenv("PG_DATABASE")
	HostPG = os.Getenv("PG_HOST")
	PortPG = os.Getenv("PG_PORT")
	PortAPI = os.Getenv("API_PORT")
	JwtKey = []byte(os.Getenv("SECRET_KEY"))

	return nil
}

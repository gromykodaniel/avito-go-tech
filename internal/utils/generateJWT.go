package utils

import (
	"fmt"

	"github.com/azoma13/avito/configs"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username})
	stringToken, err := token.SignedString(configs.JwtKey)
	if err != nil {
		return "", fmt.Errorf("error signed string for jwt token")
	}
	return stringToken, nil
}

package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/azoma13/avito/configs"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(r *http.Request) (string, error) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("missing token")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("error signing method")
		}
		return configs.JwtKey, nil

	})
	if err != nil {
		return "", fmt.Errorf("error parse jwt")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", errors.New("invalid token")
}

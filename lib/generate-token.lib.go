package lib

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username, sk string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(2 * time.Hour).Unix()

	secretKey := []byte(sk)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

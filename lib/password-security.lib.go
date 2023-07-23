package lib

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string, c string) (string, error) {
	costFactor, _ := strconv.Atoi(c)
	if costFactor < bcrypt.MinCost || costFactor > bcrypt.MaxCost {
		costFactor = bcrypt.DefaultCost
	}

	hashedString, err := bcrypt.GenerateFromPassword([]byte(p), costFactor)
	if err != nil {
		return "", err
	}

	return string(hashedString), nil
}

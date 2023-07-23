package lib

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p, c string) (string, error) {
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

func UnHashPassword(p, hashedString string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(p))

	if err == nil {
		return nil
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("wrong password")
	} else {
		return errors.New("something went wrong")
	}
}

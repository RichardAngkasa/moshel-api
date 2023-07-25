package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	result, isExist := os.LookupEnv(name)
	if isExist {
		return result
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal("Something wnet wrong went getting the environment")
	}

	return os.Getenv(name)
}

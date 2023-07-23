package lib

import (
	"os"
)

func GetEnv(name string) string {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	return os.Getenv(name)
}

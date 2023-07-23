package lib

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	result, err := HashPassword("Hello World", "10")

	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println("Hashed Password:", result)
}

func TestGenerateToken(t *testing.T) {
	result, err := GenerateToken("John Doe", "is this a testing?")

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	fmt.Println(result)

}

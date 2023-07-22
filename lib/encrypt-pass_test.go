package lib

import (
	"fmt"
	"testing"
)

func TestEncryptPass(t *testing.T) {
	encryptedPass, _ := EncryptPass("Hello, this is a secret message!")
	if encryptedPass == "" {
		t.Fatal("Function return empty string")
	}
	fmt.Println("Result", encryptedPass)
}

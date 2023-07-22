package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptPass(password string, sk string) (string, error) {
	block, err := aes.NewCipher([]byte(sk))

	if err != nil {
		return "", err
	}

	byteString := []byte(password)
	iv := make([]byte, aes.BlockSize)
	cfb := cipher.NewCFBEncrypter(block, iv)

	cipherText := make([]byte, len(byteString))
	cfb.XORKeyStream(cipherText, byteString)

	return encode(cipherText), nil
}

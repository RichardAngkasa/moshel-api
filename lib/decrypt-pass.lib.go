package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Decrypt(encryptedData, sk string) (string, error) {
	block, err := aes.NewCipher([]byte(sk))

	if err != nil {
		return "", err
	}

	cipherText, _ := decode(encryptedData)
	iv := make([]byte, aes.BlockSize)
	cfb := cipher.NewCFBDecrypter(block, iv)
	text := make([]byte, len(cipherText))
	cfb.XORKeyStream(text, cipherText)

	return string(text), nil
}

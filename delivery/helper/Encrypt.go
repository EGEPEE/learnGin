package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func GCM_encrypt(key string, plaintext string, additionalData []byte) string {

	iv := getIV()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, iv, []byte(plaintext), additionalData)
	return hex.EncodeToString(ciphertext)
}

func GCM_decrypt(key string, ct string, additionalData []byte) string {
	iv := getIV()

	ciphertext, _ := hex.DecodeString(ct)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, additionalData)
	if err != nil {
		panic(err.Error())
	}
	s := string(plaintext[:])
	return s
}

func getIV() []byte {
	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}

	return iv
}

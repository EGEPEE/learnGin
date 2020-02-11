package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
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
	iv := []byte{233, 141, 164, 199, 171, 248, 138, 252, 178, 31, 29, 147}
	return iv
}

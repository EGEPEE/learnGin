package helper

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func GCM_encrypt(key string, plaintext string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	fmt.Println(hex.EncodeToString(out))

	return hex.EncodeToString(out)
}

func GCM_decrypt(key string, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
		panic(err)
	}
	plain := make([]byte, len(ciphertext))
	c.Decrypt(plain, ciphertext)
	s := string(plain[:])
	fmt.Printf("AES Decrypyed Text:  %s\n", s)
	fmt.Println(s)
	return s
}

func GetIV() []byte {
	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}

	return iv
}

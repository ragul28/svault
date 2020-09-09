package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

// Encrypt secret using AES-GCM block ciper
func Encrypt(skey, nonce []byte, plaintext string) ([]byte, error) {

	block, err := aes.NewCipher(skey)
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return ciphertext, nil
}

// Decrypt secret using AES-GCM block ciper
func Decrypt(skey, nonce []byte, stext []byte) (string, error) {

	block, err := aes.NewCipher(skey)
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal("Error in cipher.NewGCM", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, stext, nil)
	if err != nil {
		log.Fatal("Error in aesgcm.Open", err)
	}
	return string(plaintext), nil
}

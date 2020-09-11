package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

// Encrypt secret using AES-GCM block ciper
func Encrypt(skey []byte, plaintext string) ([]byte, error) {

	block, err := aes.NewCipher(skey)
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	// Create unique 12 byte nonce for gcm
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt using aemgcm seal & append nonce to result
	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return ciphertext, nil
}

// Decrypt secret using AES-GCM block ciper
func Decrypt(skey, stext []byte) (string, error) {

	block, err := aes.NewCipher(skey)
	if err != nil {
		log.Fatal(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal("Error in cipher.NewGCM", err)
	}

	// Split nonce & secret text from encrypt data
	nonce, stext := stext[:aesgcm.NonceSize()], stext[aesgcm.NonceSize():]
	plaintext, err := aesgcm.Open(nil, nonce, stext, nil)
	if err != nil {
		log.Fatal("Error in aesgcm.Open", err)
	}
	return string(plaintext), nil
}

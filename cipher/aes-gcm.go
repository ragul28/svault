package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

// AES-GCM block ciper encryption
func Encrypt(skey, nonce []byte, plaintext string) ([]byte, error) {

	block, err := aes.NewCipher(skey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return ciphertext, err
}

// AES-GCM block ciper decryption
func Decrypt(skey, nonce []byte, stext []byte) (string, error) {

	block, err := aes.NewCipher(skey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic("Error in cipher.NewGCM", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, stext, nil)
	if err != nil {
		log.Panic("Error in aesgcm.Open", err)
	}
	return string(plaintext), err
}

package cipher

import (
	"log"
	"testing"
)

func Test_Ciper(t *testing.T) {
	encryptKey := []byte("test_master_key0")
	ciphertext, err := Encrypt(encryptKey, "secret")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Print("Encrypted text:", ciphertext)
	}

	plaintext, err := Decrypt(encryptKey, ciphertext)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Print("Decrypted text:", plaintext)
	}
}

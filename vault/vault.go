package vault

import (
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ragul28/svault/cipher"
)

func WriteVault(masterKey []byte, Key, secret string) {

	encryptKey, nonce := GenVaultKey(masterKey)

	ciphertext, _ := cipher.Encrypt(encryptKey, nonce, secret)
	fmt.Printf("%s saved in svault!\n", Key)

	KV := VaultData{time.Now().Unix(), "kv", ciphertext, 0}
	err := writeStorage(Key, KV)
	if err != nil {
		log.Panic(err)
	}
}

func ReadVault(masterKey []byte, Key string) {

	encryptKey, nonce := GenVaultKey(masterKey)

	VD, err := readStorage(Key)
	if err != nil {
		log.Fatal(err)
	}
	if len(VD.EnctyptData) == 0 {
		fmt.Println("VaultKey Not Found:", Key)
		os.Exit(0)
	}
	// fmt.Printf("%x\n", VD.EnctyptData)

	plaintextNew, _ := cipher.Decrypt(encryptKey, nonce, VD.EnctyptData)
	fmt.Printf("%s\n", plaintextNew)
}

func ListVault() {
	VDmap, _, err := getStorage()
	if err != nil {
		log.Panic(err)
	}

	counter := 0
	for mkey := range VDmap {
		if mkey != "master_key" {
			counter++
			fmt.Printf("%d. %s\n", counter, mkey)
		}
	}
}

func StatusVault() {
	VDmap, kvcount, err := getStorage()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Vault Status: initialized\nInit Time: %d\nKV Count: %d\n", VDmap["master_key"].CreatedTime, kvcount-1)
	}
}

// Helper func to get the masterkey & nonce
// TODO: Needs Better strategy
func GenVaultKey(masterKey []byte) (encryptKey, nonce []byte) {
	// Genarate sha512 of the masterkey
	masterKeysha := sha512.Sum512(masterKey)

	// Get first 32 bytes for secret key. For AES Seal/Open calls, key should be 16 bytes (AES-128) or 32 (AES-256).
	encryptKey = masterKeysha[:32]
	// fmt.Printf("secretKey: %x\n", encryptKey)

	// Get following 12 bytes of masterkeysha after secretKey.
	nonce = masterKeysha[32:44]
	// fmt.Printf("nonce: %x\n", nonce)

	return encryptKey, nonce
}

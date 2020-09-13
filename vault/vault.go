package vault

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ragul28/svault/cipher"
)

func WriteVault(encryptKey []byte, Key, secret string) {

	ciphertext, _ := cipher.Encrypt([]byte(encryptKey), secret)
	fmt.Printf("%s saved in svault!\n", Key)

	KV := VaultData{time.Now().Unix(), "kv", ciphertext, 0}
	err := KV.writeStorage(Key)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadVault(encryptKey []byte, Key string) {

	var vd VaultData
	vd, err := vd.readStorage(Key)
	if err != nil {
		log.Fatal(err)
	}
	if len(vd.EnctyptData) < 1 {
		fmt.Printf("%s not found in svault!\n", Key)
		os.Exit(0)
	}

	plaintextNew, _ := cipher.Decrypt(encryptKey, vd.EnctyptData)
	fmt.Printf("%s\n", plaintextNew)
}

func DeleteVault(encryptKey []byte, Key string) {

	err := deleteStorage(Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s deleted from svault!\n", Key)
}

func ListVault() {
	VDmap, _, err := getStorage()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	} else {
		fmt.Printf("Vault Status: initialized\nInit Time: %d\nKV Count: %d\n", VDmap["master_key"].CreatedTime, kvcount-1)
	}
}

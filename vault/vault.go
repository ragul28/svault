package vault

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ragul28/svault/cipher"
)

const bucket = "kvStore"

func WriteVault(encryptKey []byte, Key, secret string) {

	ciphertext, _ := cipher.Encrypt([]byte(encryptKey), secret)
	fmt.Printf("%s saved in svault!\n", Key)

	boltdb := open(getVautlPath())

	vd, err := json.Marshal(VaultData{time.Now().Unix(), "kv", ciphertext, 0})
	if err != nil {
		log.Fatal(err)
	}

	err = writeDB(boltdb, bucket, Key, vd)
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

	boltdb := open(getVautlPath())
	err := deleteDB(boltdb, bucket, Key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s deleted from svault!\n", Key)
}

func ListVault() {

	boltdb := open(getVautlPath())
	_, err := iterateDB(boltdb, bucket)
	if err != nil {
		log.Fatal(err)
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

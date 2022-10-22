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

	boltdb := open(getVautlPath())
	val, err := readDB(boltdb, bucket, Key)
	if err != nil {
		log.Fatalln("svault is empty!")
	}

	var vd VaultData
	err = json.Unmarshal(val, &vd)
	if err != nil {
		fmt.Printf("%s not found in svault!\n", Key)
		// log.Fatal(err)
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
		fmt.Printf("No keys found, Vault is empty!")
		// log.Fatal(err)
	}
}

func StatusVault() {
	boltdb := open(getVautlPath())
	val, err := readDB(boltdb, master_bucket, "master_key")
	if err != nil {
		log.Fatal(err)
	}
	var vd VaultData
	err = json.Unmarshal(val, &vd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vault Status: initialized\nInit Time: ", time.Unix(vd.CreatedTime, 0).Format("2006-01-02 15:04:05"))
}

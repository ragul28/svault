package vault

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/user"
	"time"
)

const master_bucket = "masterStore"

// Vault init create masterkey & storage file if not initialzed.
func Init(freshInit bool) string {

	// Remove the File to indicate a fresh init
	if freshInit {
		os.Remove(getVautlPath())
	}

	boltdb := open(getVautlPath())

	// check secretFile file exist
	val, err := readDB(boltdb, master_bucket, "master_key")
	if err != nil {

		masterKey := genRandomSecretKey(32)

		val, _ := json.Marshal(VaultData{time.Now().Unix(), "sys", []byte{}, 0})

		err = writeDB(boltdb, master_bucket, "master_key", val)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Vault initialized")
		fmt.Printf("Master Key: %s\n", masterKey)

		return string(masterKey)

	}

	var vd VaultData
	err = json.Unmarshal(val, &vd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Vault already initialized.\nMaster Key Generated at", time.Unix(vd.CreatedTime, 0).UTC())
	return ""
}

// Genrate crypto random secret key from predefined ascii set
func genRandomSecretKey(n int) []byte {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-+"

	randomKey := make([]byte, n)

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			fmt.Println(err)
		}
		randomKey = append(randomKey, letters[num.Int64()])
	}
	return randomKey
}

func getVautlPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return user.HomeDir + "/.svault/svault.data"
}

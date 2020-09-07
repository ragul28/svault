package vault

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

// Vault init create masterkey & storage file if not initialzed.
func VaultInit(freshInit bool) string {

	if freshInit {
		_ = os.Remove(Defaultpath)
	}

	// check secretFile file exist
	if secretFile, err := readStorage("master_key"); err == nil {

		fmt.Println("Vault already initialized.\nMaster Key Generated at", time.Unix(secretFile.CreatedTime, 0).UTC())
		return ""

	} else {

		fmt.Println("Vault initialized")

		masterKey := genRandomSecretKey(32)
		fmt.Printf("Master Key: %s\n", masterKey)

		RK := VaultData{time.Now().Unix(), "root", []byte("MasterKeyGenerated"), 0}
		if writeStorage("master_key", RK) != nil {
			log.Panic(err)
		}

		return string(masterKey)
	}
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

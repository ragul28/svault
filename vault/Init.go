package vault

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

// Vault init create masterkey & storage file if not initialzed.
func VaultInit(freshInit bool) string {

	// Remove the File to indicate a fresh init
	if freshInit {
		_ = os.Remove(getVautlPath())
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

		// create dot dir for svault storage
		os.MkdirAll(filepath.Base(getVautlPath()), os.ModePerm)
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

func getVautlPath() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return user.HomeDir + "/.svault/svault.data"
}

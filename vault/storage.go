package vault

import (
	"encoding/json"
	"log"
	"os"
)

type VaultData struct {
	CreatedTime int64
	Type        string
	EnctyptData []byte
	Version     int
}

func writeStorage(Key string, v VaultData) error {

	var vaultMap map[string]VaultData

	if _, err := os.Stat(getVautlPath()); err == nil {
		vaultMap, _, err = getStorage()
		if err != nil {
			return err
		}
	} else {
		vaultMap = make(map[string]VaultData)
	}

	vaultMap[Key] = v

	file, err := os.OpenFile(getVautlPath(), os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Println("File cannot open or found", err)
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(vaultMap)
	return nil
}

func readStorage(Key string) (VaultData, error) {
	vaultMap := make(map[string]VaultData)
	data, err := os.Open(getVautlPath())
	if err != nil {
		return VaultData{}, err
	}

	decoder := json.NewDecoder(data)
	decoder.Decode(&vaultMap)
	return vaultMap[Key], nil
}

func getStorage() (map[string]VaultData, int, error) {
	vaultMap := make(map[string]VaultData)
	data, err := os.Open(getVautlPath())
	if err != nil {
		return map[string]VaultData{}, 0, err
	}

	decoder := json.NewDecoder(data)
	decoder.Decode(&vaultMap)
	return vaultMap, len(vaultMap), nil
}

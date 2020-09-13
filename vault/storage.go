package vault

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type VaultData struct {
	CreatedTime int64
	Type        string
	EnctyptData []byte
	Version     int
}

type storage interface {
	writeStorage(Key string) error
	readStorage(Key string) error
}

func (vd *VaultData) writeStorage(Key string) error {

	var vaultMap map[string]VaultData

	// Check data file avilable then get storage else create new vault map
	if _, err := os.Stat(getVautlPath()); err == nil {
		vaultMap, _, err = getStorage()
		if err != nil {
			return err
		}
	} else {
		vaultMap = make(map[string]VaultData)
	}

	vaultMap[Key] = *vd

	// Save vaultmap in data file
	file, err := os.OpenFile(getVautlPath(), os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Println("File cannot open or found", err)
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(vaultMap)
	return nil
}

func (vd *VaultData) readStorage(Key string) (VaultData, error) {
	vaultMap := make(map[string]VaultData)
	data, err := os.Open(getVautlPath())
	if err != nil {
		return VaultData{}, err
	}

	decoder := json.NewDecoder(data)
	decoder.Decode(&vaultMap)
	return vaultMap[Key], nil
}

func deleteStorage(Key string) error {

	vaultMap, _, err := getStorage()
	if err != nil {
		return err
	}

	_, ok := vaultMap[Key]
	if ok {
		delete(vaultMap, Key)
	} else {
		return errors.New("Key not found!")
	}

	file, err := os.OpenFile(getVautlPath(), os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Println("File cannot open or found", err)
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.Encode(vaultMap)
	return nil
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

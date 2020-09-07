package main

import (
	"github.com/ragul28/svault/vault"
)

func main() {

	MasterKey := []byte(vault.VaultInit(true))

	vault.WriteVault(MasterKey, "secret_key", "supersecrettoken")
	vault.ReadVault(MasterKey, "secret_key")

	vault.StatusVault()
	vault.ListVault()

}

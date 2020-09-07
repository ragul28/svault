package cmd

import (
	"log"
	"os"

	"github.com/ragul28/svault/vault"

	"github.com/spf13/cobra"
)

var masterkey string

func init() {
	rootCmd.AddCommand(storeCmd, getCmd, listCmd)
	storeCmd.Flags().StringVarP(&masterkey, "masterkey", "m", "", "Pass masterkey as flag")
	getCmd.Flags().StringVarP(&masterkey, "masterkey", "m", "", "Pass masterkey as flag")
}

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store secret to vault store",
	Run: func(cmd *cobra.Command, args []string) {
		masterkey = checkMasterKey(masterkey)
		vault.WriteVault([]byte(masterkey), args[0], args[1])
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get secret from vault store",
	Run: func(cmd *cobra.Command, args []string) {
		masterkey = checkMasterKey(masterkey)
		vault.ReadVault([]byte(masterkey), args[0])
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored secret from vault store",
	Run: func(cmd *cobra.Command, args []string) {
		vault.ListVault()
	},
}

func checkMasterKey(masterkey string) string {
	if masterkey != "" {
		return masterkey
	} else if os.Getenv("MASTER_KEY") != "" {
		return os.Getenv("MASTER_KEY")
	} else {
		log.Fatal("MasterKey missing. Pass the masterkey as env/flag")
		return ""
	}
}

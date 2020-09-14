package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/ragul28/svault/vault"

	"github.com/spf13/cobra"
)

var masterkey string

func init() {
	rootCmd.AddCommand(storeCmd, getCmd, deleteCmd, listCmd)
	storeCmd.Flags().StringVarP(&masterkey, "masterkey", "m", "", "Pass masterkey as flag")
	getCmd.Flags().StringVarP(&masterkey, "masterkey", "m", "", "Pass masterkey as flag")
}

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store secret to vault store",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Missing arguments key/vaule")
		}
		masterkey = checkMasterKey(masterkey)
		vault.WriteVault([]byte(masterkey), args[0], args[1])
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get secret from vault store",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Missing argument key")
		}
		masterkey = checkMasterKey(masterkey)
		vault.ReadVault([]byte(masterkey), args[0])
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete secret from vault store",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Missing argument key")
		}
		masterkey = checkMasterKey(masterkey)
		vault.DeleteVault([]byte(masterkey), args[0])
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored secret from vault store",
	Args:  cobra.NoArgs,
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

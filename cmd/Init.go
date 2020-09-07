package cmd

import (
	"os"

	"github.com/ragul28/svault/vault"

	"github.com/spf13/cobra"
)

var freshInit bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&freshInit, "freshInit", "f", false, "Fresh Init & recreate the masterkey")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init vault secret engine",
	Run: func(cmd *cobra.Command, args []string) {
		MK := vault.VaultInit(freshInit)
		os.Setenv("MASTER_KEY", MK)
	},
}

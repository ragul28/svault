package cmd

import (
	"github.com/ragul28/svault/vault"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of the vault engine",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		vault.StatusVault()
	},
}

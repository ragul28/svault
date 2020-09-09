package cmd

import (
	"github.com/ragul28/svault/vault"

	"github.com/spf13/cobra"
)

var freshInit bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&freshInit, "freshInit", "f", false, "Clear vault & recreate the masterkey")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init vault secret engine",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		vault.Init(freshInit)
	},
}

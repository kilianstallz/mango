package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mango/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of mango",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

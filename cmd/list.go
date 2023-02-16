package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mango/config"
	"mango/internal/check"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed Go versions",
	RunE: func(cmd *cobra.Command, args []string) error {

		// get installed versions
		versions, err := check.ListInstalledVersions()
		cobra.CheckErr(err)
		fmt.Println("Installed Go versions:")
		// print versions
		for _, version := range versions {
			if version == viper.GetString(config.CurrentEnv) {
				fmt.Printf("* %s \n", version)
				continue
			}
			fmt.Printf("  %s \n", version)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

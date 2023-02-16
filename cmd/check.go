package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mango/config"
	"mango/internal/check"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates",
	RunE: func(cmd *cobra.Command, args []string) error {

		// query the Go site for the latest version
		latest, err := check.QueryLatestGoVersion()
		cobra.CheckErr(err)

		// compare the latest version with the current version
		// if the latest version is greater than the current version
		// then print a message to the user

		installedVersions, err := check.ListInstalledVersions()
		cobra.CheckErr(err)
		viper.Set(config.InstalledVersions, installedVersions)
		viper.Set("hasUpdate", true)

		// check if the latest version is installed
		latestInstalled := false
		for _, version := range installedVersions {
			if version == latest {
				latestInstalled = true
			}
		}

		if !latestInstalled {
			fmt.Printf("A new version of Go is available: %s \n", latest)
			return nil
		}

		fmt.Printf("You are running the latest version of Go: %s \n", latest)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

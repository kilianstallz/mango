package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mango/internal/check"
	"mango/internal/cli"
	"mango/internal/download"
)

var updateCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a new version of Go",
	Run: func(cmd *cobra.Command, args []string) {

		// check for new version
		latest, err := check.IsUpdateAvailable()
		cobra.CheckErr(err)
		if latest != "" {
			// ask for confirmation
			resp := cli.AskForConfirmation("Do you want to download it?", false)
			if !resp {
				return
			}

			fmt.Println("Downloading new version of Go...")
			// download the new version
			err = download.DownloadGoVersion(latest)
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

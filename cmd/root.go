package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"mango/config"
	"os"
	"strconv"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "mango",
	Short: "Mango is a CLI tool for managing your Go Versions",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "config file (default is $HOME/.mangorc.yaml)")
}

func initConfig() {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mangorc")

		// check if the file exists and if not, create it with default values
		if _, err := os.Stat(home + "/.mangorc.yaml"); os.IsNotExist(err) {
			// file does not exist
			cfg, err := config.New()
			cobra.CheckErr(err)
			cfgYaml, err := yaml.Marshal(cfg)
			fmt.Println(string(cfgYaml))
			cobra.CheckErr(err)
			viper.ReadConfig(bytes.NewBuffer(cfgYaml))
			viper.AutomaticEnv()
			err = viper.SafeWriteConfig()
			cobra.CheckErr(err)
		}

	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// run check command to check for updates
	err := checkCmd.RunE(nil, nil)
	cobra.CheckErr(err)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		printError(err)
		os.Exit(1)
	}
}

func printError(err error) {
	env := os.Getenv("DEBUG")
	debug, _ := strconv.ParseBool(env)
	if env != "" && debug {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println(err)
	}
}

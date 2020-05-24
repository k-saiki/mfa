package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Service []Services
}

type Services struct {
	Name   string
	Secret string
}

var config Config

var rootCmd = &cobra.Command{
	Use:   "mfa",
	Short: "Generate a totp token with cli.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigType("yaml")

	if os.Getenv("MFA_CONFIG") != "" {
		viper.SetConfigFile(os.Getenv("MFA_CONFIG"))
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		configPath := filepath.Join(home, ".mfa")
		viper.AddConfigPath(configPath)
		viper.SetConfigName("secrets")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if err := viper.UnmarshalExact(&config); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

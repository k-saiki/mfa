package cmd

import (
	"log"
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
	Short: "Generate a TOTP(Time-based One-time Password) token",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	// Default config is "$HOME/.mfa/secrets" or use environment variables to override config file path.
	viper.SetConfigType("yaml")
	viper.BindEnv("MFA_CONFIG")

	if viper.Get("MFA_CONFIG") != nil {
		configPath := filepath.Dir(viper.GetString("MFA_CONFIG"))
		filename := filepath.Base(viper.GetString("MFA_CONFIG"))
		viper.AddConfigPath(configPath)
		viper.SetConfigName(filename)
	} else {
		configPath := filepath.Join(home, ".mfa")
		viper.AddConfigPath(configPath)
		viper.SetConfigName("secrets")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[ERROR] %s", err)
		os.Exit(-1)
	}

	if err := viper.UnmarshalExact(&config); err != nil {
		log.Printf("[ERROR] %s", err)
		os.Exit(-1)
	}
}

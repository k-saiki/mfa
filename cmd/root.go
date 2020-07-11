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
var version string
var revision string

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mfa",
		Short: "Generate a totp token with cli.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(NewGenerateCommand())
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewVersionCommand())
	return cmd
}

func Execute() {
	cmd := NewCommand()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println("Error:", err)
		os.Exit(1)
	}
}

func LoadConfig() {
	viper.SetConfigType("yaml")

	if os.Getenv("MFA_CONFIG") != "" {
		viper.SetConfigFile(os.Getenv("MFA_CONFIG"))
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		configPath := filepath.Join(home, ".mfa")
		viper.AddConfigPath(configPath)
		viper.SetConfigName("secrets")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if err := viper.UnmarshalExact(&config); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

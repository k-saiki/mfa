package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Service []Service `yaml:"service"`
}

type Service struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

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
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func getConfigPath() (string, error) {
	if path := os.Getenv("MFA_CONFIG"); path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	// Try .yml first, then .yaml
	for _, ext := range []string{".yml", ".yaml"} {
		path := filepath.Join(home, ".mfa", "secrets"+ext)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return filepath.Join(home, ".mfa", "secrets.yml"), nil
}

func LoadConfig() (*Config, string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, "", err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, configPath, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, configPath, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, configPath, nil
}

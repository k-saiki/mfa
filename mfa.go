package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/pquerna/otp/totp"
	"github.com/urfave/cli"
)

const configDefault = ".mfa/secrets"

type Config struct {
	Service []MFAConfig `toml:"service"`
}

type MFAConfig struct {
	Name   string `toml:"name"`
	Secret string `toml:"secret"`
}

func main() {
	var config Config
	if err := loadConfig(configDefault, &config); err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "mfa"
	app.Usage = "Generate TOTP(Time-based One-time Password) token."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "List configured services.",
			Action: func(context *cli.Context) error {
				for _, cfg := range config.Service {
					fmt.Printf("%s\n", cfg.Name)
				}
				return nil
			},
		},
		{
			Name:  "gen",
			Usage: "Generate TOTP token.",
			Action: func(context *cli.Context) error {
				now := time.Now()
				service := context.Args().Get(0)
				if service == "" {
					msg := "Enter a service name."
					return errors.New(msg)
				}
				for _, cfg := range config.Service {
					if cfg.Name == service {
						token, _ := totp.GenerateCode(cfg.Secret, now)
						fmt.Println(token)
						break
					}
					msg := "Service \"" + service + "\" not found."
					return errors.New(msg)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(configPath string, config *Config) error {
	homeDir, err := homedir.Dir()
	if err != nil {
		return err
	}

	configFullPath := filepath.Join(homeDir, configPath)

	if _, err := toml.DecodeFile(configFullPath, config); err != nil {
		return err
	}

	return nil
}

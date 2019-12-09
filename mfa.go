package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/pquerna/otp/totp"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const configDefault = ".mfa/secrets"

type Config struct {
	Services []Service `yaml:"service"`
}

type Service struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

func main() {
	var config Config

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "mfa"
	app.Usage = "Generate TOTP(Time-based One-time Password) token."
	app.Version = "0.0.2"

	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "List configured services",
			Action: func(context *cli.Context) error {
				if err := loadConfig(configDefault, &config); err != nil {
					log.Fatal(err)
				}

				for _, s := range config.Services {
					fmt.Println(s.Name)
				}
				return nil
			},
		},
		{
			Name:  "gen",
			Usage: "Generate TOTP token.",
			Action: func(context *cli.Context) error {
				if err := loadConfig(configDefault, &config); err != nil {
					log.Fatal(err)
				}

				now := time.Now()
				service := context.Args().Get(0)
				if service == "" {
					msg := "Enter a service name"
					return errors.New(msg)
				}

				for _, s := range config.Services {
					if s.Name == service {
						token, _ := totp.GenerateCode(s.Secret, now)
						fmt.Println(token)
						os.Exit(0)
					}
				}

				msg := "[ERROR] Service \"" + service + "\" not found."
				return errors.New(msg)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(t string, c *Config) error {
	homeDir, err := homedir.Dir()
	if err != nil {
		return err
	}

	configFullPath := filepath.Join(homeDir, t)
	buf, err := ioutil.ReadFile(configFullPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	return nil
}

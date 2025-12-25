package cmd

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen <service>",
		Short: "Generate a totp token",
		Args:  cobra.ExactArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			config, configPath, err := LoadConfig()
			if err != nil {
				return err
			}

			service := args[0]

			for _, s := range config.Service {
				if s.Name == service {
					token, err := totp.GenerateCode(s.Secret, time.Now())
					if err != nil {
						return fmt.Errorf("failed to generate token: %w", err)
					}
					cmd.Println(token)
					return nil
				}
			}

			return fmt.Errorf("%s not found in %s", service, configPath)
		},
	}

	return cmd
}

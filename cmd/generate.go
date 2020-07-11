package cmd

import (
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen <service>",
		Short: "Generate a totp token",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			LoadConfig()
			service := args[0]

			for _, s := range config.Service {
				if s.Name == service {
					token, _ := totp.GenerateCode(s.Secret, time.Now())
					cmd.Println(token)
					os.Exit(0)
				}
			}

			cmd.Printf("Error: %s not found in %s\n", service, viper.ConfigFileUsed())
		},
	}

	return cmd
}

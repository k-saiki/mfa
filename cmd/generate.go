package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// genCmd represents the generate command
var genCmd = &cobra.Command{
	Use:   "gen <service>",
	Short: "Generate a totp token",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		service := args[0]

		for _, s := range config.Service {
			if s.Name == service {
				token, _ := totp.GenerateCode(s.Secret, time.Now())
				fmt.Println(token)
				os.Exit(0)
			}
		}

		fmt.Printf("Error: %s not found in %s\n", service, viper.ConfigFileUsed())
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

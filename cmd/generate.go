package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// genCmd represents the generate command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a totp token",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Printf("[ERROR] Missing argument. enter service name.")
			os.Exit(1)
		}

		if len(args) > 1 {
			log.Printf("[ERROR] Too many arguments.")
			os.Exit(1)
		}

		service := args[0]

		for _, s := range config.Service {
			if s.Name == service {
				token, _ := totp.GenerateCode(s.Secret, time.Now())
				fmt.Println(token)
				os.Exit(0)
			}
		}

		log.Printf("[ERROR] %s not found in %s", service, viper.ConfigFileUsed())
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

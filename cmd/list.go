package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print services",

	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range config.Service {
			fmt.Println(s.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

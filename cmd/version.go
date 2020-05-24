package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string
var revision string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Revision: %s\n", revision)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

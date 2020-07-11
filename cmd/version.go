package cmd

import (
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {
			ShowVersion()
		},
	}

	return cmd
}

func ShowVersion() {
	cmd := &cobra.Command{}
	cmd.Println("Version: ", version)
	cmd.Println("Revision: ", revision)
}

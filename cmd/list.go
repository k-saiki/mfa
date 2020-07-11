package cmd

import (
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Print services",
		Run: func(cmd *cobra.Command, args []string) {
			LoadConfig()

			for _, s := range config.Service {
				cmd.Println(s.Name)
			}
		},
	}

	return cmd
}

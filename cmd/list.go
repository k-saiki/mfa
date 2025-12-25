package cmd

import (
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Print services",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, _, err := LoadConfig()
			if err != nil {
				return err
			}

			for _, s := range config.Service {
				cmd.Println(s.Name)
			}
			return nil
		},
	}

	return cmd
}

package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {
			showVersion(cmd.OutOrStdout())
		},
	}

	return cmd
}

func showVersion(w io.Writer) {
	fmt.Fprintln(w, "Version:", version)
	fmt.Fprintln(w, "Revision:", revision)
}

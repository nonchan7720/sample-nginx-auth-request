package cmd

import "github.com/spf13/cobra"

func rootCommand() *cobra.Command {
	cmd := cobra.Command{}

	cmd.AddCommand(app1Command())
	cmd.AddCommand(app2Command())
	return &cmd
}

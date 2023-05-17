package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

func NewCmdEnvironment(s printers.IOStreams) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "environment",
		Aliases: []string{"e", "env"},
		Short:   "publishing environment subcommands",
		//Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdEnvironmentCreate(s))
	cmd.AddCommand(NewCmdEnvironmentRead(s))
	cmd.AddCommand(NewCmdEnvironmentUpdate(s))
	cmd.AddCommand(NewCmdEnvironmentDelete(s))

	return cmd
}

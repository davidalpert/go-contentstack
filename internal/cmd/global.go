package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

func NewCmdGlobal(s printers.IOStreams) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "global",
		Aliases: []string{"g"},
		Short:   "global field subcommands",
		//Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdGlobalCreate(s))
	cmd.AddCommand(NewCmdGlobalRead(s))
	cmd.AddCommand(NewCmdGlobalUpdate(s))
	cmd.AddCommand(NewCmdGlobalDelete(s))

	return cmd
}

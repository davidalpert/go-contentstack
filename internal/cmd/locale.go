package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

func NewCmdLocale(s printers.IOStreams) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "locale",
		Aliases: []string{"l", "loc"},
		Short:   "locale subcommands",
		//Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdLocaleCreate(s))
	cmd.AddCommand(NewCmdLocaleRead(s))
	cmd.AddCommand(NewCmdLocaleUpdate(s))
	cmd.AddCommand(NewCmdLocaleDelete(s))

	return cmd
}

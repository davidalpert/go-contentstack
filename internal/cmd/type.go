package cmd

import (
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

func NewCmdContentType(s printers.IOStreams) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "types",
		Aliases: []string{"type", "content-type", "t"},
		Short:   "content type subcommands",
		//Args:    cobra.NoArgs,
	}

	cmd.AddCommand(NewCmdContentTypeCreate(s))
	cmd.AddCommand(NewCmdContentTypeRead(s))
	cmd.AddCommand(NewCmdContentTypeUpdate(s))
	cmd.AddCommand(NewCmdContentTypeDelete(s))

	return cmd
}

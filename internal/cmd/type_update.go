package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type ContentTypeUpdateOptions struct {
	*printers.PrinterOptions
	Config   cfg.Config
	Client   *management.Client
	UID      string
	NewTitle string
}

func NewContentTypeUpdateOptions(s printers.IOStreams) *ContentTypeUpdateOptions {
	return &ContentTypeUpdateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdContentTypeUpdate(s printers.IOStreams) *cobra.Command {
	o := NewContentTypeUpdateOptions(s)
	var cmd = &cobra.Command{
		Use:   "update <uid> <new_title>",
		Short: "update the title of the ContentType with the given UID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run()
		},
	}

	o.AddPrinterFlags(cmd.Flags())

	return cmd
}

// Complete the options
func (o *ContentTypeUpdateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.UID = args[0]
	o.NewTitle = args[1]

	if err := cfg.ReadMergedInto(&o.Config); err != nil {
		return err
	}

	c, err := management.NewClient(&o.Config.ContentStack.ManagementApi)
	if err != nil {
		return err
	}
	o.Client = c

	return nil
}

// Validate the options
func (o *ContentTypeUpdateOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *ContentTypeUpdateOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	g, err := o.Client.GetOneContentType(o.UID)

	g.Title = o.NewTitle

	r, err := o.Client.UpdateContentType(g)
	if err != nil {
		return err
	}

	return o.WriteOutput(r)
}

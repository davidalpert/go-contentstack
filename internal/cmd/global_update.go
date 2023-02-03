package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type GlobalUpdateOptions struct {
	*printers.PrinterOptions
	Config   cfg.Config
	Client   *management.Client
	UID      string
	NewTitle string
}

func NewGlobalUpdateOptions(s printers.IOStreams) *GlobalUpdateOptions {
	return &GlobalUpdateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdGlobalUpdate(s printers.IOStreams) *cobra.Command {
	o := NewGlobalUpdateOptions(s)
	var cmd = &cobra.Command{
		Use:   "update <uid> <new_title>",
		Short: "update the title of the GlobalField with the given UID",
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
func (o *GlobalUpdateOptions) Complete(cmd *cobra.Command, args []string) error {
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
func (o *GlobalUpdateOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *GlobalUpdateOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	g, err := o.Client.GetOneGlobalField(o.UID)

	g.Title = o.NewTitle

	r, err := o.Client.UpdateGlobalField(g)
	if err != nil {
		return err
	}

	return o.WriteOutput(r)
}

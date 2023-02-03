package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type GlobalReadOptions struct {
	*printers.PrinterOptions
	Config cfg.Config
	Client *management.Client
	UID    string
}

func NewGlobalReadOptions(s printers.IOStreams) *GlobalReadOptions {
	return &GlobalReadOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdGlobalRead(s printers.IOStreams) *cobra.Command {
	o := NewGlobalReadOptions(s)
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "show current configuration values",
		Args:  cobra.NoArgs,
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
	cmd.Flags().StringVarP(&o.UID, "uid", "u", "", "uid of a single global field")

	return cmd
}

// Complete the options
func (o *GlobalReadOptions) Complete(cmd *cobra.Command, args []string) error {
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
func (o *GlobalReadOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *GlobalReadOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	if o.UID != "" {
		r, err := o.Client.GetOneGlobalField(o.UID)
		if err != nil {
			return err
		}
		return o.WriteOutput(r)
	}

	r, err := o.Client.GetAllGlobalFields()
	if err != nil {
		return err
	}
	return o.WriteOutput(r)
}

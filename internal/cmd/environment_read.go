package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type EnvironmentReadOptions struct {
	*printers.PrinterOptions
	Config cfg.Config
	Client *management.Client
	Name   string
}

func NewEnvironmentReadOptions(s printers.IOStreams) *EnvironmentReadOptions {
	return &EnvironmentReadOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdEnvironmentRead(s printers.IOStreams) *cobra.Command {
	o := NewEnvironmentReadOptions(s)
	var cmd = &cobra.Command{
		Use:   "get",
		Short: "show all environments",
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
	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "name of a single publishing environment")

	return cmd
}

// Complete the options
func (o *EnvironmentReadOptions) Complete(cmd *cobra.Command, args []string) error {
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
func (o *EnvironmentReadOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *EnvironmentReadOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	if o.Name != "" {
		r, err := o.Client.GetOnePublishingEnvironment(o.Name)
		if err != nil {
			return err
		}
		return o.WriteOutput(r)
	}

	r, err := o.Client.GetAllPublishingEnvironments()
	if err != nil {
		return err
	}
	return o.WriteOutput(r)
}

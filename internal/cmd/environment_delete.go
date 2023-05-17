package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type EnvironmentDeleteOptions struct {
	*printers.PrinterOptions
	Config cfg.Config
	Client *management.Client
	Name   string
	Force  bool
}

func NewEnvironmentDeleteOptions(s printers.IOStreams) *EnvironmentDeleteOptions {
	return &EnvironmentDeleteOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdEnvironmentDelete(s printers.IOStreams) *cobra.Command {
	o := NewEnvironmentDeleteOptions(s)
	var cmd = &cobra.Command{
		Use:     "delete <name>",
		Short:   "delete a publishing environment with the given name",
		Aliases: []string{"d", "del", "rm"},
		Args:    cobra.ExactArgs(1),
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
	cmd.Flags().BoolVarP(&o.Force, "force", "f", false, "force (bypass confirmation)")

	return cmd
}

// Complete the options
func (o *EnvironmentDeleteOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Name = args[0]

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
func (o *EnvironmentDeleteOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *EnvironmentDeleteOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	if !o.Force {
		confirmed := false
		if err := askOneWithStreams(o.IOStreams, &survey.Confirm{
			Message: fmt.Sprintf("delete publishing environment %#v", o.Name),
		}, &confirmed); err != nil {
			return err
		}

		if !confirmed {
			return fmt.Errorf("aborted delete")
		}
	}

	err := o.Client.DeletePublishingEnvironment(o.Name)
	if err != nil {
		return err
	}

	return o.WriteOutput(fmt.Sprintf("deleted publishing environment %s", o.Name))
}

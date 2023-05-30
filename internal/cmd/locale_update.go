package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type LocaleUpdateOptions struct {
	*printers.PrinterOptions
	Config             cfg.Config
	Client             *management.Client
	Code               string
	Name               string
	FallbackLocaleCode string
	Force              bool
}

func NewLocaleUpdateOptions(s printers.IOStreams) *LocaleUpdateOptions {
	return &LocaleUpdateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdLocaleUpdate(s printers.IOStreams) *cobra.Command {
	o := NewLocaleUpdateOptions(s)
	var cmd = &cobra.Command{
		Use:     "update <code>",
		Short:   "update the details for a locale for the given code",
		Aliases: []string{"u", "up"},
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
	cmd.Flags().StringVarP(&o.Name, "name", "n", "", "optional name for the locale; defaults to the code")
	cmd.Flags().StringVar(&o.FallbackLocaleCode, "fallback-code", "", "optional fallback code")
	cmd.Flags().BoolVarP(&o.Force, "force", "f", false, "force (bypass confirmation)")

	return cmd
}

// Complete the options
func (o *LocaleUpdateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Code = args[0]

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
func (o *LocaleUpdateOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocaleUpdateOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}
	l, err := o.Client.GetOneLocale(o.Code)
	if err != nil {
		return err
	}

	if o.Name != "" {
		l.Name = o.Name
	}

	if o.FallbackLocaleCode != "" {
		l.FallbackLocaleCode = o.FallbackLocaleCode
	}

	if !o.Force {
		confirmed := false
		if err := askOneWithStreams(o.IOStreams, &survey.Confirm{
			Message: fmt.Sprintf("update locale %#v to %#v (fallback to: %#v)", l.Code, l.Name, l.FallbackLocaleCode),
		}, &confirmed); err != nil {
			return err
		}

		if !confirmed {
			return fmt.Errorf("aborted update")
		}
	}

	r, err := o.Client.UpdateLocale(l)
	if err != nil {
		return err
	}

	return o.WriteOutput(r)
}

package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type LocaleCreateOptions struct {
	*printers.PrinterOptions
	Config       cfg.Config
	Client       *management.Client
	Code         string
	Name         string
	FallbackCode string
}

func NewLocaleCreateOptions(s printers.IOStreams) *LocaleCreateOptions {
	return &LocaleCreateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdLocaleCreate(s printers.IOStreams) *cobra.Command {
	o := NewLocaleCreateOptions(s)
	var cmd = &cobra.Command{
		Use:     "create <code>",
		Short:   "create a new locale with the given name",
		Aliases: []string{"c", "cr"},
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
	cmd.Flags().StringVar(&o.Name, "fallback-code", "en-us", "optional fallback code")

	return cmd
}

// Complete the options
func (o *LocaleCreateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Code = args[0]
	if o.Name == "" {
		o.Name = o.Code
	}

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
func (o *LocaleCreateOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocaleCreateOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	r, err := o.Client.CreateLocale(o.Code, o.Name, o.FallbackCode)
	if err != nil {
		return err
	}

	return o.WriteOutput(r)
}

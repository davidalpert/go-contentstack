package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"github.com/davidalpert/go-printers/v1"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type LocaleReadOptions struct {
	*printers.PrinterOptions
	Config cfg.Config
	Client *management.Client
	Code   string
}

func NewLocaleReadOptions(s printers.IOStreams) *LocaleReadOptions {
	return &LocaleReadOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultTableWriter(),
	}
}

func NewCmdLocaleRead(s printers.IOStreams) *cobra.Command {
	o := NewLocaleReadOptions(s)
	var cmd = &cobra.Command{
		Use:     "get",
		Short:   "show all locales",
		Aliases: []string{"g", "l", "list"},
		Args:    cobra.NoArgs,
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
	//cmd.Flags().StringVar(&o.Code, "code", "", "code of a single locale")

	return cmd
}

// Complete the options
func (o *LocaleReadOptions) Complete(cmd *cobra.Command, args []string) error {
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
func (o *LocaleReadOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *LocaleReadOptions) Run() error {
	if o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}
	if o.FormatCategory() == "text" {
		o.WithDefaultOutput("table")
	}

	var result []schema.Locale

	//if o.Code != "" {
	//	r, err := o.Client.GetOneLocale(o.Code)
	//	if err != nil {
	//		return err
	//	}
	//	result = make([]schema.Locale, 1)
	//	result[0] = *r
	//} else {
	r, err := o.Client.GetAllLocales()
	if err != nil {
		return err
	}
	result = make([]schema.Locale, len(r))
	copy(result, r)
	//}

	return o.WithTableWriter("locales", func(t *tablewriter.Table) {
		t.SetHeader([]string{"uid", "code", "name", "fallback locale"})
		for _, row := range result {
			t.Append([]string{row.Uid, row.Code, row.Name, row.FallbackLocaleCode})
		}

	}).WriteOutput(result)
}

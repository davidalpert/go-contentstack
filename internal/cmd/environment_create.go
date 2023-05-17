package cmd

import (
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
	"strings"
)

type EnvironmentCreateOptions struct {
	*printers.PrinterOptions
	Config       cfg.Config
	Client       *management.Client
	Name         string
	RawURLParams []string
	URLsByLocale map[string]string
}

func NewEnvironmentCreateOptions(s printers.IOStreams) *EnvironmentCreateOptions {
	return &EnvironmentCreateOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
		URLsByLocale:   map[string]string{},
	}
}

func NewCmdEnvironmentCreate(s printers.IOStreams) *cobra.Command {
	o := NewEnvironmentCreateOptions(s)
	var cmd = &cobra.Command{
		Use:     "create <name>",
		Short:   "create a new publishing environment with the given name",
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
	cmd.Flags().StringSliceVarP(&o.RawURLParams, "url", "u", []string{}, "locale;url pairs describing render url by locale")

	return cmd
}

// Complete the options
func (o *EnvironmentCreateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.Name = args[0]

	for _, s := range o.RawURLParams {
		parts := strings.Split(s, ";")
		if len(parts) > 1 {
			o.URLsByLocale[parts[0]] = parts[1]
		}
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
func (o *EnvironmentCreateOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *EnvironmentCreateOptions) Run() error {
	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	w := schema.SingleEnvironmentWrapper{
		Environment: schema.Environment{
			Name: o.Name,
			Urls: []schema.LocaleURLPair{},
		},
	}

	for l, u := range o.URLsByLocale {
		w.Environment.Urls = append(w.Environment.Urls, schema.LocaleURLPair{Locale: l, Url: u})
	}

	r, err := o.Client.CreatePublishingEnvironment(&w)
	if err != nil {
		return err
	}

	return o.WriteOutput(r)
}

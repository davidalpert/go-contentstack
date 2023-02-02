package cmd

import (
	"context"
	"github.com/davidalpert/go-contentstack/internal/cfg"
	"github.com/davidalpert/go-contentstack/v1/management"
	"github.com/davidalpert/go-printers/v1"
	"github.com/spf13/cobra"
)

type GlobalGetOptions struct {
	*printers.PrinterOptions
	Config cfg.Config
}

func NewGlobalGetOptions(s printers.IOStreams) *GlobalGetOptions {
	return &GlobalGetOptions{
		PrinterOptions: printers.NewPrinterOptions().WithStreams(s).WithDefaultOutput("yaml"),
	}
}

func NewCmdGlobalGet(s printers.IOStreams) *cobra.Command {
	o := NewGlobalGetOptions(s)
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

	return cmd
}

// Complete the options
func (o *GlobalGetOptions) Complete(cmd *cobra.Command, args []string) error {
	if err := cfg.ReadMergedInto(&o.Config); err != nil {
		return err
	}

	return nil
}

// Validate the options
func (o *GlobalGetOptions) Validate() error {
	if err := o.Config.Validate(); err != nil {
		return err
	}

	return o.PrinterOptions.Validate()
}

// Run the command
func (o *GlobalGetOptions) Run() error {
	clientConfig := management.NewConfiguration()
	clientConfig.Servers = make([]management.ServerConfiguration, 1)
	clientConfig.Servers[0].URL = o.Config.ContentStack.ManagementApi.Host

	client := management.NewAPIClient(clientConfig)

	ctx := context.Background()
	ctx = context.WithValue(ctx, management.ContextAPIKeys, map[string]management.APIKey{
		"access_token": {Key: o.Config.ContentStack.ManagementApi.Token},
		"api_key":      {Key: o.Config.ContentStack.ManagementApi.Token},
	})

	client.GlobalFieldsApi.Getallglobalfields(ctx).Execute()

	if o.FormatCategory() == "table" || o.FormatCategory() == "csv" {
		o.WithDefaultOutput("json")
	}

	return o.WriteOutput(clientConfig)
}
